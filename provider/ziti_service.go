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
	"github.com/openziti/edge-api/rest_management_api_client/service"
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
type Service struct{}

// Each resource has in input struct, defining what arguments it accepts.
type ServiceArgs struct {
	BaseArgsEntity

	// configs
	Configs []string `pulumi:"configs"`

	// Describes whether connections must support end-to-end encryption on both sides of the connection.
	// Required: true
	EncryptionRequired bool `pulumi:"encryptionRequired"`

	// role attributes
	RoleAttributes []string `pulumi:"roleAttributes"`

	// terminator strategy
	TerminatorStrategy string `pulumi:"terminatorStrategy,optional"`
}

type PostureQueriesType struct {

	// is passing
	// Required: true
	IsPassing bool `pulumi:"isPassing"`

	// policy Id
	// Required: true
	PolicyID string `pulumi:"policyId"`

	// policy type
	PolicyType rest_model.DialBind `pulumi:"policyType,optional"`

	// posture queries
	// Required: true
	// TODO - we have to map this to a pulumi type..
	PostureQueries []*rest_model.PostureQuery `pulumi:"postureQueries"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type ServiceState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	ServiceArgs
	// Here we define a required output called result.
	// Data interface{} `pulumi:"data"`
	BaseStateEntity

	// map of config data for this service keyed by the config type name. Only configs of the types requested will be returned.
	// Required: true
	Config map[string]map[string]interface{} `pulumi:"config"`

	// configs
	// Required: true
	Configs []string `pulumi:"configs"`

	// Describes whether connections must support end-to-end encryption on both sides of the connection. Read-only property, set at create.
	// Required: true
	EncryptionRequired bool `pulumi:"encryptionRequired"`

	// name
	// Required: true
	Name string `pulumi:"name"`

	// permissions
	// Required: true
	Permissions rest_model.DialBindArray `pulumi:"permissions"`

	// posture queries
	// Required: true
	PostureQueries []PostureQueriesType `pulumi:"postureQueries"`

	// role attributes
	// Required: true
	RoleAttributes rest_model.Attributes `pulumi:"roleAttributes"`

	// terminator strategy
	// Required: true
	TerminatorStrategy string `pulumi:"terminatorStrategy"`
}

// All resources must implement Create at a minumum.
func (thiz *Service) Create(ctx p.Context, name string, input ServiceArgs, preview bool) (string, ServiceState, error) {
	// bail out early when we are in preview mode
	if preview {
		return IdPreviewPrefix + name, ServiceState{
			ServiceArgs:        input,
			BaseStateEntity:    buildBaseStatePreviewEntity(name, input.BaseArgsEntity),
			Config:             make(map[string]map[string]interface{}),
			Configs:            input.Configs,
			EncryptionRequired: input.EncryptionRequired,
			Name:               input.Name,
			Permissions:        nil,
			PostureQueries:     nil,
			RoleAttributes:     input.RoleAttributes,
			TerminatorStrategy: input.TerminatorStrategy,
		}, nil
	}

	retErr := func(err error) (string, ServiceState, error) {
		return "", ServiceState{ServiceArgs: input}, err
	}
	ce, c, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	data := &rest_model.ServiceCreate{
		Configs:            input.Configs,
		EncryptionRequired: &input.EncryptionRequired,
		Name:               &input.Name,
		RoleAttributes:     input.RoleAttributes,
		Tags:               buildZitiTags(input.Tags),
		TerminatorStrategy: input.TerminatorStrategy,
	}
	createParams := &service.CreateServiceParams{
		Service: data,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	resp, err := ce.client.Service.CreateService(createParams, nil)
	if err != nil {
		var badReq *service.CreateServiceBadRequest
		if errors.As(err, &badReq) {
			err2, dupe := formatApiErrDupeCheck(ctx, badReq, badReq.Payload)
			if dupe && c.assimilate {
				// find identity by name...
				findParams := &service.ListServicesParams{
					Filter:  buildNameFilter(input.Name),
					Context: context.Background(),
				}
				findRet, err3 := ce.client.Service.ListServices(findParams, nil)
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
				state, err := readService(ce, existingId, input, true)
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
	state, err := readService(ce, createdId, input, false)
	if err != nil {
		return createdId, state, err
	}
	return createdId, state, nil
}

func (*Service) Diff(ctx p.Context, id string, olds ServiceState, news ServiceArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))
	if news.EncryptionRequired != olds.EncryptionRequired {
		diff["encryptionRequired"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.TerminatorStrategy != olds.TerminatorStrategy && news.TerminatorStrategy != "" && olds.TerminatorStrategy != "terminatorStrategy" {
		diff["terminatorStrategy"] = p.PropertyDiff{Kind: p.Update}
	}
	diffStrArrayIgnoreOrder(ctx, diff, "configs", olds.Configs, news.Configs)
	diffStrArrayIgnoreOrder(ctx, diff, "roleAttributes", olds.RoleAttributes, news.RoleAttributes)
	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on Service %s/%s: Found %d diffs: %v", news.Name, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func readService(ce *CacheEntry, id string, input ServiceArgs, assimilated bool) (ServiceState, error) {
	params := &service.DetailServiceParams{
		ID:      id,
		Context: context.Background(),
	}
	detailResp, err := ce.client.Service.DetailService(params, nil)
	if err != nil {
		return ServiceState{ServiceArgs: input}, err
	}
	respPayload := detailResp.GetPayload()

	pql := len(respPayload.Data.PostureQueries)
	postureQueries := make([]PostureQueriesType, pql)
	if respPayload.Data.PostureQueries != nil {
		for i, query := range respPayload.Data.PostureQueries {
			postureQueries[i] = PostureQueriesType{
				IsPassing:      *query.IsPassing,
				PolicyID:       *query.PolicyID,
				PolicyType:     query.PolicyType,
				PostureQueries: query.PostureQueries,
			}
		}
	}
	// fmt.Printf("get  output: %+v\n", respPayload)
	return ServiceState{
		ServiceArgs:        input,
		BaseStateEntity:    buildBaseState(respPayload.Data.BaseEntity, assimilated),
		Config:             ifte(respPayload.Data.Config != nil, respPayload.Data.Config, make(map[string]map[string]interface{})),
		Configs:            respPayload.Data.Configs,
		EncryptionRequired: *respPayload.Data.EncryptionRequired,
		Name:               *respPayload.Data.Name,
		Permissions:        respPayload.Data.Permissions,
		PostureQueries:     postureQueries,
		RoleAttributes:     *respPayload.Data.RoleAttributes,
		TerminatorStrategy: *respPayload.Data.TerminatorStrategy,
	}, nil
}

func (*Service) Read(ctx p.Context, id string, inputs ServiceArgs, state ServiceState) (string, ServiceArgs, ServiceState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	readState, err := readService(ce, id, inputs, state.Assimilated)
	if err != nil {
		return id, inputs, readState, err
	}
	return id, inputs, readState, nil
}

func (*Service) Update(ctx p.Context, id string, olds ServiceState, news ServiceArgs, preview bool) (ServiceState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}

	updateData := &rest_model.ServiceUpdate{
		Configs:            news.Configs,
		EncryptionRequired: news.EncryptionRequired,
		Name:               &news.Name,
		RoleAttributes:     news.RoleAttributes,
		Tags:               buildZitiTags(news.Tags),
		TerminatorStrategy: news.TerminatorStrategy,
	}
	updateParams := &service.UpdateServiceParams{
		Service: updateData,
		ID:      id,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return olds, nil
	}

	_, err = ce.client.Service.UpdateService(updateParams, nil)
	if err != nil {
		var badReq *service.UpdateServiceBadRequest
		if errors.As(err, &badReq) {
			return olds, formatApiErr(ctx, badReq, badReq.Payload)
		}
		return olds, err
	}

	readState, err := readService(ce, id, news, olds.Assimilated)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*Service) Delete(ctx p.Context, id string, state ServiceState) error {
	ce, c, err := initClient(ctx)
	if err != nil {
		return err
	}
	if state.Assimilated && !c.deleteAssimilated {
		ctx.Logf(diag.Info, "DELETE on %s[%s]: Keeping on OpenZiti as this object was assimilated!", "Service", id)
		return nil
	}
	deleteParams := &service.DeleteServiceParams{
		ID:      id,
		Context: context.Background(),
	}
	deleteParams.SetTimeout(30 * time.Second)
	// ctx.Logf(diag.Info, "Calling delete on %s; output: %#v\n", id, *deleteParams)
	_, err = ce.client.Service.DeleteService(deleteParams, nil)
	if err != nil {
		return handleDeleteErr(ctx, err, id, "Config")
	}
	return nil
}
