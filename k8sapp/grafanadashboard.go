package k8sapp

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
	"time"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"golang.org/x/exp/slices"
)

type GrafanaDashboardProps struct {
	URL          string            `json:"url"`
	GnetID       *int              `json:"gnet_id"`
	Title        *string           `json:"title"`
	Replacements map[string]string `json:"replacements"`
	Folder       string            `json:"folder"`
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%v", h.Sum32())
}

func GetCachedDashboard(url string, cacheDir string) []byte {
	if strings.HasPrefix(url, "file://") {
		filePath := strings.TrimPrefix(url, "file://")
		return []byte(infrahelpers.GetFileContents(filePath))
	}

	dashboardsCacheDir := fmt.Sprintf("%s/%s", cacheDir, "dashboards")
	if err := os.MkdirAll(dashboardsCacheDir, os.ModePerm); err != nil {
		log.Fatalln("GetCachedDashboard MkdirAll failed", err)
	}

	date := time.Now().Format("2006-01-02")
	fileName := hash(date+url) + ".json"
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
	return []byte(infrahelpers.GetFileContents(dashboardsCacheDir + "/" + fileName))
}

func NewGrafanaDashboards(scope kubegogen.Construct, props map[string]GrafanaDashboardProps) kubegogen.Construct {
	for _, dashboardID := range infrahelpers.MapKeys(props) {
		dashboardProps := props[dashboardID]
		NewGrafanaDashboard(scope, dashboardID, dashboardProps)
	}
	return scope
}

func NewGrafanaDashboard(scope kubegogen.Construct, dashboardID string, props GrafanaDashboardProps) kubegogen.Construct {
	cacheDir := GetGlobalContext(scope).CacheDir
	dashboardContents := GetCachedDashboard(props.URL, cacheDir)
	dashboard := map[string]interface{}{}
	if err := json.Unmarshal(dashboardContents, &dashboard); err != nil {
		panic(err)
	}
	if props.GnetID != nil {
		dashboard["gnet_id"] = *props.GnetID
	}
	if props.Title != nil {
		dashboard["title"] = *props.Title
	}
	if dashboardID == "" {
		dashboardID = dashboard["uid"].(string)
	} else {
		dashboard["uid"] = dashboardID
	}
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
				// fmt.Println(input)
				query := input["pluginId"]
				label := input["label"]
				inputType := input["type"]
				templatingList = append(templatingList, map[string]any{
					"hide":    0,
					"label":   label,
					"name":    inputName,
					"options": []any{},
					"query":   query,
					"refresh": 1,
					"regex":   "",
					"type":    inputType,
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
	for k, v := range props.Replacements {
		outStr = strings.ReplaceAll(outStr, k, v)
	}

	return NewConfigMap(scope, dashboardID, &ConfigmapProps{
		Name: "grafana-dashboard-" + dashboardID + "-json",
		// Name: ("grafana-dashboards-" + id + "-" + baseName),
		Labels: map[string]string{
			"grafana_dashboard": "1",
		},
		Annotations: map[string]string{
			"grafana_folder": props.Folder,
		},
		Data: map[string]string{
			dashboardID + ".json": outStr,
		},
	})
}

func GetFilePaths(globPath string) []string {
	paths, err := filepath.Glob(globPath)
	if err != nil {
		panic(err)
	}
	slices.Sort(paths)
	return paths
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
