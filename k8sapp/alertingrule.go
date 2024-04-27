package k8sapp

import (
	"path"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"golang.org/x/exp/slices"
)

type AlertingRules struct {
	URL               string                       `json:"url"`
	SkipGroups        []string                     `json:"skipGroups"`
	SkipAlerts        []string                     `json:"skipAlerts"`
	Replacements      map[string]string            `json:"replacements"`
	AlertReplacements map[string]map[string]string `json:"alertReplacements"`
}

func NewAlertingRules(scope kubegogen.Scope, props map[string]AlertingRules) kubegogen.Scope {
	for _, alertingRuleID := range infrahelpers.MapKeys(props) {
		dashboardProps := props[alertingRuleID]
		NewAlertingRule(scope, alertingRuleID, dashboardProps)
	}
	return scope
}

func NewAlertingRule(scope kubegogen.Scope, alertingRuleID string, props AlertingRules) {
	cacheDir := GetConfig(scope).CacheDir
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
			if alertReplacements, ok := props.AlertReplacements[ruleName]; ok {
				ruleStr := infrahelpers.ToYamlString(rule)
				for k, v := range alertReplacements {
					ruleStr = strings.ReplaceAll(ruleStr, k, v)
				}
				rule = infrahelpers.FromYamlString[map[string]any](ruleStr)
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

	NewConfigMap(scope, &ConfigmapProps{
		Name: "alerting-rule-" + alertingRuleID,
		Labels: map[string]string{
			"alerting_rule": "1",
		},
		Data: map[string]string{
			alertingRuleID + ".yaml": outStr,
		},
	})
}
