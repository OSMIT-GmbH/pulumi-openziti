// Copyright 2025, Pulumi Corporation.
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

// Package provider implements a simple random resource and component.
package provider

import (
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

// Name controls how this provider is referenced in package names and elsewhere.
const Name string = "openziti"

// Provider creates a new instance of the provider.
func Provider() p.Provider {
	p, err := infer.NewProviderBuilder().
		WithDisplayName("OpenZITI").
		WithDescription("A Pulumi provider for managing OpenZITI resources.").
		WithHomepage("https://github.com/OSMIT-GmbH/pulumi-openziti").
		WithNamespace("osmit-gmbh").
		WithGoImportPath("github.com/OSMIT-GmbH/pulumi-openziti/sdk/go/pulumi-openziti").
		WithRepository("https://github.com/OSMIT-GmbH/pulumi-openziti").
		WithResources(
			// we have to use ConfigObj as name otherwise we have a name
			// clash with ProviderConfig in dotnet module... :-/
			infer.Resource[*ConfigObj, ConfigArgs, ConfigState](&ConfigObj{}),
			infer.Resource[*EdgeRouter, EdgeRouterArgs, EdgeRouterState](&EdgeRouter{}),
			infer.Resource[*EdgeRouterPolicy, EdgeRouterPolicyArgs, EdgeRouterPolicyState](&EdgeRouterPolicy{}),
			infer.Resource[*EnrolledIdentity, EnrolledIdentityArgs, EnrolledIdentityState](&EnrolledIdentity{}),
			infer.Resource[*Identity, IdentityArgs, IdentityState](&Identity{}),
			infer.Resource[*Service, ServiceArgs, ServiceState](&Service{}),
			infer.Resource[*ServiceEdgeRouterPolicy, ServiceEdgeRouterPolicyArgs, ServiceEdgeRouterPolicyState](&ServiceEdgeRouterPolicy{}),
			infer.Resource[*ServicePolicy, ServicePolicyArgs, ServicePolicyState](&ServicePolicy{}),
		).
		// WithComponents(infer.ComponentF(NewRandomComponent)).
		WithConfig(infer.Config(&OpenZitiProviderConfig{})).
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		}).Build()
	if err != nil {
		panic(fmt.Errorf("unable to build provider: %w", err))
	}
	return p
}
