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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	edgeapis "github.com/openziti/edge-api/rest_management_api_client"
	"github.com/openziti/edge-api/rest_management_api_client/config"
	"github.com/openziti/edge-api/rest_model"
	"github.com/openziti/edge-api/rest_util"
	"github.com/openziti/sdk-golang/ziti"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"reflect"
	"time"

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
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Config: infer.Config[*OpenZitiProviderConfig](),
	})
}

type OpenZitiProviderConfig struct {
	User     string `pulumi:"user"`
	Password string `pulumi:"password" provider:"secret"`
	Uri      string `pulumi:"uri"`
	Insecure bool   `pulumi:"insecure,optional"`
	Version  string `pulumi:"version,optional"` // version seems to be provided automatically
	cacheKey string
}

var _ = (infer.Annotated)((*OpenZitiProviderConfig)(nil))

func (c *OpenZitiProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.User, "The username. It's important but not secret.")
	a.Describe(&c.Password, "The password. It is very secret.")
	a.Describe(&c.Uri, `The URI to the API`)
	// a.SetDefault(&c.Insecure, false)
}

var _ = (infer.CustomConfigure)((*OpenZitiProviderConfig)(nil))

func (c *OpenZitiProviderConfig) Configure(ctx p.Context) error {
	// apiUrl, err := url.Parse(c.Uri)

	//if err != nil {
	//	// return nil, errors.Wrap(err, "could not parse ZtAPI from configuration as URI")
	//	// fmt.Errorf("no session token returned from login request to %v. Received: %v", c.Uri, zitiLogin.String())
	//	return err
	//}
	c.cacheKey = fmt.Sprintf("%s:%s:%s", c.Uri, c.User, c.Password)

	//ctx.Log(diag.Info, msg)
	return nil
}

type Link struct {

	// comment
	Comment string `pulumi:"comment,optional"`

	// href
	// Required: true
	// Format: uri
	Href string `pulumi:"href"`

	// method
	Method string `pulumi:"method,optional"`
}

type Links map[string]Link

type Tags map[string]interface{}

type EntityRef struct {

	// links
	Links Links `pulumi:"_links"`

	// entity
	Entity string `pulumi:"entity,optional"`

	// id
	ID string `pulumi:"id,optional"`

	// name
	Name string `pulumi:"name,optional"`
}

type BaseArgsEntity struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Name string `pulumi:"name"`

	// tags
	Tags Tags `pulumi:"tags,optional"`
}

type BaseStateEntity struct {

	// links
	// Required: true
	Links Links `pulumi:"_links"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt string `pulumi:"createdAt"`

	// ID
	// Required: true
	ID string `pulumi:"id"`

	// tags
	Tags Tags `pulumi:"tags,optional"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt string `pulumi:"updatedAt"`
}

func buildLinks(src rest_model.Links) Links {
	ret := make(Links)
	for key, link := range src {
		ret[key] = Link{Comment: link.Comment, Method: link.Method, Href: link.Href.String()}
	}
	return ret
}

func buildBaseState(src rest_model.BaseEntity) BaseStateEntity {
	tags := buildTags(*src.Tags)
	return BaseStateEntity{Links: buildLinks(src.Links),
		ID:        *src.ID,
		CreatedAt: src.CreatedAt.String(),
		Tags:      tags,
		UpdatedAt: src.UpdatedAt.String(),
	}
}

func buildEntityRef(src *rest_model.EntityRef) EntityRef {
	return EntityRef{Links: buildLinks(src.Links),
		Entity: src.Entity,
		ID:     src.ID,
		Name:   src.Name,
	}
}

func buildTags(tags rest_model.Tags) Tags {
	//if len(tags.SubTags) == 0 {
	//	return rest_model.Tags{}
	//}
	out := make(Tags)
	for key, value := range tags.SubTags {
		out[key] = value
	}
	return out
}

func buildZitiTags(tags Tags) rest_model.Tags {
	//if len(tags.SubTags) == 0 {
	//	return rest_model.Tags{}
	//}
	out := make(rest_model.SubTags)
	for key, value := range tags {
		out[key] = value
	}
	return rest_model.Tags{SubTags: out}
}

func dumpStruct(ctx p.Context, data interface{}) {
	out, err := json.Marshal(data)
	if err != nil {
		ctx.Logf(diag.Info, "ERROR: failed serializing data: %s => %#v", err.Error(), data)
	} else {
		ctx.Logf(diag.Info, "OK: createConf(json)=%s", string(out))
	}
}

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

type CacheEntry struct {
	//apiSession *rest_model.CurrentAPISessionDetail
	client      *edgeapis.ZitiEdgeManagement
	configTypes map[string]string
	identities  map[string]string
}

var cache = make(map[string]CacheEntry)

func getConfigTypeId(cache CacheEntry, name string) (string, error) {
	if len(cache.configTypes) == 0 {
		// initial load
		// filter := fmt.Sprintf("name=\"%v\"", sdk.ZrokProxyConfig)
		limit := int64(100)
		offset := int64(0)
		listReq := &config.ListConfigTypesParams{
			// Filter:  &filter,
			Limit:   &limit,
			Offset:  &offset,
			Context: context.Background(),
		}
		ctResp, err := cache.client.Config.ListConfigTypes(listReq, nil)
		if err != nil {
			return "", err
		}

		cache.configTypes = make(map[string]string)
		for _, entity := range ctResp.GetPayload().Data {
			// wrapper := api.Wrap(entity)
			// configTypes[wrapper.String("name")] = wrapper.String("id")
			cache.configTypes[*entity.Name] = *entity.ID
		}
		// fmt.Printf("retrived configValues: %+v\n", cache.configTypes)
	}
	typeId := cache.configTypes[name]
	if typeId == "" {
		return typeId, errors.New(fmt.Sprintf("No configType declared for `%s` - known types: %+v", name, cache.configTypes))
	}
	return typeId, nil
}

func initClient(ctx p.Context) (CacheEntry, error) {
	c := infer.GetConfig[OpenZitiProviderConfig](ctx)
	ce := cache[c.cacheKey]
	if ce.client == nil {
		// creds := edge_apis.New([]*x509.Certificate{testIdCerts.cert}, testIdCerts.key)
		caPool, err := ziti.GetControllerWellKnownCaPool(c.Uri)
		if err != nil {
			return ce, err
		}

		client, err := rest_util.NewEdgeManagementClientWithUpdb(c.User, c.Password, c.Uri, caPool)
		if err != nil {
			return ce, err
		}

		//creds := edgeapis.NewUpdbCredentials(c.User, c.Password)
		//// creds.CaPool = caPool
		//
		//client := edgeapis.NewManagementApiClient(apiUrl, caPool)
		//apiSession, err := client.Authenticate(creds, nil)
		//if err != nil {
		//	return err
		//}

		// fmt.Printf("identity name: %#v; token: s\n", client)
		ce.client = client
	}
	return ce, nil
}

func formatApiErr(ctx p.Context, err error, apiError *rest_model.APIErrorEnvelope) error {
	// if errors.Is(err, config.CreateConfigBadRequest) {
	errOut, err2 := json.Marshal(apiError.Error)
	if err2 != nil {
		ctx.Logf(diag.Error, "ERROR: type: ErrorString: %v, MarshallingError: %s Payload.Error=%+v  PalyloadErrorCause=%#v PayloadMetadata: %+v\n", err.Error(), err2.Error(), apiError.Error, apiError.Error.Cause, apiError.Meta)
		return errors.Join(err, err2)
	}
	return fmt.Errorf("ERROR: type: ErrorString: %s, Payload=%s\n", err.Error(), string(errOut))
}
func handleDeleteErr(ctx p.Context, err error, id string, typeName string) error {
	var apiError *runtime.APIError
	if errors.As(err, &apiError) {
		if apiError.Code == 404 {
			ctx.Logf(diag.Warning, "DELETE on %s[%s] returned 404 - assuming already deleted!", typeName, id)
			return nil
		}
	}
	return err
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

func diffWalk(ctx p.Context, diff map[string]p.PropertyDiff, path string, old reflect.Value, new reflect.Value) {
	ctx.Log(diag.Debug, fmt.Sprintf("diffWalk: visiting %s: old: %s new: %s", path, old.String(), new.String()))
	// Indirect through pointers and interfaces
	for old.Kind() == reflect.Ptr || old.Kind() == reflect.Interface {
		old = old.Elem()
	}
	for new.Kind() == reflect.Ptr || new.Kind() == reflect.Interface {
		new = new.Elem()
	}
	if new.Kind() != old.Kind() {
		ctx.Log(diag.Info, fmt.Sprintf("diffWalk: visiting %s: Kind changed: old: %s new: %s", path, old.Kind().String(), new.Kind().String()))
		diff[path] = p.PropertyDiff{Kind: p.Update}
		return
	}
	switch old.Kind() {
	case reflect.Array, reflect.Slice:
		mv := min(old.Len(), new.Len())
		if old.Len() != new.Len() {
			diff[fmt.Sprintf("%s[%d]", path, mv+1)] = p.PropertyDiff{Kind: p.Update}
		}
		for i := 0; i < mv; i++ {
			diffWalk(ctx, diff, fmt.Sprintf("%s[%d]", path, i), old.Index(i), new.Index(i))
		}
	case reflect.Map:
		if old.Len() != new.Len() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
		for _, k := range old.MapKeys() {
			diffWalk(ctx, diff, fmt.Sprintf("%s.%s", path, k.String()), old.MapIndex(k), new.MapIndex(k))
		}
	case reflect.String:
		if old.String() != new.String() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	case reflect.Int:
		if old.Int() != new.Int() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	case reflect.Bool:
		if old.Bool() != new.Bool() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	case reflect.Float64:
		if old.Float() != new.Float() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	default:
		// handle other types
		diff[path] = p.PropertyDiff{Kind: p.Update}
		ctx.Log(diag.Warning, fmt.Sprintf("Unhandled types comparing %s: %s<>%s, %s != %s", path, old.Kind().String(), new.Kind().String(), old.String(), new.String()))
	}
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
