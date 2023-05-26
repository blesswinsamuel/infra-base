import { Helm } from "cdk8s";
import { Construct } from "constructs";
import * as fs from "https://deno.land/std@0.177.1/fs/mod.ts";
import { execSync } from "https://deno.land/std@0.177.1/node/child_process.ts";
import process from "https://deno.land/std@0.177.1/node/process.ts";
import * as path from "https://deno.land/std@0.177.1/path/mod.ts";

export interface ImageInfo {
  readonly repository: string;
  readonly tag: string;
}

export interface ChartInfo {
  readonly repo: string;
  readonly chart: string;
  readonly version: string;
}

export interface HelmProps {
  readonly chartInfo: ChartInfo;
  readonly chartFileNamePrefix?: string;
  readonly releaseName: string;
  readonly namespace?: string;
  // deno-lint-ignore no-explicit-any
  readonly values?: { [key: string]: any };
}

const cacheDir = process.env.CACHE_DIR || "./cache";

export class HelmCached extends Construct {
  constructor(scope: Construct, id: string, props: HelmProps) {
    super(scope, id);

    const chartsCacheDir = path.join(cacheDir, "charts");
    fs.ensureDirSync(chartsCacheDir);

    if (props.chartInfo.repo === undefined) {
      throw new Error(
        `props.chartInfo.repo is undefined for ${props.releaseName}`,
      );
    }

    let chartFileName =
      `${props.chartInfo.chart}-${props.chartInfo.version}.tgz`;
    if (props.chartFileNamePrefix) {
      chartFileName =
        `${props.chartFileNamePrefix}-${props.chartInfo.version}.tgz`;
    }

    const chartPath = `${chartsCacheDir}/${chartFileName}`;
    const cmd =
      `helm pull ${props.chartInfo.chart} --repo ${props.chartInfo.repo} --destination ${chartsCacheDir} --version ${props.chartInfo.version}`;
    console.log(`cmd: ${cmd}`);
    try {
      const out = execSync(cmd, {});
      console.log(
        `Fetching chart '${props.chartInfo.chart}' from repo '${props.chartInfo.repo}' version '${props.chartInfo.version}' ...`,
      );
      if (out != null && out.length > 0) {
        console.log(out.toString());
      }
    } catch (err) {
      throw new Error(`Error occured during helm pull command: ${err}`);
    }

    new Helm(this, id, {
      chart: chartPath,
      releaseName: props.releaseName,
      namespace: props.namespace,
      values: props.values,
      helmFlags: ["--include-crds", "--skip-tests", "--no-hooks"],
    });
  }
}
