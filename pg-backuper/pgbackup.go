package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"text/template"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/goccy/go-yaml"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Schedule string `json:"schedule"`
}

type Notify struct {
	URL string `json:"url"`
}

type BackupDestinationFileSystem struct {
	PathTemplate       string `json:"pathTemplate"`
	pathTemplateParsed *template.Template
}

type BackupDestinationS3 struct {
	Bucket             string `json:"bucket"`
	Endpoint           string `json:"endpoint"`
	AccessKey          string `json:"accessKey"`
	SecretKey          string `json:"secretKey"`
	PathTemplate       string `json:"pathTemplate"`
	pathTemplateParsed *template.Template
}

type BackupDestination struct {
	FileSystem    *BackupDestinationFileSystem `json:"filesystem"`
	S3            *BackupDestinationS3         `json:"s3"`
	EncryptionKey string                       `json:"encryptionKey"`
}

type Config struct {
	BackupDestinations []BackupDestination `json:"backupDestinations"`
	Databases          map[string]Database `json:"databases"`
	Notify             []Notify            `json:"notify"`
}

var (
	backupCountTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pg_backuper_backup_count_total",
		Help: "The total number of backups",
	}, []string{"database", "status"})
	backupDurationSeconds = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pg_backuper_backup_duration_seconds",
		Help:    "The duration of backups",
		Buckets: prometheus.DefBuckets,
	}, []string{"database"})
)

type jobStatusUpdate struct {
	name   string
	status gocron.JobStatus
}

type SchedulerMonitor struct {
	Notify             []Notify
	jobStatusUpdatesCh chan jobStatusUpdate
}

func NewSchedulerMonitor(notify []Notify) *SchedulerMonitor {
	return &SchedulerMonitor{
		Notify:             notify,
		jobStatusUpdatesCh: make(chan jobStatusUpdate),
	}
}

var _ gocron.Monitor = (*SchedulerMonitor)(nil)

func (m *SchedulerMonitor) IncrementJob(id uuid.UUID, name string, tags []string, status gocron.JobStatus) {
	backupCountTotal.WithLabelValues(name, string(status)).Inc()
	m.jobStatusUpdatesCh <- jobStatusUpdate{name: name, status: status}
}

func (m *SchedulerMonitor) RecordJobTiming(startTime, endTime time.Time, id uuid.UUID, name string, tags []string) {
	backupDurationSeconds.WithLabelValues(name).Observe(endTime.Sub(startTime).Seconds())
}

func (m *SchedulerMonitor) StartNotifyJob() {
	jobStatus := map[string]gocron.JobStatus{}
	// lastSuccessfulTime := map[string]time.Time{}
	for update := range m.jobStatusUpdatesCh {
		log.Info().Str("name", update.name).Str("status", string(update.status)).Msg("Job status update")
		// if update.status == gocron.Success {
		// 	lastSuccessfulTime[update.name] = time.Now()
		// }
		// jobStatus[update.name] = update.status
		isAllSuccessful := true
		for _, status := range jobStatus {
			// lastTime := lastSuccessfulTime[name]
			// if time.Since(lastTime) > 12*time.Hour {
			if status == gocron.Fail {
				isAllSuccessful = false
			}
		}
		if isAllSuccessful {
			for _, notify := range m.Notify {
				log.Info().Str("url", notify.URL).Msg("Sending notification")
				if err := m.SendNotification(notify.URL); err != nil {
					log.Error().Err(err).Msg("Failed to send notification")
				}
			}
		}
	}
}

func (m *SchedulerMonitor) SendNotification(url string) error {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	app := &cli.App{
		Name:  "pg-backuper",
		Usage: "Backup PostgreSQL databases",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config-file",
				Value:   "config.yaml",
				EnvVars: []string{"CONFIG_FILE"},
				Usage:   "Config file path",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start scheduled backups",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "db",
						Usage: "Database name",
					},
					&cli.StringFlag{
						Name:  "metrics-port",
						Value: "2112",
						Usage: "Metrics port",
					},
				},
				Action: func(cCtx *cli.Context) error {
					cfg, err := parseConfig(cCtx.String("config-file"))
					if err != nil {
						return err
					}
					databases := cfg.Databases
					if cCtx.IsSet("db") {
						dbName := cCtx.String("db")
						db, ok := cfg.Databases[dbName]
						if !ok {
							return fmt.Errorf("database %s not found", dbName)
						}
						databases = map[string]Database{dbName: db}
					}
					jobMonitor := NewSchedulerMonitor(cfg.Notify)
					go jobMonitor.StartNotifyJob()
					// create a scheduler
					s, err := gocron.NewScheduler(
						gocron.WithMonitor(jobMonitor),
						gocron.WithLocation(time.UTC),
						gocron.WithLimitConcurrentJobs(1, gocron.LimitModeWait),
					)
					if err != nil {
						return err
					}
					jobs := make(map[string]gocron.Job)
					printNextRuns := func() {
						for name, j := range jobs {
							nextRun, err := j.NextRun()
							if err != nil {
								log.Error().Str("job", name).Err(err).Msg("Failed to get next run")
								continue
							}
							log.Info().
								Str("job_name", j.Name()).
								Stringer("job_id", j.ID()).
								Str("database", name).
								Str("schedule", databases[name].Schedule).
								Time("next_run", nextRun).
								Stringer("next_run_in", time.Until(nextRun)).
								Msg("Next run in")
						}
					}
					for name, db := range databases {
						j, err := s.NewJob(
							gocron.CronJob(db.Schedule, false),
							gocron.NewTask(
								func(db Database) {
									if err := createDbDump(db, cfg.BackupDestinations); err != nil {
										log.Error().Str("job", name).Err(err).Msg("Failed to create dump")
										return
									}
								},
								db,
							),
							gocron.WithName(db.Database),
						)
						if err != nil {
							return err
						}
						jobs[name] = j
						nextRun, err := j.NextRun()
						if err != nil {
							return err
						}
						log.Info().
							Str("job_name", j.Name()).
							Stringer("job_id", j.ID()).
							Str("database", name).
							Str("schedule", db.Schedule).
							Time("next_run", nextRun).
							Stringer("next_run_in", time.Since(nextRun)).
							Msg("Job scheduled")
					}
					s.Start()
					printNextRuns()

					http.Handle("/metrics", promhttp.Handler())
					http.Handle("/run-now", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						if r.Method != http.MethodPost {
							http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
							return
						}
						for name, j := range jobs {
							log.Info().Str("job_name", j.Name()).Stringer("job_id", j.ID()).Msg("Manually started job")
							if err := j.RunNow(); err != nil {
								log.Error().Str("job", name).Err(err).Msg("Failed to run job")
							}
						}
					}))
					go http.ListenAndServe(":"+cCtx.String("metrics-port"), nil)

					cancelChan := make(chan os.Signal, 1)
					signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
					<-cancelChan
					err = s.Shutdown()
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "backup",
				Usage:   "Backup databases",
				Aliases: []string{"b"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "db",
						Usage: "Database name",
					},
				},
				Action: func(cCtx *cli.Context) error {
					cfg, err := parseConfig(cCtx.String("config-file"))
					if err != nil {
						return err
					}
					databases := cfg.Databases
					if cCtx.IsSet("db") {
						dbName := cCtx.String("db")
						db, ok := cfg.Databases[dbName]
						if !ok {
							return fmt.Errorf("database %s not found", dbName)
						}
						databases = map[string]Database{dbName: db}
					}
					for _, db := range databases {
						log.Info().
							Any("database", db.Database).
							Any("username", db.Username).
							Any("password_length", len(db.Password)).
							Any("host", db.Host).
							Any("port", db.Port).
							Any("schedule", db.Schedule).
							Msg("Backing up")
						if err := createDbDump(db, cfg.BackupDestinations); err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Failed to run app")
	}
}

func parseConfig(configFilePath string) (Config, error) {
	configBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}
	cfg := Config{}
	if err := yaml.UnmarshalWithOptions(configBytes, &cfg, yaml.UseJSONUnmarshaler()); err != nil {
		return Config{}, err
	}
	for name, db := range cfg.Databases {
		if db.Database == "" {
			db.Database = name
		}
		cfg.Databases[name] = db
	}
	for _, dest := range cfg.BackupDestinations {
		if dest.FileSystem != nil {
			tmpl, err := template.New("pathTemplate").Parse(dest.FileSystem.PathTemplate)
			if err != nil {
				return Config{}, err
			}
			dest.FileSystem.pathTemplateParsed = tmpl
		}
		if dest.S3 != nil {
			tmpl, err := template.New("pathTemplate").Parse(dest.S3.PathTemplate)
			if err != nil {
				return Config{}, err
			}
			dest.S3.pathTemplateParsed = tmpl
		}
	}
	return cfg, nil
}

func createDbDump(db Database, backupDestinations []BackupDestination) error {
	log.Info().Str("database", db.Database).Msg("Creating dump")
	// https://stackoverflow.com/questions/74708724/dump-and-restore-a-postgres-database-using-pg-dump-and-pg-restore
	tmpFilePath := path.Join(os.TempDir(), db.Database+".pgdump")
	defer os.Remove(tmpFilePath)
	cmd := exec.Command("pg_dump", "-Fc", "-f", tmpFilePath)
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("PGUSER=%s", db.Username),
		fmt.Sprintf("PGPASSWORD=%s", db.Password),
		fmt.Sprintf("PGDATABASE=%s", db.Database),
		fmt.Sprintf("PGHOST=%s", db.Host),
		fmt.Sprintf("PGPORT=%d", db.Port),
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to create dump for %s: %s: %v", db.Database, out, err)
	} else {
		log.Info().Bytes("out", out).Str("database", db.Database).Str("filename", tmpFilePath).Msg("Dump created")
	}
	for _, dest := range backupDestinations {
		if dest.EncryptionKey != "" {
			log.Info().Str("database", db.Database).Str("path", tmpFilePath).Msg("Encrypting dump")
			if err := encryptFile(tmpFilePath, tmpFilePath+".enc", dest.EncryptionKey); err != nil {
				return fmt.Errorf("failed to encrypt dump: %v", err)
			}
			tmpFilePath = tmpFilePath + ".enc"
		}

		if dest.FileSystem != nil {
			if err := writeToFilesystemBackupDestination(tmpFilePath, db, *dest.FileSystem); err != nil {
				return err
			}
		}
		if dest.S3 != nil {
			if err := writeToS3BackupDestination(tmpFilePath, db, *dest.S3); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeToFilesystemBackupDestination(sourceFilePath string, db Database, dest BackupDestinationFileSystem) error {
	templatePathBuf := &bytes.Buffer{}
	err := dest.pathTemplateParsed.Execute(templatePathBuf, map[string]string{
		"Database":  db.Database,
		"Timestamp": time.Now().Format("2006-01-02T15:04:05"),
	})
	if err != nil {
		return fmt.Errorf("failed to execute dump path template: %v", err)
	}
	dumpFilePath := templatePathBuf.String()
	log.Info().Str("database", db.Database).Str("path", dumpFilePath).Msg("Copying dump to filesystem")

	// check if directory exists
	parentDir := filepath.Dir(dumpFilePath)
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		log.Info().Str("path", parentDir).Msg("Creating parent directory")
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return fmt.Errorf("failed to create parent directory %s: %v", parentDir, err)
		}
	} else {
		log.Debug().Str("path", parentDir).Msg("Parent directory exists")
	}

	// copy file to destination
	size, err := copyFile(sourceFilePath, dumpFilePath)
	if err != nil {
		return fmt.Errorf("failed to copy dump to filesystem: %v", err)
	}
	log.Info().Str("database", db.Database).Str("path", dumpFilePath).Any("size", size).Msg("Dump copied to filesystem")
	return nil
}

func writeToS3BackupDestination(sourceFilePath string, db Database, dest BackupDestinationS3) error {
	templatePathBuf := &bytes.Buffer{}
	err := dest.pathTemplateParsed.Execute(templatePathBuf, map[string]string{
		"Database":  db.Database,
		"Timestamp": time.Now().Format("2006-01-02T15:04:05"),
	})
	if err != nil {
		return fmt.Errorf("failed to execute dump path template: %v", err)
	}
	dumpFilePath := templatePathBuf.String()
	log.Info().Str("database", db.Database).Str("path", dumpFilePath).Msg("Copying dump to S3")
	// copy file to destination

	// Initialize minio client object.
	minioClient, err := minio.New(dest.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(dest.AccessKey, dest.SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		return err
	}
	info, err := minioClient.FPutObject(context.TODO(), dest.Bucket, dumpFilePath, sourceFilePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	log.Info().Str("database", db.Database).Str("path", dumpFilePath).Str("etag", info.ETag).Any("size", info.Size).Msg("Dump copied to S3")
	return nil
}

func copyFile(src, dst string) (int64, error) {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return 0, err
	}
	// Sync to disk
	err = dstFile.Sync()
	if err != nil {
		return 0, err
	}
	return n, nil
}

func encryptFile(src, dst, key string) error {
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-salt", "-in", src, "-out", dst, "-k", key)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to encrypt file: %s: %v", out, err)
	}
	return nil
}
