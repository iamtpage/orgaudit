# orgaudit

Assign the SpaceAuditor role to Cloud Foundry OrgAuditors.

`orgaudit` assigns the SpaceAuditor role to all OrgAuditors in a Cloud Foundry Organization.

**Usage:**

Provide `orgaudit` with your system domain, a user ID and client secret, and the
name of the organization to run against:

`$ orgaudit -d local.pcfdev.io --user-id cf-mgmt --client-secret cf-mgmt-secret -o foo`

The `--user-id` and `--client-secret` flags can be omitted by using the
`USER_ID` and `CLIENT_SECRET` environment variables.

**Setup:**

You'll need a client ID and client secret with the appropriate scopes in order
to perform these actions.  To create a UAA client, use the
[`uaac` Command Line](https://github.com/cloudfoundry/cf-uaac):

```
$ CLIENT_ID=user
$ CLIENT_SECRET=s3cr3t
$ UAA_DOMAIN=uaa.local.pcfdev.io    # typically uaa.<system_domain>

$ uaac target $UAA_DOMAIN
$ uaac token client get admin -s admin-client-secret
$ uaac client add $CLIENT_ID \
  --name $CLIENT_ID \
  --secret $CLIENT_SECRET \
  --authorized_grant_types client_credentials,refresh_token \
  --authorities cloud_controller.admin,scim.read,scim.write
```

