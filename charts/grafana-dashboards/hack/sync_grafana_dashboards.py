#!/usr/bin/env python3
"""Fetch dashboards from provided urls into this chart."""
"""https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/hack/sync_grafana_dashboards.py"""
"""https://github.com/VictoriaMetrics/helm-charts/blob/master/charts/victoria-metrics-k8s-stack/hack/sync_grafana_dashboards.py"""
import json
import re
import textwrap
from os import makedirs, path
import shutil

import requests
import yaml
from yaml.representer import SafeRepresenter


# https://stackoverflow.com/a/20863889/961092
class LiteralStr(str):
    pass


def change_style(style, representer):
    def new_representer(dumper, data):
        scalar = representer(dumper, data)
        scalar.style = style
        return scalar

    return new_representer


# Source files list
charts = [
    {
        'source': 'https://raw.githubusercontent.com/prometheus-operator/kube-prometheus/main/manifests/grafana-dashboardDefinitions.yaml',
        'destination': '../dashboards/monitoring',
        'type': 'yaml',
    },
    {
        'source': 'https://raw.githubusercontent.com/etcd-io/website/master/content/en/docs/v3.5/op-guide/grafana.json',
        'destination': '../dashboards/other',
        'type': 'json',
    },
    {
        'source': 'https://raw.githubusercontent.com/VictoriaMetrics/VictoriaMetrics/master/dashboards/victoriametrics.json',
        'destination': '../dashboards/monitoring',
        'type': 'json',
    },
    {
        'source': 'https://raw.githubusercontent.com/VictoriaMetrics/VictoriaMetrics/master/dashboards/vmagent.json',
        'name': 'vmagent',
        'destination': '../dashboards/monitoring',
        'type': 'json',
    },
    {
        'source': 'https://raw.githubusercontent.com/VictoriaMetrics/VictoriaMetrics/master/dashboards/vmalert.json',
        'name': 'vmalert',
        'destination': '../dashboards/monitoring',
        'type': 'json',
    },
]

skip_list = [
    "prometheus.json",
    "prometheus-remote-write.json"
]

# Additional conditions map
# condition_map = {
#     'alertmanager-overview': '.Values.coreDns.enabled',
#     'grafana-coredns-k8s': '.Values.coreDns.enabled',
#     'etcd': '.Values.kubeEtcd.enabled',
#     'apiserver': '.Values.kubeApiServer.enabled',
#     'controller-manager': '.Values.kubeControllerManager.enabled',
#     'kubelet': '.Values.kubelet.enabled',
#     'proxy': '.Values.kubeProxy.enabled',
#     'scheduler': '.Values.kubeScheduler.enabled',
#     'node-rsrc-use': '.Values.nodeExporter.enabled',
#     'node-cluster-rsrc-use': '.Values.nodeExporter.enabled',
#     'nodes': '.Values.nodeExporter.enabled',
#     'prometheus-remote-write': '.Values.prometheus.prometheusSpec.remoteWriteDashboards'
# }

def write_group_to_file(resource_name, content, url, destination):
    # recreate the file
    filename = resource_name + '.json'
    new_filename = "%s/%s" % (destination, filename)

    # make sure directories to store the file exist
    makedirs(destination, exist_ok=True)

    with open(new_filename, 'w') as f:
        f.write(content)

    print("Generated %s" % new_filename)


def main():
    for destination in set(chart['destination'] for chart in charts):
        shutil.rmtree(destination, ignore_errors=True)

    # read the rules, create a new template file per group
    for chart in charts:
        print("Generating rules from %s" % chart['source'])
        response = requests.get(chart['source'])
        if response.status_code != 200:
            print('Skipping the file, response code %s not equals 200' % response.status_code)
            continue
        raw_text = response.text

        if ('max_kubernetes' not in chart):
            chart['max_kubernetes']="9.9.9-9"

        if chart['type'] == 'yaml':
            yaml_text = yaml.full_load(raw_text)
            groups = yaml_text['items']
            for group in groups:
                for resource, content in group['data'].items():
                    if resource in skip_list:
                        continue
                    write_group_to_file(resource.replace('.json', ''), content, chart['source'], chart['destination'])
        elif chart['type'] == 'json':
            json_text = json.loads(raw_text)
            # is it already a dashboard structure or is it nested (etcd case)?
            flat_structure = bool(json_text.get('annotations'))
            if flat_structure:
                resource = path.basename(chart['source']).replace('.json', '')
                write_group_to_file(resource, json.dumps(json_text, indent=4), chart['source'], chart['destination'])
            else:
                for resource, content in json_text.items():
                    write_group_to_file(resource.replace('.json', ''), json.dumps(content, indent=4), chart['source'], chart['destination'])
    print("Finished")


if __name__ == '__main__':
    main()