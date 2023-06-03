// generated by cdk8s
import { ApiObject, ApiObjectMetadata, GroupVersionKind } from 'cdk8s';
import { Construct } from 'constructs';


/**
 * MiddlewareTCP is the CRD implementation of a Traefik TCP middleware. More info: https://doc.traefik.io/traefik/v2.10/middlewares/overview/
 *
 * @schema MiddlewareTCP
 */
export class MiddlewareTcp extends ApiObject {
  /**
   * Returns the apiVersion and kind for "MiddlewareTCP"
   */
  public static readonly GVK: GroupVersionKind = {
    apiVersion: 'traefik.io/v1alpha1',
    kind: 'MiddlewareTCP',
  }

  /**
   * Renders a Kubernetes manifest for "MiddlewareTCP".
   *
   * This can be used to inline resource manifests inside other objects (e.g. as templates).
   *
   * @param props initialization props
   */
  public static manifest(props: MiddlewareTcpProps): any {
    return {
      ...MiddlewareTcp.GVK,
      ...toJson_MiddlewareTcpProps(props),
    };
  }

  /**
   * Defines a "MiddlewareTCP" API object
   * @param scope the scope in which to define this object
   * @param id a scope-local name for the object
   * @param props initialization props
   */
  public constructor(scope: Construct, id: string, props: MiddlewareTcpProps) {
    super(scope, id, {
      ...MiddlewareTcp.GVK,
      ...props,
    });
  }

  /**
   * Renders the object to Kubernetes JSON.
   */
  public toJson(): any {
    const resolved = super.toJson();

    return {
      ...MiddlewareTcp.GVK,
      ...toJson_MiddlewareTcpProps(resolved),
    };
  }
}

/**
 * MiddlewareTCP is the CRD implementation of a Traefik TCP middleware. More info: https://doc.traefik.io/traefik/v2.10/middlewares/overview/
 *
 * @schema MiddlewareTCP
 */
export interface MiddlewareTcpProps {
  /**
   * @schema MiddlewareTCP#metadata
   */
  readonly metadata: ApiObjectMetadata;

  /**
   * MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.
   *
   * @schema MiddlewareTCP#spec
   */
  readonly spec: MiddlewareTcpSpec;

}

/**
 * Converts an object of type 'MiddlewareTcpProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MiddlewareTcpProps(obj: MiddlewareTcpProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'metadata': obj.metadata,
    'spec': toJson_MiddlewareTcpSpec(obj.spec),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.
 *
 * @schema MiddlewareTcpSpec
 */
export interface MiddlewareTcpSpec {
  /**
   * InFlightConn defines the InFlightConn middleware configuration.
   *
   * @schema MiddlewareTcpSpec#inFlightConn
   */
  readonly inFlightConn?: MiddlewareTcpSpecInFlightConn;

  /**
   * IPWhiteList defines the IPWhiteList middleware configuration.
   *
   * @schema MiddlewareTcpSpec#ipWhiteList
   */
  readonly ipWhiteList?: MiddlewareTcpSpecIpWhiteList;

}

/**
 * Converts an object of type 'MiddlewareTcpSpec' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MiddlewareTcpSpec(obj: MiddlewareTcpSpec | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'inFlightConn': toJson_MiddlewareTcpSpecInFlightConn(obj.inFlightConn),
    'ipWhiteList': toJson_MiddlewareTcpSpecIpWhiteList(obj.ipWhiteList),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * InFlightConn defines the InFlightConn middleware configuration.
 *
 * @schema MiddlewareTcpSpecInFlightConn
 */
export interface MiddlewareTcpSpecInFlightConn {
  /**
   * Amount defines the maximum amount of allowed simultaneous connections. The middleware closes the connection if there are already amount connections opened.
   *
   * @schema MiddlewareTcpSpecInFlightConn#amount
   */
  readonly amount?: number;

}

/**
 * Converts an object of type 'MiddlewareTcpSpecInFlightConn' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MiddlewareTcpSpecInFlightConn(obj: MiddlewareTcpSpecInFlightConn | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'amount': obj.amount,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * IPWhiteList defines the IPWhiteList middleware configuration.
 *
 * @schema MiddlewareTcpSpecIpWhiteList
 */
export interface MiddlewareTcpSpecIpWhiteList {
  /**
   * SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
   *
   * @schema MiddlewareTcpSpecIpWhiteList#sourceRange
   */
  readonly sourceRange?: string[];

}

/**
 * Converts an object of type 'MiddlewareTcpSpecIpWhiteList' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MiddlewareTcpSpecIpWhiteList(obj: MiddlewareTcpSpecIpWhiteList | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'sourceRange': obj.sourceRange?.map(y => y),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */
