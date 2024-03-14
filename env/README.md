# edgio/env

This package groups Edgio Environment specific funcs.

## `env.NewClient(params common.ClientParams) (ClientStruct, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

envClient, err := env.NewClient(common.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{AccessToken: "some-access-token"},
})
```

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's environments.

### `env.NewClient` Mandatory Params

- `params.Credentials`
  - `params.Credentials.Key`: Edgio REST API Client Key. This is the API Client's identifier on Edgio. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
  - `params.Credentials.Secret`: Edgio REST API Client Secret. This is the API Client' secret, unique for each API Client. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.

### `env.NewClient` Optional Params & Default Values

- `params.Credentials`
  - `params.Credentials.Scopes`: Edgio REST API Client scopes requested by the client. Different APIs needs different scopes. Refer to the [REST API docs](https://docs.edg.io/rest_api) to figure which ones you need.
    - Default value: `app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config` (all scopes).
  - `params.Credentials.AuthURL`: Edgio REST API auth url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://id.edgio.app/connect/token` (Edgio's default auth API url).
- `params.Config`
  - `params.Config.ApiVersion`: Intended REST API Version. Each one of the Edgio REST APIs has its own Version, that must be provided when creating the client.
    - Default value: `v0.1`
  - `params.Config.Service`: Intended REST API Service. Each one of the Edgio REST APIs has its own Service, that must be provided when creating the client.
    - Default value: `accounts`
  - `params.Config.Scope`: Intended REST API Scope. Each one of the Edgio REST APIs has its own Scope, that must be provided when creating the client.
    - Default value: `properties`
  - `params.Config.Url`: Edgio REST API resources url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://edgioapis.com` (Edgio's default resources API url).

## `env.List(propertyID string) (ListResultType, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

client, err := env.NewClient(env.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{},
})

client.List("some-property-id") // [{ "id": "some-id", "name": "some-env-name", "legacy_account_number": "", "default_domain_name": "", "dns_domain_name": "", "can_members_deploy": true, "only_maintainers_can_deploy": true, "http_request_logging": true, "pci_compliance": true, "created_at": "2019-08-24T14:15:22Z", "updated_at": "2019-08-24T14:15:22Z" }]
```

This func list environments for a given Edgio Property. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all environments until actual pagination is implemented. Returns a list of environments for a given Property or an error if anything goes wrong.

### `env.List` Mandatory Params

- `propertyID`: Property ID from the property which owns the environments to be retrieved

### `env.List` Optional Params & Default Values

There is no optional parameters for that function

<p align="right"><em><a href="../#edgio-organizations-api">back to the main README</a></em></p>
