import * as openziti from "@pulumi/openziti";

const obj1 = new openziti.ConfigObj("my-config-1",
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

export const output = obj1.data;
export const name = obj1.name;
export const id = obj1.id;
