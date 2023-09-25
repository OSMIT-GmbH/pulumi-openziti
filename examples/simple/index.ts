import * as openziti from "@pulumi/openziti";

const random = new openziti.Random("my-random", { length: 24 });

export const output = random.result;
