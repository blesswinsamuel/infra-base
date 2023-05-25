import { App, YamlOutputType } from "cdk8s";

const app = new App({
  // Instead of the default "dist"
  outdir: "output",
  // Instead of ".k8s.yaml"
  outputFileExtension: ".generated.yaml",
  // Divide every resource into its own file, instead of grouping by Chart
  yamlOutputType: YamlOutputType.FILE_PER_RESOURCE,
});
app.synth();
