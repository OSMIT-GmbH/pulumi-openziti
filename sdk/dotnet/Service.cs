// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openziti
{
    [OpenzitiResourceType("openziti:index:Service")]
    public partial class Service : global::Pulumi.CustomResource
    {
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("_links")]
        public Output<ImmutableDictionary<string, Outputs.Link>> _links { get; private set; } = null!;

        [Output("config")]
        public Output<ImmutableDictionary<string, ImmutableDictionary<string, object>>> Config { get; private set; } = null!;

        [Output("configs")]
        public Output<ImmutableArray<string>> Configs { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("encryptionRequired")]
        public Output<bool> EncryptionRequired { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("permissions")]
        public Output<ImmutableArray<string>> Permissions { get; private set; } = null!;

        [Output("postureQueries")]
        public Output<ImmutableArray<Outputs.PostureQueriesType>> PostureQueries { get; private set; } = null!;

        [Output("roleAttributes")]
        public Output<ImmutableArray<string>> RoleAttributes { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("terminatorStrategy")]
        public Output<string> TerminatorStrategy { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("openziti:index:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("openziti:index:Service", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Service(name, id, options);
        }
    }

    public sealed class ServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("configs", required: true)]
        private InputList<string>? _configs;
        public InputList<string> Configs
        {
            get => _configs ?? (_configs = new InputList<string>());
            set => _configs = value;
        }

        [Input("encryptionRequired", required: true)]
        public Input<bool> EncryptionRequired { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("roleAttributes", required: true)]
        private InputList<string>? _roleAttributes;
        public InputList<string> RoleAttributes
        {
            get => _roleAttributes ?? (_roleAttributes = new InputList<string>());
            set => _roleAttributes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("terminatorStrategy")]
        public Input<string>? TerminatorStrategy { get; set; }

        public ServiceArgs()
        {
        }
        public static new ServiceArgs Empty => new ServiceArgs();
    }
}
