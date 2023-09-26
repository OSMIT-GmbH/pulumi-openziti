import * as openziti from "@pulumi/openziti";

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
    },
);
const obj2 = new openziti.ConfigObj('oz-test-hostv1-config',
    {
        name: 'testconfig.host.v1',
        configTypeName: 'host.v1',
        data: {
            address: 'test.ziti',
            port: 443,
            protocol: 'tcp',
        },
    },
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
    },
);
export const output = obj1.data;
export const name = obj1.name;
export const id = obj1.id;
