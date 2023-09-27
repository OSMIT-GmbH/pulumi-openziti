import * as pulumi from "@pulumi/pulumi";
import * as openziti from "@pulumi/openziti";

const openzitiConfig = new pulumi.Config("openziti");
const invokeOptions: pulumi.ResourceOptions = {
    provider: new openziti.Provider('openziti-provider', {
        uri: openzitiConfig.require("uri"),
        user: openzitiConfig.require("user"),
        password: openzitiConfig.require("password"),
        assimilate: "true"
    }),
    // ignoreChanges: ["*"]
};

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