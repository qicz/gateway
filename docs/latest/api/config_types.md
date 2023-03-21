# API Reference

## Packages
- [config.gateway.envoyproxy.io/v1alpha1](#configgatewayenvoyproxyiov1alpha1)


## config.gateway.envoyproxy.io/v1alpha1

Package v1alpha1 contains API schema definitions for the config.gateway.envoyproxy.io
API group.


### Resource Types
- [EnvoyGateway](#envoygateway)
- [EnvoyProxy](#envoyproxy)



## EnvoyGateway



EnvoyGateway is the schema for the envoygateways API.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `config.gateway.envoyproxy.io/v1alpha1`
| `kind` _string_ | `EnvoyGateway`
| `EnvoyGatewaySpec` _[EnvoyGatewaySpec](#envoygatewayspec)_ | EnvoyGatewaySpec defines the desired state of EnvoyGateway. |


## EnvoyGatewaySpec



EnvoyGatewaySpec defines the desired state of Envoy Gateway.

_Appears in:_
- [EnvoyGateway](#envoygateway)

| Field | Description |
| --- | --- |
| `gateway` _[Gateway](#gateway)_ | Gateway defines desired Gateway API specific configuration. If unset, default configuration parameters will apply. |
| `provider` _[Provider](#provider)_ | Provider defines the desired provider and provider-specific configuration. If unspecified, the Kubernetes provider is used with default configuration parameters. |
| `rateLimit` _[RateLimit](#ratelimit)_ | RateLimit defines the configuration associated with the Rate Limit service deployed by Envoy Gateway required to implement the Global Rate limiting functionality. The specific rate limit service used here is the reference implementation in Envoy. For more details visit https://github.com/envoyproxy/ratelimit. This configuration is unneeded for "Local" rate limiting. |


## EnvoyProxy



EnvoyProxy is the schema for the envoyproxies API.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `config.gateway.envoyproxy.io/v1alpha1`
| `kind` _string_ | `EnvoyProxy`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.26/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[EnvoyProxySpec](#envoyproxyspec)_ | EnvoyProxySpec defines the desired state of EnvoyProxy. |


## EnvoyProxySpec



EnvoyProxySpec defines the desired state of EnvoyProxy.

_Appears in:_
- [EnvoyProxy](#envoyproxy)

| Field                                              | Description                                                                                                                                                                                   |
|----------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `provider` _[ResourceProvider](#resourceprovider)_ | Provider defines the desired resource provider and provider-specific configuration. If unspecified, the "Kubernetes" resource provider is used with default configuration parameters.         |
| `logging` _[ProxyLogging](#proxylogging)_          | Logging defines logging parameters for managed proxies. If unspecified, default settings apply. This type is not implemented until https://github.com/envoyproxy/gateway/issues/280 is fixed. |
| `bootstrap` _[ProxyBootstrap](#proxybootstrap)_    | Bootstrap defines the Envoy Bootstrap as a YAML string. Visit https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/bootstrap/v3/bootstrap.proto#envoy-v3-api-msg-config-bootstrap-v3-bootstrap to learn more about the syntax. |

> If set, this is the Bootstrap configuration used for the managed Envoy Proxy fleet instead of the default Bootstrap configuration
set by Envoy Gateway. Some fields within the Bootstrap that are required to communicate with the xDS Server (Envoy Gateway) and receive xDS resources
from it are not configurable and will result in the `EnvoyProxy` resource being rejected. Backward compatibility across minor versions is not guaranteed.
We strongly recommend using `egctl x translate` to generate a `EnvoyProxy` resource with the `Bootstrap` field set to the default
Bootstrap configuration used. You can edit this configuration, and rerun `egctl x translate` to ensure there are no validation errors.




## FileProvider



FileProvider defines configuration for the File provider.

_Appears in:_
- [Provider](#provider)



## Gateway



Gateway defines the desired Gateway API configuration of Envoy Gateway.

_Appears in:_
- [EnvoyGatewaySpec](#envoygatewayspec)

| Field | Description |
| --- | --- |
| `controllerName` _string_ | ControllerName defines the name of the Gateway API controller. If unspecified, defaults to "gateway.envoyproxy.io/gatewayclass-controller". See the following for additional details: https://gateway-api.sigs.k8s.io/v1alpha2/references/spec/#gateway.networking.k8s.io/v1alpha2.GatewayClass |


## KubernetesDeploymentSpec



KubernetesDeploymentSpec defines the desired state of the Kubernetes deployment resource.

_Appears in:_
- [KubernetesResourceProvider](#kubernetesresourceprovider)

| Field                                  | Description                                                                                                                                                                             |
|----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `replicas` _integer_                   | Replicas is the number of desired pods. Defaults to 1.                                                                                                                                  |
| `podAnnotations` _map[string][string]_ | PodAnnotations are the annotations that should be appended to the pods.                                                                                                                 |
| `resources` _corev1.ResourceRequirements_ | Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ `corev1 "k8s.io/api/core/v1"`  |


## KubernetesServiceSpec



EnvoyService defines the desired state of the Envoy service resource.

_Appears in:_
- [KubernetesResourceProvider](#kubernetesresourceprovider)

| Field                                  | Description                                                                                                                                                                             |
|----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `annotations` _map[string][string]_ | Annotations that should be appended to the service.By default, no annotations are appended.                                                                                                |


## KubernetesProvider



KubernetesProvider defines configuration for the Kubernetes provider.

_Appears in:_
- [Provider](#provider)



## KubernetesResourceProvider



KubernetesResourceProvider defines configuration for the Kubernetes resource provider.

_Appears in:_
- [ResourceProvider](#resourceprovider)

| Field | Description |
| --- | --- |
| `envoyDeployment` _[KubernetesDeploymentSpec](#kubernetesdeploymentspec)_ | EnvoyDeployment defines the desired state of the Envoy deployment resource. If unspecified, default settings for the manged Envoy deployment resource are applied. |
| `envoyService` _[KubernetesServiceSpec](#KubernetesServiceSpec)_ | EnvoyService defines the desired state of the Envoy service resource. If unspecified, default settings for the manged Envoy service resource are applied. |


## LogComponent

_Underlying type:_ `string`

LogComponent defines a component that supports a configured logging level. This type is not implemented until https://github.com/envoyproxy/gateway/issues/280 is fixed.

_Appears in:_
- [ProxyLogging](#proxylogging)



## LogLevel

_Underlying type:_ `string`

LogLevel defines a log level for system logs. This type is not implemented until https://github.com/envoyproxy/gateway/issues/280 is fixed.

_Appears in:_
- [ProxyLogging](#proxylogging)



## Provider



Provider defines the desired configuration of a provider.

_Appears in:_
- [EnvoyGatewaySpec](#envoygatewayspec)

| Field | Description |
| --- | --- |
| `type` _[ProviderType](#providertype)_ | Type is the type of provider to use. Supported types are "Kubernetes". |
| `kubernetes` _[KubernetesProvider](#kubernetesprovider)_ | Kubernetes defines the configuration of the Kubernetes provider. Kubernetes provides runtime configuration via the Kubernetes API. |
| `file` _[FileProvider](#fileprovider)_ | File defines the configuration of the File provider. File provides runtime configuration defined by one or more files. This type is not implemented until https://github.com/envoyproxy/gateway/issues/1001 is fixed. |


## ProviderType

_Underlying type:_ `string`

ProviderType defines the types of providers supported by Envoy Gateway.

_Appears in:_
- [Provider](#provider)
- [ResourceProvider](#resourceprovider)



## ProxyLogging



ProxyLogging defines logging parameters for managed proxies. This type is not implemented until https://github.com/envoyproxy/gateway/issues/280 is fixed.

_Appears in:_
- [EnvoyProxySpec](#envoyproxyspec)

| Field | Description |
| --- | --- |
| `level` _object (keys:[LogComponent](#logcomponent), values:[LogLevel](#loglevel))_ | Level is a map of logging level per component, where the component is the key and the log level is the value. If unspecified, defaults to "System: Info". |


## ProxyBootstrap

Bootstrap defines the Envoy Bootstrap as a YAML string. Visit https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/bootstrap/v3/bootstrap.proto#envoy-v3-api-msg-config-bootstrap-v3-bootstrap to learn more about the syntax.
If set, this is the Bootstrap configuration used for the managed Envoy Proxy fleet instead of the default Bootstrap configuration set by Envoy Gateway.
Some fields within the Bootstrap that are required to communicate with the xDS Server (Envoy Gateway) and receive xDS resources
from it are not configurable and will result in the `EnvoyProxy` resource being rejected.
Backward compatibility across minor versions is not guaranteed.
We strongly recommend using `egctl x translate` to generate a `EnvoyProxy` resource with the `Bootstrap` field set to the default
Bootstrap configuration used. You can edit this configuration, and rerun `egctl x translate` to ensure there are no validation errors.


_Appears in:_
- [EnvoyProxySpec](#envoyproxyspec)



## RateLimit



RateLimit defines the configuration associated with the Rate Limit Service used for Global Rate Limiting.

_Appears in:_
- [EnvoyGatewaySpec](#envoygatewayspec)

| Field | Description |
| --- | --- |
| `backend` _[RateLimitDatabaseBackend](#ratelimitdatabasebackend)_ | Backend holds the configuration associated with the database backend used by the rate limit service to store state associated with global ratelimiting. |


## RateLimitDatabaseBackend



RateLimitDatabaseBackend defines the configuration associated with the database backend used by the rate limit service.

_Appears in:_
- [RateLimit](#ratelimit)

| Field | Description |
| --- | --- |
| `type` _[RateLimitDatabaseBackendType](#ratelimitdatabasebackendtype)_ | Type is the type of database backend to use. Supported types are: * Redis: Connects to a Redis database. |
| `redis` _[RateLimitRedisSettings](#ratelimitredissettings)_ | Redis defines the settings needed to connect to a Redis database. |


## RateLimitDatabaseBackendType

_Underlying type:_ `string`

RateLimitDatabaseBackendType specifies the types of database backend to be used by the rate limit service.

_Appears in:_
- [RateLimitDatabaseBackend](#ratelimitdatabasebackend)



## RateLimitRedisSettings



RateLimitRedisSettings defines the configuration for connecting to a Redis database.

_Appears in:_
- [RateLimitDatabaseBackend](#ratelimitdatabasebackend)

| Field | Description |
| --- | --- |
| `url` _string_ | URL of the Redis Database. |


## ResourceProvider



ResourceProvider defines the desired state of a resource provider.

_Appears in:_
- [EnvoyProxySpec](#envoyproxyspec)

| Field | Description |
| --- | --- |
| `type` _[ProviderType](#providertype)_ | Type is the type of resource provider to use. A resource provider provides infrastructure resources for running the data plane, e.g. Envoy proxy, and optional auxiliary control planes. Supported types are "Kubernetes". |
| `kubernetes` _[KubernetesResourceProvider](#kubernetesresourceprovider)_ | Kubernetes defines the desired state of the Kubernetes resource provider. Kubernetes provides infrastructure resources for running the data plane, e.g. Envoy proxy. If unspecified and type is "Kubernetes", default settings for managed Kubernetes resources are applied. |


