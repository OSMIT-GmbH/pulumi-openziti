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

__all__ = [
    'EntityRef',
    'EnvInfo',
    'IdentityEnrollments',
    'IdentityEnrollmentsOtt',
    'IdentityEnrollmentsOttca',
    'IdentityEnrollmentsUpdb',
    'Link',
    'NamedRole',
    'PostureQueriesType',
    'SdkInfo',
]

@pulumi.output_type
class EntityRef(dict):
    def __init__(__self__, *,
                 _links: Mapping[str, 'outputs.Link'],
                 entity: Optional[str] = None,
                 id: Optional[str] = None,
                 name: Optional[str] = None):
        EntityRef._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            _links=_links,
            entity=entity,
            id=id,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             _links: Mapping[str, 'outputs.Link'],
             entity: Optional[str] = None,
             id: Optional[str] = None,
             name: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("_links", _links)
        if entity is not None:
            _setter("entity", entity)
        if id is not None:
            _setter("id", id)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def _links(self) -> Mapping[str, 'outputs.Link']:
        return pulumi.get(self, "_links")

    @property
    @pulumi.getter
    def entity(self) -> Optional[str]:
        return pulumi.get(self, "entity")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


@pulumi.output_type
class EnvInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "osRelease":
            suggest = "os_release"
        elif key == "osVersion":
            suggest = "os_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arch: Optional[str] = None,
                 os: Optional[str] = None,
                 os_release: Optional[str] = None,
                 os_version: Optional[str] = None):
        EnvInfo._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            arch=arch,
            os=os,
            os_release=os_release,
            os_version=os_version,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             arch: Optional[str] = None,
             os: Optional[str] = None,
             os_release: Optional[str] = None,
             os_version: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if arch is not None:
            _setter("arch", arch)
        if os is not None:
            _setter("os", os)
        if os_release is not None:
            _setter("os_release", os_release)
        if os_version is not None:
            _setter("os_version", os_version)

    @property
    @pulumi.getter
    def arch(self) -> Optional[str]:
        return pulumi.get(self, "arch")

    @property
    @pulumi.getter
    def os(self) -> Optional[str]:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="osRelease")
    def os_release(self) -> Optional[str]:
        return pulumi.get(self, "os_release")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> Optional[str]:
        return pulumi.get(self, "os_version")


@pulumi.output_type
class IdentityEnrollments(dict):
    def __init__(__self__, *,
                 ott: Optional['outputs.IdentityEnrollmentsOtt'] = None,
                 ottca: Optional['outputs.IdentityEnrollmentsOttca'] = None,
                 updb: Optional['outputs.IdentityEnrollmentsUpdb'] = None):
        IdentityEnrollments._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ott=ott,
            ottca=ottca,
            updb=updb,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ott: Optional['outputs.IdentityEnrollmentsOtt'] = None,
             ottca: Optional['outputs.IdentityEnrollmentsOttca'] = None,
             updb: Optional['outputs.IdentityEnrollmentsUpdb'] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if ott is not None:
            _setter("ott", ott)
        if ottca is not None:
            _setter("ottca", ottca)
        if updb is not None:
            _setter("updb", updb)

    @property
    @pulumi.getter
    def ott(self) -> Optional['outputs.IdentityEnrollmentsOtt']:
        return pulumi.get(self, "ott")

    @property
    @pulumi.getter
    def ottca(self) -> Optional['outputs.IdentityEnrollmentsOttca']:
        return pulumi.get(self, "ottca")

    @property
    @pulumi.getter
    def updb(self) -> Optional['outputs.IdentityEnrollmentsUpdb']:
        return pulumi.get(self, "updb")


@pulumi.output_type
class IdentityEnrollmentsOtt(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expiresAt":
            suggest = "expires_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IdentityEnrollmentsOtt. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IdentityEnrollmentsOtt.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IdentityEnrollmentsOtt.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expires_at: Optional[str] = None,
                 id: Optional[str] = None,
                 jwt: Optional[str] = None,
                 token: Optional[str] = None):
        IdentityEnrollmentsOtt._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            expires_at=expires_at,
            id=id,
            jwt=jwt,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             expires_at: Optional[str] = None,
             id: Optional[str] = None,
             jwt: Optional[str] = None,
             token: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if id is not None:
            _setter("id", id)
        if jwt is not None:
            _setter("jwt", jwt)
        if token is not None:
            _setter("token", token)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[str]:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def jwt(self) -> Optional[str]:
        return pulumi.get(self, "jwt")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        return pulumi.get(self, "token")


@pulumi.output_type
class IdentityEnrollmentsOttca(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "caId":
            suggest = "ca_id"
        elif key == "expiresAt":
            suggest = "expires_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IdentityEnrollmentsOttca. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IdentityEnrollmentsOttca.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IdentityEnrollmentsOttca.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ca: Optional['outputs.EntityRef'] = None,
                 ca_id: Optional[str] = None,
                 expires_at: Optional[str] = None,
                 id: Optional[str] = None,
                 jwt: Optional[str] = None,
                 token: Optional[str] = None):
        IdentityEnrollmentsOttca._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ca=ca,
            ca_id=ca_id,
            expires_at=expires_at,
            id=id,
            jwt=jwt,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ca: Optional['outputs.EntityRef'] = None,
             ca_id: Optional[str] = None,
             expires_at: Optional[str] = None,
             id: Optional[str] = None,
             jwt: Optional[str] = None,
             token: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if ca is not None:
            _setter("ca", ca)
        if ca_id is not None:
            _setter("ca_id", ca_id)
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if id is not None:
            _setter("id", id)
        if jwt is not None:
            _setter("jwt", jwt)
        if token is not None:
            _setter("token", token)

    @property
    @pulumi.getter
    def ca(self) -> Optional['outputs.EntityRef']:
        return pulumi.get(self, "ca")

    @property
    @pulumi.getter(name="caId")
    def ca_id(self) -> Optional[str]:
        return pulumi.get(self, "ca_id")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[str]:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def jwt(self) -> Optional[str]:
        return pulumi.get(self, "jwt")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        return pulumi.get(self, "token")


@pulumi.output_type
class IdentityEnrollmentsUpdb(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expiresAt":
            suggest = "expires_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IdentityEnrollmentsUpdb. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IdentityEnrollmentsUpdb.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IdentityEnrollmentsUpdb.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 expires_at: Optional[str] = None,
                 id: Optional[str] = None,
                 jwt: Optional[str] = None,
                 token: Optional[str] = None):
        IdentityEnrollmentsUpdb._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            expires_at=expires_at,
            id=id,
            jwt=jwt,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             expires_at: Optional[str] = None,
             id: Optional[str] = None,
             jwt: Optional[str] = None,
             token: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if id is not None:
            _setter("id", id)
        if jwt is not None:
            _setter("jwt", jwt)
        if token is not None:
            _setter("token", token)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[str]:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def jwt(self) -> Optional[str]:
        return pulumi.get(self, "jwt")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        return pulumi.get(self, "token")


@pulumi.output_type
class Link(dict):
    def __init__(__self__, *,
                 href: str,
                 comment: Optional[str] = None,
                 method: Optional[str] = None):
        Link._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            href=href,
            comment=comment,
            method=method,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             href: str,
             comment: Optional[str] = None,
             method: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("href", href)
        if comment is not None:
            _setter("comment", comment)
        if method is not None:
            _setter("method", method)

    @property
    @pulumi.getter
    def href(self) -> str:
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def method(self) -> Optional[str]:
        return pulumi.get(self, "method")


@pulumi.output_type
class NamedRole(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 role: Optional[str] = None):
        NamedRole._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[str] = None,
             role: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if name is not None:
            _setter("name", name)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        return pulumi.get(self, "role")


@pulumi.output_type
class PostureQueriesType(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "isPassing":
            suggest = "is_passing"
        elif key == "policyId":
            suggest = "policy_id"
        elif key == "postureQueries":
            suggest = "posture_queries"
        elif key == "policyType":
            suggest = "policy_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PostureQueriesType. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PostureQueriesType.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PostureQueriesType.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 is_passing: bool,
                 policy_id: str,
                 posture_queries: Sequence['_rest_model.outputs.PostureQuery'],
                 policy_type: Optional[str] = None):
        PostureQueriesType._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            is_passing=is_passing,
            policy_id=policy_id,
            posture_queries=posture_queries,
            policy_type=policy_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             is_passing: bool,
             policy_id: str,
             posture_queries: Sequence['_rest_model.outputs.PostureQuery'],
             policy_type: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("is_passing", is_passing)
        _setter("policy_id", policy_id)
        _setter("posture_queries", posture_queries)
        if policy_type is not None:
            _setter("policy_type", policy_type)

    @property
    @pulumi.getter(name="isPassing")
    def is_passing(self) -> bool:
        return pulumi.get(self, "is_passing")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="postureQueries")
    def posture_queries(self) -> Sequence['_rest_model.outputs.PostureQuery']:
        return pulumi.get(self, "posture_queries")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[str]:
        return pulumi.get(self, "policy_type")


@pulumi.output_type
class SdkInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "appId":
            suggest = "app_id"
        elif key == "appVersion":
            suggest = "app_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SdkInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SdkInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SdkInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 app_id: Optional[str] = None,
                 app_version: Optional[str] = None,
                 branch: Optional[str] = None,
                 revision: Optional[str] = None,
                 type: Optional[str] = None,
                 version: Optional[str] = None):
        SdkInfo._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            app_id=app_id,
            app_version=app_version,
            branch=branch,
            revision=revision,
            type=type,
            version=version,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             app_id: Optional[str] = None,
             app_version: Optional[str] = None,
             branch: Optional[str] = None,
             revision: Optional[str] = None,
             type: Optional[str] = None,
             version: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if app_id is not None:
            _setter("app_id", app_id)
        if app_version is not None:
            _setter("app_version", app_version)
        if branch is not None:
            _setter("branch", branch)
        if revision is not None:
            _setter("revision", revision)
        if type is not None:
            _setter("type", type)
        if version is not None:
            _setter("version", version)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[str]:
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="appVersion")
    def app_version(self) -> Optional[str]:
        return pulumi.get(self, "app_version")

    @property
    @pulumi.getter
    def branch(self) -> Optional[str]:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def revision(self) -> Optional[str]:
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


