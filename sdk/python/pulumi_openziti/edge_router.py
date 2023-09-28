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

__all__ = ['EdgeRouterArgs', 'EdgeRouter']

@pulumi.input_type
class EdgeRouterArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 cost: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 is_tunneler_enabled: Optional[pulumi.Input[bool]] = None,
                 no_traversal: Optional[pulumi.Input[bool]] = None,
                 role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a EdgeRouter resource.
        """
        EdgeRouterArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            app_data=app_data,
            cost=cost,
            disabled=disabled,
            is_tunneler_enabled=is_tunneler_enabled,
            no_traversal=no_traversal,
            role_attributes=role_attributes,
            tags=tags,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: pulumi.Input[str],
             app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             cost: Optional[pulumi.Input[int]] = None,
             disabled: Optional[pulumi.Input[bool]] = None,
             is_tunneler_enabled: Optional[pulumi.Input[bool]] = None,
             no_traversal: Optional[pulumi.Input[bool]] = None,
             role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("name", name)
        if app_data is not None:
            _setter("app_data", app_data)
        if cost is not None:
            _setter("cost", cost)
        if disabled is not None:
            _setter("disabled", disabled)
        if is_tunneler_enabled is not None:
            _setter("is_tunneler_enabled", is_tunneler_enabled)
        if no_traversal is not None:
            _setter("no_traversal", no_traversal)
        if role_attributes is not None:
            _setter("role_attributes", role_attributes)
        if tags is not None:
            _setter("tags", tags)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="appData")
    def app_data(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "app_data")

    @app_data.setter
    def app_data(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "app_data", value)

    @property
    @pulumi.getter
    def cost(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "cost")

    @cost.setter
    def cost(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cost", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="isTunnelerEnabled")
    def is_tunneler_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_tunneler_enabled")

    @is_tunneler_enabled.setter
    def is_tunneler_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_tunneler_enabled", value)

    @property
    @pulumi.getter(name="noTraversal")
    def no_traversal(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_traversal")

    @no_traversal.setter
    def no_traversal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_traversal", value)

    @property
    @pulumi.getter(name="roleAttributes")
    def role_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "role_attributes")

    @role_attributes.setter
    def role_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_attributes", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class EdgeRouter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 cost: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 is_tunneler_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_traversal: Optional[pulumi.Input[bool]] = None,
                 role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Create a EdgeRouter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EdgeRouterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EdgeRouter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EdgeRouterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EdgeRouterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            EdgeRouterArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 cost: Optional[pulumi.Input[int]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 is_tunneler_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_traversal: Optional[pulumi.Input[bool]] = None,
                 role_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EdgeRouterArgs.__new__(EdgeRouterArgs)

            __props__.__dict__["app_data"] = app_data
            __props__.__dict__["cost"] = cost
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["is_tunneler_enabled"] = is_tunneler_enabled
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["no_traversal"] = no_traversal
            __props__.__dict__["role_attributes"] = role_attributes
            __props__.__dict__["tags"] = tags
            __props__.__dict__["_links"] = None
            __props__.__dict__["cert_pem"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["enrollment_created_at"] = None
            __props__.__dict__["enrollment_expires_at"] = None
            __props__.__dict__["enrollment_jwt"] = None
            __props__.__dict__["enrollment_token"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["hostname"] = None
            __props__.__dict__["id"] = None
            __props__.__dict__["is_online"] = None
            __props__.__dict__["is_verified"] = None
            __props__.__dict__["supported_protocols"] = None
            __props__.__dict__["sync_status"] = None
            __props__.__dict__["unverified_cert_pem"] = None
            __props__.__dict__["unverified_fingerprint"] = None
            __props__.__dict__["updated_at"] = None
        super(EdgeRouter, __self__).__init__(
            'openziti:index:EdgeRouter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EdgeRouter':
        """
        Get an existing EdgeRouter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EdgeRouterArgs.__new__(EdgeRouterArgs)

        __props__.__dict__["_links"] = None
        __props__.__dict__["app_data"] = None
        __props__.__dict__["cert_pem"] = None
        __props__.__dict__["cost"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["disabled"] = None
        __props__.__dict__["enrollment_created_at"] = None
        __props__.__dict__["enrollment_expires_at"] = None
        __props__.__dict__["enrollment_jwt"] = None
        __props__.__dict__["enrollment_token"] = None
        __props__.__dict__["fingerprint"] = None
        __props__.__dict__["hostname"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["is_online"] = None
        __props__.__dict__["is_tunneler_enabled"] = None
        __props__.__dict__["is_verified"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["no_traversal"] = None
        __props__.__dict__["role_attributes"] = None
        __props__.__dict__["supported_protocols"] = None
        __props__.__dict__["sync_status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["unverified_cert_pem"] = None
        __props__.__dict__["unverified_fingerprint"] = None
        __props__.__dict__["updated_at"] = None
        return EdgeRouter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def _links(self) -> pulumi.Output[Mapping[str, 'outputs.Link']]:
        return pulumi.get(self, "_links")

    @property
    @pulumi.getter(name="appData")
    def app_data(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "app_data")

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cert_pem")

    @property
    @pulumi.getter
    def cost(self) -> pulumi.Output[int]:
        return pulumi.get(self, "cost")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="enrollmentCreatedAt")
    def enrollment_created_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enrollment_created_at")

    @property
    @pulumi.getter(name="enrollmentExpiresAt")
    def enrollment_expires_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enrollment_expires_at")

    @property
    @pulumi.getter(name="enrollmentJwt")
    def enrollment_jwt(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enrollment_jwt")

    @property
    @pulumi.getter(name="enrollmentToken")
    def enrollment_token(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "enrollment_token")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isOnline")
    def is_online(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_online")

    @property
    @pulumi.getter(name="isTunnelerEnabled")
    def is_tunneler_enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_tunneler_enabled")

    @property
    @pulumi.getter(name="isVerified")
    def is_verified(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_verified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noTraversal")
    def no_traversal(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "no_traversal")

    @property
    @pulumi.getter(name="roleAttributes")
    def role_attributes(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "role_attributes")

    @property
    @pulumi.getter(name="supportedProtocols")
    def supported_protocols(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "supported_protocols")

    @property
    @pulumi.getter(name="syncStatus")
    def sync_status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sync_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="unverifiedCertPem")
    def unverified_cert_pem(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "unverified_cert_pem")

    @property
    @pulumi.getter(name="unverifiedFingerprint")
    def unverified_fingerprint(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "unverified_fingerprint")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

