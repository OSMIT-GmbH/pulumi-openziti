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

type Provider struct {
	pulumi.ProviderResourceState

	// The password. It is very secret.
	Password pulumi.StringOutput `pulumi:"password"`
	// The URI to the API
	Uri pulumi.StringOutput `pulumi:"uri"`
	// The username. It's important but not secret.
	User    pulumi.StringOutput    `pulumi:"user"`
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Uri == nil {
		return nil, errors.New("invalid value for required argument 'Uri'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:openziti", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	Insecure *bool `pulumi:"insecure"`
	// The password. It is very secret.
	Password string `pulumi:"password"`
	// The URI to the API
	Uri string `pulumi:"uri"`
	// The username. It's important but not secret.
	User    string  `pulumi:"user"`
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	Insecure pulumi.BoolPtrInput
	// The password. It is very secret.
	Password pulumi.StringInput
	// The URI to the API
	Uri pulumi.StringInput
	// The username. It's important but not secret.
	User    pulumi.StringInput
	Version pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: i.ToProviderOutputWithContext(ctx).OutputState,
	}
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: o.OutputState,
	}
}

// The password. It is very secret.
func (o ProviderOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The URI to the API
func (o ProviderOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

// The username. It's important but not secret.
func (o ProviderOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

func (o ProviderOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
