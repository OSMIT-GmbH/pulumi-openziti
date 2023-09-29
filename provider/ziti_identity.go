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
type Identity struct{}

type IdentityCreateEnrollment struct {

	// ott
	Ott bool `pulumi:"ott,optional"`

	// ottca
	Ottca string `pulumi:"ottca,optional"`

	// updb
	Updb string `pulumi:"updb,optional"`
}

// Each resource has in input struct, defining what arguments it accepts.
type IdentityArgs struct {
	BaseArgsEntity

	// app data
	AppData Tags `pulumi:"appData,optional"`

	// auth policy Id
	AuthPolicyID string `pulumi:"authPolicyId,optional"`

	// default hosting cost
	DefaultHostingCost rest_model.TerminatorCost `pulumi:"defaultHostingCost,optional"`

	// default hosting precedence
	DefaultHostingPrecedence rest_model.TerminatorPrecedence `pulumi:"defaultHostingPrecedence,optional"`

	// enrollment
	Enrollment IdentityCreateEnrollment `pulumi:"enrollment,optional"`

	// external Id
	ExternalID *string `pulumi:"externalId,optional"`

	// is admin
	// Required: true
	IsAdmin bool `pulumi:"isAdmin"`

	// role attributes
	RoleAttributes rest_model.Attributes `pulumi:"roleAttributes,optional"`

	// service hosting costs
	ServiceHostingCosts rest_model.TerminatorCostMap `pulumi:"serviceHostingCosts,optional"`

	// service hosting precedences
	ServiceHostingPrecedences rest_model.TerminatorPrecedenceMap `pulumi:"serviceHostingPrecedences,optional"`

	// type
	// Required: true
	Type rest_model.IdentityType `pulumi:"type"`
}

type IdentityEnrollmentsOtt struct {

	// expires at
	// Format: date-time
	ExpiresAt string `pulumi:"expiresAt,optional"`

	// id
	ID string `pulumi:"id,optional"`

	// jwt
	JWT string `pulumi:"jwt,optional"`

	// token
	Token string `pulumi:"token,optional"`
}

type IdentityEnrollmentsOttca struct {

	// ca
	Ca EntityRef `pulumi:"ca,optional"`

	// ca Id
	CaID string `pulumi:"caId,optional"`

	// expires at
	// Format: date-time
	ExpiresAt string `pulumi:"expiresAt,optional"`

	// id
	ID string `pulumi:"id,optional"`

	// jwt
	JWT string `pulumi:"jwt,optional"`

	// token
	Token string `pulumi:"token,optional"`
}
type IdentityEnrollmentsUpdb struct {

	// expires at
	// Format: date-time
	ExpiresAt string `pulumi:"expiresAt,optional"`

	// id
	ID string `pulumi:"id,optional"`

	// jwt
	JWT string `pulumi:"jwt,optional"`

	// token
	Token string `pulumi:"token,optional"`
}
type IdentityEnrollments struct {

	// ott
	Ott *IdentityEnrollmentsOtt `pulumi:"ott,optional"`

	// ottca
	Ottca *IdentityEnrollmentsOttca `pulumi:"ottca,optional"`

	// updb
	Updb *IdentityEnrollmentsUpdb `pulumi:"updb,optional"`
}

type IdentityAuthenticatorsCert struct {

	// fingerprint
	Fingerprint string `pulumi:"fingerprint,optional"`

	// id
	ID string `pulumi:"id,optional"`
}
type IdentityAuthenticatorsUpdb struct {

	// id
	ID string `pulumi:"id,optional"`

	// username
	Username string `pulumi:"username,optional"`
}
type IdentityAuthenticators struct {

	// cert
	Cert IdentityAuthenticatorsCert `pulumi:"cert,optional"`

	// updb
	Updb IdentityAuthenticatorsUpdb `pulumi:"updb,optional"`
}

type EnvInfo struct {

	// arch
	Arch string `pulumi:"arch,optional"`

	// os
	Os string `pulumi:"os,optional"`

	// os release
	OsRelease string `pulumi:"osRelease,optional"`

	// os version
	OsVersion string `pulumi:"osVersion,optional"`
}

type SdkInfo struct {

	// app Id
	AppID string `pulumi:"appId,optional"`

	// app version
	AppVersion string `pulumi:"appVersion,optional"`

	// branch
	Branch string `pulumi:"branch,optional"`

	// revision
	Revision string `pulumi:"revision,optional"`

	// type
	Type string `pulumi:"type,optional"`

	// version
	Version string `pulumi:"version,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type IdentityState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	IdentityArgs
	// Here we define a required output called result.
	// Data interface{} `pulumi:"data"`
	BaseStateEntity

	// app data
	AppData Tags `pulumi:"appData,optional"`

	// auth policy
	// Required: true
	AuthPolicy EntityRef `pulumi:"authPolicy"`

	// auth policy Id
	// Required: true
	AuthPolicyID string `pulumi:"authPolicyId"`

	// authenticators
	// Required: true
	Authenticators rest_model.IdentityAuthenticators `pulumi:"authenticators"`

	// default hosting cost
	// Required: true
	DefaultHostingCost rest_model.TerminatorCost `pulumi:"defaultHostingCost"`

	// default hosting precedence
	DefaultHostingPrecedence rest_model.TerminatorPrecedence `pulumi:"defaultHostingPrecedence,optional"`

	// disabled
	// Required: true
	Disabled bool `pulumi:"disabled"`

	// disabled at
	// Format: date-time
	DisabledAt *string `pulumi:"disabledAt,optional"`

	// disabled until
	// Format: date-time
	DisabledUntil *string `pulumi:"disabledUntil,optional"`

	// enrollment
	// Required: true
	Enrollment IdentityEnrollments `pulumi:"enrollment"`

	// env info
	// Required: true
	EnvInfo EnvInfo `pulumi:"envInfo"`

	// external Id
	ExternalID *string `pulumi:"externalId,optional"`

	// has Api session
	// Required: true
	HasAPISession bool `pulumi:"hasApiSession"`

	// has edge router connection
	// Required: true
	HasEdgeRouterConnection bool `pulumi:"hasEdgeRouterConnection"`

	// is admin
	// Required: true
	IsAdmin bool `pulumi:"isAdmin"`

	// is default admin
	// Required: true
	IsDefaultAdmin bool `pulumi:"isDefaultAdmin"`

	// is mfa enabled
	// Required: true
	IsMfaEnabled bool `pulumi:"isMfaEnabled"`

	// name
	// Required: true
	Name string `pulumi:"name"`

	// role attributes
	// Required: true
	RoleAttributes []string `pulumi:"roleAttributes"`

	// sdk info
	// Required: true
	SdkInfo SdkInfo `pulumi:"sdkInfo"`

	// service hosting costs
	// Required: true
	ServiceHostingCosts rest_model.TerminatorCostMap `pulumi:"serviceHostingCosts"`

	// service hosting precedences
	// Required: true
	ServiceHostingPrecedences rest_model.TerminatorPrecedenceMap `pulumi:"serviceHostingPrecedences"`

	// type
	// Required: true
	Type EntityRef `pulumi:"type"`

	// type Id
	// Required: true
	TypeID string `pulumi:"typeId"`
}

// All resources must implement Create at a minumum.
func (thiz *Identity) Create(ctx p.Context, name string, input IdentityArgs, preview bool) (string, IdentityState, error) {
	// bail out now when we are in preview mode
	if preview {
		return IdPreviewPrefix + name, IdentityState{
			IdentityArgs:              input,
			BaseStateEntity:           buildBaseStatePreviewEntity(name, input.BaseArgsEntity),
			AppData:                   input.AppData,
			AuthPolicy:                EntityRef{},
			AuthPolicyID:              "",
			Authenticators:            rest_model.IdentityAuthenticators{},
			DefaultHostingCost:        input.DefaultHostingCost,
			DefaultHostingPrecedence:  input.DefaultHostingPrecedence,
			Disabled:                  false,
			DisabledAt:                nil,
			DisabledUntil:             nil,
			Enrollment:                IdentityEnrollments{},
			EnvInfo:                   EnvInfo{},
			ExternalID:                input.ExternalID,
			HasAPISession:             false,
			HasEdgeRouterConnection:   false,
			IsAdmin:                   input.IsAdmin,
			IsDefaultAdmin:            false,
			IsMfaEnabled:              false,
			Name:                      input.Name,
			RoleAttributes:            input.RoleAttributes,
			SdkInfo:                   SdkInfo{},
			ServiceHostingCosts:       input.ServiceHostingCosts,
			ServiceHostingPrecedences: input.ServiceHostingPrecedences,
			Type: EntityRef{
				Links:  nil,
				Entity: "",
				ID:     "",
				Name:   fmt.Sprintf("%s", input.Type),
			},
			TypeID: "",
		}, nil
	}

	retErr := func(err error) (string, IdentityState, error) {
		return "", IdentityState{IdentityArgs: input}, err
	}
	ce, c, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	confCreate := &rest_model.IdentityCreate{
		AppData:                  buildZitiTags(input.AppData),
		AuthPolicyID:             &input.AuthPolicyID,
		DefaultHostingCost:       &input.DefaultHostingCost,
		DefaultHostingPrecedence: input.DefaultHostingPrecedence,
		Enrollment: &rest_model.IdentityCreateEnrollment{
			Ott:   input.Enrollment.Ott,
			Ottca: input.Enrollment.Ottca,
			Updb:  input.Enrollment.Updb,
		},
		ExternalID:                input.ExternalID,
		IsAdmin:                   &input.IsAdmin,
		Name:                      &input.Name,
		RoleAttributes:            &input.RoleAttributes,
		ServiceHostingCosts:       input.ServiceHostingCosts,
		ServiceHostingPrecedences: input.ServiceHostingPrecedences,
		Tags:                      buildZitiTags(input.Tags),
		Type:                      &input.Type,
	}
	createParams := &identity.CreateIdentityParams{
		Identity: confCreate,
		Context:  context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	resp, err := ce.client.Identity.CreateIdentity(createParams, nil)
	if err != nil {
		var badReq *identity.CreateIdentityBadRequest
		if errors.As(err, &badReq) {
			err2, dupe := formatApiErrDupeCheck(ctx, badReq, badReq.Payload)
			if dupe && c.assimilate {
				// find identity by name...
				findParams := &identity.ListIdentitiesParams{
					Filter:  buildNameFilter(input.Name),
					Context: context.Background(),
				}
				findRet, err3 := ce.client.Identity.ListIdentities(findParams, nil)
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
				state, err := readIdentity(ce, existingId, input, true)
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
	state, err := readIdentity(ce, createdId, input, false)
	if err != nil {
		return createdId, state, err
	}
	return createdId, state, nil
}

func (*Identity) Diff(ctx p.Context, id string, olds IdentityState, news IdentityArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))

	diffWalk(ctx, diff, "appData", reflect.ValueOf(olds.AppData), reflect.ValueOf(news.AppData))
	if news.AuthPolicyID != olds.AuthPolicyID && news.AuthPolicyID != "" && olds.AuthPolicyID != "default" {
		diff["authPolicyId"] = p.PropertyDiff{Kind: p.Update}
	}

	if news.DefaultHostingCost != olds.DefaultHostingCost {
		diff["defaultHostingCost"] = p.PropertyDiff{Kind: p.Update}
	}

	if news.DefaultHostingPrecedence != olds.DefaultHostingPrecedence && news.DefaultHostingPrecedence != "" && olds.DefaultHostingPrecedence != "default" {
		diff["defaultHostingPrecedence"] = p.PropertyDiff{Kind: p.Update}
	}

	// enrollment is only changeable on creation
	//if news.Enrollment.Ott != olds.IdentityArgs.Enrollment.Ott {
	//	diff["enrollment.ott"] = p.PropertyDiff{Kind: p.Update}
	//}
	//if news.Enrollment.Ottca != olds.IdentityArgs.Enrollment.Ottca {
	//	diff["enrollment.ottCa"] = p.PropertyDiff{Kind: p.Update}
	//}
	//if news.Enrollment.Updb != olds.IdentityArgs.Enrollment.Updb {
	//	diff["enrollment.updb"] = p.PropertyDiff{Kind: p.Update}
	//}

	if news.ExternalID != olds.ExternalID {
		diff["externalId"] = p.PropertyDiff{Kind: p.Update}
	}

	if news.IsAdmin != olds.IsAdmin {
		diff["isAdmin"] = p.PropertyDiff{Kind: p.Update}
	}

	diffWalk(ctx, diff, "roleAttributes", reflect.ValueOf(olds.RoleAttributes), reflect.ValueOf(news.RoleAttributes))

	diffWalk(ctx, diff, "serviceHostingCosts", reflect.ValueOf(olds.ServiceHostingCosts), reflect.ValueOf(news.ServiceHostingCosts))

	diffWalk(ctx, diff, "serviceHostingPrecedences", reflect.ValueOf(olds.ServiceHostingPrecedences), reflect.ValueOf(news.ServiceHostingPrecedences))

	if string(news.Type) != olds.Type.Name {
		diff["type"] = p.PropertyDiff{Kind: p.Update}
	}

	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on Identity %s/%s: Found %d diffs: %v", news.Name, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func readIdentity(ce *CacheEntry, id string, input IdentityArgs, assimilated bool) (IdentityState, error) {
	params := &identity.DetailIdentityParams{
		ID:      id,
		Context: context.Background(),
	}
	detailResp, err := ce.client.Identity.DetailIdentity(params, nil)
	if err != nil {
		return IdentityState{IdentityArgs: input}, err
	}
	respPayload := detailResp.GetPayload()
	data := respPayload.Data
	// fmt.Printf("get  output: %+v\n", respPayload)
	return IdentityState{
		IdentityArgs:             input,
		BaseStateEntity:          buildBaseState(data.BaseEntity, assimilated),
		AppData:                  buildTags(*data.AppData),
		AuthPolicy:               buildEntityRef(data.AuthPolicy),
		AuthPolicyID:             *data.AuthPolicyID,
		Authenticators:           *data.Authenticators,
		DefaultHostingCost:       *data.DefaultHostingCost,
		DefaultHostingPrecedence: data.DefaultHostingPrecedence,
		Disabled:                 *data.Disabled,
		DisabledAt:               iftfden(data.DisabledAt != nil, func() string { return data.DisabledAt.String() }),
		DisabledUntil:            iftfden(data.DisabledUntil != nil, func() string { return data.DisabledUntil.String() }),
		Enrollment: IdentityEnrollments{
			Ott: iftfden(data.Enrollment.Ott != nil, func() IdentityEnrollmentsOtt {
				return IdentityEnrollmentsOtt{
					ExpiresAt: data.Enrollment.Ott.ExpiresAt.String(),
					ID:        data.Enrollment.Ott.ID,
					JWT:       data.Enrollment.Ott.JWT,
					Token:     data.Enrollment.Ott.Token,
				}
			}),
			Ottca: iftfden(data.Enrollment.Ottca != nil, func() IdentityEnrollmentsOttca {
				return IdentityEnrollmentsOttca{
					Ca:        buildEntityRef(data.Enrollment.Ottca.Ca),
					CaID:      data.Enrollment.Ottca.CaID,
					ExpiresAt: data.Enrollment.Ottca.ExpiresAt.String(),
					ID:        data.Enrollment.Ottca.ID,
					JWT:       data.Enrollment.Ottca.JWT,
					Token:     data.Enrollment.Ottca.Token,
				}
			}),
			Updb: iftfden(data.Enrollment.Updb != nil, func() IdentityEnrollmentsUpdb {
				return IdentityEnrollmentsUpdb{
					ExpiresAt: data.Enrollment.Updb.ExpiresAt.String(),
					ID:        data.Enrollment.Updb.ID,
					JWT:       data.Enrollment.Updb.JWT,
					Token:     data.Enrollment.Updb.Token,
				}
			}),
		},
		EnvInfo: EnvInfo{
			Arch:      data.EnvInfo.Arch,
			Os:        data.EnvInfo.Os,
			OsRelease: data.EnvInfo.OsRelease,
			OsVersion: data.EnvInfo.OsVersion,
		},
		ExternalID:              data.ExternalID,
		HasAPISession:           *data.HasAPISession,
		HasEdgeRouterConnection: *data.HasEdgeRouterConnection,
		IsAdmin:                 *data.IsAdmin,
		IsDefaultAdmin:          *data.IsDefaultAdmin,
		IsMfaEnabled:            *data.IsMfaEnabled,
		Name:                    *data.Name,
		RoleAttributes:          iftfe(data.RoleAttributes != nil, func() []string { return *data.RoleAttributes }, []string{}),
		SdkInfo: SdkInfo{
			AppID:      data.SdkInfo.AppID,
			AppVersion: data.SdkInfo.AppVersion,
			Branch:     data.SdkInfo.Branch,
			Revision:   data.SdkInfo.Revision,
			Type:       data.SdkInfo.Type,
			Version:    data.SdkInfo.Version,
		},
		ServiceHostingCosts:       data.ServiceHostingCosts,
		ServiceHostingPrecedences: data.ServiceHostingPrecedences,
		Type:                      buildEntityRef(data.Type),
		TypeID:                    *data.TypeID,
	}, nil
}

func (*Identity) Read(ctx p.Context, id string, inputs IdentityArgs, state IdentityState) (string, IdentityArgs, IdentityState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	readState, err := readIdentity(ce, id, inputs, state.Assimilated)
	if err != nil {
		return id, inputs, readState, err
	}
	return id, inputs, readState, nil
}

func (*Identity) Update(ctx p.Context, id string, olds IdentityState, news IdentityArgs, preview bool) (IdentityState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}
	if err != nil {
		return olds, err
	}
	updateData := &rest_model.IdentityUpdate{
		AppData:                  buildZitiTags(news.AppData),
		AuthPolicyID:             &news.AuthPolicyID,
		DefaultHostingCost:       &news.DefaultHostingCost,
		DefaultHostingPrecedence: news.DefaultHostingPrecedence,

		ExternalID:                news.ExternalID,
		IsAdmin:                   &news.IsAdmin,
		Name:                      &news.Name,
		RoleAttributes:            &news.RoleAttributes,
		ServiceHostingCosts:       news.ServiceHostingCosts,
		ServiceHostingPrecedences: news.ServiceHostingPrecedences,
		Tags:                      buildZitiTags(news.Tags),
		Type:                      &news.Type,
	}
	updateParams := &identity.UpdateIdentityParams{
		Identity: updateData,
		ID:       id,
		Context:  context.Background(),
	}
	// dumpStruct(ctx, confCreate)

	// bail out now when we are in preview mode
	if preview {
		return olds, nil
	}

	_, err = ce.client.Identity.UpdateIdentity(updateParams, nil)
	if err != nil {
		var badReq *identity.UpdateIdentityBadRequest
		if errors.As(err, &badReq) {
			return olds, formatApiErr(ctx, badReq, badReq.Payload)
		}
		return olds, err
	}

	readState, err := readIdentity(ce, id, news, olds.Assimilated)
	if err != nil {
		return readState, err
	}
	return readState, nil
}

func (*Identity) Delete(ctx p.Context, id string, state IdentityState) error {
	ce, c, err := initClient(ctx)
	if err != nil {
		return err
	}
	if state.Assimilated && !c.deleteAssimilated {
		ctx.Logf(diag.Info, "DELETE on %s[%s]: Keeping on OpenZiti as this object was assimilated!", "Identity", id)
		return nil
	}
	deleteParams := &identity.DeleteIdentityParams{
		ID:      id,
		Context: context.Background(),
	}
	deleteParams.SetTimeout(30 * time.Second)
	// ctx.Logf(diag.Info, "Calling delete on %s; output: %#v\n", id, *deleteParams)
	_, err = ce.client.Identity.DeleteIdentity(deleteParams, nil)
	if err != nil {
		return handleDeleteErr(ctx, err, id, "Config")
	}
	return nil
}
