package k8sbase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"golang.org/x/exp/slices"
)

type GrafanaDashboardsProps struct {
	Enabled    bool                                    `yaml:"enabled"`
	Dashboards map[string]GrafanaDashboardsConfigProps `yaml:"dashboards"`
}

type GrafanaDashboardsConfigProps struct {
	Type     string               `yaml:"type"` // local or remote
	GlobPath *string              `yaml:"globPath"`
	URL      *[]DashboardURLProps `yaml:"urls"`
	Folder   string               `yaml:"folder"`
}

type DashboardURLProps struct {
	URL          string            `yaml:"url"`
	GnetID       *int              `yaml:"gnetId"`
	ID           string            `yaml:"id"`
	Title        *string           `yaml:"title"`
	Replacements map[string]string `yaml:"replacements"`
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%v", h.Sum32())
}

func GetCachedDashboard(url string) []byte {
	fileName := hash(url) + ".json"
	dashboardsCacheDir := fmt.Sprintf("%s/%s", helpers.CacheDir, "dashboards")
	if err := os.MkdirAll(dashboardsCacheDir, os.ModePerm); err != nil {
		log.Fatalln("GetCachedDashboard MkdirAll failed", err)
	}

	if _, err := os.Stat(dashboardsCacheDir + "/" + fileName); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("GetCachedDashboard downloading", url)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				panic(resp.Status)
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			if err := os.WriteFile(dashboardsCacheDir+"/"+fileName, data, 0644); err != nil {
				panic(err)
			}
		} else {
			log.Fatalln("GetCachedDashboard Stat failed", err)
		}
	}
	data, err := os.ReadFile(dashboardsCacheDir + "/" + fileName)
	if err != nil {
		panic(err)
	}
	return data
}

func NewGrafanaDashboards(scope constructs.Construct, props GrafanaDashboardsProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("grafana-dashboards"), &cprops)

	type dashboardItem struct {
		id              string
		dashboardConfig GrafanaDashboardsConfigProps
	}
	dashboardList := []dashboardItem{}
	for k, v := range props.Dashboards {
		dashboardList = append(dashboardList, dashboardItem{
			id:              k,
			dashboardConfig: v,
		})
	}
	slices.SortFunc(dashboardList, func(a dashboardItem, b dashboardItem) bool {
		return a.id < b.id
	})
	for _, item := range dashboardList {
		id := item.id
		dashboardConfig := item.dashboardConfig
		if dashboardConfig.GlobPath != nil {
			for _, filePath := range GetFilePaths(*dashboardConfig.GlobPath) {
				fileContents := GetFileContents(filePath)
				baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
				k8s.NewKubeConfigMap(chart, jsii.String(id+"-"+baseName), &k8s.KubeConfigMapProps{
					Metadata: &k8s.ObjectMeta{
						Name: jsii.String("grafana-dashboard-" + "-" + baseName + "-json"),
						// Name: jsii.String("grafana-dashboards-" + id + "-" + baseName),
						Labels: &map[string]*string{
							"grafana_dashboard": jsii.String("1"),
						},
						Annotations: &map[string]*string{
							"grafana_folder": jsii.String(dashboardConfig.Folder),
						},
					},
					Data: &map[string]*string{
						filepath.Base(filePath): jsii.String(fileContents),
					},
				})
			}
		}
		if dashboardConfig.URL != nil {
			for _, url := range *dashboardConfig.URL {
				dashboardContents := GetCachedDashboard(url.URL)
				dashboard := map[string]interface{}{}
				if err := json.Unmarshal(dashboardContents, &dashboard); err != nil {
					panic(err)
				}
				if url.GnetID != nil {
					dashboard["gnet_id"] = *url.GnetID
				}
				if url.Title != nil {
					dashboard["title"] = *url.Title
				}
				dashboard["uid"] = url.ID
				if dashboard["__inputs"] != nil {
					inputs := dashboard["__inputs"].([]any)
					for _, input := range inputs {
						input := input.(map[string]any)
						inputName := input["name"].(string)
						templating := dashboard["templating"].(map[string]any)
						templatingList := templating["list"].([]any)
						isAlreadyTemplated := false
						for _, templatingItem := range templatingList {
							templatingItem := templatingItem.(map[string]any)
							if templatingItem["name"] == inputName {
								isAlreadyTemplated = true
							}
						}
						if !isAlreadyTemplated {
							templatingList = append(templatingList, map[string]any{
								"hide":    0,
								"label":   "datasource",
								"name":    inputName,
								"options": []any{},
								"query":   "prometheus",
								"refresh": 1,
								"regex":   "",
								"type":    "datasource",
							})
						}
						templating["list"] = templatingList
					}
				}

				buf := new(bytes.Buffer)
				enc := json.NewEncoder(buf)
				enc.SetIndent("", "  ")
				if err := enc.Encode(dashboard); err != nil {
					panic(err)
				}
				outStr := buf.String()
				for k, v := range url.Replacements {
					outStr = strings.ReplaceAll(outStr, k, v)
				}

				k8s.NewKubeConfigMap(chart, jsii.String(url.ID), &k8s.KubeConfigMapProps{
					Metadata: &k8s.ObjectMeta{
						Name: jsii.String("grafana-dashboard-" + url.ID + "-json"),
						// Name: jsii.String("grafana-dashboards-" + id + "-" + baseName),
						Labels: &map[string]*string{
							"grafana_dashboard": jsii.String("1"),
						},
						Annotations: &map[string]*string{
							"grafana_folder": jsii.String(dashboardConfig.Folder),
						},
					},
					Data: &map[string]*string{
						url.ID + ".json": jsii.String(outStr),
					},
				})
			}
		}
	}

	return chart
}

func GetFilePaths(globPath string) []string {
	paths, err := filepath.Glob(globPath)
	if err != nil {
		panic(err)
	}
	slices.Sort(paths)
	return paths
}

func GetFileContents(path string) string {
	valuesFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(valuesFile)
}

// # [x] alertmanager-overview.json - "Alertmanager / Overview"
// # [x] apiserver.json - "Kubernetes / API server"
// # [ ] cluster-total.json - "Kubernetes / Networking / Cluster"
// # [ ] controller-manager.json - "Kubernetes / Controller Manager"
// # [x] grafana-overview.json - "Grafana Overview"
// # [ ] k8s-resources-cluster.json - "Kubernetes / Compute Resources / Cluster"
// # [ ] k8s-resources-namespace.json - "Kubernetes / Compute Resources / Namespace (Pods)"
// # [ ] k8s-resources-node.json - "Kubernetes / Compute Resources / Node (Pods)"
// # [ ] k8s-resources-pod.json - "Kubernetes / Compute Resources / Pod"
// # [ ] k8s-resources-workload.json - "Kubernetes / Compute Resources / Workload"
// # [ ] k8s-resources-workloads-namespace.json - "Kubernetes / Compute Resources / Namespace (Workloads)"
// # [ ] kubelet.json - "Kubernetes / Kubelet"
// # [ ] namespace-by-pod.json - "Kubernetes / Networking / Namespace (Pods)"
// # [ ] namespace-by-workload.json - "Kubernetes / Networking / Namespace (Workload)"
// # [ ] node-cluster-rsrc-use.json - "Node Exporter / USE Method / Cluster"
// # [ ] node-rsrc-use.json - "Node Exporter / USE Method / Node"
// # [x] nodes-darwin.json - "Node Exporter / MacOS"
// # [x] nodes.json - "Node Exporter / Nodes"
// # [x] persistentvolumesusage.json - "Kubernetes / Persistent Volumes"
// # [ ] pod-total.json - "Kubernetes / Networking / Pod"
// # [ ] proxy.json - "Kubernetes / Proxy"
// # [ ] scheduler.json - "Kubernetes / Scheduler"
// # [x] victoriametrics.json - "VictoriaMetrics"
// # [x] vmagent.json - "vmagent"
// # [x] vmalert.json - "vmalert"
// # [ ] workload-total.json - "Kubernetes / Networking / Workload"

// # {
// #     'source': 'https://raw.githubusercontent.com/prometheus-operator/kube-prometheus/main/manifests/grafana-dashboardDefinitions.yaml',
// #     'destination': '../dashboards/monitoring',
// #     'type': 'yaml',
// #     'dashboards': [
// #         'alertmanager-overview',
// #         'apiserver',
// #         'cluster-total',
// #         'controller-manager',
// #         'grafana-overview',
// #         'k8s-resources-cluster',
// #         'k8s-resources-namespace',
// #         'k8s-resources-node',
// #         'k8s-resources-pod',
// #         'k8s-resources-workload',
// #         'k8s-resources-workloads-namespace',
// #         'kubelet',
// #         'namespace-by-pod',
// #         'namespace-by-workload',
// #         'node-cluster-rsrc-use',
// #         'node-rsrc-use',
// #         'nodes-darwin',
// #         'nodes',
// #         'persistentvolumesusage',
// #         'pod-total',
// #         'prometheus-remote-write',
// #         'prometheus',
// #         'proxy',
// #         'scheduler',
// #         'workload-total',
// #     ],
// # },
