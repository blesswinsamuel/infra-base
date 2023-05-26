import { Construct } from "constructs";
import { GlobalProps } from "./global.ts";
import { Secrets, SecretsProps } from "./secrets.ts";

export interface BaseProps {
  // Global     GlobalProps     `yaml:"global"`
  // Ingress    IngressProps    `yaml:"ingress"`
  // System     SystemProps     `yaml:"system"`
  // Secrets    SecretsProps    `yaml:"secrets"`
  // Auth       AuthProps       `yaml:"auth"`
  // Monitoring MonitoringProps `yaml:"monitoring"`
  // Databases  DatabaseProps   `yaml:"databases"`

  readonly global: GlobalProps;
  readonly secrets: SecretsProps;
}

export class Base extends Construct {
  constructor(scope: Construct, id: string, props: BaseProps) {
    super(scope, id);

    // secrets
    new Secrets(this, "secrets", props.secrets);

    // // ingress
    // new Ingress(this, "ingress", props.ingress);

    // // system
    // new System(this, "system", props.system);

    // // monitoring
    // new Monitoring(this, "monitoring", props.monitoring);

    // // database
    // new Database(this, "databases", props.databases);

    // // auth
    // new Auth(this, "auth", props.auth);
  }
}
