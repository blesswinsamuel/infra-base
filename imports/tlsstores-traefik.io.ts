// generated by cdk8s
import { ApiObject, ApiObjectMetadata, GroupVersionKind } from 'cdk8s';
import { Construct } from 'constructs';


/**
 * TLSStore is the CRD implementation of a Traefik TLS Store. For the time being, only the TLSStore named default is supported. This means that you cannot have two stores that are named default in different Kubernetes namespaces. More info: https://doc.traefik.io/traefik/v2.10/https/tls/#certificates-stores
 *
 * @schema TLSStore
 */
export class TlsStore extends ApiObject {
  /**
   * Returns the apiVersion and kind for "TLSStore"
   */
  public static readonly GVK: GroupVersionKind = {
    apiVersion: 'traefik.io/v1alpha1',
    kind: 'TLSStore',
  }

  /**
   * Renders a Kubernetes manifest for "TLSStore".
   *
   * This can be used to inline resource manifests inside other objects (e.g. as templates).
   *
   * @param props initialization props
   */
  public static manifest(props: TlsStoreProps): any {
    return {
      ...TlsStore.GVK,
      ...toJson_TlsStoreProps(props),
    };
  }

  /**
   * Defines a "TLSStore" API object
   * @param scope the scope in which to define this object
   * @param id a scope-local name for the object
   * @param props initialization props
   */
  public constructor(scope: Construct, id: string, props: TlsStoreProps) {
    super(scope, id, {
      ...TlsStore.GVK,
      ...props,
    });
  }

  /**
   * Renders the object to Kubernetes JSON.
   */
  public toJson(): any {
    const resolved = super.toJson();

    return {
      ...TlsStore.GVK,
      ...toJson_TlsStoreProps(resolved),
    };
  }
}

/**
 * TLSStore is the CRD implementation of a Traefik TLS Store. For the time being, only the TLSStore named default is supported. This means that you cannot have two stores that are named default in different Kubernetes namespaces. More info: https://doc.traefik.io/traefik/v2.10/https/tls/#certificates-stores
 *
 * @schema TLSStore
 */
export interface TlsStoreProps {
  /**
   * @schema TLSStore#metadata
   */
  readonly metadata: ApiObjectMetadata;

  /**
   * TLSStoreSpec defines the desired state of a TLSStore.
   *
   * @schema TLSStore#spec
   */
  readonly spec: TlsStoreSpec;

}

/**
 * Converts an object of type 'TlsStoreProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_TlsStoreProps(obj: TlsStoreProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'metadata': obj.metadata,
    'spec': toJson_TlsStoreSpec(obj.spec),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * TLSStoreSpec defines the desired state of a TLSStore.
 *
 * @schema TlsStoreSpec
 */
export interface TlsStoreSpec {
  /**
   * Certificates is a list of secret names, each secret holding a key/certificate pair to add to the store.
   *
   * @schema TlsStoreSpec#certificates
   */
  readonly certificates?: TlsStoreSpecCertificates[];

  /**
   * DefaultCertificate defines the default certificate configuration.
   *
   * @schema TlsStoreSpec#defaultCertificate
   */
  readonly defaultCertificate?: TlsStoreSpecDefaultCertificate;

  /**
   * DefaultGeneratedCert defines the default generated certificate configuration.
   *
   * @schema TlsStoreSpec#defaultGeneratedCert
   */
  readonly defaultGeneratedCert?: TlsStoreSpecDefaultGeneratedCert;

}

/**
 * Converts an object of type 'TlsStoreSpec' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_TlsStoreSpec(obj: TlsStoreSpec | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'certificates': obj.certificates?.map(y => toJson_TlsStoreSpecCertificates(y)),
    'defaultCertificate': toJson_TlsStoreSpecDefaultCertificate(obj.defaultCertificate),
    'defaultGeneratedCert': toJson_TlsStoreSpecDefaultGeneratedCert(obj.defaultGeneratedCert),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * Certificate holds a secret name for the TLSStore resource.
 *
 * @schema TlsStoreSpecCertificates
 */
export interface TlsStoreSpecCertificates {
  /**
   * SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.
   *
   * @schema TlsStoreSpecCertificates#secretName
   */
  readonly secretName: string;

}

/**
 * Converts an object of type 'TlsStoreSpecCertificates' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_TlsStoreSpecCertificates(obj: TlsStoreSpecCertificates | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'secretName': obj.secretName,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * DefaultCertificate defines the default certificate configuration.
 *
 * @schema TlsStoreSpecDefaultCertificate
 */
export interface TlsStoreSpecDefaultCertificate {
  /**
   * SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.
   *
   * @schema TlsStoreSpecDefaultCertificate#secretName
   */
  readonly secretName: string;

}

/**
 * Converts an object of type 'TlsStoreSpecDefaultCertificate' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_TlsStoreSpecDefaultCertificate(obj: TlsStoreSpecDefaultCertificate | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'secretName': obj.secretName,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * DefaultGeneratedCert defines the default generated certificate configuration.
 *
 * @schema TlsStoreSpecDefaultGeneratedCert
 */
export interface TlsStoreSpecDefaultGeneratedCert {
  /**
   * Domain is the domain definition for the DefaultCertificate.
   *
   * @schema TlsStoreSpecDefaultGeneratedCert#domain
   */
  readonly domain?: TlsStoreSpecDefaultGeneratedCertDomain;

  /**
   * Resolver is the name of the resolver that will be used to issue the DefaultCertificate.
   *
   * @schema TlsStoreSpecDefaultGeneratedCert#resolver
   */
  readonly resolver?: string;

}

/**
 * Converts an object of type 'TlsStoreSpecDefaultGeneratedCert' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_TlsStoreSpecDefaultGeneratedCert(obj: TlsStoreSpecDefaultGeneratedCert | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'domain': toJson_TlsStoreSpecDefaultGeneratedCertDomain(obj.domain),
    'resolver': obj.resolver,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * Domain is the domain definition for the DefaultCertificate.
 *
 * @schema TlsStoreSpecDefaultGeneratedCertDomain
 */
export interface TlsStoreSpecDefaultGeneratedCertDomain {
  /**
   * Main defines the main domain name.
   *
   * @schema TlsStoreSpecDefaultGeneratedCertDomain#main
   */
  readonly main?: string;

  /**
   * SANs defines the subject alternative domain names.
   *
   * @schema TlsStoreSpecDefaultGeneratedCertDomain#sans
   */
  readonly sans?: string[];

}

/**
 * Converts an object of type 'TlsStoreSpecDefaultGeneratedCertDomain' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_TlsStoreSpecDefaultGeneratedCertDomain(obj: TlsStoreSpecDefaultGeneratedCertDomain | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'main': obj.main,
    'sans': obj.sans?.map(y => y),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

