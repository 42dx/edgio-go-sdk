# edgio/common

This package expose public interfaces, structs and other functions that could not be hosted on the internal packages, since in some cases they are required by both other internal packages and some public ones, leading to cyclic imports. Due to that, they were outsourced to their own package.

## `common.ClientConfig.Merge(other common.ClientConfig{}) ClientConfig`

Merges the `other` ClientConfig into the default one, overwriting the current values with the other's if they are not empty.

```go
baseConfig := common.ClientConfig{
  Url: "a",
  ApiVersion: "a",
  Service: "a",
  Scope: "a",
  OrgId: "a",
}

newConfig := common.ClientConfig{
  ApiVersion: "b",
  Service: "b",
}

baseConfig.Merge(newConfig) // returns common.ClientConfig{Url: "a", ApiVersion: "b", Service: "b", Scope: "a", OrgId:  "a"}
```

### `common.ClientConfig.Merge` Mandatory Params

- `other common.ClientConfig`: An `other` param must be provided to this func, even though no `common.ClientConfig` key is mandatory. If no `common.ClientConfig` is provided, the default value will be returned.

### `common.ClientConfig.Merge` Optional Params & Default Values

There is no optional parameters for that function

<p align="right"><em><a href="../#edgio-common-lib">back to the main README</a></em></p>
