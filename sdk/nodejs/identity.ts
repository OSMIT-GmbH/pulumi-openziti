// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Identity extends pulumi.CustomResource {
    /**
     * Get an existing Identity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Identity {
        return new Identity(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openziti:index:Identity';

    /**
     * Returns true if the given object is an instance of Identity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Identity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Identity.__pulumiType;
    }

    public /*out*/ readonly _assimilated!: pulumi.Output<boolean>;
    public /*out*/ readonly _links!: pulumi.Output<{[key: string]: outputs.Link}>;
    public readonly appData!: pulumi.Output<{[key: string]: any} | undefined>;
    public /*out*/ readonly authPolicy!: pulumi.Output<outputs.EntityRef>;
    public readonly authPolicyId!: pulumi.Output<string>;
    public /*out*/ readonly authenticators!: pulumi.Output<outputs.rest_model.IdentityAuthenticators>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly defaultHostingCost!: pulumi.Output<number>;
    public readonly defaultHostingPrecedence!: pulumi.Output<string | undefined>;
    public /*out*/ readonly disabled!: pulumi.Output<boolean>;
    public /*out*/ readonly disabledAt!: pulumi.Output<string | undefined>;
    public /*out*/ readonly disabledUntil!: pulumi.Output<string | undefined>;
    public readonly enrollment!: pulumi.Output<outputs.IdentityEnrollments>;
    public /*out*/ readonly envInfo!: pulumi.Output<outputs.EnvInfo>;
    public readonly externalId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly hasApiSession!: pulumi.Output<boolean>;
    public /*out*/ readonly hasEdgeRouterConnection!: pulumi.Output<boolean>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    public readonly isAdmin!: pulumi.Output<boolean>;
    public /*out*/ readonly isDefaultAdmin!: pulumi.Output<boolean>;
    public /*out*/ readonly isMfaEnabled!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    public readonly roleAttributes!: pulumi.Output<string[]>;
    public /*out*/ readonly sdkInfo!: pulumi.Output<outputs.SdkInfo>;
    public readonly serviceHostingCosts!: pulumi.Output<{[key: string]: number}>;
    public readonly serviceHostingPrecedences!: pulumi.Output<{[key: string]: string}>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly type!: pulumi.Output<outputs.EntityRef>;
    public /*out*/ readonly typeId!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Identity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.isAdmin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isAdmin'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["appData"] = args ? args.appData : undefined;
            resourceInputs["authPolicyId"] = args ? args.authPolicyId : undefined;
            resourceInputs["defaultHostingCost"] = args ? args.defaultHostingCost : undefined;
            resourceInputs["defaultHostingPrecedence"] = args ? args.defaultHostingPrecedence : undefined;
            resourceInputs["enrollment"] = args ? args.enrollment : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["isAdmin"] = args ? args.isAdmin : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["roleAttributes"] = args ? args.roleAttributes : undefined;
            resourceInputs["serviceHostingCosts"] = args ? args.serviceHostingCosts : undefined;
            resourceInputs["serviceHostingPrecedences"] = args ? args.serviceHostingPrecedences : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["_assimilated"] = undefined /*out*/;
            resourceInputs["_links"] = undefined /*out*/;
            resourceInputs["authPolicy"] = undefined /*out*/;
            resourceInputs["authenticators"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["disabled"] = undefined /*out*/;
            resourceInputs["disabledAt"] = undefined /*out*/;
            resourceInputs["disabledUntil"] = undefined /*out*/;
            resourceInputs["envInfo"] = undefined /*out*/;
            resourceInputs["hasApiSession"] = undefined /*out*/;
            resourceInputs["hasEdgeRouterConnection"] = undefined /*out*/;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["isDefaultAdmin"] = undefined /*out*/;
            resourceInputs["isMfaEnabled"] = undefined /*out*/;
            resourceInputs["sdkInfo"] = undefined /*out*/;
            resourceInputs["typeId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["_assimilated"] = undefined /*out*/;
            resourceInputs["_links"] = undefined /*out*/;
            resourceInputs["appData"] = undefined /*out*/;
            resourceInputs["authPolicy"] = undefined /*out*/;
            resourceInputs["authPolicyId"] = undefined /*out*/;
            resourceInputs["authenticators"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["defaultHostingCost"] = undefined /*out*/;
            resourceInputs["defaultHostingPrecedence"] = undefined /*out*/;
            resourceInputs["disabled"] = undefined /*out*/;
            resourceInputs["disabledAt"] = undefined /*out*/;
            resourceInputs["disabledUntil"] = undefined /*out*/;
            resourceInputs["enrollment"] = undefined /*out*/;
            resourceInputs["envInfo"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["hasApiSession"] = undefined /*out*/;
            resourceInputs["hasEdgeRouterConnection"] = undefined /*out*/;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["isAdmin"] = undefined /*out*/;
            resourceInputs["isDefaultAdmin"] = undefined /*out*/;
            resourceInputs["isMfaEnabled"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["roleAttributes"] = undefined /*out*/;
            resourceInputs["sdkInfo"] = undefined /*out*/;
            resourceInputs["serviceHostingCosts"] = undefined /*out*/;
            resourceInputs["serviceHostingPrecedences"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["typeId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["enrollment"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Identity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Identity resource.
 */
export interface IdentityArgs {
    appData?: pulumi.Input<{[key: string]: any}>;
    authPolicyId?: pulumi.Input<string>;
    defaultHostingCost?: pulumi.Input<number>;
    defaultHostingPrecedence?: pulumi.Input<string>;
    enrollment?: pulumi.Input<inputs.IdentityCreateEnrollmentArgs>;
    externalId?: pulumi.Input<string>;
    isAdmin: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    roleAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    serviceHostingCosts?: pulumi.Input<{[key: string]: pulumi.Input<number>}>;
    serviceHostingPrecedences?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tags?: pulumi.Input<{[key: string]: any}>;
    type: pulumi.Input<string>;
}
