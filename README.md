# orgaudit

Assign the SpaceAuditor role to Cloud Foundry OrgAuditors.

`orgaudit` assigns the SpaceAuditor role to all OrgAuditors in a Cloud Foundry Organization.

**Releases:**

Download a release for your platform from the [releases](https://github.com/iamtpage/orgaudit/releases) page.

**Usage:**
Note: pre-v0.0.5 supports the `-o` option instead of `-w`/`-b`, with `-o` allowing you to specify one Org. v0.0.5+ takes all the Orgs in the foundation and allows you fine-grain control using the whitelist/blacklist functionality

Provide `orgaudit` with your system domain, a user ID and client secret, and the
whitelist/blacklist configuration to suite your use-case:

`$ orgaudit -d local.pcfdev.io --user-id cf-orgaudit --client-secret cf-orgaudit-secret -w "org-1,org-2" -b "system,services"`

The `--user-id` and `--client-secret` flags can be omitted by using the
`USER_ID` and `CLIENT_SECRET` environment variables.

**Setup:**

You'll need a client ID and client secret with the appropriate scopes in order
to perform these actions. To create a UAA client, use the
[`uaac` Command Line](https://github.com/cloudfoundry/cf-uaac):

```
$ CLIENT_ID=user
$ CLIENT_SECRET=s3cr3t
$ UAA_DOMAIN=uaa.local.pcfdev.io # typically uaa.<system_domain>

$ uaac target $UAA_DOMAIN
$ uaac token client get admin -s admin-client-secret
$ uaac client add $CLIENT_ID \
--name $CLIENT_ID \
--secret $CLIENT_SECRET \
--authorized_grant_types client_credentials,refresh_token \
--authorities cloud_controller.admin,scim.read,scim.write
```
