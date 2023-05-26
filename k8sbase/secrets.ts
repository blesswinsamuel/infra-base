import { Construct } from "constructs";
// import { ClusterSecretStoreProps } from "./cluster-secret-store";
// import { ExternalSecretsProps } from "./external-secrets";
import { KubeNamespace } from "./imports/k8s.ts";
import {
  ExternalSecrets,
  ExternalSecretsProps,
} from "./secrets-external-secrets.ts";
// import { SecretsDockerCredsProps } from "./secrets-docker-creds";

export interface SecretsProps {
  readonly externalSecrets: ExternalSecretsProps;
  //   readonly dockerCreds: SecretsDockerCredsProps;
  //   readonly clusterSecretStore: ClusterSecretStoreProps;
}

export class Secrets extends Construct {
  constructor(scope: Construct, id: string, props: SecretsProps) {
    super(scope, id);

    new KubeNamespace(this, "secrets");

    new ExternalSecrets(this, "externalSecrets", props.externalSecrets);

    // new SecretsDockerCreds(this, "dockerCreds", props.dockerCreds);
    // new ClusterSecretStore(
    //   this,
    //   "clusterSecretStore",
    //   props.clusterSecretStore,
    // );
  }
}
