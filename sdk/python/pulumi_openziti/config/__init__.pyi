# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

assimilate: Optional[str]
"""
Assimilate an existing object during create
"""

deleteAssimilated: Optional[str]
"""
Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
"""

password: Optional[str]
"""
The password. It is very secret.
"""

uri: Optional[str]
"""
The URI to the API
"""

user: Optional[str]
"""
The username. It's important but not secret.
"""

version: Optional[str]

