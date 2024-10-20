// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package resource

import (
	"fmt"
	"sync"

	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwapiv1a3 "sigs.k8s.io/gateway-api/apis/v1alpha3"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	mcsapiv1a1 "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
)

var lock sync.RWMutex

// Append the resource to Resources' slice.
// For some resources that are unsupported will be inserted directly.
// ref internal/gatewayapi/resource/resource_cache.go:8
func (r *Resources) Append(obj client.Object) bool {
	lock.Lock()
	defer lock.Unlock()

	objUID := string(obj.GetUID())
	if objUID == "" {
		objUID = fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName())
	}

	resourceCache := r.resourceCache(obj.GetObjectKind().GroupVersionKind().Kind)
	if resourceCache.Has(objUID) {
		return false
	}

	switch res := obj.(type) {
	case *egv1a1.EnvoyProxy:
		r.EnvoyProxiesForGateways = append(r.EnvoyProxiesForGateways, res)
	case *gwapiv1.Gateway:
		r.Gateways = append(r.Gateways, res)
	case *gwapiv1b1.ReferenceGrant:
		r.ReferenceGrants = append(r.ReferenceGrants, res)
	case *corev1.Namespace:
		r.Namespaces = append(r.Namespaces, res)
	case *mcsapiv1a1.ServiceImport:
		r.ServiceImports = append(r.ServiceImports, res)
	case *discoveryv1.EndpointSlice:
		r.EndpointSlices = append(r.EndpointSlices, res)
	case *corev1.Secret:
		r.Secrets = append(r.Secrets, res)
	case *corev1.ConfigMap:
		r.ConfigMaps = append(r.ConfigMaps, res)
	case *egv1a1.EnvoyPatchPolicy:
		r.EnvoyPatchPolicies = append(r.EnvoyPatchPolicies, res)
	case *egv1a1.ClientTrafficPolicy:
		r.ClientTrafficPolicies = append(r.ClientTrafficPolicies, res)
	case *egv1a1.BackendTrafficPolicy:
		r.BackendTrafficPolicies = append(r.BackendTrafficPolicies, res)
	case *egv1a1.SecurityPolicy:
		r.SecurityPolicies = append(r.SecurityPolicies, res)
	case *gwapiv1a3.BackendTLSPolicy:
		r.BackendTLSPolicies = append(r.BackendTLSPolicies, res)
	case *egv1a1.EnvoyExtensionPolicy:
		r.EnvoyExtensionPolicies = append(r.EnvoyExtensionPolicies, res)
	case *egv1a1.HTTPRouteFilter:
		r.HTTPRouteFilters = append(r.HTTPRouteFilters, res)
	case *corev1.Service:
		r.Services = append(r.Services, res)
	case *egv1a1.Backend:
		r.Backends = append(r.Backends, res)
	case *gwapiv1.HTTPRoute:
		r.HTTPRoutes = append(r.HTTPRoutes, res)
	case *gwapiv1.GRPCRoute:
		r.GRPCRoutes = append(r.GRPCRoutes, res)
	case *gwapiv1a2.TLSRoute:
		r.TLSRoutes = append(r.TLSRoutes, res)
	case *gwapiv1a2.TCPRoute:
		r.TCPRoutes = append(r.TCPRoutes, res)
	case *gwapiv1a2.UDPRoute:
		r.UDPRoutes = append(r.UDPRoutes, res)
	}

	resourceCache.Insert(objUID)
	return true
}

func (r *Resources) GetNamespace(name string) *corev1.Namespace {
	for _, ns := range r.Namespaces {
		if ns.Name == name {
			return ns
		}
	}

	return nil
}

func (r *Resources) GetEnvoyProxy(namespace, name string) *egv1a1.EnvoyProxy {
	for _, ep := range r.EnvoyProxiesForGateways {
		if ep.Namespace == namespace && ep.Name == name {
			return ep
		}
	}

	return nil
}

func (r *Resources) GetService(namespace, name string) *corev1.Service {
	for _, svc := range r.Services {
		if svc.Namespace == namespace && svc.Name == name {
			return svc
		}
	}

	return nil
}

func (r *Resources) GetServiceImport(namespace, name string) *mcsapiv1a1.ServiceImport {
	for _, svcImp := range r.ServiceImports {
		if svcImp.Namespace == namespace && svcImp.Name == name {
			return svcImp
		}
	}

	return nil
}

func (r *Resources) GetBackend(namespace, name string) *egv1a1.Backend {
	for _, be := range r.Backends {
		if be.Namespace == namespace && be.Name == name {
			return be
		}
	}

	return nil
}

func (r *Resources) GetSecret(namespace, name string) *corev1.Secret {
	for _, secret := range r.Secrets {
		if secret.Namespace == namespace && secret.Name == name {
			return secret
		}
	}

	return nil
}

func (r *Resources) GetConfigMap(namespace, name string) *corev1.ConfigMap {
	for _, configMap := range r.ConfigMaps {
		if configMap.Namespace == namespace && configMap.Name == name {
			return configMap
		}
	}

	return nil
}

func (r *Resources) GetEndpointSlicesForBackend(svcNamespace, svcName string, backendKind string) []*discoveryv1.EndpointSlice {
	var endpointSlices []*discoveryv1.EndpointSlice
	for _, endpointSlice := range r.EndpointSlices {
		var backendSelectorLabel string
		switch backendKind {
		case KindService:
			backendSelectorLabel = discoveryv1.LabelServiceName
		case KindServiceImport:
			backendSelectorLabel = mcsapiv1a1.LabelServiceName
		}
		if svcNamespace == endpointSlice.Namespace &&
			endpointSlice.GetLabels()[backendSelectorLabel] == svcName {
			endpointSlices = append(endpointSlices, endpointSlice)
		}
	}
	return endpointSlices
}
