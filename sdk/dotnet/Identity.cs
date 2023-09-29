// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openziti
{
    [OpenzitiResourceType("openziti:index:Identity")]
    public partial class Identity : global::Pulumi.CustomResource
    {
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("_links")]
        public Output<ImmutableDictionary<string, Outputs.Link>> _links { get; private set; } = null!;

        [Output("appData")]
        public Output<ImmutableDictionary<string, object>?> AppData { get; private set; } = null!;

        [Output("authPolicy")]
        public Output<Outputs.EntityRef> AuthPolicy { get; private set; } = null!;

        [Output("authPolicyId")]
        public Output<string> AuthPolicyId { get; private set; } = null!;

        [Output("authenticators")]
        public Output<Pulumi.Openziti.Rest_model.Outputs.IdentityAuthenticators> Authenticators { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("defaultHostingCost")]
        public Output<int> DefaultHostingCost { get; private set; } = null!;

        [Output("defaultHostingPrecedence")]
        public Output<string?> DefaultHostingPrecedence { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        [Output("disabledAt")]
        public Output<string?> DisabledAt { get; private set; } = null!;

        [Output("disabledUntil")]
        public Output<string?> DisabledUntil { get; private set; } = null!;

        [Output("enrollment")]
        public Output<Outputs.IdentityEnrollments> Enrollment { get; private set; } = null!;

        [Output("envInfo")]
        public Output<Outputs.EnvInfo> EnvInfo { get; private set; } = null!;

        [Output("externalId")]
        public Output<string?> ExternalId { get; private set; } = null!;

        [Output("hasApiSession")]
        public Output<bool> HasApiSession { get; private set; } = null!;

        [Output("hasEdgeRouterConnection")]
        public Output<bool> HasEdgeRouterConnection { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("isAdmin")]
        public Output<bool> IsAdmin { get; private set; } = null!;

        [Output("isDefaultAdmin")]
        public Output<bool> IsDefaultAdmin { get; private set; } = null!;

        [Output("isMfaEnabled")]
        public Output<bool> IsMfaEnabled { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("roleAttributes")]
        public Output<ImmutableArray<string>> RoleAttributes { get; private set; } = null!;

        [Output("sdkInfo")]
        public Output<Outputs.SdkInfo> SdkInfo { get; private set; } = null!;

        [Output("serviceHostingCosts")]
        public Output<ImmutableDictionary<string, int>> ServiceHostingCosts { get; private set; } = null!;

        [Output("serviceHostingPrecedences")]
        public Output<ImmutableDictionary<string, string>> ServiceHostingPrecedences { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Outputs.EntityRef> Type { get; private set; } = null!;

        [Output("typeId")]
        public Output<string> TypeId { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Identity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Identity(string name, IdentityArgs args, CustomResourceOptions? options = null)
            : base("openziti:index:Identity", name, args ?? new IdentityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Identity(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("openziti:index:Identity", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "enrollment",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Identity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Identity Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Identity(name, id, options);
        }
    }

    public sealed class IdentityArgs : global::Pulumi.ResourceArgs
    {
        [Input("appData")]
        private InputMap<object>? _appData;
        public InputMap<object> AppData
        {
            get => _appData ?? (_appData = new InputMap<object>());
            set => _appData = value;
        }

        [Input("authPolicyId")]
        public Input<string>? AuthPolicyId { get; set; }

        [Input("defaultHostingCost")]
        public Input<int>? DefaultHostingCost { get; set; }

        [Input("defaultHostingPrecedence")]
        public Input<string>? DefaultHostingPrecedence { get; set; }

        [Input("enrollment")]
        public Input<Inputs.IdentityCreateEnrollmentArgs>? Enrollment { get; set; }

        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("isAdmin", required: true)]
        public Input<bool> IsAdmin { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("roleAttributes")]
        private InputList<string>? _roleAttributes;
        public InputList<string> RoleAttributes
        {
            get => _roleAttributes ?? (_roleAttributes = new InputList<string>());
            set => _roleAttributes = value;
        }

        [Input("serviceHostingCosts")]
        private InputMap<int>? _serviceHostingCosts;
        public InputMap<int> ServiceHostingCosts
        {
            get => _serviceHostingCosts ?? (_serviceHostingCosts = new InputMap<int>());
            set => _serviceHostingCosts = value;
        }

        [Input("serviceHostingPrecedences")]
        private InputMap<string>? _serviceHostingPrecedences;
        public InputMap<string> ServiceHostingPrecedences
        {
            get => _serviceHostingPrecedences ?? (_serviceHostingPrecedences = new InputMap<string>());
            set => _serviceHostingPrecedences = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IdentityArgs()
        {
        }
        public static new IdentityArgs Empty => new IdentityArgs();
    }
}
