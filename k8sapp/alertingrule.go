package k8sapp

import (
	"path"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"golang.org/x/exp/slices"
)

type AlertingRuleProps struct {
	// 	// URLs *string              `json:"globPath"`
	// 	URLs map[string]RuleURLProps `json:"urls"`
	// }

	// type RuleURLProps struct {
	URL          string            `json:"url"`
	SkipGroups   []string          `json:"skipGroups"`
	SkipAlerts   []string          `json:"skipAlerts"`
	Replacements map[string]string `json:"replacements"`
}

func NewAlertingRules(scope kubegogen.Construct, props map[string]AlertingRuleProps) kubegogen.Construct {
	for _, alertingRuleID := range infrahelpers.MapKeys(props) {
		dashboardProps := props[alertingRuleID]
		NewAlertingRule(scope, alertingRuleID, dashboardProps)
	}
	return scope
}

func NewAlertingRule(scope kubegogen.Construct, alertingRuleID string, props AlertingRuleProps) kubegogen.Construct {
	cacheDir := GetGlobalContext(scope).CacheDir
	groups := []any{}
	data := GetCachedFile(props.URL, path.Join(cacheDir, "alerting-rules"))
	for k, v := range props.Replacements {
		data = []byte(strings.ReplaceAll(string(data), k, v))
	}
	yamlRules := infrahelpers.FromYamlString[map[string]any](string(data))
	if g, ok := yamlRules["groups"]; ok {
		groups = g.([]any)
	} else if g, ok := yamlRules["spec"]; ok {
		if g, ok := g.(map[string]any)["groups"]; ok {
			groups = g.([]any)
		}
	}
	groupsFiltered := []any{}
	for _, group := range groups {
		rulesFiltered := []any{}
		group := group.(map[string]any)
		groupName := group["name"].(string)
		if slices.Contains(props.SkipGroups, groupName) {
			continue
		}
		rules := group["rules"].([]any)
		for _, rule := range rules {
			rule := rule.(map[string]any)
			ruleName, _ := rule["alert"].(string)
			if slices.Contains(props.SkipAlerts, ruleName) {
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
	outStr := infrahelpers.ToYamlString(map[string]any{"groups": groupsFiltered})

	return NewConfigMap(scope, alertingRuleID, &ConfigmapProps{
		Name: "alerting-rule-" + alertingRuleID,
		Labels: map[string]string{
			"alerting_rule": "1",
		},
		Data: map[string]string{
			alertingRuleID + ".yaml": outStr,
		},
	})
}
