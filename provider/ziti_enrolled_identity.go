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
	"encoding/json"
	"github.com/openziti/sdk-golang/ziti/enroll"
	p "github.com/pulumi/pulumi-go-provider"
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
func (thiz *EnrolledIdentity) Create(ctx p.Context, name string, input EnrolledIdentityArgs, preview bool) (string, EnrolledIdentityState, error) {
	// bail out now when we are in preview mode
	if preview {
		return IdPreviewPrefix + name, EnrolledIdentityState{
			EnrolledIdentityArgs: input,
			IdentityJson:         "",
		}, nil
	}

	state, err := enrollIdentity(input)
	return name, state, err
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

func (*EnrolledIdentity) Diff(ctx p.Context, id string, olds EnrolledIdentityState, news EnrolledIdentityArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.EnrollmentJwt != olds.EnrollmentJwt {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*EnrolledIdentity) Read(ctx p.Context, id string, inputs EnrolledIdentityArgs, state EnrolledIdentityState) (string, EnrolledIdentityArgs, EnrolledIdentityState, error) {
	// noop...
	return id, inputs, state, nil
}

func (*EnrolledIdentity) Update(ctx p.Context, id string, olds EnrolledIdentityState, news EnrolledIdentityArgs, preview bool) (EnrolledIdentityState, error) {
	if preview || olds.EnrollmentJwt == news.EnrollmentJwt || news.EnrollmentJwt == "" {
		// noop
		return EnrolledIdentityState{
			EnrolledIdentityArgs: news,
			IdentityJson:         olds.IdentityJson,
		}, nil
	}

	// news.EnrolmentJwt value has changed - ren-enroll
	state, err := enrollIdentity(news)
	return state, err
}

func (*EnrolledIdentity) Delete(ctx p.Context, id string, state EnrolledIdentityState) error {
	// noop..
	return nil
}
