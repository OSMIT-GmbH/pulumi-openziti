// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openziti
{
    [OpenzitiResourceType("openziti:index:ServicePolicy")]
    public partial class ServicePolicy : global::Pulumi.CustomResource
    {
        [Output("_links")]
        public Output<ImmutableDictionary<string, Outputs.Link>> _links { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("identityRoles")]
        public Output<ImmutableArray<string>> IdentityRoles { get; private set; } = null!;

        [Output("identityRolesDisplay")]
        public Output<ImmutableArray<Outputs.NamedRole>> IdentityRolesDisplay { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("postureCheckRoles")]
        public Output<ImmutableArray<string>> PostureCheckRoles { get; private set; } = null!;

        [Output("postureCheckRolesDisplay")]
        public Output<ImmutableArray<Outputs.NamedRole>> PostureCheckRolesDisplay { get; private set; } = null!;

        [Output("semantic")]
        public Output<string> Semantic { get; private set; } = null!;

        [Output("serviceRoles")]
        public Output<ImmutableArray<string>> ServiceRoles { get; private set; } = null!;

        [Output("serviceRolesDisplay")]
        public Output<ImmutableArray<Outputs.NamedRole>> ServiceRolesDisplay { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePolicy(string name, ServicePolicyArgs args, CustomResourceOptions? options = null)
            : base("openziti:index:ServicePolicy", name, args ?? new ServicePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("openziti:index:ServicePolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServicePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServicePolicy(name, id, options);
        }
    }

    public sealed class ServicePolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("identityRoles", required: true)]
        private InputList<string>? _identityRoles;
        public InputList<string> IdentityRoles
        {
            get => _identityRoles ?? (_identityRoles = new InputList<string>());
            set => _identityRoles = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("postureCheckRoles")]
        private InputList<string>? _postureCheckRoles;
        public InputList<string> PostureCheckRoles
        {
            get => _postureCheckRoles ?? (_postureCheckRoles = new InputList<string>());
            set => _postureCheckRoles = value;
        }

        [Input("semantic", required: true)]
        public Input<string> Semantic { get; set; } = null!;

        [Input("serviceRoles", required: true)]
        private InputList<string>? _serviceRoles;
        public InputList<string> ServiceRoles
        {
            get => _serviceRoles ?? (_serviceRoles = new InputList<string>());
            set => _serviceRoles = value;
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

        public ServicePolicyArgs()
        {
        }
        public static new ServicePolicyArgs Empty => new ServicePolicyArgs();
    }
}