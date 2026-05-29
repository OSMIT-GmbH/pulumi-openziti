// Copyright 2023, OSMIT GmbH
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
	"bytes"
	"context"
	"encoding/json"
	"github.com/openziti/sdk-golang/ziti/enroll"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type EnrolledIdentity struct{}

// Each resource has in input struct, defining what arguments it accepts.
type EnrolledIdentityArgs struct {
	// enrollmentJwt
	EnrollmentJwt string `pulumi:"enrollmentJwt" provider:"secret,input"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type EnrolledIdentityState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	EnrolledIdentityArgs

	// enrollment
	// Required: true
	IdentityJson string `pulumi:"identityJson" provider:"secret,output"`
}

// All resources must implement Create at a minumum.
func (EnrolledIdentity) Create(ctx context.Context, req infer.CreateRequest[EnrolledIdentityArgs]) (infer.CreateResponse[EnrolledIdentityState], error) {
	input := req.Inputs
	name := req.Name
	preview := req.DryRun
	// bail out now when we are in preview mode
	if preview {
		return infer.CreateResponse[EnrolledIdentityState]{
			ID: IdPreviewPrefix + name,
			Output: EnrolledIdentityState{
				EnrolledIdentityArgs: input,
				IdentityJson:         "",
			},
		}, nil
	}

	state, err := enrollIdentity(input)
	return infer.CreateResponse[EnrolledIdentityState]{ID: name, Output: state}, err
}

func enrollIdentity(input EnrolledIdentityArgs) (EnrolledIdentityState, error) {
	retErr := func(err error) (EnrolledIdentityState, error) {
		return EnrolledIdentityState{EnrolledIdentityArgs: input}, err
	}
	// Enroll the identity
	tkn, _, err := enroll.ParseToken(input.EnrollmentJwt)
	if err != nil {
		return retErr(err)
	}

	flags := enroll.EnrollmentFlags{
		Token:  tkn,
		KeyAlg: "RSA",
	}
	conf, err := enroll.Enroll(flags)
	if err != nil {
		return retErr(err)
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(conf)

	return EnrolledIdentityState{EnrolledIdentityArgs: input, IdentityJson: string(buffer.Bytes())}, err
}

func (EnrolledIdentity) Diff(ctx context.Context, req infer.DiffRequest[EnrolledIdentityArgs, EnrolledIdentityState]) (infer.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	news := req.Inputs
	olds := req.State
	if news.EnrollmentJwt != olds.EnrollmentJwt {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (EnrolledIdentity) Read(ctx context.Context, req infer.ReadRequest[EnrolledIdentityArgs, EnrolledIdentityState]) (infer.ReadResponse[EnrolledIdentityArgs, EnrolledIdentityState], error) {
	// noop...
	return infer.ReadResponse[EnrolledIdentityArgs, EnrolledIdentityState]{ID: req.ID, Inputs: req.Inputs, State: req.State}, nil
}

func (EnrolledIdentity) Update(ctx context.Context, req infer.UpdateRequest[EnrolledIdentityArgs, EnrolledIdentityState]) (infer.UpdateResponse[EnrolledIdentityState], error) {
	if req.DryRun || req.State.EnrollmentJwt == req.Inputs.EnrollmentJwt || req.Inputs.EnrollmentJwt == "" || req.Inputs.EnrollmentJwt == "~~used~~" {
		// noop
		return infer.UpdateResponse[EnrolledIdentityState]{
			Output: EnrolledIdentityState{
				EnrolledIdentityArgs: req.Inputs,
				IdentityJson:         req.State.IdentityJson,
			},
		}, nil
	}

	// news.EnrolmentJwt value has changed - ren-enroll
	state, err := enrollIdentity(req.Inputs)
	return infer.UpdateResponse[EnrolledIdentityState]{Output: state}, err
}

func (EnrolledIdentity) Delete(ctx context.Context, req infer.DeleteRequest[EnrolledIdentityState]) (infer.DeleteResponse, error) {
	// noop..
	return infer.DeleteResponse{}, nil
}
