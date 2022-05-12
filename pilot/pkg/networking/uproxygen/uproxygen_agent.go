//go:build agent
// +build agent

// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// TODO this file just exists to avoid pulling all of client-go into the agent binary

package uproxygen

import (
	"istio.io/istio/pilot/pkg/model"
)

var _ model.XdsResourceGenerator = &UProxyConfigGenerator{}

type UProxyConfigGenerator struct {
	EndpointIndex    *model.EndpointIndex
	ServiceDiscovery model.ServiceDiscovery
}

func (g *UProxyConfigGenerator) Generate(
	proxy *model.Proxy,
	w *model.WatchedResource,
	req *model.PushRequest,
) (model.Resources, model.XdsLogDetails, error) {
	return nil, model.DefaultXdsLogDetails, nil
}