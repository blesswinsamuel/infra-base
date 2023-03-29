#!/usr/bin/env python3
import json
import re
import textwrap
import os
import sys
import shutil
from pathlib import Path

import requests
import yaml
from yaml.representer import SafeRepresenter


def write_dashboard_to_file(resource_name, content, destination: Path):
    # recreate the file
    filename = resource_name + '.json'
    new_filename = destination / filename

    # make sure directories to store the file exist
    new_filename.parent.mkdir(parents=True, exist_ok=True)

    with open(destination / filename, 'w') as f:
        f.write(content)

    print("Generated %s" % new_filename)


def download_dashboards(dashboards, destination_dir):
    shutil.rmtree(destination_dir, ignore_errors=True)

    # read the rules, create a new template file per group
    for dashboard in dashboards:
        from_url = dashboard['from']
        name = dashboard['to']
        print("Generating dashboard from %s" % from_url)
        response = requests.get(from_url)
        if response.status_code != 200:
            print('Skipping the file, response code %s not equals 200' % response.status_code)
            continue
        raw_text = response.text

        dashboard_parsed = json.loads(raw_text)
        if 'gnet_id' in dashboard:
            dashboard_parsed['gnetId'] = dashboard['gnet_id']
        if 'title' in dashboard:
            dashboard_parsed['title'] = dashboard['title']
        dashboard_parsed['uid'] = dashboard['to'].replace('/', '-')
        dashboard_json = json.dumps(dashboard_parsed, sort_keys=True, indent=2)
        if 'replacements' in dashboard:
            for old, new in dashboard['replacements'].items():
                dashboard_json = dashboard_json.replace(old, new)
        write_dashboard_to_file(name, dashboard_json, destination_dir)
    print("Finished")


def main():
    with open(sys.argv[1]) as f:
        config = yaml.load(f, Loader=yaml.FullLoader)
    destination_dir = config['dashboards_destination_dir']
    
    print('Downloading dashboards')
    download_dashboards(config['dashboards'], Path(destination_dir))


if __name__ == '__main__':
    main()