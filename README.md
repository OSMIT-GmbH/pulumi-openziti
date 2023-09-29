# Pulumi Provider For OpenZiti

This provider supports access to some of the OpenZiti management APIs.
Currently, it only supports some of the APIs we need for our purpose.
Feel free to extend it ;)

**WARNING:** This is in a quite an early stage - things my change ;)

## Supported APIs

Following API's are supported (CRUD Operations)

| API-Name                | Link to API Description                                                                                                                            | Comments                                                                                                                             |
|-------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
| ConfigObj               | [Management API: Config](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/Config)                                   | Named as `ConfigObj` to avoid naming issues in C# SDK; Additional `configTypeName` field automatically mapping to the `configTypeId` |
| Identity                | [Management API: Identity](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/Identity)                               |                                                                                                                                      |
| EdgeRouter              | [Management API: Edge Router](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/Edge-Router)                         |                                                                                                                                      |
| EdgeRouterPolicy        | [Management API: Edge Router Policy](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/Edge-Router-Policy)           |                                                                                                                                      |
| EnrolledIdentity        | [Management API: Enroll](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/Enroll)                                   | This enrolls an Identity by its JWT. It takes the Identity's JWT from `enrollment.ott.jwt` as input and generates an `identityJson`  |
| Service                 | [Management API: Service](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/Service)                                 |                                                                                                                                      |
| ServiceEdgeRouterPolicy | [Management API: ServiceEdgeRouterPolicy](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/ServiceEdgeRouterPolicy) |                                                                                                                                      |
| ServicePolicy           | [Management API: ServicePolicy](https://openziti.io/docs/reference/developer/api/edge-management-reference/#tag/ServicePolicy)                     | `identityRoles` has a lookup: `@name`-identities are mapped from it's name to the id.                                                |

## Usage

You have to set following config options:

```bash
pulumi config set openziti:uri https://localhost:1280
pulumi config set openziti:user admin
pulumi config set openziti:password <myTopSecretZitiPassword>
# `assimilate` allows to integrate existing objects. Great for migration
# pulumi config set openziti:assimilate true
```

For samples how to create the OpenZiti Objects have a look on the example [index.ts](examples/simple/index.ts)

## Developing

This provider is based on the [pulumi-provider-boilerplate](https://github.com/pulumi/pulumi-provider-boilerplate).
Just use the README.md documentation from the boilerplate to set up your development environment.

This boilerplate creates a working Pulumi-owned provider named `openziti`.
It implements a random number generator that you can [build and test out for yourself](#test-against-the-example) and
then replace the Random code with code specific to your provider.

