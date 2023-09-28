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
	"github.com/openziti/edge-api/rest_management_api_client/edge_router"
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
type EdgeRouter struct{}

// Each resource has in input struct, defining what arguments it accepts.
type EdgeRouterArgs struct {
	BaseArgsEntity

	// app data
	AppData Tags `pulumi:"appData,optional"`

	// cost
	// Maximum: 65535
	// Minimum: 0
	Cost int64 `pulumi:"cost,optional"`

	// disabled
	Disabled bool `pulumi:"disabled,optional"`

	// is tunneler enabled
	IsTunnelerEnabled bool `pulumi:"isTunnelerEnabled,optional"`

	// no traversal
	NoTraversal bool `pulumi:"noTraversal,optional"`

	// role attributes
	RoleAttributes rest_model.Attributes `pulumi:"roleAttributes,optional"`
}

type VersionInfo struct {

	// arch
	// Required: true
	Arch string `pulumi:"arch"`

	// build date
	// Required: true
	BuildDate string `pulumi:"buildDate"`

	// os
	// Required: true
	Os string `pulumi:"os"`

	// revision
	// Required: true
	Revision string `pulumi:"revision"`

	// version
	// Required: true
	Version string `pulumi:"version"`
}

type CommonEdgeRouterProperties struct {

	// app data
	AppData Tags `pulumi:"appData,optional"`

	// cost
	// Required: true
	// Maximum: 65535
	// Minimum: 0
	Cost int64 `pulumi:"cost"`

	// disabled
	// Required: true
	Disabled bool `pulumi:"disabled"`

	// hostname
	// Required: true
	Hostname string `pulumi:"hostname"`

	// is online
	// Required: true
	IsOnline bool `pulumi:"isOnline"`

	// name
	// Required: true
	Name string `pulumi:"name"`

	// no traversal
	// Required: true
	NoTraversal bool `pulumi:"noTraversal"`

	// supported protocols
	// Required: true
	SupportedProtocols map[string]string `pulumi:"supportedProtocols"`

	// sync status
	// Required: true
	SyncStatus string `pulumi:"syncStatus"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type EdgeRouterState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	// EdgeRouterArgs
	// Here we define a required output called result.
	// Data interface{} `pulumi:"data"`
	BaseStateEntity

	CommonEdgeRouterProperties

	// cert pem
	CertPem *string `pulumi:"certPem,optional"`

	// enrollment created at
	// Format: date-time
	EnrollmentCreatedAt *string `pulumi:"enrollmentCreatedAt,optional"`

	// enrollment expires at
	// Format: date-time
	EnrollmentExpiresAt *string `pulumi:"enrollmentExpiresAt,optional"`

	// enrollment Jwt
	EnrollmentJWT *string `pulumi:"enrollmentJwt,optional"`

	// enrollment token
	EnrollmentToken *string `pulumi:"enrollmentToken,optional"`

	// fingerprint
	Fingerprint string `pulumi:"fingerprint,optional"`

	// is tunneler enabled
	// Required: true
	IsTunnelerEnabled bool `pulumi:"isTunnelerEnabled"`

	// is verified
	// Required: true
	IsVerified bool `pulumi:"isVerified"`

	// role attributes
	// Required: true
	RoleAttributes rest_model.Attributes `pulumi:"roleAttributes"`

	// unverified cert pem
	UnverifiedCertPem *string `pulumi:"unverifiedCertPem,optional"`

	// unverified fingerprint
	UnverifiedFingerprint *string `pulumi:"unverifiedFingerprint,optional"`

	// version info
	VersionInfo VersionInfo `json:"versionInfo"`
}

// All resources must implement Create at a minumum.
func (thiz *EdgeRouter) Create(ctx p.Context, name string, input EdgeRouterArgs, preview bool) (string, EdgeRouterState, error) {
	retErr := func(err error) (string, EdgeRouterState, error) {
		return "", EdgeRouterState{}, err
	}
	ce, c, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	confCreate := &rest_model.EdgeRouterCreate{
		AppData:           buildZitiTags(input.AppData),
		Cost:              &input.Cost,
		Disabled:          &input.Disabled,
		IsTunnelerEnabled: input.IsTunnelerEnabled,
		Name:              &input.Name,
		NoTraversal:       &input.NoTraversal,
		RoleAttributes:    &input.RoleAttributes,
		Tags:              buildZitiTags(input.Tags),
	}
	createParams := &edge_router.CreateEdgeRouterParams{
		EdgeRouter: confCreate,
		Context:    context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		previewState := EdgeRouterState{
			// EdgeRouterArgs: input,
			BaseStateEntity: BaseStateEntity{
				Links:     nil,
				CreatedAt: "",
				ID:        "",
				Tags:      input.Tags,
				UpdatedAt: "",
			},
			CommonEdgeRouterProperties: CommonEdgeRouterProperties{
				AppData:            input.AppData,
				Cost:               input.Cost,
				Disabled:           input.Disabled,
				Hostname:           input.Name,
				IsOnline:           false,
				Name:               input.Name,
				NoTraversal:        input.NoTraversal,
				SupportedProtocols: make(map[string]string),
				SyncStatus:         "",
			},
			CertPem:               nil,
			EnrollmentCreatedAt:   nil,
			EnrollmentExpiresAt:   nil,
			EnrollmentJWT:         nil,
			EnrollmentToken:       nil,
			Fingerprint:           "",
			IsTunnelerEnabled:     input.IsTunnelerEnabled,
			IsVerified:            false,
			RoleAttributes:        input.RoleAttributes,
			UnverifiedCertPem:     nil,
			UnverifiedFingerprint: nil,
			VersionInfo:           VersionInfo{},
		}
		// dumpStruct(ctx, "previewState", previewState)

		return name, previewState, nil
	}

	resp, err := ce.client.EdgeRouter.CreateEdgeRouter(createParams, nil)
	if err != nil {
		var badReq *edge_router.CreateEdgeRouterBadRequest
		if errors.As(err, &badReq) {
			err2, dupe := formatApiErrDupeCheck(ctx, badReq, badReq.Payload)
			if dupe && c.assimilate {
				// find identity by name...
				findParams := &edge_router.ListEdgeRoutersParams{
					Filter:  buildNameFilter(input.Name),
					Context: context.Background(),
				}
				findRet, err3 := ce.client.EdgeRouter.ListEdgeRouters(findParams, nil)
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
				state, err := readEdgeRouter(ce, existingId, input)
				if err != nil {
					ctx.Logf(diag.Error, "Assimilate failed: Fetch failed with: %s", err2)
					return retErr(err2)
				}
				updatedState, err := thiz.Update(ctx, existingId, state, input, preview)
				// dumpStruct(ctx, "response", updatedState)
				return existingId, updatedState, err
			}
			return retErr(err2)
		}

		return retErr(err)
	}
	createdId := resp.GetPayload().Data.ID
	state, err := readEdgeRouter(ce, createdId, input)
	if err != nil {
		return createdId, state, err
	}
	return createdId, state, nil
}

func (*EdgeRouter) Diff(ctx p.Context, id string, olds EdgeRouterState, news EdgeRouterArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
	}
	diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))

	diffWalk(ctx, diff, "appData", reflect.ValueOf(olds.CommonEdgeRouterProperties.AppData), reflect.ValueOf(news.AppData))
	if news.Cost != olds.CommonEdgeRouterProperties.Cost {
		diff["cost"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Disabled != olds.CommonEdgeRouterProperties.Disabled {
		diff["disabled"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.IsTunnelerEnabled != olds.IsTunnelerEnabled {
		diff["isTunnelerEnabled"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.NoTraversal != olds.CommonEdgeRouterProperties.NoTraversal {
		diff["noTraversal"] = p.PropertyDiff{Kind: p.Update}
	}
	diffStrArrayIgnoreOrder(ctx, diff, "roleAttributes", olds.RoleAttributes, news.RoleAttributes)

	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on EdgeRouter %s/%s: Found %d diffs: %v", news.Name, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func readEdgeRouter(ce *CacheEntry, id string, input EdgeRouterArgs) (EdgeRouterState, error) {
	params := &edge_router.DetailEdgeRouterParams{
		ID:      id,
		Context: context.Background(),
	}
	detailResp, err := ce.client.EdgeRouter.DetailEdgeRouter(params, nil)
	if err != nil {
		return EdgeRouterState{}, err
	}
	data := detailResp.GetPayload().Data
	// fmt.Printf("get  output: %+v\n", respPayload)
	return EdgeRouterState{
		// EdgeRouterArgs:  input,
		BaseStateEntity: buildBaseState(data.BaseEntity),
		CertPem:         data.CertPem,
		CommonEdgeRouterProperties: CommonEdgeRouterProperties{
			AppData:            buildTags(*data.AppData),
			Cost:               *data.Cost,
			Disabled:           *data.Disabled,
			Hostname:           *data.Hostname,
			IsOnline:           *data.IsOnline,
			Name:               *data.Name,
			NoTraversal:        *data.NoTraversal,
			SupportedProtocols: data.SupportedProtocols,
			SyncStatus:         *data.SyncStatus,
		},
		EnrollmentCreatedAt:   iftfden(data.EnrollmentCreatedAt != nil, func() string { return data.EnrollmentCreatedAt.String() }),
		EnrollmentExpiresAt:   iftfden(data.EnrollmentExpiresAt != nil, func() string { return data.EnrollmentCreatedAt.String() }),
		EnrollmentJWT:         data.EnrollmentJWT,
		EnrollmentToken:       data.EnrollmentToken,
		Fingerprint:           data.Fingerprint,
		IsTunnelerEnabled:     *data.IsTunnelerEnabled,
		IsVerified:            *data.IsVerified,
		RoleAttributes:        iftfe(data.RoleAttributes != nil, func() []string { return *data.RoleAttributes }, []string{}),
		UnverifiedCertPem:     data.UnverifiedCertPem,
		UnverifiedFingerprint: data.UnverifiedFingerprint,
		VersionInfo: VersionInfo{
			Arch:      *data.VersionInfo.Arch,
			BuildDate: *data.VersionInfo.BuildDate,
			Os:        *data.VersionInfo.Os,
			Revision:  *data.VersionInfo.Revision,
			Version:   *data.VersionInfo.Version,
		},
	}, nil
}

func (*EdgeRouter) Read(ctx p.Context, id string, inputs EdgeRouterArgs, state EdgeRouterState) (string, EdgeRouterArgs, EdgeRouterState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	readState, err := readEdgeRouter(ce, id, inputs)
	if err != nil {
		return id, inputs, readState, err
	}
	return id, inputs, readState, nil
}

func (*EdgeRouter) Update(ctx p.Context, id string, olds EdgeRouterState, news EdgeRouterArgs, preview bool) (EdgeRouterState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}
	if err != nil {
		return olds, err
	}
	updateData := &rest_model.EdgeRouterUpdate{
		AppData:           buildZitiTags(news.AppData),
		Cost:              &news.Cost,
		Disabled:          &news.Disabled,
		IsTunnelerEnabled: news.IsTunnelerEnabled,
		Name:              &news.Name,
		NoTraversal:       &news.NoTraversal,
		RoleAttributes:    &news.RoleAttributes,
		Tags:              buildZitiTags(news.Tags),
	}
	updateParams := &edge_router.UpdateEdgeRouterParams{
		EdgeRouter: updateData,
		ID:         id,
		Context:    context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return olds, nil
	}

	_, err = ce.client.EdgeRouter.UpdateEdgeRouter(updateParams, nil)
	if err != nil {
		var badReq *edge_router.UpdateEdgeRouterBadRequest
		if errors.As(err, &badReq) {
			return olds, formatApiErr(ctx, badReq, badReq.Payload)
		}
		return olds, err
	}

	readState, err := readEdgeRouter(ce, id, news)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*EdgeRouter) Delete(ctx p.Context, id string, _ EdgeRouterState) error {
	ce, _, err := initClient(ctx)
	if err != nil {
		return err
	}
	deleteParams := &edge_router.DeleteEdgeRouterParams{
		ID:      id,
		Context: context.Background(),
	}
	deleteParams.SetTimeout(30 * time.Second)
	// ctx.Logf(diag.Info, "Calling delete on %s; output: %#v\n", id, *deleteParams)
	_, err = ce.client.EdgeRouter.DeleteEdgeRouter(deleteParams, nil)
	if err != nil {
		return handleDeleteErr(ctx, err, id, "Config")
	}
	return nil
}
