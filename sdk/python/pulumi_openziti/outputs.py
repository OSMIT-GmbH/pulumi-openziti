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

__all__ = [
    'EntityRef',
    'Link',
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


