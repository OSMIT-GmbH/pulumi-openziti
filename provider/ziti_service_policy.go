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
	"context"
	"errors"
	"fmt"
	"github.com/openziti/edge-api/rest_management_api_client/service_policy"
	"github.com/openziti/edge-api/rest_model"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"reflect"
	"time"

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
type ServicePolicy struct{}

// Each resource has in input struct, defining what arguments it accepts.
type ServicePolicyArgs struct {
	BaseArgsEntity

	// identity roles
	IdentityRoles rest_model.Roles `pulumi:"identityRoles"`

	// posture check roles
	PostureCheckRoles rest_model.Roles `pulumi:"postureCheckRoles,optional"`

	// semantic
	// Required: true
	Semantic rest_model.Semantic `pulumi:"semantic"`

	// service roles
	ServiceRoles rest_model.Roles `pulumi:"serviceRoles"`

	// type
	// Required: true
	Type rest_model.DialBind `pulumi:"type"`
}

type NamedRole struct {

	// name
	Name string `pulumi:"name,optional"`

	// role
	Role string `pulumi:"role,optional"`
}
type NamedRoles []NamedRole

func buildRoleDisplay(in rest_model.NamedRoles) NamedRoles {
	namedRoles := make([]NamedRole, len(in))
	if in != nil {
		for i, nr := range in {
			namedRoles[i] = NamedRole{
				Name: nr.Name,
				Role: nr.Role,
			}
		}
	}
	return namedRoles
}

// Each resource has a state, describing the fields that exist on the created resource.
type ServicePolicyState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	ServicePolicyArgs
	// Here we define a required output called result.
	// Data interface{} `pulumi:"data"`
	BaseStateEntity

	// identity roles
	// Required: true
	IdentityRoles rest_model.Roles `pulumi:"identityRoles"`

	// identity roles display
	// Required: true
	IdentityRolesDisplay NamedRoles `pulumi:"identityRolesDisplay"`

	// posture check roles
	// Required: true
	PostureCheckRoles rest_model.Roles `pulumi:"postureCheckRoles"`

	// posture check roles display
	// Required: true
	PostureCheckRolesDisplay NamedRoles `pulumi:"postureCheckRolesDisplay"`

	// semantic
	// Required: true
	Semantic rest_model.Semantic `pulumi:"semantic"`

	// service roles
	// Required: true
	ServiceRoles rest_model.Roles `pulumi:"serviceRoles"`

	// service roles display
	// Required: true
	ServiceRolesDisplay NamedRoles `pulumi:"serviceRolesDisplay"`

	// type
	// Required: true
	Type rest_model.DialBind `pulumi:"type"`
}

func (*ServicePolicy) Check(ctx p.Context, name string, oldInputs ServicePolicyArgs, newInputs resource.PropertyMap) (ServicePolicyArgs, []p.CheckFailure, error) {
	if _, ok := newInputs["postureCheckRoles"]; !ok {
		newInputs["postureCheckRoles"] = resource.NewArrayProperty([]resource.PropertyValue{})
	}
	return infer.DefaultCheck[ServicePolicyArgs](newInputs)
}

// All resources must implement Create at a minumum.
func (thiz *ServicePolicy) Create(ctx p.Context, name string, input ServicePolicyArgs, preview bool) (string, ServicePolicyState, error) {
	retErr := func(err error) (string, ServicePolicyState, error) {
		return "", ServicePolicyState{ServicePolicyArgs: input}, err
	}
	ce, c, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	data := &rest_model.ServicePolicyCreate{
		Name:              &input.Name,
		Tags:              buildZitiTags(input.Tags),
		IdentityRoles:     input.IdentityRoles,
		PostureCheckRoles: input.PostureCheckRoles,
		Semantic:          &input.Semantic,
		ServiceRoles:      input.ServiceRoles,
		Type:              &input.Type,
	}
	createParams := &service_policy.CreateServicePolicyParams{
		Policy:  data,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return name, ServicePolicyState{ServicePolicyArgs: input}, nil
	}

	resp, err := ce.client.ServicePolicy.CreateServicePolicy(createParams, nil)
	if err != nil {
		var badReq *service_policy.CreateServicePolicyBadRequest
		if errors.As(err, &badReq) {
			err2, dupe := formatApiErrDupeCheck(ctx, badReq, badReq.Payload)
			fmt.Printf("DupeCheck: %b %b %s", dupe, c.assimilate, c.Assimilate)

			if dupe && c.assimilate {
				// find identity by name...
				findParams := &service_policy.ListServicePoliciesParams{
					Filter:  buildNameFilter(input.Name),
					Context: context.Background(),
				}
				findRet, err3 := ce.client.ServicePolicy.ListServicePolicies(findParams, nil)
				if err3 != nil {
					ctx.Logf(diag.Error, "Assimilate failed: List failed with %s", err3.Error())
					return retErr(err2)
				}
				if len(findRet.Payload.Data) != 1 {
					ctx.Logf(diag.Error, "Assimilate failed: List returned unexpected result count: %v", findRet.Payload.Data)
					return retErr(err2)
				}
				existingId := *findRet.Payload.Data[0].ID
				ctx.Logf(diag.Info, "Assimilating existing ID: %s", existingId)
				state, err := readServicePolicy(ce, existingId, input)
				if err != nil {
					ctx.Logf(diag.Error, "Assimilate failed: Fetch failed with: %s", err2)
					return retErr(err2)
				}
				updatedState, err := thiz.Update(ctx, existingId, state, input, preview)
				return existingId, updatedState, err
			}
			return retErr(err2)
		}

		return retErr(err)
	}
	createdId := resp.GetPayload().Data.ID
	state, err := readServicePolicy(ce, createdId, input)
	if err != nil {
		return createdId, state, err
	}
	return createdId, state, nil
}

func (*ServicePolicy) Diff(ctx p.Context, id string, olds ServicePolicyState, news ServicePolicyArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))
	if news.Type != olds.Type {
		diff["type"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.Semantic != olds.Semantic {
		diff["semantic"] = p.PropertyDiff{Kind: p.Update}
	}
	diffStrArrayIgnoreOrder(ctx, diff, "serviceRoles", olds.ServiceRoles, news.ServiceRoles)
	diffStrArrayIgnoreOrder(ctx, diff, "identityRoles", olds.IdentityRoles, news.IdentityRoles)
	diffStrArrayIgnoreOrder(ctx, diff, "postureCheckRoles", olds.PostureCheckRoles, news.PostureCheckRoles)
	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on ServicePolicy %s/%s: Found %d diffs: %v", news.Name, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func readServicePolicy(ce CacheEntry, id string, input ServicePolicyArgs) (ServicePolicyState, error) {
	params := &service_policy.DetailServicePolicyParams{
		ID:      id,
		Context: context.Background(),
	}
	detailResp, err := ce.client.ServicePolicy.DetailServicePolicy(params, nil)
	if err != nil {
		return ServicePolicyState{ServicePolicyArgs: input}, err
	}
	rd := detailResp.GetPayload().Data

	return ServicePolicyState{
		ServicePolicyArgs:        input,
		BaseStateEntity:          buildBaseState(rd.BaseEntity),
		IdentityRoles:            rd.IdentityRoles,
		IdentityRolesDisplay:     buildRoleDisplay(rd.IdentityRolesDisplay),
		PostureCheckRoles:        ifte(rd.PostureCheckRoles != nil, rd.PostureCheckRoles, []string{}),
		PostureCheckRolesDisplay: buildRoleDisplay(rd.PostureCheckRolesDisplay),
		Semantic:                 *rd.Semantic,
		ServiceRoles:             rd.ServiceRoles,
		ServiceRolesDisplay:      buildRoleDisplay(rd.ServiceRolesDisplay),
		Type:                     *rd.Type,
	}, nil
}

func (*ServicePolicy) Read(ctx p.Context, id string, inputs ServicePolicyArgs, state ServicePolicyState) (string, ServicePolicyArgs, ServicePolicyState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	readState, err := readServicePolicy(ce, id, inputs)
	if err != nil {
		return id, inputs, readState, err
	}
	return id, inputs, readState, nil
}

func (*ServicePolicy) Update(ctx p.Context, id string, olds ServicePolicyState, news ServicePolicyArgs, preview bool) (ServicePolicyState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}

	updateData := &rest_model.ServicePolicyUpdate{
		Name:              &news.Name,
		Tags:              buildZitiTags(news.Tags),
		IdentityRoles:     news.IdentityRoles,
		PostureCheckRoles: news.PostureCheckRoles,
		Semantic:          &news.Semantic,
		ServiceRoles:      news.ServiceRoles,
		Type:              &news.Type,
	}
	updateParams := &service_policy.UpdateServicePolicyParams{
		Policy:  updateData,
		ID:      id,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return olds, nil
	}

	_, err = ce.client.ServicePolicy.UpdateServicePolicy(updateParams, nil)
	if err != nil {
		var badReq *service_policy.UpdateServicePolicyBadRequest
		if errors.As(err, &badReq) {
			return olds, formatApiErr(ctx, badReq, badReq.Payload)
		}
		return olds, err
	}

	readState, err := readServicePolicy(ce, id, news)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*ServicePolicy) Delete(ctx p.Context, id string, _ ServicePolicyState) error {
	ce, _, err := initClient(ctx)
	if err != nil {
		return err
	}
	deleteParams := &service_policy.DeleteServicePolicyParams{
		ID:      id,
		Context: context.Background(),
	}
	deleteParams.SetTimeout(30 * time.Second)
	// ctx.Logf(diag.Info, "Calling delete on %s; output: %#v\n", id, *deleteParams)
	_, err = ce.client.ServicePolicy.DeleteServicePolicy(deleteParams, nil)
	if err != nil {
		return handleDeleteErr(ctx, err, id, "Config")
	}
	return nil
}
