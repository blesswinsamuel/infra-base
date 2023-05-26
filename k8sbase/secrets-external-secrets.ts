import { Chart } from "cdk8s";
import { Construct } from "constructs";
import { ChartInfo, HelmCached } from "./utils/helm.ts";

export interface ExternalSecretsProps {
  enabled: boolean;
  helm: ChartInfo;
}

// https://github.com/external-secrets/external-secrets/tree/main/deploy/charts/external-secrets
export class ExternalSecrets extends Construct {
  constructor(scope: Construct, id: string, props: ExternalSecretsProps) {
    super(scope, id);
    if (!props.enabled) {
      return;
    }

    const chart = new Chart(this, "external-secrets", {
      namespace: "external-secrets",
    });

    new HelmCached(this, "helm", {
      chartInfo: props.helm,
      releaseName: "external-secrets",
      namespace: chart.namespace,
      values: {
        installCRDs: true,
      },
    });
  }
}
