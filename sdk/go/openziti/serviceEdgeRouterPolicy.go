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

type ServiceEdgeRouterPolicy struct {
	pulumi.CustomResourceState

	_assimilated           pulumi.BoolOutput        `pulumi:"_assimilated"`
	_links                 LinkMapOutput            `pulumi:"_links"`
	CreatedAt              pulumi.StringOutput      `pulumi:"createdAt"`
	EdgeRouterRoles        pulumi.StringArrayOutput `pulumi:"edgeRouterRoles"`
	EdgeRouterRolesDisplay NamedRoleArrayOutput     `pulumi:"edgeRouterRolesDisplay"`
	Id                     pulumi.StringOutput      `pulumi:"id"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	Semantic               pulumi.StringOutput      `pulumi:"semantic"`
	ServiceRoles           pulumi.StringArrayOutput `pulumi:"serviceRoles"`
	ServiceRolesDisplay    NamedRoleArrayOutput     `pulumi:"serviceRolesDisplay"`
	Tags                   pulumi.MapOutput         `pulumi:"tags"`
	UpdatedAt              pulumi.StringOutput      `pulumi:"updatedAt"`
}

// NewServiceEdgeRouterPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceEdgeRouterPolicy(ctx *pulumi.Context,
	name string, args *ServiceEdgeRouterPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceEdgeRouterPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Semantic == nil {
		return nil, errors.New("invalid value for required argument 'Semantic'")
	}
	if args.ServiceRoles == nil {
		return nil, errors.New("invalid value for required argument 'ServiceRoles'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceEdgeRouterPolicy
	err := ctx.RegisterResource("openziti:index:ServiceEdgeRouterPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEdgeRouterPolicy gets an existing ServiceEdgeRouterPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEdgeRouterPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEdgeRouterPolicyState, opts ...pulumi.ResourceOption) (*ServiceEdgeRouterPolicy, error) {
	var resource ServiceEdgeRouterPolicy
	err := ctx.ReadResource("openziti:index:ServiceEdgeRouterPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEdgeRouterPolicy resources.
type serviceEdgeRouterPolicyState struct {
}

type ServiceEdgeRouterPolicyState struct {
}

func (ServiceEdgeRouterPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeRouterPolicyState)(nil)).Elem()
}

type serviceEdgeRouterPolicyArgs struct {
	EdgeRouterRoles []string               `pulumi:"edgeRouterRoles"`
	Name            string                 `pulumi:"name"`
	Semantic        string                 `pulumi:"semantic"`
	ServiceRoles    []string               `pulumi:"serviceRoles"`
	Tags            map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceEdgeRouterPolicy resource.
type ServiceEdgeRouterPolicyArgs struct {
	EdgeRouterRoles pulumi.StringArrayInput
	Name            pulumi.StringInput
	Semantic        pulumi.StringInput
	ServiceRoles    pulumi.StringArrayInput
	Tags            pulumi.MapInput
}

func (ServiceEdgeRouterPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEdgeRouterPolicyArgs)(nil)).Elem()
}

type ServiceEdgeRouterPolicyInput interface {
	pulumi.Input

	ToServiceEdgeRouterPolicyOutput() ServiceEdgeRouterPolicyOutput
	ToServiceEdgeRouterPolicyOutputWithContext(ctx context.Context) ServiceEdgeRouterPolicyOutput
}

func (*ServiceEdgeRouterPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEdgeRouterPolicy)(nil)).Elem()
}

func (i *ServiceEdgeRouterPolicy) ToServiceEdgeRouterPolicyOutput() ServiceEdgeRouterPolicyOutput {
	return i.ToServiceEdgeRouterPolicyOutputWithContext(context.Background())
}

func (i *ServiceEdgeRouterPolicy) ToServiceEdgeRouterPolicyOutputWithContext(ctx context.Context) ServiceEdgeRouterPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEdgeRouterPolicyOutput)
}

func (i *ServiceEdgeRouterPolicy) ToOutput(ctx context.Context) pulumix.Output[*ServiceEdgeRouterPolicy] {
	return pulumix.Output[*ServiceEdgeRouterPolicy]{
		OutputState: i.ToServiceEdgeRouterPolicyOutputWithContext(ctx).OutputState,
	}
}

type ServiceEdgeRouterPolicyOutput struct{ *pulumi.OutputState }

func (ServiceEdgeRouterPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEdgeRouterPolicy)(nil)).Elem()
}

func (o ServiceEdgeRouterPolicyOutput) ToServiceEdgeRouterPolicyOutput() ServiceEdgeRouterPolicyOutput {
	return o
}

func (o ServiceEdgeRouterPolicyOutput) ToServiceEdgeRouterPolicyOutputWithContext(ctx context.Context) ServiceEdgeRouterPolicyOutput {
	return o
}

func (o ServiceEdgeRouterPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceEdgeRouterPolicy] {
	return pulumix.Output[*ServiceEdgeRouterPolicy]{
		OutputState: o.OutputState,
	}
}

func (o ServiceEdgeRouterPolicyOutput) _assimilated() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.BoolOutput { return v._assimilated }).(pulumi.BoolOutput)
}

func (o ServiceEdgeRouterPolicyOutput) _links() LinkMapOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) LinkMapOutput { return v._links }).(LinkMapOutput)
}

func (o ServiceEdgeRouterPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ServiceEdgeRouterPolicyOutput) EdgeRouterRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringArrayOutput { return v.EdgeRouterRoles }).(pulumi.StringArrayOutput)
}

func (o ServiceEdgeRouterPolicyOutput) EdgeRouterRolesDisplay() NamedRoleArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) NamedRoleArrayOutput { return v.EdgeRouterRolesDisplay }).(NamedRoleArrayOutput)
}

func (o ServiceEdgeRouterPolicyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringOutput { return v.Id }).(pulumi.StringOutput)
}

func (o ServiceEdgeRouterPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceEdgeRouterPolicyOutput) Semantic() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringOutput { return v.Semantic }).(pulumi.StringOutput)
}

func (o ServiceEdgeRouterPolicyOutput) ServiceRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringArrayOutput { return v.ServiceRoles }).(pulumi.StringArrayOutput)
}

func (o ServiceEdgeRouterPolicyOutput) ServiceRolesDisplay() NamedRoleArrayOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) NamedRoleArrayOutput { return v.ServiceRolesDisplay }).(NamedRoleArrayOutput)
}

func (o ServiceEdgeRouterPolicyOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

func (o ServiceEdgeRouterPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEdgeRouterPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceEdgeRouterPolicyInput)(nil)).Elem(), &ServiceEdgeRouterPolicy{})
	pulumi.RegisterOutputType(ServiceEdgeRouterPolicyOutput{})
}
