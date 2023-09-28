import * as pulumi from "@pulumi/pulumi";
import * as openziti from "@pulumi/openziti";

// This is optional - when provider is not specified the openziti:[uri/user/password] keys are taken directly.
const openzitiConfig = new pulumi.Config("openziti");
const invokeOptions: pulumi.ResourceOptions = {
    provider: new openziti.Provider('openziti-provider', {
        uri: openzitiConfig.require("uri"),
        user: openzitiConfig.require("user"),
        password: openzitiConfig.require("password"),
        assimilate: "true",
    }, {
        // provider config changes currently enforce a Create/Delete on all items
        // associated with this provider. This is not a good idea for Identity's and
        // EdgeRouter's. So we better ignore changes on the provider. ;)
        ignoreChanges: ["version", "uri", "assimilate", "user", "password", "insecure"]
    }),
};

const router = new openziti.EdgeRouter('oz-test-edge-router',
    {
        name: 'pulumi-edge-router',
        roleAttributes: ['public'],
        isTunnelerEnabled: true,
    }, invokeOptions
);

const routerPolicy = new openziti.EdgeRouterPolicy('oz-test-edge-router-policy',
    {
        name: 'pulumi-all-endpoints-public-routers',
        edgeRouterRoles: ['#public'],
        identityRoles: ['#all'],
        semantic: 'AnyOf',
    }, invokeOptions
);

const routerServicePolicy = new openziti.ServiceEdgeRouterPolicy('oz-test-service-edge-router-policy',
    {
        name: 'pulumi-all-routers-all-services',
        edgeRouterRoles: ['#all'],
        serviceRoles: ['#all'],
        semantic: 'AnyOf',
    }, invokeOptions
);

const obj1 = new openziti.ConfigObj('oz-test-interceptv1-config',
    {
        name: 'testconfig.intercept.v1',
        configTypeName: 'intercept.v1',
        data: {
            addresses: ['test.ziti'],
            portRanges: [
                {
                    high: 80,
                    low: 80,
                },
                {
                    high: 443,
                    low: 443,
                },
            ],
            protocols: ['tcp'],
        },
        tags: {
            hello: "world"
        }
    }, invokeOptions
);
export const output = obj1.data;
export const name = obj1.name;
export const id = obj1.id;

const obj2 = new openziti.ConfigObj('oz-test-hostv1-config',
    {
        name: 'testconfig.host.v1',
        configTypeName: 'host.v1',
        data: {
            address: 'test.ziti',
            port: 443,
            protocol: 'tcp',
        },
    }, invokeOptions
);


const svc = new openziti.Service('oz-test-service',
    {
        name: 'test-service',
        configs: [ obj1.id, obj2.id ],
        encryptionRequired: true,
        roleAttributes: ['test1'],
        tags: { pulumi: "yes!" }
    },invokeOptions
);

const id1 = new openziti.Identity('oz-test-identity',
    {
        name: 'pulumi-test',
        type: 'Default',
        isAdmin: false,
        enrollment: {
            ott: true
        },
        tags: {
            pulumi: "true"
        }
    }, invokeOptions
);
export const id1exp = id1;
export const id1id = id1.id;


const svcBind = new openziti.ServicePolicy('oz-test-service-pol-bind',
    {
        name: 'test-service.bind',
        // identityRoles: ['#test', id1.id.apply((id) => `@${id}`)],
        identityRoles: ['#test',],
        semantic: 'AnyOf',
        serviceRoles: ['#test1'],
        tags: {pulumi: "yes!"},
        type: 'Bind',
    }, invokeOptions
);


export const routerId = router.id;
export const routerEnrolmentToken = router.enrollmentToken;
