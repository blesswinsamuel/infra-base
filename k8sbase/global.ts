export interface GlobalProps {
  readonly domain: string;
  readonly certIssuer: string;
  readonly clusterExternalSecretStoreName: string;
  readonly internetAuthType: string;
}
