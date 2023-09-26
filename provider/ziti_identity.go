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
	"context"
	"errors"
	"fmt"
	"github.com/openziti/edge-api/rest_management_api_client/config"
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
type ConfigObj struct{}

// Each resource has in input struct, defining what arguments it accepts.
type ConfigArgs struct {
	BaseArgsEntity

	ConfigTypeName string `pulumi:"configTypeName"`

	// The data section of a config is based on the schema of its type
	// Required: true
	Data interface{} `pulumi:"data"`

	// Data map[string]interface{} `pulumi:"data"`
	//Data string `pulumi:"data"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type ConfigState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	ConfigArgs
	// Here we define a required output called result.
	// Data interface{} `pulumi:"data"`
	BaseStateEntity
	// config type
	// Required: true
	ConfigType EntityRef `pulumi:"configType"`

	// config type Id
	// Required: true
	ConfigTypeID *string `pulumi:"configTypeId"`

	// The data section of a config is based on the schema of its type
	// Required: true
	Data interface{} `pulumi:"data"`
}

// All resources must implement Create at a minumum.
func (*ConfigObj) Create(ctx p.Context, name string, input ConfigArgs, preview bool) (string, ConfigState, error) {
	retErr := func(err error) (string, ConfigState, error) {
		return "", ConfigState{ConfigArgs: input}, err
	}
	ce, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}
	configTypeID, err := getConfigTypeId(ce, input.ConfigTypeName)
	if err != nil {
		return retErr(err)
	}

	tags := buildZitiTags(input.Tags)
	confCreate := &rest_model.ConfigCreate{
		ConfigTypeID: &configTypeID,
		Data:         &input.Data,
		Name:         &input.Name,
		Tags:         &tags,
	}
	confParams := &config.CreateConfigParams{
		Config:  confCreate,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return name, ConfigState{ConfigArgs: input}, nil
	}

	resp, err := ce.client.Config.CreateConfig(confParams, nil)
	if err != nil {
		var badReq *config.CreateConfigBadRequest
		if errors.As(err, &badReq) {
			return retErr(formatApiErr(ctx, badReq, badReq.Payload))
		}

		return retErr(err)
	}
	createdId := resp.GetPayload().Data.ID
	state, err := readConfig(ce, createdId, input)
	if err != nil {
		return createdId, state, err
	}
	return createdId, state, nil
}

func (*ConfigObj) Diff(ctx p.Context, id string, olds ConfigState, news ConfigArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))
	if news.ConfigTypeName != olds.ConfigTypeName {
		diff["configTypeName"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	diffWalk(ctx, diff, "data", reflect.ValueOf(olds.Data), reflect.ValueOf(news.Data))
	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("Found %d diffs: %v", len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func readConfig(ce CacheEntry, id string, input ConfigArgs) (ConfigState, error) {
	detailConfParams := &config.DetailConfigParams{
		ID:      id,
		Context: context.Background(),
	}
	detailResp, err := ce.client.Config.DetailConfig(detailConfParams, nil)
	if err != nil {
		return ConfigState{ConfigArgs: input}, err
	}
	respPayload := detailResp.GetPayload()
	// fmt.Printf("get  output: %+v\n", respPayload)
	return ConfigState{ConfigArgs: input,
		BaseStateEntity: buildBaseState(respPayload.Data.BaseEntity),
		ConfigType:      buildEntityRef(respPayload.Data.ConfigType),
		ConfigTypeID:    respPayload.Data.ConfigTypeID,
		Data:            respPayload.Data.Data,
	}, nil
}

func (*ConfigObj) Read(ctx p.Context, id string, inputs ConfigArgs, state ConfigState) (string, ConfigArgs, ConfigState, error) {
	ce, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	readState, err := readConfig(ce, id, inputs)
	if err != nil {
		return id, inputs, readState, err
	}
	return id, inputs, readState, nil
}

func (*ConfigObj) Update(ctx p.Context, id string, olds ConfigState, news ConfigArgs, preview bool) (ConfigState, error) {
	ce, err := initClient(ctx)
	if err != nil {
		return olds, err
	}
	if err != nil {
		return olds, err
	}
	tags := buildZitiTags(news.Tags)
	confCreate := &rest_model.ConfigUpdate{
		Data: &news.Data,
		Name: &news.Name,
		Tags: &tags,
	}
	confParams := &config.UpdateConfigParams{
		Config:  confCreate,
		ID:      id,
		Context: context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return olds, nil
	}

	_, err = ce.client.Config.UpdateConfig(confParams, nil)
	if err != nil {
		var badReq *config.UpdateConfigBadRequest
		if errors.As(err, &badReq) {
			return olds, formatApiErr(ctx, badReq, badReq.Payload)
		}
		return olds, err
	}

	readState, err := readConfig(ce, id, news)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*ConfigObj) Delete(ctx p.Context, id string, _ ConfigState) error {
	ce, err := initClient(ctx)
	if err != nil {
		return err
	}
	deleteParams := &config.DeleteConfigParams{
		ID: id,
	}
	deleteParams.SetTimeout(30 * time.Second)
	// ctx.Logf(diag.Info, "Calling delete on %s; output: %#v\n", id, *deleteParams)
	_, err = ce.client.Config.DeleteConfig(deleteParams, nil)
	if err != nil {
		return handleDeleteErr(ctx, err, id, "Config")
	}
	return nil
}
