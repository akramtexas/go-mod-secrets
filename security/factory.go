/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

// Package security defines the core functionality of the application and how to interact with the secret store.
package security

import (
	"fmt"

	"github.com/edgexfoundry-holding/go-mod-core-security/pkg/interfaces"
	"github.com/edgexfoundry-holding/go-mod-core-security/pkg/types"

	"github.com/edgexfoundry-holding/go-mod-core-security/internal/pkg/http"
	"github.com/edgexfoundry-holding/go-mod-core-security/internal/pkg/vault"
)

func NewSecurityClient(config types.Config) (interfaces.Client, error) {
	switch config.Provider {
	case types.VaultProvider:
		return vault.Client{}, nil
	case types.HTTPProvider:
		return http.Client{
			HttpConfig: http.Configuration{
				Host:           config.Target.Host,
				Port:           config.Target.Port,
				Path:           config.Target.Path,
				Authentication: config.Authentication,
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown provider type '%s' requested", config.Provider)
	}

}
