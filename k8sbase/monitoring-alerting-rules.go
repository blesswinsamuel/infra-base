package k8sbase

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type AlertingRulesProps struct {
	Rules map[string]AlertingRuleConfigProps `json:"rules"`
}

type AlertingRuleConfigProps struct {
	Type string `json:"type"` // local or remote
	// URLs *string              `json:"globPath"`
	URLs *[]RuleURLProps `json:"urls"`
}

type RuleURLProps struct {
	ID           string            `json:"id"`
	URL          string            `json:"url"`
	SkipGroups   []string          `json:"skipGroups"`
	SkipAlerts   []string          `json:"skipAlerts"`
	Replacements map[string]string `json:"replacements"`
}

func GetCachedAlertingRule(url string, cacheDir string) []byte {
	alertsCacheDir := fmt.Sprintf("%s/%s", cacheDir, "alerts")
	if err := os.MkdirAll(alertsCacheDir, os.ModePerm); err != nil {
		log.Fatalln("GetCachedAlertingRule MkdirAll failed", err)
	}

	baseName := filepath.Base(url)
	if _, err := os.Stat(alertsCacheDir + "/" + baseName); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("GetCachedAlertingRule downloading", url)
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
			if err := os.WriteFile(alertsCacheDir+"/"+baseName, data, 0644); err != nil {
				panic(err)
			}
		} else {
			log.Fatalln("GetCachedAlertingRule Stat failed", err)
		}
	}
	data, err := os.ReadFile(alertsCacheDir + "/" + baseName)
	if err != nil {
		panic(err)
	}
	return data
}

func NewAlertingRules(scope packager.Construct, props AlertingRulesProps) packager.Chart {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("alerting-rules", cprops)

	rules := map[string]string{}
	cacheDir := k8sapp.GetGlobalContext(scope).CacheDir

	for _, rulesConfig := range props.Rules {
		if rulesConfig.URLs != nil {
			for _, urlConfig := range *rulesConfig.URLs {
				groups := []interface{}{}
				data := GetCachedAlertingRule(urlConfig.URL, cacheDir)
				for k, v := range urlConfig.Replacements {
					data = []byte(strings.ReplaceAll(string(data), k, v))
				}
				yamlRules := map[string]interface{}{}
				if err := yaml.Unmarshal(data, &yamlRules); err != nil {
					panic(err)
				}
				if g, ok := yamlRules["groups"]; ok {
					groups = g.([]interface{})
				} else if g, ok := yamlRules["spec"]; ok {
					if g, ok := g.(map[string]interface{})["groups"]; ok {
						groups = g.([]interface{})
					}
				}
				groupsFiltered := []interface{}{}
				for _, group := range groups {
					rulesFiltered := []interface{}{}
					group := group.(map[string]interface{})
					groupName := group["name"].(string)
					if slices.Contains(urlConfig.SkipGroups, groupName) {
						continue
					}
					rules := group["rules"].([]interface{})
					for _, rule := range rules {
						rule := rule.(map[string]interface{})
						ruleName, _ := rule["alert"].(string)
						if slices.Contains(urlConfig.SkipAlerts, ruleName) {
							continue
						}
						rulesFiltered = append(rulesFiltered, rule)
					}
					group["rules"] = rulesFiltered
					if len(rulesFiltered) == 0 {
						continue
					}
					groupsFiltered = append(groupsFiltered, group)
					// buf := bytes.NewBufferString("")
					// enc := yaml.NewEncoder(buf)
					// enc.
				}
				outBytes, err := yaml.Marshal(map[string]interface{}{"groups": groupsFiltered})
				if err != nil {
					panic(err)
				}
				rules[urlConfig.ID+".yaml"] = string(outBytes)
			}
		}
	}
	k8sapp.NewConfigMap(chart, jsii.String("config-map"), &k8sapp.ConfigmapProps{
		Name: "alerting-rules",
		Data: rules,
	})
	return chart
}
