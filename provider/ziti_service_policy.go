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
	"github.com/openziti/edge-api/rest_management_api_client/identity"
	"github.com/openziti/edge-api/rest_management_api_client/service_policy"
	"github.com/openziti/edge-api/rest_model"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"reflect"
	"strings"
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

func resolveIdentity(ce *CacheEntry, name string) (string, *string, error) {
	if len(ce.identities) == 0 {
		ce.identitiesMutex.Lock()

		if len(ce.identities) == 0 {
			// initial load
			limit := int64(5000)
			offset := int64(0)
			listReq := &identity.ListIdentitiesParams{
				// Filter:  &filter,
				Limit:   &limit,
				Offset:  &offset,
				Context: context.Background(),
			}
			ctResp, err := ce.client.Identity.ListIdentities(listReq, nil)
			if err != nil {
				ce.identitiesMutex.Unlock()
				return "", nil, err
			}

			identities := make(map[string]string)
			identitiesReverse := make(map[string]string)
			for _, entity := range ctResp.GetPayload().Data {
				identities[*entity.Name] = *entity.ID
				identitiesReverse[*entity.ID] = *entity.Name
			}
			ce.identities = identities
			ce.identitiesReverse = identitiesReverse
		}
		ce.identitiesMutex.Unlock()
	}
	// try to find ID for name
	if identityId, ok := ce.identities[name]; ok {
		return identityId, nil, nil
	}
	// try to check if name is an ID
	if _, ok := ce.identitiesReverse[name]; ok {
		return name, nil, nil
	}

	ce.identitiesMutex.Lock()
	// re-check from lock....
	// try to find ID for name
	if identityId, ok := ce.identities[name]; ok {
		ce.identitiesMutex.Unlock()
		return identityId, nil, nil
	}
	// try to check if name is an ID
	if _, ok := ce.identitiesReverse[name]; ok {
		ce.identitiesMutex.Unlock()
		return name, nil, nil
	}

	// okay, we have to do api lookups... first: try by ID...
	fetchParams := &identity.DetailIdentityParams{
		ID:      name,
		Context: context.Background(),
	}
	fetchResp, err := ce.client.Identity.DetailIdentity(fetchParams, nil)
	if err == nil {
		// hit
		ce.identities[*fetchResp.Payload.Data.Name] = *fetchResp.Payload.Data.ID
		ce.identitiesReverse[*fetchResp.Payload.Data.ID] = *fetchResp.Payload.Data.Name
		ce.identitiesMutex.Unlock()
		return name, nil, nil
	}

	// next try: find by name...
	limit := int64(1)
	offset := int64(0)
	listReq := &identity.ListIdentitiesParams{
		Filter:  buildNameFilter(name),
		Limit:   &limit,
		Offset:  &offset,
		Context: context.Background(),
	}
	ctResp, err := ce.client.Identity.ListIdentities(listReq, nil)
	if err == nil && len(ctResp.Payload.Data) == 1 {
		// found!
		item := ctResp.Payload.Data[0]
		ce.identities[*item.Name] = *item.ID
		ce.identitiesReverse[*item.ID] = *item.Name
		ce.identitiesMutex.Unlock()
		return *item.ID, nil, err
	}

	ce.identitiesMutex.Unlock()
	msg := fmt.Sprintf("No Identity found matching `%s`", name)
	return name, &msg, nil
}

func (*ServicePolicy) Check(ctx p.Context, name string, oldInputs, newInputs resource.PropertyMap) (ServicePolicyArgs, []p.CheckFailure, error) {
	//if _, ok := newInputs["postureCheckRoles"]; !ok {
	//	newInputs["postureCheckRoles"] = resource.NewArrayProperty([]resource.PropertyValue{})
	//}
	var failures []p.CheckFailure
	if identityRoles, ok := newInputs["identityRoles"]; ok {
		ce, _, err := initClient(ctx)
		if err != nil {
			return ServicePolicyArgs{}, nil, err
		}
		identityRolesA := identityRoles.ArrayValue()
		identityRolesN := make([]resource.PropertyValue, len(identityRolesA))
		for idx, identityRole := range identityRolesA {
			if irs := identityRole.StringValue(); strings.HasPrefix(irs, "@") && !strings.HasPrefix(irs, "@"+IdPreviewPrefix) {
				resolved, msg, err := resolveIdentity(ce, irs[1:])
				if err != nil {
					return ServicePolicyArgs{}, nil, err
				}
				if msg != nil {
					failures = append(failures, p.CheckFailure{
						Property: "identityRoles",
						Reason:   *msg,
					})
				}
				identityRolesN[idx] = resource.NewStringProperty("@" + resolved)
			} else {
				identityRolesN[idx] = identityRole
			}
		}
		newInputs["identityRoles"] = resource.NewArrayProperty(identityRolesN)
	}
	ret, failures2, err := infer.DefaultCheck[ServicePolicyArgs](newInputs)
	if failures == nil {
		failures = failures2
	} else if failures2 == nil {
		// just keep failures
	} else {
		// merge both
		for _, failure := range failures2 {
			failures = append(failures, failure)
		}
	}
	return ret, failures, err
}

// All resources must implement Create at a minumum.
func (thiz *ServicePolicy) Create(ctx p.Context, name string, input ServicePolicyArgs, preview bool) (string, ServicePolicyState, error) {
	// bail out now when we are in preview mode
	if preview {
		return name, ServicePolicyState{
			ServicePolicyArgs:        input,
			BaseStateEntity:          buildBaseStatePreviewEntity(name, input.BaseArgsEntity),
			IdentityRoles:            input.IdentityRoles,
			IdentityRolesDisplay:     NamedRoles{},
			PostureCheckRoles:        input.PostureCheckRoles,
			PostureCheckRolesDisplay: NamedRoles{},
			Semantic:                 input.Semantic,
			ServiceRoles:             input.ServiceRoles,
			ServiceRolesDisplay:      nil,
			Type:                     input.Type,
		}, nil
	}

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

	resp, err := ce.client.ServicePolicy.CreateServicePolicy(createParams, nil)
	if err != nil {
		var badReq *service_policy.CreateServicePolicyBadRequest
		if errors.As(err, &badReq) {
			err2, dupe := formatApiErrDupeCheck(ctx, badReq, badReq.Payload)
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
				state, err := readServicePolicy(ce, existingId, input, true)
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
	state, err := readServicePolicy(ce, createdId, input, false)
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

func readServicePolicy(ce *CacheEntry, id string, input ServicePolicyArgs, assimilated bool) (ServicePolicyState, error) {
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
		BaseStateEntity:          buildBaseState(rd.BaseEntity, assimilated),
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
	readState, err := readServicePolicy(ce, id, inputs, state.Assimilated)
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

	readState, err := readServicePolicy(ce, id, news, olds.Assimilated)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*ServicePolicy) Delete(ctx p.Context, id string, state ServicePolicyState) error {
	ce, c, err := initClient(ctx)
	if err != nil {
		return err
	}
	if state.Assimilated && !c.deleteAssimilated {
		ctx.Logf(diag.Info, "DELETE on %s[%s]: Keeping on OpenZiti as this object was assimilated!", "ServicePolicy", id)
		return nil
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
