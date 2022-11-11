#!/usr/bin/env python3
import json
import re
import textwrap
from os import makedirs, path
import shutil
from pathlib import Path

import requests
import yaml
from yaml.representer import SafeRepresenter


dashboards = [
    dict(name='monitoring/alertmanager-overview', title='Alertmanager / Overview', id=9578, revision=4),
    # dict(name='monitoring/cluster-total', title='Kubernetes / Networking / Cluster', id=15761, revision=7),
    dict(name='monitoring/node-exporter', title='Node Exporter', id=1860, revision=29),
    dict(name='monitoring/grafana-internals', title='Grafana Internals', id=3590, revision=3),
    dict(name='monitoring/k8s-persistent-volumes', title='Kubernetes / Persistent Volumes', id=13646, revision=2),

    dict(name='monitoring/k8s-system-api-server', title='Kubernetes / System / API Server', id=15761, revision=10),
    dict(name='monitoring/k8s-system-coredns', title='Kubernetes / System / CoreDNS', id=15762, revision=7),
    dict(name='monitoring/k8s-views-global', title='Kubernetes / Views / Global', id=15757, revision=19),
    dict(name='monitoring/k8s-views-namespaces', title='Kubernetes / Views / Namespaces', id=15758, revision=14),
    dict(name='monitoring/k8s-views-nodes', title='Kubernetes / Views / Nodes', id=15759, revision=13),
    dict(name='monitoring/k8s-views-pods', title='Kubernetes / Views / Pods', id=15760, revision=14),
    
    # VictoriaMetrics - https://grafana.com/VictoriaMetrics
    dict(name='monitoring/victoriametrics-single', title='VictoriaMetrics / single', id=10229, revision=24),
    dict(name='monitoring/victoriametrics-vmagent', title='VictoriaMetrics / vmagent', id=12683, revision=10),
    dict(name='monitoring/victoriametrics-vmalert', title='VictoriaMetrics / vmalert', id=14950, revision=2),
]
# [x] alertmanager-overview.json - "Alertmanager / Overview"
# [x] apiserver.json - "Kubernetes / API server"
# [ ] cluster-total.json - "Kubernetes / Networking / Cluster"
# [ ] controller-manager.json - "Kubernetes / Controller Manager"
# [x] grafana-overview.json - "Grafana Overview"
# [ ] k8s-resources-cluster.json - "Kubernetes / Compute Resources / Cluster"
# [ ] k8s-resources-namespace.json - "Kubernetes / Compute Resources / Namespace (Pods)"
# [ ] k8s-resources-node.json - "Kubernetes / Compute Resources / Node (Pods)"
# [ ] k8s-resources-pod.json - "Kubernetes / Compute Resources / Pod"
# [ ] k8s-resources-workload.json - "Kubernetes / Compute Resources / Workload"
# [ ] k8s-resources-workloads-namespace.json - "Kubernetes / Compute Resources / Namespace (Workloads)"
# [ ] kubelet.json - "Kubernetes / Kubelet"
# [ ] namespace-by-pod.json - "Kubernetes / Networking / Namespace (Pods)"
# [ ] namespace-by-workload.json - "Kubernetes / Networking / Namespace (Workload)"
# [ ] node-cluster-rsrc-use.json - "Node Exporter / USE Method / Cluster"
# [ ] node-rsrc-use.json - "Node Exporter / USE Method / Node"
# [x] nodes-darwin.json - "Node Exporter / MacOS"
# [x] nodes.json - "Node Exporter / Nodes"
# [x] persistentvolumesusage.json - "Kubernetes / Persistent Volumes"
# [ ] pod-total.json - "Kubernetes / Networking / Pod"
# [ ] proxy.json - "Kubernetes / Proxy"
# [ ] scheduler.json - "Kubernetes / Scheduler"
# [x] victoriametrics.json - "VictoriaMetrics"
# [x] vmagent.json - "vmagent"
# [x] vmalert.json - "vmalert"
# [ ] workload-total.json - "Kubernetes / Networking / Workload"

def write_dashboard_to_file(resource_name, content, destination: Path):
    # recreate the file
    filename = resource_name + '.json'
    new_filename = destination / filename

    # make sure directories to store the file exist
    new_filename.parent.mkdir(parents=True, exist_ok=True)

    with open(destination / filename, 'w') as f:
        f.write(content)

    print("Generated %s" % new_filename)


def download_dashboards(destination_dir):
    shutil.rmtree(destination_dir, ignore_errors=True)

    # read the rules, create a new template file per group
    for dashboard in dashboards:
        name = dashboard['name']
        print("Generating dashboard from %s" % name)
        url = f"https://grafana.com/api/dashboards/{dashboard['id']}/revisions/{dashboard['revision']}/download"
        response = requests.get(url)
        if response.status_code != 200:
            print('Skipping the file, response code %s not equals 200' % response.status_code)
            continue
        raw_text = response.text

        dashboard_parsed = json.loads(raw_text)
        if 'title' in dashboard:
            dashboard_parsed['title'] = dashboard['title']
        dashboard_parsed['uid'] = dashboard['name'].replace('/', '-')
        dashboard_json = json.dumps(dashboard_parsed, sort_keys=True, indent=2)

        write_dashboard_to_file(name, dashboard_json, destination_dir)
    print("Finished")


def check_dashboard_updates(destination_dir):
    shutil.rmtree(destination_dir, ignore_errors=True)

    # read the rules, create a new template file per group
    for dashboard in dashboards:
        name = dashboard['name']
        url = f"https://grafana.com/api/dashboards/{dashboard['id']}"
        response = requests.get(url)
        if response.status_code != 200:
            print('Skipping the file, response code %s not equals 200' % response.status_code)
            continue
        raw_text = response.text

        dashboard_parsed = json.loads(raw_text)
        if dashboard['revision'] != dashboard_parsed['revision']:
            print(f"Dashboard '{dashboard['name']}' ({dashboard['id']}): {dashboard['revision']} -> {dashboard_parsed['revision']}")

    print("Finished")


def main():
    destination_dir = Path('../dashboards/').resolve()
    print('Updates Available:')
    check_dashboard_updates(destination_dir)
    print()
    
    print('Downloading dashboards')
    download_dashboards(destination_dir)


if __name__ == '__main__':
    main()