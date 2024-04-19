package k8sbase

import (
	"path"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type AlertingRulesProps struct {
	Rules map[string]AlertingRuleConfigProps `json:"rules"`
}

type AlertingRuleConfigProps struct {
	Type string `json:"type"` // local or remote
	// URLs *string              `json:"globPath"`
	URLs map[string]RuleURLProps `json:"urls"`
}

type RuleURLProps struct {
	URL          string            `json:"url"`
	SkipGroups   []string          `json:"skipGroups"`
	SkipAlerts   []string          `json:"skipAlerts"`
	Replacements map[string]string `json:"replacements"`
}

func (props *AlertingRulesProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	cprops := kubegogen.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("alerting-rules", cprops)

	rules := map[string]string{}
	cacheDir := k8sapp.GetGlobalContext(scope).CacheDir

	for _, rulesConfig := range props.Rules {
		if rulesConfig.URLs != nil {
			for _, rulesID := range infrahelpers.MapKeys(rulesConfig.URLs) {
				urlConfig := rulesConfig.URLs[rulesID]
				groups := []interface{}{}
				data := k8sapp.GetCachedFile(urlConfig.URL, path.Join(cacheDir, "alerts"))
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
				rules[rulesID+".yaml"] = string(outBytes)
			}
		}
	}
	k8sapp.NewConfigMap(chart, "config-map", &k8sapp.ConfigmapProps{
		Name: "alerting-rules",
		Data: rules,
	})
	return chart
}
