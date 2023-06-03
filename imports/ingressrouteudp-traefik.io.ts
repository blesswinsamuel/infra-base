// generated by cdk8s
import { ApiObject, ApiObjectMetadata, GroupVersionKind } from 'cdk8s';
import { Construct } from 'constructs';


/**
 * IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
 *
 * @schema IngressRouteUDP
 */
export class IngressRouteUdp extends ApiObject {
  /**
   * Returns the apiVersion and kind for "IngressRouteUDP"
   */
  public static readonly GVK: GroupVersionKind = {
    apiVersion: 'traefik.io/v1alpha1',
    kind: 'IngressRouteUDP',
  }

  /**
   * Renders a Kubernetes manifest for "IngressRouteUDP".
   *
   * This can be used to inline resource manifests inside other objects (e.g. as templates).
   *
   * @param props initialization props
   */
  public static manifest(props: IngressRouteUdpProps): any {
    return {
      ...IngressRouteUdp.GVK,
      ...toJson_IngressRouteUdpProps(props),
    };
  }

  /**
   * Defines a "IngressRouteUDP" API object
   * @param scope the scope in which to define this object
   * @param id a scope-local name for the object
   * @param props initialization props
   */
  public constructor(scope: Construct, id: string, props: IngressRouteUdpProps) {
    super(scope, id, {
      ...IngressRouteUdp.GVK,
      ...props,
    });
  }

  /**
   * Renders the object to Kubernetes JSON.
   */
  public toJson(): any {
    const resolved = super.toJson();

    return {
      ...IngressRouteUdp.GVK,
      ...toJson_IngressRouteUdpProps(resolved),
    };
  }
}

/**
 * IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
 *
 * @schema IngressRouteUDP
 */
export interface IngressRouteUdpProps {
  /**
   * @schema IngressRouteUDP#metadata
   */
  readonly metadata: ApiObjectMetadata;

  /**
   * IngressRouteUDPSpec defines the desired state of a IngressRouteUDP.
   *
   * @schema IngressRouteUDP#spec
   */
  readonly spec: IngressRouteUdpSpec;

}

/**
 * Converts an object of type 'IngressRouteUdpProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_IngressRouteUdpProps(obj: IngressRouteUdpProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'metadata': obj.metadata,
    'spec': toJson_IngressRouteUdpSpec(obj.spec),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * IngressRouteUDPSpec defines the desired state of a IngressRouteUDP.
 *
 * @schema IngressRouteUdpSpec
 */
export interface IngressRouteUdpSpec {
  /**
   * EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.10/routing/entrypoints/ Default: all.
   *
   * @schema IngressRouteUdpSpec#entryPoints
   */
  readonly entryPoints?: string[];

  /**
   * Routes defines the list of routes.
   *
   * @schema IngressRouteUdpSpec#routes
   */
  readonly routes: IngressRouteUdpSpecRoutes[];

}

/**
 * Converts an object of type 'IngressRouteUdpSpec' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_IngressRouteUdpSpec(obj: IngressRouteUdpSpec | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'entryPoints': obj.entryPoints?.map(y => y),
    'routes': obj.routes?.map(y => toJson_IngressRouteUdpSpecRoutes(y)),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * RouteUDP holds the UDP route configuration.
 *
 * @schema IngressRouteUdpSpecRoutes
 */
export interface IngressRouteUdpSpecRoutes {
  /**
   * Services defines the list of UDP services.
   *
   * @schema IngressRouteUdpSpecRoutes#services
   */
  readonly services?: IngressRouteUdpSpecRoutesServices[];

}

/**
 * Converts an object of type 'IngressRouteUdpSpecRoutes' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_IngressRouteUdpSpecRoutes(obj: IngressRouteUdpSpecRoutes | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'services': obj.services?.map(y => toJson_IngressRouteUdpSpecRoutesServices(y)),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * ServiceUDP defines an upstream UDP service to proxy traffic to.
 *
 * @schema IngressRouteUdpSpecRoutesServices
 */
export interface IngressRouteUdpSpecRoutesServices {
  /**
   * Name defines the name of the referenced Kubernetes Service.
   *
   * @schema IngressRouteUdpSpecRoutesServices#name
   */
  readonly name: string;

  /**
   * Namespace defines the namespace of the referenced Kubernetes Service.
   *
   * @schema IngressRouteUdpSpecRoutesServices#namespace
   */
  readonly namespace?: string;

  /**
   * NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP. The Kubernetes Service itself does load-balance to the pods. By default, NativeLB is false.
   *
   * @schema IngressRouteUdpSpecRoutesServices#nativeLB
   */
  readonly nativeLb?: boolean;

  /**
   * Port defines the port of a Kubernetes Service. This can be a reference to a named port.
   *
   * @schema IngressRouteUdpSpecRoutesServices#port
   */
  readonly port: IngressRouteUdpSpecRoutesServicesPort;

  /**
   * Weight defines the weight used when balancing requests between multiple Kubernetes Service.
   *
   * @schema IngressRouteUdpSpecRoutesServices#weight
   */
  readonly weight?: number;

}

/**
 * Converts an object of type 'IngressRouteUdpSpecRoutesServices' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_IngressRouteUdpSpecRoutesServices(obj: IngressRouteUdpSpecRoutesServices | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'name': obj.name,
    'namespace': obj.namespace,
    'nativeLB': obj.nativeLb,
    'port': obj.port?.value,
    'weight': obj.weight,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * Port defines the port of a Kubernetes Service. This can be a reference to a named port.
 *
 * @schema IngressRouteUdpSpecRoutesServicesPort
 */
export class IngressRouteUdpSpecRoutesServicesPort {
  public static fromNumber(value: number): IngressRouteUdpSpecRoutesServicesPort {
    return new IngressRouteUdpSpecRoutesServicesPort(value);
  }
  public static fromString(value: string): IngressRouteUdpSpecRoutesServicesPort {
    return new IngressRouteUdpSpecRoutesServicesPort(value);
  }
  private constructor(public readonly value: number | string) {
  }
}

