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
	"github.com/openziti/edge-api/rest_management_api_client/edge_router_policy"
	"github.com/openziti/edge-api/rest_model"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
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
type EdgeRouterPolicy struct{}

// Each resource has in input struct, defining what arguments it accepts.
type EdgeRouterPolicyArgs struct {
	BaseArgsEntity

	// edge router check roles
	EdgeRouterRoles rest_model.Roles `pulumi:"edgeRouterRoles,optional"`

	// identity roles
	IdentityRoles rest_model.Roles `pulumi:"identityRoles"`

	// semantic
	// Required: true
	Semantic rest_model.Semantic `pulumi:"semantic"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type EdgeRouterPolicyState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	EdgeRouterPolicyArgs
	// Here we define a required output called result.
	// Data interface{} `pulumi:"data"`
	BaseStateEntity

	// edge router roles
	// Required: true
	EdgeRouterRoles rest_model.Roles `pulumi:"edgeRouterRoles"`

	// edge router roles display
	// Required: true
	EdgeRouterRolesDisplay NamedRoles `pulumi:"edgeRouterRolesDisplay"`

	// identity roles
	// Required: true
	IdentityRoles rest_model.Roles `pulumi:"identityRoles"`

	// identity roles display
	// Required: true
	IdentityRolesDisplay NamedRoles `pulumi:"identityRolesDisplay"`

	// semantic
	// Required: true
	Semantic rest_model.Semantic `pulumi:"semantic"`
}

// All resources must implement Create at a minumum.
func (thiz *EdgeRouterPolicy) Create(ctx p.Context, name string, input EdgeRouterPolicyArgs, preview bool) (string, EdgeRouterPolicyState, error) {
	retErr := func(err error) (string, EdgeRouterPolicyState, error) {
		return "", EdgeRouterPolicyState{EdgeRouterPolicyArgs: input}, err
	}
	ce, c, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	data := &rest_model.EdgeRouterPolicyCreate{
		Name:            &input.Name,
		Tags:            buildZitiTags(input.Tags),
		EdgeRouterRoles: input.EdgeRouterRoles,
		IdentityRoles:   input.IdentityRoles,
		Semantic:        &input.Semantic,
	}
	createParams := &edge_router_policy.CreateEdgeRouterPolicyParams{
		Policy:  data,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return name, EdgeRouterPolicyState{EdgeRouterPolicyArgs: input}, nil
	}

	resp, err := ce.client.EdgeRouterPolicy.CreateEdgeRouterPolicy(createParams, nil)
	if err != nil {
		var badReq *edge_router_policy.CreateEdgeRouterPolicyBadRequest
		if errors.As(err, &badReq) {
			err2, dupe := formatApiErrDupeCheck(ctx, badReq, badReq.Payload)
			if dupe && c.assimilate {
				// find identity by name...
				findParams := &edge_router_policy.ListEdgeRouterPoliciesParams{
					Filter:  buildNameFilter(input.Name),
					Context: context.Background(),
				}
				findRet, err3 := ce.client.EdgeRouterPolicy.ListEdgeRouterPolicies(findParams, nil)
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
				state, err := readEdgeRouterPolicy(ce, existingId, input, true)
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
	state, err := readEdgeRouterPolicy(ce, createdId, input, false)
	if err != nil {
		return createdId, state, err
	}
	return createdId, state, nil
}

func (*EdgeRouterPolicy) Diff(ctx p.Context, id string, olds EdgeRouterPolicyState, news EdgeRouterPolicyArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))
	if news.Semantic != olds.Semantic {
		diff["semantic"] = p.PropertyDiff{Kind: p.Update}
	}
	diffStrArrayIgnoreOrder(ctx, diff, "edgeRouterRoles", olds.EdgeRouterRoles, news.EdgeRouterRoles)
	diffStrArrayIgnoreOrder(ctx, diff, "identityRoles", olds.IdentityRoles, news.IdentityRoles)
	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on EdgeRouterPolicy %s/%s: Found %d diffs: %v", news.Name, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func readEdgeRouterPolicy(ce *CacheEntry, id string, input EdgeRouterPolicyArgs, assimilated bool) (EdgeRouterPolicyState, error) {
	params := &edge_router_policy.DetailEdgeRouterPolicyParams{
		ID:      id,
		Context: context.Background(),
	}
	detailResp, err := ce.client.EdgeRouterPolicy.DetailEdgeRouterPolicy(params, nil)
	if err != nil {
		return EdgeRouterPolicyState{EdgeRouterPolicyArgs: input}, err
	}
	rd := detailResp.GetPayload().Data

	return EdgeRouterPolicyState{
		EdgeRouterPolicyArgs:   input,
		BaseStateEntity:        buildBaseState(rd.BaseEntity, assimilated),
		IdentityRoles:          rd.IdentityRoles,
		IdentityRolesDisplay:   buildRoleDisplay(rd.IdentityRolesDisplay),
		EdgeRouterRoles:        ifte(rd.EdgeRouterRoles != nil, rd.EdgeRouterRoles, []string{}),
		EdgeRouterRolesDisplay: buildRoleDisplay(rd.EdgeRouterRolesDisplay),
		Semantic:               *rd.Semantic,
	}, nil
}

func (*EdgeRouterPolicy) Read(ctx p.Context, id string, inputs EdgeRouterPolicyArgs, state EdgeRouterPolicyState) (string, EdgeRouterPolicyArgs, EdgeRouterPolicyState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	readState, err := readEdgeRouterPolicy(ce, id, inputs, state.Assimilated)
	if err != nil {
		return id, inputs, readState, err
	}
	return id, inputs, readState, nil
}

func (*EdgeRouterPolicy) Update(ctx p.Context, id string, olds EdgeRouterPolicyState, news EdgeRouterPolicyArgs, preview bool) (EdgeRouterPolicyState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}

	updateData := &rest_model.EdgeRouterPolicyUpdate{
		Name:            &news.Name,
		Tags:            buildZitiTags(news.Tags),
		EdgeRouterRoles: news.EdgeRouterRoles,
		IdentityRoles:   news.IdentityRoles,
		Semantic:        &news.Semantic,
	}
	updateParams := &edge_router_policy.UpdateEdgeRouterPolicyParams{
		Policy:  updateData,
		ID:      id,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return olds, nil
	}

	_, err = ce.client.EdgeRouterPolicy.UpdateEdgeRouterPolicy(updateParams, nil)
	if err != nil {
		var badReq *edge_router_policy.UpdateEdgeRouterPolicyBadRequest
		if errors.As(err, &badReq) {
			return olds, formatApiErr(ctx, badReq, badReq.Payload)
		}
		return olds, err
	}

	readState, err := readEdgeRouterPolicy(ce, id, news, olds.Assimilated)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*EdgeRouterPolicy) Delete(ctx p.Context, id string, state EdgeRouterPolicyState) error {
	ce, c, err := initClient(ctx)
	if err != nil {
		return err
	}
	if state.Assimilated && !c.deleteAssimilated {
		ctx.Logf(diag.Info, "DELETE on %s[%s]: Keeping on OpenZiti as this object was assimilated!", "EdgeRouterPolicy", id)
		return nil
	}
	deleteParams := &edge_router_policy.DeleteEdgeRouterPolicyParams{
		ID:      id,
		Context: context.Background(),
	}
	deleteParams.SetTimeout(30 * time.Second)
	// ctx.Logf(diag.Info, "Calling delete on %s; output: %#v\n", id, *deleteParams)
	_, err = ce.client.EdgeRouterPolicy.DeleteEdgeRouterPolicy(deleteParams, nil)
	if err != nil {
		return handleDeleteErr(ctx, err, id, "Config")
	}
	return nil
}
