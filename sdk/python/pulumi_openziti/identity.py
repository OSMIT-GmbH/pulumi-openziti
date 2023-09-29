# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from . import rest_model as _rest_model
from ._inputs import *

__all__ = ['IdentityArgs', 'Identity']

@pulumi.input_type
class IdentityArgs:
    def __init__(__self__, *,
                 is_admin: pulumi.Input[bool],
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 auth_policy_id: Optional[pulumi.Input[str]] = None,
                 default_hosting_cost: Optional[pulumi.Input[int]] = None,
                 default_hosting_precedence: Optional[pulumi.Input[str]] = None,
                 enrollment: Optional[pulumi.Input['IdentityCreateEnrollmentArgs']] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_hosting_costs: Optional[pulumi.Input[Mapping[str, pulumi.Input[int]]]] = None,
                 service_hosting_precedences: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Identity resource.
        """
        IdentityArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            is_admin=is_admin,
            name=name,
            type=type,
            app_data=app_data,
            auth_policy_id=auth_policy_id,
            default_hosting_cost=default_hosting_cost,
            default_hosting_precedence=default_hosting_precedence,
            enrollment=enrollment,
            external_id=external_id,
            role_attributes=role_attributes,
            service_hosting_costs=service_hosting_costs,
            service_hosting_precedences=service_hosting_precedences,
            tags=tags,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             is_admin: pulumi.Input[bool],
             name: pulumi.Input[str],
             type: pulumi.Input[str],
             app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             auth_policy_id: Optional[pulumi.Input[str]] = None,
             default_hosting_cost: Optional[pulumi.Input[int]] = None,
             default_hosting_precedence: Optional[pulumi.Input[str]] = None,
             enrollment: Optional[pulumi.Input['IdentityCreateEnrollmentArgs']] = None,
             external_id: Optional[pulumi.Input[str]] = None,
             role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             service_hosting_costs: Optional[pulumi.Input[Mapping[str, pulumi.Input[int]]]] = None,
             service_hosting_precedences: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("is_admin", is_admin)
        _setter("name", name)
        _setter("type", type)
        if app_data is not None:
            _setter("app_data", app_data)
        if auth_policy_id is not None:
            _setter("auth_policy_id", auth_policy_id)
        if default_hosting_cost is not None:
            _setter("default_hosting_cost", default_hosting_cost)
        if default_hosting_precedence is not None:
            _setter("default_hosting_precedence", default_hosting_precedence)
        if enrollment is not None:
            _setter("enrollment", enrollment)
        if external_id is not None:
            _setter("external_id", external_id)
        if role_attributes is not None:
            _setter("role_attributes", role_attributes)
        if service_hosting_costs is not None:
            _setter("service_hosting_costs", service_hosting_costs)
        if service_hosting_precedences is not None:
            _setter("service_hosting_precedences", service_hosting_precedences)
        if tags is not None:
            _setter("tags", tags)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="appData")
    def app_data(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "app_data")

    @app_data.setter
    def app_data(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "app_data", value)

    @property
    @pulumi.getter(name="authPolicyId")
    def auth_policy_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auth_policy_id")

    @auth_policy_id.setter
    def auth_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_policy_id", value)

    @property
    @pulumi.getter(name="defaultHostingCost")
    def default_hosting_cost(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_hosting_cost")

    @default_hosting_cost.setter
    def default_hosting_cost(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_hosting_cost", value)

    @property
    @pulumi.getter(name="defaultHostingPrecedence")
    def default_hosting_precedence(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_hosting_precedence")

    @default_hosting_precedence.setter
    def default_hosting_precedence(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_hosting_precedence", value)

    @property
    @pulumi.getter
    def enrollment(self) -> Optional[pulumi.Input['IdentityCreateEnrollmentArgs']]:
        return pulumi.get(self, "enrollment")

    @enrollment.setter
    def enrollment(self, value: Optional[pulumi.Input['IdentityCreateEnrollmentArgs']]):
        pulumi.set(self, "enrollment", value)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter(name="roleAttributes")
    def role_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "role_attributes")

    @role_attributes.setter
    def role_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_attributes", value)

    @property
    @pulumi.getter(name="serviceHostingCosts")
    def service_hosting_costs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[int]]]]:
        return pulumi.get(self, "service_hosting_costs")

    @service_hosting_costs.setter
    def service_hosting_costs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[int]]]]):
        pulumi.set(self, "service_hosting_costs", value)

    @property
    @pulumi.getter(name="serviceHostingPrecedences")
    def service_hosting_precedences(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "service_hosting_precedences")

    @service_hosting_precedences.setter
    def service_hosting_precedences(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "service_hosting_precedences", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Identity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 auth_policy_id: Optional[pulumi.Input[str]] = None,
                 default_hosting_cost: Optional[pulumi.Input[int]] = None,
                 default_hosting_precedence: Optional[pulumi.Input[str]] = None,
                 enrollment: Optional[pulumi.Input[pulumi.InputType['IdentityCreateEnrollmentArgs']]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_hosting_costs: Optional[pulumi.Input[Mapping[str, pulumi.Input[int]]]] = None,
                 service_hosting_precedences: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Identity resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Identity resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IdentityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IdentityArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 auth_policy_id: Optional[pulumi.Input[str]] = None,
                 default_hosting_cost: Optional[pulumi.Input[int]] = None,
                 default_hosting_precedence: Optional[pulumi.Input[str]] = None,
                 enrollment: Optional[pulumi.Input[pulumi.InputType['IdentityCreateEnrollmentArgs']]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_hosting_costs: Optional[pulumi.Input[Mapping[str, pulumi.Input[int]]]] = None,
                 service_hosting_precedences: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityArgs.__new__(IdentityArgs)

            __props__.__dict__["app_data"] = app_data
            __props__.__dict__["auth_policy_id"] = auth_policy_id
            __props__.__dict__["default_hosting_cost"] = default_hosting_cost
            __props__.__dict__["default_hosting_precedence"] = default_hosting_precedence
            if enrollment is not None and not isinstance(enrollment, IdentityCreateEnrollmentArgs):
                enrollment = enrollment or {}
                def _setter(key, value):
                    enrollment[key] = value
                IdentityCreateEnrollmentArgs._configure(_setter, **enrollment)
            __props__.__dict__["enrollment"] = enrollment
            __props__.__dict__["external_id"] = external_id
            if is_admin is None and not opts.urn:
                raise TypeError("Missing required property 'is_admin'")
            __props__.__dict__["is_admin"] = is_admin
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["role_attributes"] = role_attributes
            __props__.__dict__["service_hosting_costs"] = service_hosting_costs
            __props__.__dict__["service_hosting_precedences"] = service_hosting_precedences
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["_assimilated"] = None
            __props__.__dict__["_links"] = None
            __props__.__dict__["auth_policy"] = None
            __props__.__dict__["authenticators"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["disabled"] = None
            __props__.__dict__["disabled_at"] = None
            __props__.__dict__["disabled_until"] = None
            __props__.__dict__["env_info"] = None
            __props__.__dict__["has_api_session"] = None
            __props__.__dict__["has_edge_router_connection"] = None
            __props__.__dict__["id"] = None
            __props__.__dict__["is_default_admin"] = None
            __props__.__dict__["is_mfa_enabled"] = None
            __props__.__dict__["sdk_info"] = None
            __props__.__dict__["type_id"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["enrollment"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Identity, __self__).__init__(
            'openziti:index:Identity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Identity':
        """
        Get an existing Identity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IdentityArgs.__new__(IdentityArgs)

        __props__.__dict__["_assimilated"] = None
        __props__.__dict__["_links"] = None
        __props__.__dict__["app_data"] = None
        __props__.__dict__["auth_policy"] = None
        __props__.__dict__["auth_policy_id"] = None
        __props__.__dict__["authenticators"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["default_hosting_cost"] = None
        __props__.__dict__["default_hosting_precedence"] = None
        __props__.__dict__["disabled"] = None
        __props__.__dict__["disabled_at"] = None
        __props__.__dict__["disabled_until"] = None
        __props__.__dict__["enrollment"] = None
        __props__.__dict__["env_info"] = None
        __props__.__dict__["external_id"] = None
        __props__.__dict__["has_api_session"] = None
        __props__.__dict__["has_edge_router_connection"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["is_admin"] = None
        __props__.__dict__["is_default_admin"] = None
        __props__.__dict__["is_mfa_enabled"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["role_attributes"] = None
        __props__.__dict__["sdk_info"] = None
        __props__.__dict__["service_hosting_costs"] = None
        __props__.__dict__["service_hosting_precedences"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["type_id"] = None
        __props__.__dict__["updated_at"] = None
        return Identity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def _assimilated(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "_assimilated")

    @property
    @pulumi.getter
    def _links(self) -> pulumi.Output[Mapping[str, 'outputs.Link']]:
        return pulumi.get(self, "_links")

    @property
    @pulumi.getter(name="appData")
    def app_data(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "app_data")

    @property
    @pulumi.getter(name="authPolicy")
    def auth_policy(self) -> pulumi.Output['outputs.EntityRef']:
        return pulumi.get(self, "auth_policy")

    @property
    @pulumi.getter(name="authPolicyId")
    def auth_policy_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auth_policy_id")

    @property
    @pulumi.getter
    def authenticators(self) -> pulumi.Output['_rest_model.outputs.IdentityAuthenticators']:
        return pulumi.get(self, "authenticators")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultHostingCost")
    def default_hosting_cost(self) -> pulumi.Output[int]:
        return pulumi.get(self, "default_hosting_cost")

    @property
    @pulumi.getter(name="defaultHostingPrecedence")
    def default_hosting_precedence(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "default_hosting_precedence")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="disabledAt")
    def disabled_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "disabled_at")

    @property
    @pulumi.getter(name="disabledUntil")
    def disabled_until(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "disabled_until")

    @property
    @pulumi.getter
    def enrollment(self) -> pulumi.Output['outputs.IdentityEnrollments']:
        return pulumi.get(self, "enrollment")

    @property
    @pulumi.getter(name="envInfo")
    def env_info(self) -> pulumi.Output['outputs.EnvInfo']:
        return pulumi.get(self, "env_info")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="hasApiSession")
    def has_api_session(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "has_api_session")

    @property
    @pulumi.getter(name="hasEdgeRouterConnection")
    def has_edge_router_connection(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "has_edge_router_connection")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter(name="isDefaultAdmin")
    def is_default_admin(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_default_admin")

    @property
    @pulumi.getter(name="isMfaEnabled")
    def is_mfa_enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_mfa_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleAttributes")
    def role_attributes(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "role_attributes")

    @property
    @pulumi.getter(name="sdkInfo")
    def sdk_info(self) -> pulumi.Output['outputs.SdkInfo']:
        return pulumi.get(self, "sdk_info")

    @property
    @pulumi.getter(name="serviceHostingCosts")
    def service_hosting_costs(self) -> pulumi.Output[Mapping[str, int]]:
        return pulumi.get(self, "service_hosting_costs")

    @property
    @pulumi.getter(name="serviceHostingPrecedences")
    def service_hosting_precedences(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "service_hosting_precedences")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['outputs.EntityRef']:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="typeId")
    def type_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

