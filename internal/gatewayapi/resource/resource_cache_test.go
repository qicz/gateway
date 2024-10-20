// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package resource

import (
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwapiv1a3 "sigs.k8s.io/gateway-api/apis/v1alpha3"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
	mcsapiv1a1 "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
)

func TestResourcesCache(t *testing.T) {
	resources := []client.Object{
		&egv1a1.EnvoyProxy{
			TypeMeta: metav1.TypeMeta{
				Kind: "EnvoyProxy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "EnvoyProxy",
			},
		},
		&gwapiv1.Gateway{
			TypeMeta: metav1.TypeMeta{
				Kind: "Gateway",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "Gateway",
			},
		},
		&mcsapiv1a1.ServiceImport{
			TypeMeta: metav1.TypeMeta{
				Kind: "ServiceImport",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ServiceImport",
			},
		},
		&corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind: "Secret",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "Secret",
			},
		},
		&corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind: "ConfigMap",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ConfigMap",
			},
		},
		&gwapiv1b1.ReferenceGrant{
			TypeMeta: metav1.TypeMeta{
				Kind: "ReferenceGrant",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ReferenceGrant",
			},
		},
		&corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind: "Namespace",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "Namespace",
			},
		},
		&discoveryv1.EndpointSlice{
			TypeMeta: metav1.TypeMeta{
				Kind: "EndpointSlice",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "endpointSlice",
			},
		},
		&egv1a1.EnvoyPatchPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind: "EnvoyPatchPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "EnvoyPatchPolicy",
			},
		},
		&egv1a1.ClientTrafficPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind: "ClientTrafficPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "ClientTrafficPolicy",
			},
		},
		&egv1a1.BackendTrafficPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind: "BackendTrafficPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "BackendTrafficPolicy",
			},
		},
		&egv1a1.SecurityPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind: "SecurityPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "SecurityPolicy",
			},
		},
		&gwapiv1a3.BackendTLSPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind: "BackendTLSPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "BackendTLSPolicy",
			},
		},
		&egv1a1.EnvoyExtensionPolicy{
			TypeMeta: metav1.TypeMeta{
				Kind: "EnvoyExtensionPolicy",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "EnvoyExtensionPolicy",
			},
		},
		&egv1a1.Backend{
			TypeMeta: metav1.TypeMeta{
				Kind: "Backend",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "Backend",
			},
		},
		&egv1a1.HTTPRouteFilter{
			TypeMeta: metav1.TypeMeta{
				Kind: "HTTPRouteFilter",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "HTTPRouteFilter",
			},
		},
		&corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind: "Service",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "Service",
			},
		},
		&gwapiv1.HTTPRoute{
			TypeMeta: metav1.TypeMeta{
				Kind: "HTTPRoute",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "HTTPRoute",
			},
		},
		&gwapiv1.GRPCRoute{
			TypeMeta: metav1.TypeMeta{
				Kind: "GRPCRoute",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "GRPCRoute",
			},
		},
		&gwapiv1a2.TLSRoute{
			TypeMeta: metav1.TypeMeta{
				Kind: "TLSRoute",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "TLSRoute",
			},
		},
		&gwapiv1a2.TCPRoute{
			TypeMeta: metav1.TypeMeta{
				Kind: "TCPRoute",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "TCPRoute",
			},
		},
		&gwapiv1a2.UDPRoute{
			TypeMeta: metav1.TypeMeta{
				Kind: "UDPRoute",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "UDPRoute",
			},
		},
	}
	for _, res := range resources {
		r := NewResources()
		kind := res.GetObjectKind().GroupVersionKind().Kind
		t.Run(kind, func(t *testing.T) {
			name := res.GetName()
			resourceUID := fmt.Sprintf("%s/%s", res.GetNamespace(), name)
			if got := r.resourceCache(kind).Has(resourceUID); !got {
				t.Logf("resource with kind %v exist? %v", kind, got)
			}
			t.Logf("appending resource with kind %v", kind)
			switch kind {
			case "Namespace":
				if r.GetNamespace(name) == nil {
					t.Logf("%s is not exist", kind)
				}
			case "Service":
				if r.GetService("", name) == nil {
					t.Logf("%s is not exist", kind)
				}
			case "ServiceImport":
				if r.GetServiceImport("", name) == nil {
					t.Logf("%s is not exist", kind)
				}
			case "Backend":
				if r.GetBackend("", name) == nil {
					t.Logf("%s is not exist", kind)
				}
			case "Secret":
				if r.GetSecret("", name) == nil {
					t.Logf("%s is not exist", kind)
				}
			case "ConfigMap":
				if r.GetConfigMap("", name) == nil {
					t.Logf("%s is not exist", kind)
				}
			case "EnvoyProxy":
				if r.GetEnvoyProxy("", name) == nil {
					t.Logf("%s is not exist", kind)
				}
			}
			if r.Append(res) {
				t.Logf("appends resource with kind %v", kind)
			}
			switch kind {
			case "Namespace":
				if r.GetNamespace(name) != nil && r.GetNamespace(name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			case "Service":
				if r.GetService("", name) != nil && r.GetService("", name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			case "ServiceImport":
				if r.GetServiceImport("", name) != nil && r.GetServiceImport("", name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			case "Backend":
				if r.GetBackend("", name) != nil && r.GetBackend("", name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			case "Secret":
				if r.GetSecret("", name) != nil && r.GetSecret("", name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			case "ConfigMap":
				if r.GetConfigMap("", name) != nil && r.GetConfigMap("", name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			case "EnvoyProxy":
				if r.GetEnvoyProxy("", name) != nil && r.GetEnvoyProxy("", name).Name == name {
					t.Logf("%s has aleary appended", kind)
				}
			}
			if got := r.resourceCache(kind).Has(resourceUID); got {
				t.Logf("resource with kind %v exist? %v", kind, got)
			}
			if !r.Append(res) {
				t.Logf("already exist resource with kind %v", kind)
			}
		})
	}
}

func TestSet(t *testing.T) {
	type fields struct {
		Values map[string]string
	}
	type args struct {
		item string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       bool
		insertWant bool
	}{
		{
			name: "item exist",
			fields: fields{Values: map[string]string{
				"item": "",
			}},
			args:       args{item: "item"},
			want:       true,
			insertWant: true,
		},
		{
			name: "item not exist",
			fields: fields{Values: map[string]string{
				"item0": "",
			}},
			args:       args{item: "item"},
			want:       false,
			insertWant: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{Values: tt.fields.Values}
			if got := s.Has(tt.args.item); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})

		t.Run(fmt.Sprintf("%s-insert", tt.name), func(t *testing.T) {
			s := newSet()
			s.Insert(tt.args.item)
			if got := s.Has(tt.args.item); got != tt.insertWant {
				t.Errorf("Insert Has() = %v, insert want %v", got, tt.insertWant)
			}
		})
	}
}
