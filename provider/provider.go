// Copyright 2023, Pulumi Corporation.
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

package provider

import (
	_ "github.com/motemen/go-loghttp/global"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

// Some hints:
// OpenZiti Samples: https://github.com/openziti/ziti/blob/fd2f3bf7092105391494985727ecc612cfc3e13b/ziti/cmd/edge/quickstart_shared_test.go

const Name string = "openziti"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			// we have to use ConfigObj as name otherwise we have a name
			// clash with ProviderConfig in dotnet module... :-/
			infer.Resource[*ConfigObj, ConfigArgs, ConfigState](),
			infer.Resource[*EdgeRouter, EdgeRouterArgs, EdgeRouterState](),
			infer.Resource[*EdgeRouterPolicy, EdgeRouterPolicyArgs, EdgeRouterPolicyState](),
			infer.Resource[*EnrolledIdentity, EnrolledIdentityArgs, EnrolledIdentityState](),
			infer.Resource[*Identity, IdentityArgs, IdentityState](),
			infer.Resource[*Service, ServiceArgs, ServiceState](),
			infer.Resource[*ServiceEdgeRouterPolicy, ServiceEdgeRouterPolicyArgs, ServiceEdgeRouterPolicyState](),
			infer.Resource[*ServicePolicy, ServicePolicyArgs, ServicePolicyState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Config: infer.Config[*OpenZitiProviderConfig](),
	})
}
