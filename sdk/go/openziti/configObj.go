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

type ConfigObj struct {
	pulumi.CustomResourceState

	_assimilated   pulumi.BoolOutput   `pulumi:"_assimilated"`
	_links         LinkMapOutput       `pulumi:"_links"`
	ConfigType     EntityRefOutput     `pulumi:"configType"`
	ConfigTypeId   pulumi.StringOutput `pulumi:"configTypeId"`
	ConfigTypeName pulumi.StringOutput `pulumi:"configTypeName"`
	CreatedAt      pulumi.StringOutput `pulumi:"createdAt"`
	Data           pulumi.AnyOutput    `pulumi:"data"`
	Id             pulumi.StringOutput `pulumi:"id"`
	Name           pulumi.StringOutput `pulumi:"name"`
	Tags           pulumi.MapOutput    `pulumi:"tags"`
	UpdatedAt      pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewConfigObj registers a new resource with the given unique name, arguments, and options.
func NewConfigObj(ctx *pulumi.Context,
	name string, args *ConfigObjArgs, opts ...pulumi.ResourceOption) (*ConfigObj, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigTypeName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigTypeName'")
	}
	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigObj
	err := ctx.RegisterResource("openziti:index:ConfigObj", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigObj gets an existing ConfigObj resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigObj(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigObjState, opts ...pulumi.ResourceOption) (*ConfigObj, error) {
	var resource ConfigObj
	err := ctx.ReadResource("openziti:index:ConfigObj", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigObj resources.
type configObjState struct {
}

type ConfigObjState struct {
}

func (ConfigObjState) ElementType() reflect.Type {
	return reflect.TypeOf((*configObjState)(nil)).Elem()
}

type configObjArgs struct {
	ConfigTypeName string                 `pulumi:"configTypeName"`
	Data           interface{}            `pulumi:"data"`
	Name           string                 `pulumi:"name"`
	Tags           map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigObj resource.
type ConfigObjArgs struct {
	ConfigTypeName pulumi.StringInput
	Data           pulumi.Input
	Name           pulumi.StringInput
	Tags           pulumi.MapInput
}

func (ConfigObjArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configObjArgs)(nil)).Elem()
}

type ConfigObjInput interface {
	pulumi.Input

	ToConfigObjOutput() ConfigObjOutput
	ToConfigObjOutputWithContext(ctx context.Context) ConfigObjOutput
}

func (*ConfigObj) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigObj)(nil)).Elem()
}

func (i *ConfigObj) ToConfigObjOutput() ConfigObjOutput {
	return i.ToConfigObjOutputWithContext(context.Background())
}

func (i *ConfigObj) ToConfigObjOutputWithContext(ctx context.Context) ConfigObjOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigObjOutput)
}

func (i *ConfigObj) ToOutput(ctx context.Context) pulumix.Output[*ConfigObj] {
	return pulumix.Output[*ConfigObj]{
		OutputState: i.ToConfigObjOutputWithContext(ctx).OutputState,
	}
}

type ConfigObjOutput struct{ *pulumi.OutputState }

func (ConfigObjOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigObj)(nil)).Elem()
}

func (o ConfigObjOutput) ToConfigObjOutput() ConfigObjOutput {
	return o
}

func (o ConfigObjOutput) ToConfigObjOutputWithContext(ctx context.Context) ConfigObjOutput {
	return o
}

func (o ConfigObjOutput) ToOutput(ctx context.Context) pulumix.Output[*ConfigObj] {
	return pulumix.Output[*ConfigObj]{
		OutputState: o.OutputState,
	}
}

func (o ConfigObjOutput) _assimilated() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.BoolOutput { return v._assimilated }).(pulumi.BoolOutput)
}

func (o ConfigObjOutput) _links() LinkMapOutput {
	return o.ApplyT(func(v *ConfigObj) LinkMapOutput { return v._links }).(LinkMapOutput)
}

func (o ConfigObjOutput) ConfigType() EntityRefOutput {
	return o.ApplyT(func(v *ConfigObj) EntityRefOutput { return v.ConfigType }).(EntityRefOutput)
}

func (o ConfigObjOutput) ConfigTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.StringOutput { return v.ConfigTypeId }).(pulumi.StringOutput)
}

func (o ConfigObjOutput) ConfigTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.StringOutput { return v.ConfigTypeName }).(pulumi.StringOutput)
}

func (o ConfigObjOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ConfigObjOutput) Data() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.AnyOutput { return v.Data }).(pulumi.AnyOutput)
}

func (o ConfigObjOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.StringOutput { return v.Id }).(pulumi.StringOutput)
}

func (o ConfigObjOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigObjOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

func (o ConfigObjOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigObj) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigObjInput)(nil)).Elem(), &ConfigObj{})
	pulumi.RegisterOutputType(ConfigObjOutput{})
}
