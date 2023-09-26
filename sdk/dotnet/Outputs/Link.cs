// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openziti.Outputs
{

    [OutputType]
    public sealed class Link
    {
        public readonly string? Comment;
        public readonly string Href;
        public readonly string? Method;

        [OutputConstructor]
        private Link(
            string? comment,

            string href,

            string? method)
        {
            Comment = comment;
            Href = href;
            Method = method;
        }
    }
}