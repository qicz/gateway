// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	"github.com/envoyproxy/gateway/api/config/v1alpha1"
)

var (
	inPath = "./testdata/decoder/in/"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		in     string
		out    *v1alpha1.EnvoyGateway
		expect bool
	}{
		{
			in: inPath + "kube-provider.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Provider: v1alpha1.DefaultGatewayProvider(),
				},
			},
			expect: true,
		},
		{
			in: inPath + "gateway-controller-name.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Gateway: v1alpha1.DefaultGateway(),
				},
			},
			expect: true,
		},
		{
			in: inPath + "provider-with-gateway.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Gateway:  v1alpha1.DefaultGateway(),
					Provider: v1alpha1.DefaultGatewayProvider(),
				},
			},
			expect: true,
		},
		{
			in: inPath + "provider-mixing-gateway.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Provider: v1alpha1.DefaultGatewayProvider(),
				},
			},
			expect: true,
		},
		{
			in: inPath + "gateway-mixing-provider.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Gateway: v1alpha1.DefaultGateway(),
				},
			},
			expect: true,
		},
		{
			in: inPath + "provider-mixing-gateway.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Provider: v1alpha1.DefaultGatewayProvider(),
					Gateway:  v1alpha1.DefaultGateway(),
				},
			},
			expect: false,
		},
		{
			in: inPath + "gateway-mixing-provider.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Provider: v1alpha1.DefaultGatewayProvider(),
					Gateway:  v1alpha1.DefaultGateway(),
				},
			},
			expect: false,
		},
		{
			in: inPath + "gateway-ratelimit.yaml",
			out: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{
					Kind:       v1alpha1.KindEnvoyGateway,
					APIVersion: v1alpha1.GroupVersion.String(),
				},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Gateway: v1alpha1.DefaultGateway(),
					Provider: &v1alpha1.GatewayProvider{
						Kubernetes: &v1alpha1.GatewayKubernetesProvider{
							RateLimitDeployment: &v1alpha1.KubernetesDeploymentSpec{
								Replicas: v1alpha1.DefaultKubernetesDeploymentReplicas(),
								Container: &v1alpha1.KubernetesContainerSpec{
									Resources: v1alpha1.DefaultResourceRequirements(),
									SecurityContext: &corev1.SecurityContext{
										RunAsUser:                pointer.Int64(2000),
										AllowPrivilegeEscalation: pointer.Bool(false),
									},
								},
								Pod: &v1alpha1.KubernetesPodSpec{
									Annotations: map[string]string{
										"key1": "val1",
										"key2": "val2",
									},
									SecurityContext: &corev1.PodSecurityContext{
										RunAsUser:           pointer.Int64(1000),
										RunAsGroup:          pointer.Int64(3000),
										FSGroup:             pointer.Int64(2000),
										FSGroupChangePolicy: func(s corev1.PodFSGroupChangePolicy) *corev1.PodFSGroupChangePolicy { return &s }(corev1.FSGroupChangeOnRootMismatch),
									},
								},
							},
						},
					},
					RateLimit: &v1alpha1.RateLimit{
						Backend: v1alpha1.RateLimitDatabaseBackend{
							Type:  v1alpha1.RedisBackendType,
							Redis: &v1alpha1.RateLimitRedisSettings{URL: "localhost:6379"},
						},
					},
				},
			},
			expect: true,
		},
		{
			in:     inPath + "no-api-version.yaml",
			expect: false,
		},
		{
			in:     inPath + "no-kind.yaml",
			expect: false,
		},
		{
			in:     "/non/existent/config.yaml",
			expect: false,
		},
		{
			in:     inPath + "invalid-gateway-group.yaml",
			expect: false,
		},
		{
			in:     inPath + "invalid-gateway-kind.yaml",
			expect: false,
		},
		{
			in:     inPath + "invalid-gateway-version.yaml",
			expect: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			eg, err := Decode(tc.in)
			if tc.expect {
				require.NoError(t, err)
				require.Equal(t, tc.out, eg)
			} else {
				require.Equal(t, !reflect.DeepEqual(tc.out, eg) || err != nil, true)
			}
		})
	}
}
