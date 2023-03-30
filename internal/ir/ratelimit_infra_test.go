// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package ir

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/envoyproxy/gateway/api/config/v1alpha1"
)

func TestRateLimitInfra_GetRateLimitConfig(t *testing.T) {
	type fields struct {
		ServiceConfigs []*RateLimitServiceConfig
		Config         *v1alpha1.EnvoyGateway
	}
	tests := []struct {
		name               string
		fields             fields
		wantConfig         *v1alpha1.EnvoyGateway
		wantFixedConfig    *v1alpha1.EnvoyGateway
		wantServiceConfigs []*RateLimitServiceConfig
	}{
		{
			name: "nil ratelimit infra config",
			fields: fields{
				ServiceConfigs: nil,
				Config:         nil,
			},
			wantConfig:         nil,
			wantFixedConfig:    &v1alpha1.EnvoyGateway{},
			wantServiceConfigs: nil,
		},
		{
			name: "with provider ratelimit infra config",
			fields: fields{
				ServiceConfigs: []*RateLimitServiceConfig{
					{
						Name:   "name",
						Config: "config",
					},
				},
				Config: &v1alpha1.EnvoyGateway{
					TypeMeta: metav1.TypeMeta{},
					EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
						Provider: v1alpha1.DefaultGatewayProvider(),
					},
				},
			},
			wantConfig: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Provider: v1alpha1.DefaultGatewayProvider(),
				},
			},
			wantFixedConfig: &v1alpha1.EnvoyGateway{
				TypeMeta: metav1.TypeMeta{},
				EnvoyGatewaySpec: v1alpha1.EnvoyGatewaySpec{
					Provider: v1alpha1.DefaultGatewayProvider(),
				},
			},
			wantServiceConfigs: []*RateLimitServiceConfig{
				{
					Name:   "name",
					Config: "config",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RateLimitInfra{
				ServiceConfigs: tt.fields.ServiceConfigs,
				Config:         tt.fields.Config,
			}
			assert.True(t, reflect.DeepEqual(tt.wantConfig, p.Config))
			assert.True(t, reflect.DeepEqual(tt.wantServiceConfigs, p.ServiceConfigs))
			assert.Equalf(t, tt.wantFixedConfig, p.GetRateLimitConfig(), "GetRateLimitConfig()")
		})
	}
}
