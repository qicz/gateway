// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package resource

// supportedCacheResourceKinds that are supported cache manager.
// if some resources are supported remove the code slash comment.
var supportedCacheResourceKinds = []string{
	"EnvoyProxy",
	"Gateway",
	"ReferenceGrant",
	"Namespace",
	"ServiceImport",
	"EndpointSlice",
	"Secret",
	"ConfigMap",
	"EnvoyPatchPolicy",
	"ClientTrafficPolicy",
	"BackendTrafficPolicy",
	"SecurityPolicy",
	"BackendTLSPolicy",
	"EnvoyExtensionPolicy",
	"HTTPRouteFilter",
	// "Service", // Unsupported
	// "Backend", // Unsupported
	// "HTTPRoute", // Unsupported
	// "GRPCRoute", // Unsupported
	// "TLSRoute", // Unsupported
	// "TCPRoute", // Unsupported
	// "UDPRoute", // Unsupported
}

// Set holds the resources with the kind.
// +k8s:deepcopy-gen=true
type Set struct {
	Values map[string]string `json:"-" yaml:"-"`
}

func newSet() *Set {
	return &Set{
		Values: map[string]string{},
	}
}

func (s *Set) Has(item string) bool {
	_, contained := s.Values[item]
	return contained
}

func (s *Set) Insert(items ...string) {
	for _, item := range items {
		s.Values[item] = ""
	}
}

// Cache holds some duplicate resources in memory.
// +k8s:deepcopy-gen=true
type Cache struct {
	ResourceSet map[string]*Set `json:"-" yaml:"-"`
}

func (r *Resources) InitCache() {
	r.Cache = newCache()
}

func newCache() *Cache {
	resourceSets := make(map[string]*Set, len(supportedCacheResourceKinds))
	for _, kind := range supportedCacheResourceKinds {
		resourceSets[kind] = newSet()
	}

	return &Cache{
		ResourceSet: resourceSets,
	}
}

func (r *Resources) resourceCache(kind string) *Set {
	if rs, ok := r.Cache.ResourceSet[kind]; ok {
		return rs
	}

	return newSet()
}
