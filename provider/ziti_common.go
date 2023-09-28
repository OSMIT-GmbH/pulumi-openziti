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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	edgeapis "github.com/openziti/edge-api/rest_management_api_client"
	"github.com/openziti/edge-api/rest_management_api_client/config"
	"github.com/openziti/edge-api/rest_model"
	"github.com/openziti/edge-api/rest_util"
	"github.com/openziti/sdk-golang/ziti"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type OpenZitiProviderConfig struct {
	User     string `pulumi:"user"`
	Password string `pulumi:"password" provider:"secret"`
	Uri      string `pulumi:"uri"`
	// TODO finalize insecure!
	// Insecure   string `pulumi:"insecure,optional"`
	// insecure   bool
	// I'm not sure what's wrong with boolean - see following error:
	// error: pulumi:providers:openziti resource 'openziti-provider': property assimilate value {false} has a problem: Field 'assimilate' on 'provider.OpenZitiProviderConfig' must be a 'bool'; got 'string' instead
	Assimilate        string `pulumi:"assimilate,optional"`
	assimilate        bool
	DeleteAssimilated string `pulumi:"deleteAssimilated,optional"`
	deleteAssimilated bool
	Version           string `pulumi:"version,optional"` // version seems to be provided automatically
	cacheKey          string
}

var _ = (infer.Annotated)((*OpenZitiProviderConfig)(nil))

func (c *OpenZitiProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.User, "The username. It's important but not secret.")
	a.Describe(&c.Password, "The password. It is very secret.")
	a.Describe(&c.Uri, `The URI to the API`)
	a.Describe(&c.Assimilate, `Assimilate an existing object during create`)
	a.Describe(&c.DeleteAssimilated, `Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)`)
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
	c.assimilate = strings.EqualFold(c.Assimilate, "true")
	c.deleteAssimilated = strings.EqualFold(c.DeleteAssimilated, "true")
	// c.insecure = strings.EqualFold(c.Insecure, "true")

	//ctx.Log(diag.Info, msg)
	return nil
}

// TODO https://github.com/pulumi/pulumi-go-provider/issues/121
//var _ = (infer.CustomDiff[*OpenZitiProviderConfig, *OpenZitiProviderConfig])((*OpenZitiProviderConfig)(nil))
//
//func (*OpenZitiProviderConfig) Diff(ctx p.Context, id string, olds *OpenZitiProviderConfig, news *OpenZitiProviderConfig) (p.DiffResponse, error) {
//	fmt.Printf("Config Diff called: %v => %v", olds, news)
//	return p.DiffResponse{
//		DeleteBeforeReplace: true,
//		HasChanges:          false,
//		DetailedDiff:        nil,
//	}, nil
//}

var IdPreviewPrefix = "~~preview~~"

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

	// this config element was assimilated...
	// Required: true
	Assimilated bool `pulumi:"_assimilated"`
}

func buildLinks(src rest_model.Links) Links {
	ret := make(Links)
	for key, link := range src {
		ret[key] = Link{Comment: link.Comment, Method: link.Method, Href: link.Href.String()}
	}
	return ret
}

func buildBaseState(src rest_model.BaseEntity, assimilated bool) BaseStateEntity {
	tags := buildTags(*src.Tags)
	return BaseStateEntity{Links: buildLinks(src.Links),
		Assimilated: assimilated,
		ID:          *src.ID,
		CreatedAt:   src.CreatedAt.String(),
		Tags:        tags,
		UpdatedAt:   src.UpdatedAt.String(),
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

func buildZitiTags(tags Tags) *rest_model.Tags {
	//if len(tags.SubTags) == 0 {
	//	return rest_model.Tags{}
	//}
	out := make(rest_model.SubTags)
	for key, value := range tags {
		out[key] = value
	}
	return &rest_model.Tags{SubTags: out}
}

func dumpStruct(ctx p.Context, name string, data interface{}) {
	out, err := json.Marshal(data)
	if err != nil {
		ctx.Logf(diag.Info, "ERROR: failed serializing data: %s => %#v", err.Error(), data)
	} else {
		ctx.Logf(diag.Info, "OK: %s(json)=%s", name, string(out))
	}
}

type CacheEntry struct {
	//apiSession *rest_model.CurrentAPISessionDetail
	client            *edgeapis.ZitiEdgeManagement
	configTypesMutex  sync.Mutex
	configTypes       map[string]string
	identitiesMutex   sync.Mutex
	identities        map[string]string
	identitiesReverse map[string]string
}

var cache = make(map[string]*CacheEntry)

func getConfigTypeId(cache *CacheEntry, name string) (string, error) {
	if len(cache.configTypes) == 0 {
		cache.configTypesMutex.Lock()
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
				cache.configTypesMutex.Unlock()
				return "", err
			}

			configTypes := make(map[string]string)
			for _, entity := range ctResp.GetPayload().Data {
				// wrapper := api.Wrap(entity)
				// configTypes[wrapper.String("name")] = wrapper.String("id")
				configTypes[*entity.Name] = *entity.ID
			}
			cache.configTypes = configTypes
		}
		cache.configTypesMutex.Unlock()
		// fmt.Printf("retrived configValues: %+v\n", cache.configTypes)
	}
	typeId := cache.configTypes[name]
	if typeId == "" {
		return typeId, errors.New(fmt.Sprintf("No configType declared for `%s` - known types: %+v", name, cache.configTypes))
	}
	return typeId, nil
}

var clientMutex sync.Mutex = sync.Mutex{}

func initClient(ctx p.Context) (*CacheEntry, OpenZitiProviderConfig, error) {
	c := infer.GetConfig[OpenZitiProviderConfig](ctx)
	ce, ok := cache[c.cacheKey]
	if !ok {
		// new entry - use mutex to limit to one session

		clientMutex.Lock()
		ce, ok = cache[c.cacheKey]
		if !ok {

			// creds := edge_apis.New([]*x509.Certificate{testIdCerts.cert}, testIdCerts.key)
			caPool, err := ziti.GetControllerWellKnownCaPool(c.Uri)
			if err != nil {
				clientMutex.Unlock()
				return nil, c, err
			}

			client, err := rest_util.NewEdgeManagementClientWithUpdb(c.User, c.Password, c.Uri, caPool)
			if err != nil {
				clientMutex.Unlock()
				return nil, c, err
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
			ce = &CacheEntry{
				client: client,
			}
			cache[c.cacheKey] = ce
		}
		clientMutex.Unlock()
	}
	return ce, c, nil
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
func formatApiErrDupeCheck(ctx p.Context, err error, apiError *rest_model.APIErrorEnvelope) (error, bool) {
	errRet := formatApiErr(ctx, err, apiError)
	match, _ := regexp.MatchString(" Payload=\\{\"cause\":\\{\"field\":\"name\",\"reason\":\"duplicate value '[^\"']+' in unique index on identities store\",\"value\":\"[^\"']+\"},", errRet.Error())
	if !match {
		match, _ = regexp.MatchString(" Payload={\"cause\":{\"field\":\"name\",\"reason\":\"name is must be unique\",\"value\":\"[^\"']+\"},\"code\":\"COULD_NOT_VALIDATE\",", errRet.Error())
	}
	return errRet, match
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

func diffStrArrayIgnoreOrder(ctx p.Context, diff map[string]p.PropertyDiff, path string, old []string, new []string) {
	oldLen := len(old)
	newLen := len(new)
	minLen := min(oldLen, newLen)
	if oldLen > newLen {
		//noop
		diff[fmt.Sprintf("%s[%d]", path, minLen+1)] = p.PropertyDiff{Kind: p.Add}
	}
	if oldLen < newLen {
		diff[fmt.Sprintf("%s[%d]", path, minLen+1)] = p.PropertyDiff{Kind: p.Delete}
	}
	sort.Strings(old)
	sort.Strings(new)
	for i := 0; i < minLen; i++ {
		if old[i] != new[i] {
			diff[fmt.Sprintf("%s[%d]", path, i)] = p.PropertyDiff{Kind: p.Update}
		}
	}
}

func buildNameFilter(name string) *string {
	filter := "name=\"" + url.QueryEscape(name) + "\""
	return &filter
}

func ifte[T interface{}](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}

func iftfe[T interface{}](cond bool, trueFunc func() T, falseVal T) T {
	if cond {
		return trueFunc()
	} else {
		return falseVal
	}
}

func ifted[T interface{}](cond bool, trueVal T, falseVal T) *T {
	if cond {
		return &trueVal
	} else {
		return &falseVal
	}
}
func iftden[T interface{}](cond bool, trueVal T) *T {
	if cond {
		return &trueVal
	} else {
		return nil
	}
}

func iftfden[T interface{}](cond bool, trueFunc func() T) *T {
	if cond {
		ret := trueFunc()
		return &ret
	} else {
		return nil
	}
}
