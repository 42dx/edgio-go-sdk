# edgio/property

This package groups Edgio Property specific funcs.

## `property.NewClient(params ClientParams) (ClientStruct, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

propertyClient, err := property.NewClient(property.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{OrgID: "some-org-id", AccessToken: "some-access-token"},
})
```

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's properties.

### `property.NewClient` Mandatory Params

- `params.Credentials`
  - `params.Credentials.Key`: Edgio REST API Client Key. This is the API Client's identifier on Edgio. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
  - `params.Credentials.Secret`: Edgio REST API Client Secret. This is the API Client' secret, unique for each API Client. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
- `params.Config`
  - `params.Config.OrdID`: Edgio's OrganizationID where the properties-to-be-accessed lives under.

### `property.NewClient` Optional Params & Default Values

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

## `property.List() (ListResultType, error)`

```go
credentials := common.Creds{
  Key:    "some-api-key",
  Secret: "some-api-secret",
}

client, err := property.NewClient(property.ClientParams{
  Credentials: credentials,
  Config:      common.ClientConfig{OrgID: "some-org-id"},
})

client.List(common.URLParams{Path: "some-org-id"}) // [{ "ID": "prop-id", "Slug": "prop-slug", "CreatedAt": "2019-08-24T14:15:22Z", "UpdatedAt": "2019-08-24T14:15:22Z" }]
```

This func lists properties for a given Edgio Organization. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all properties until actual pagination is implemented. Returns a list of properties for a given Organization or an error if anything goes wrong.

### `property.List` Mandatory Params

There is no mandatory parameters for that function

### `property.List` Optional Params & Default Values

There is no optional parameters for that function

<p align="right"><em><a href="../#edgio-organizations-api">back to the main README</a></em></p>
