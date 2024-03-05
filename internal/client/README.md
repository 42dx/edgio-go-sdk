# internal/client

This package provides a base client configuration and connection to Edgio's REST API, as well as configuration validation. The public packages under `edgio` namespace uses this client under the hood to perform their API calls.

## `client.New(creds common.Creds, config common.ClientConfig) (client.Client, error)`

This constructor validates and assing default valued (if applicable) to the provided credentials and configurations and returns a client instance with a valid access token, or an error if anything goes wrong.

```go
  client.New(
    common.Creds{Key: string, Secret: string, Scopes: string, AuthURL string},
    common.ClientConfig{Url: string, ApiVersion string, Service: string, Scope: string, OrgId: string}
)
```

### `client.New` Mandatory Params

- `creds common.Creds`
  - `common.Creds.Key`: Edgio REST API Client Key. This is the API Client's identifier on Edgio. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
  - `common.Creds.Secret`: Edgio REST API Client Secret. This is the API Client' secret, unique for each API Client. You can generate your api client accessing `https://edgio.app/<your-org-name>/clients`.
- `config common.ClientConfig`
  - `common.ClientConfig.ApiVersion`: Intended REST API Version. Each one of the Edgio REST APIs has its own Version, that must be provided when creating the client.
  - `common.ClientConfig.Service`: Intended REST API Service. Each one of the Edgio REST APIs has its own Service, that must be provided when creating the client.
  - `common.ClientConfig.Scope`: Intended REST API Scope. Each one of the Edgio REST APIs has its own Scope, that must be provided when creating the client.

### `client.New` Optional Params & Default Values

- `creds common.Creds`
  - `common.Creds.Scopes`: Edgio REST API Client scopes requested by the client. Different APIs needs different scopes. Refer to the [REST API docs](https://docs.edg.io/rest_api) to figure which ones you need.
    - Default value: `app.cache+app.cache.purge+app.waf+app.waf:edit+app.waf:read+app.accounts+app.config` (all scopes).
  - `common.Creds.AuthURL`: Edgio REST API auth url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://id.edgio.app/connect/token` (Edgio's default auth API url).
- `config common.ClientConfig`
  - `common.ClientConfig.Url`: Edgio REST API resources url. You will probably never need to change this, but we included the option just in case (e. g.: future enterprise self-hosted option).
    - Default value: `https://edgioapis.com` (Edgio's default resources API url).

## `client.GetServiceURL(params common.URLParams) string`

This function generates the fully formatted Edgio REST API's url for the desired resource, identified by its `service`, `scope` and `apiVersion`.

```go
// edgioClient returned from  client.New()
orgPropertyUrl := edgioClient.GetServiceURL(common.URLParams{Path: "your-property-id"}) // https://edgioapis.com/accounts/v0.1/properties/your-property-id
```

### `client.GetServiceURL` Mandatory Params

Since this function inherits all its configuration from the created (by the `client.New` func) client, there are no mandatory parameters. All of them were already validated on the creation of the client.

### `client.GetServiceURL` Optional Params & Default Values

- `params common.URLParams`
  - `common.URLParams.Path`: Path of the resource you need. Some of the Edgio REST API's endpoint require them, some do not. You will need to check on [its docs](https://docs.edg.io/rest_api) to make sure you include this param when needed.

<p align="right"><em><a href="../../#client">back to the main README</a></em></p>
