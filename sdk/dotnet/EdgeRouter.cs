// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openziti
{
    [OpenzitiResourceType("openziti:index:EdgeRouter")]
    public partial class EdgeRouter : global::Pulumi.CustomResource
    {
        [Output("_assimilated")]
        public Output<bool> _assimilated { get; private set; } = null!;

        [Output("_links")]
        public Output<ImmutableDictionary<string, Outputs.Link>> _links { get; private set; } = null!;

        [Output("appData")]
        public Output<ImmutableDictionary<string, object>?> AppData { get; private set; } = null!;

        [Output("certPem")]
        public Output<string?> CertPem { get; private set; } = null!;

        [Output("cost")]
        public Output<int> Cost { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        [Output("enrollmentCreatedAt")]
        public Output<string?> EnrollmentCreatedAt { get; private set; } = null!;

        [Output("enrollmentExpiresAt")]
        public Output<string?> EnrollmentExpiresAt { get; private set; } = null!;

        [Output("enrollmentJwt")]
        public Output<string?> EnrollmentJwt { get; private set; } = null!;

        [Output("enrollmentToken")]
        public Output<string?> EnrollmentToken { get; private set; } = null!;

        [Output("fingerprint")]
        public Output<string?> Fingerprint { get; private set; } = null!;

        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("isOnline")]
        public Output<bool> IsOnline { get; private set; } = null!;

        [Output("isTunnelerEnabled")]
        public Output<bool> IsTunnelerEnabled { get; private set; } = null!;

        [Output("isVerified")]
        public Output<bool> IsVerified { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("noTraversal")]
        public Output<bool> NoTraversal { get; private set; } = null!;

        [Output("roleAttributes")]
        public Output<ImmutableArray<string>> RoleAttributes { get; private set; } = null!;

        [Output("supportedProtocols")]
        public Output<ImmutableDictionary<string, string>> SupportedProtocols { get; private set; } = null!;

        [Output("syncStatus")]
        public Output<string> SyncStatus { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("unverifiedCertPem")]
        public Output<string?> UnverifiedCertPem { get; private set; } = null!;

        [Output("unverifiedFingerprint")]
        public Output<string?> UnverifiedFingerprint { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeRouter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeRouter(string name, EdgeRouterArgs args, CustomResourceOptions? options = null)
            : base("openziti:index:EdgeRouter", name, args ?? new EdgeRouterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeRouter(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("openziti:index:EdgeRouter", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "enrollmentJwt",
                    "enrollmentToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EdgeRouter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeRouter Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EdgeRouter(name, id, options);
        }
    }

    public sealed class EdgeRouterArgs : global::Pulumi.ResourceArgs
    {
        [Input("appData")]
        private InputMap<object>? _appData;
        public InputMap<object> AppData
        {
            get => _appData ?? (_appData = new InputMap<object>());
            set => _appData = value;
        }

        [Input("cost")]
        public Input<int>? Cost { get; set; }

        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("isTunnelerEnabled")]
        public Input<bool>? IsTunnelerEnabled { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("noTraversal")]
        public Input<bool>? NoTraversal { get; set; }

        [Input("roleAttributes")]
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

        public EdgeRouterArgs()
        {
        }
        public static new EdgeRouterArgs Empty => new EdgeRouterArgs();
    }
}
