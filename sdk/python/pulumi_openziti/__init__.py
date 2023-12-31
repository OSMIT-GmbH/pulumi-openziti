# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .config_obj import *
from .edge_router import *
from .edge_router_policy import *
from .enrolled_identity import *
from .identity import *
from .provider import *
from .service import *
from .service_edge_router_policy import *
from .service_policy import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_openziti.config as __config
    config = __config
    import pulumi_openziti.rest_model as __rest_model
    rest_model = __rest_model
else:
    config = _utilities.lazy_import('pulumi_openziti.config')
    rest_model = _utilities.lazy_import('pulumi_openziti.rest_model')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "openziti",
  "mod": "index",
  "fqn": "pulumi_openziti",
  "classes": {
   "openziti:index:ConfigObj": "ConfigObj",
   "openziti:index:EdgeRouter": "EdgeRouter",
   "openziti:index:EdgeRouterPolicy": "EdgeRouterPolicy",
   "openziti:index:EnrolledIdentity": "EnrolledIdentity",
   "openziti:index:Identity": "Identity",
   "openziti:index:Service": "Service",
   "openziti:index:ServiceEdgeRouterPolicy": "ServiceEdgeRouterPolicy",
   "openziti:index:ServicePolicy": "ServicePolicy"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "openziti",
  "token": "pulumi:providers:openziti",
  "fqn": "pulumi_openziti",
  "class": "Provider"
 }
]
"""
)
