// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openziti
{
    [OpenzitiResourceType("openziti:index:ConfigObj")]
    public partial class ConfigObj : global::Pulumi.CustomResource
    {
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("_links")]
        public Output<ImmutableDictionary<string, Outputs.Link>> _links { get; private set; } = null!;

        [Output("configType")]
        public Output<Outputs.EntityRef> ConfigType { get; private set; } = null!;

        [Output("configTypeId")]
        public Output<string> ConfigTypeId { get; private set; } = null!;

        [Output("configTypeName")]
        public Output<string> ConfigTypeName { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("data")]
        public Output<object> Data { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigObj resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigObj(string name, ConfigObjArgs args, CustomResourceOptions? options = null)
            : base("openziti:index:ConfigObj", name, args ?? new ConfigObjArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigObj(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("openziti:index:ConfigObj", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigObj resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigObj Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigObj(name, id, options);
        }
    }

    public sealed class ConfigObjArgs : global::Pulumi.ResourceArgs
    {
        [Input("configTypeName", required: true)]
        public Input<string> ConfigTypeName { get; set; } = null!;

        [Input("data", required: true)]
        public Input<object> Data { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ConfigObjArgs()
        {
        }
        public static new ConfigObjArgs Empty => new ConfigObjArgs();
    }
}
