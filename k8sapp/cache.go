package k8sapp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
)

func GetCachedFile(fileUrl string, cacheDir string) []byte {
	// Parse the URL
	parsedURL, err := url.Parse(fileUrl)
	if err != nil {
		log.Panic().Err(err).Msg("GetCachedFile url.Parse failed")
	}

	if parsedURL.Scheme == "file" {
		return []byte(infrahelpers.GetFileContents(strings.TrimPrefix(parsedURL.Path, "/")))
	}

	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		log.Panic().Err(err).Msg("GetCachedFile MkdirAll failed")
	}

	lockedUrl := getLockedUrl(fileUrl)
	if lockedUrl == "" {
		lockedUrl = fileUrl
		if parsedURL.Host == "github.com" {
			pattern := `^/([^\/]+)\/([^\/]+)\/raw\/([^\/]+)\/(.+)`
			re := regexp.MustCompile(pattern)

			matches := re.FindStringSubmatch(parsedURL.Path)

			// Extract org, repo, branch, and filepath
			org := matches[1]
			repo := matches[2]
			branch := matches[3]
			filepath := matches[4]
			// log.Info().Str("org", org).Str("repo", repo).Str("branch", branch).Str("filepath", filepath).Msg("locking github file")
			branchSha := getGithubBranchSha(org, repo, branch)

			lockedUrl = fmt.Sprintf("%s://%s/%s/%s/raw/%s/%s", parsedURL.Scheme, parsedURL.Host, org, repo, branchSha, filepath)
			lockFile(fileUrl, lockedUrl)
		}
	}

	fileName := hash(lockedUrl) + ".json"
	cachedFilepath := path.Join(cacheDir, fileName)
	if _, err := os.Stat(cachedFilepath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Info().Str("url", lockedUrl).Msg("downloading")
			resp, err := http.Get(lockedUrl)
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
			if err := os.WriteFile(cachedFilepath, data, 0644); err != nil {
				panic(err)
			}
		} else {
			log.Panic().Err(err).Msg("GetCachedFile Stat failed")
		}
	}
	return []byte(infrahelpers.GetFileContents(cachedFilepath))
}

func getLockedUrl(fileUrl string) string {
	lockFilePath := path.Join("kube.lock")
	lockFile, err := os.Open(lockFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ""
		} else {
			log.Panic().Err(err).Msg("lockFile Open failed")
		}
	}
	defer lockFile.Close()
	lockFileContents, err := io.ReadAll(lockFile)
	if err != nil {
		log.Panic().Err(err).Msg("lockFile ReadAll failed")
	}
	lockFileParsed := infrahelpers.FromYamlString[map[string]string](string(lockFileContents))
	return lockFileParsed[fileUrl]
}

func lockFile(url string, lockedUrl string) {
	lockFilePath := path.Join("kube.lock")
	lockFile, err := os.Open(lockFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			lockFile, err = os.Create(lockFilePath)
			if err != nil {
				log.Panic().Err(err).Msg("lockFile Create failed")
			}
		} else {
			log.Panic().Err(err).Msg("lockFile Open failed")
		}
	}
	defer lockFile.Close()
	lockFileContents, err := io.ReadAll(lockFile)
	if err != nil {
		log.Panic().Err(err).Msg("lockFile ReadAll failed")
	}
	lockFileParsed := infrahelpers.FromYamlString[map[string]string](string(lockFileContents))
	if lockFileParsed[url] != lockedUrl {
		if lockFileParsed == nil {
			lockFileParsed = map[string]string{}
		}
		lockFileParsed[url] = lockedUrl
		lockFileContents := infrahelpers.ToYamlString(lockFileParsed)
		if err := os.WriteFile(lockFilePath, []byte(lockFileContents), 0644); err != nil {
			log.Panic().Err(err).Msg("lockFile WriteFile failed")
		}
	}
}

func getGithubBranchSha(org string, repo string, branch string) string {
	// Get the sha of the branch
	branchUrl := "https://api.github.com/repos/" + org + "/" + repo + "/branches/" + branch
	resp, err := http.Get(branchUrl)
	if err != nil {
		log.Panic().Err(err).Msg("getGithubBranchSha http.Get failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Panic().Str("status", resp.Status).Msg("getGithubBranchSha status code not 200")
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic().Err(err).Msg("getGithubBranchSha ReadAll failed")
	}
	branchData := infrahelpers.FromJSONString[map[string]any](string(data))
	return branchData["commit"].(map[string]any)["sha"].(string)
}
