// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openziti

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"internal"
)

type EdgeRouter struct {
	pulumi.CustomResourceState

	_links                LinkMapOutput            `pulumi:"_links"`
	AppData               pulumi.MapOutput         `pulumi:"appData"`
	CertPem               pulumi.StringPtrOutput   `pulumi:"certPem"`
	Cost                  pulumi.IntOutput         `pulumi:"cost"`
	CreatedAt             pulumi.StringOutput      `pulumi:"createdAt"`
	Disabled              pulumi.BoolOutput        `pulumi:"disabled"`
	EnrollmentCreatedAt   pulumi.StringPtrOutput   `pulumi:"enrollmentCreatedAt"`
	EnrollmentExpiresAt   pulumi.StringPtrOutput   `pulumi:"enrollmentExpiresAt"`
	EnrollmentJwt         pulumi.StringPtrOutput   `pulumi:"enrollmentJwt"`
	EnrollmentToken       pulumi.StringPtrOutput   `pulumi:"enrollmentToken"`
	Fingerprint           pulumi.StringPtrOutput   `pulumi:"fingerprint"`
	Hostname              pulumi.StringOutput      `pulumi:"hostname"`
	Id                    pulumi.StringOutput      `pulumi:"id"`
	IsOnline              pulumi.BoolOutput        `pulumi:"isOnline"`
	IsTunnelerEnabled     pulumi.BoolOutput        `pulumi:"isTunnelerEnabled"`
	IsVerified            pulumi.BoolOutput        `pulumi:"isVerified"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	NoTraversal           pulumi.BoolOutput        `pulumi:"noTraversal"`
	RoleAttributes        pulumi.StringArrayOutput `pulumi:"roleAttributes"`
	SupportedProtocols    pulumi.StringMapOutput   `pulumi:"supportedProtocols"`
	SyncStatus            pulumi.StringOutput      `pulumi:"syncStatus"`
	Tags                  pulumi.MapOutput         `pulumi:"tags"`
	UnverifiedCertPem     pulumi.StringPtrOutput   `pulumi:"unverifiedCertPem"`
	UnverifiedFingerprint pulumi.StringPtrOutput   `pulumi:"unverifiedFingerprint"`
	UpdatedAt             pulumi.StringOutput      `pulumi:"updatedAt"`
}

// NewEdgeRouter registers a new resource with the given unique name, arguments, and options.
func NewEdgeRouter(ctx *pulumi.Context,
	name string, args *EdgeRouterArgs, opts ...pulumi.ResourceOption) (*EdgeRouter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EdgeRouter
	err := ctx.RegisterResource("openziti:index:EdgeRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeRouter gets an existing EdgeRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeRouterState, opts ...pulumi.ResourceOption) (*EdgeRouter, error) {
	var resource EdgeRouter
	err := ctx.ReadResource("openziti:index:EdgeRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeRouter resources.
type edgeRouterState struct {
}

type EdgeRouterState struct {
}

func (EdgeRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeRouterState)(nil)).Elem()
}

type edgeRouterArgs struct {
	AppData           map[string]interface{} `pulumi:"appData"`
	Cost              *int                   `pulumi:"cost"`
	Disabled          *bool                  `pulumi:"disabled"`
	IsTunnelerEnabled *bool                  `pulumi:"isTunnelerEnabled"`
	Name              string                 `pulumi:"name"`
	NoTraversal       *bool                  `pulumi:"noTraversal"`
	RoleAttributes    []string               `pulumi:"roleAttributes"`
	Tags              map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EdgeRouter resource.
type EdgeRouterArgs struct {
	AppData           pulumi.MapInput
	Cost              pulumi.IntPtrInput
	Disabled          pulumi.BoolPtrInput
	IsTunnelerEnabled pulumi.BoolPtrInput
	Name              pulumi.StringInput
	NoTraversal       pulumi.BoolPtrInput
	RoleAttributes    pulumi.StringArrayInput
	Tags              pulumi.MapInput
}

func (EdgeRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeRouterArgs)(nil)).Elem()
}

type EdgeRouterInput interface {
	pulumi.Input

	ToEdgeRouterOutput() EdgeRouterOutput
	ToEdgeRouterOutputWithContext(ctx context.Context) EdgeRouterOutput
}

func (*EdgeRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeRouter)(nil)).Elem()
}

func (i *EdgeRouter) ToEdgeRouterOutput() EdgeRouterOutput {
	return i.ToEdgeRouterOutputWithContext(context.Background())
}

func (i *EdgeRouter) ToEdgeRouterOutputWithContext(ctx context.Context) EdgeRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeRouterOutput)
}

func (i *EdgeRouter) ToOutput(ctx context.Context) pulumix.Output[*EdgeRouter] {
	return pulumix.Output[*EdgeRouter]{
		OutputState: i.ToEdgeRouterOutputWithContext(ctx).OutputState,
	}
}

type EdgeRouterOutput struct{ *pulumi.OutputState }

func (EdgeRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeRouter)(nil)).Elem()
}

func (o EdgeRouterOutput) ToEdgeRouterOutput() EdgeRouterOutput {
	return o
}

func (o EdgeRouterOutput) ToEdgeRouterOutputWithContext(ctx context.Context) EdgeRouterOutput {
	return o
}

func (o EdgeRouterOutput) ToOutput(ctx context.Context) pulumix.Output[*EdgeRouter] {
	return pulumix.Output[*EdgeRouter]{
		OutputState: o.OutputState,
	}
}

func (o EdgeRouterOutput) _links() LinkMapOutput {
	return o.ApplyT(func(v *EdgeRouter) LinkMapOutput { return v._links }).(LinkMapOutput)
}

func (o EdgeRouterOutput) AppData() pulumi.MapOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.MapOutput { return v.AppData }).(pulumi.MapOutput)
}

func (o EdgeRouterOutput) CertPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.CertPem }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.IntOutput { return v.Cost }).(pulumi.IntOutput)
}

func (o EdgeRouterOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o EdgeRouterOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

func (o EdgeRouterOutput) EnrollmentCreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.EnrollmentCreatedAt }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) EnrollmentExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.EnrollmentExpiresAt }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) EnrollmentJwt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.EnrollmentJwt }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) EnrollmentToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.EnrollmentToken }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) Fingerprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.Fingerprint }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o EdgeRouterOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringOutput { return v.Id }).(pulumi.StringOutput)
}

func (o EdgeRouterOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

func (o EdgeRouterOutput) IsTunnelerEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.BoolOutput { return v.IsTunnelerEnabled }).(pulumi.BoolOutput)
}

func (o EdgeRouterOutput) IsVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.BoolOutput { return v.IsVerified }).(pulumi.BoolOutput)
}

func (o EdgeRouterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EdgeRouterOutput) NoTraversal() pulumi.BoolOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.BoolOutput { return v.NoTraversal }).(pulumi.BoolOutput)
}

func (o EdgeRouterOutput) RoleAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringArrayOutput { return v.RoleAttributes }).(pulumi.StringArrayOutput)
}

func (o EdgeRouterOutput) SupportedProtocols() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringMapOutput { return v.SupportedProtocols }).(pulumi.StringMapOutput)
}

func (o EdgeRouterOutput) SyncStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringOutput { return v.SyncStatus }).(pulumi.StringOutput)
}

func (o EdgeRouterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

func (o EdgeRouterOutput) UnverifiedCertPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.UnverifiedCertPem }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) UnverifiedFingerprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringPtrOutput { return v.UnverifiedFingerprint }).(pulumi.StringPtrOutput)
}

func (o EdgeRouterOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeRouter) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgeRouterInput)(nil)).Elem(), &EdgeRouter{})
	pulumi.RegisterOutputType(EdgeRouterOutput{})
}
