# Edgio Go SDK

This project goal is to implement a GO SDK wrapper for the [Edgio's REST API](https://docs.edg.io/rest_api). This SDK is a starting point for more advanced projects like a CLI and a [Terraform Provider](https://developer.hashicorp.com/terraform/language/providers).

## Project Standards

![License](https://img.shields.io/github/license/42dx/edgio-go-sdk?logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGZpbGw9Im5vbmUiIHN0cm9rZT0iI0ZGRiIgdmlld0JveD0iMCAwIDI0IDI0Ij48cGF0aCBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiIHN0cm9rZS13aWR0aD0iMiIgZD0ibTMgNiAzIDFtMCAwLTMgOWE1LjAwMiA1LjAwMiAwIDAgMCA2LjAwMSAwTTYgN2wzIDlNNiA3bDYtMm02IDIgMy0xbS0zIDEtMyA5YTUuMDAyIDUuMDAyIDAgMCAwIDYuMDAxIDBNMTggN2wzIDltLTMtOS02LTJtMC0ydjJtMCAxNlY1bTAgMTZIOW0zIDBoMyIvPjwvc3ZnPg==)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/42dx/edgio-go-sdk?logo=go)
[![semantic-release](https://img.shields.io/badge/Semantic_Release-Conventional_Commits-77f?logo=semantic-release)](https://github.com/semantic-release/semantic-release)
[![conventional-commits](https://img.shields.io/badge/Conventional_Commits-1.0.0-blue.svg?logo=conventionalcommits)](https://conventionalcommits.org)
![Static Badge](https://img.shields.io/badge/Code_Style-gofmt-blue)

## Project Meta Data

[![gh-tag](https://img.shields.io/github/v/tag/42dx/edgio-go-sdk?logo=github&label=Latest%20Version&color=orange)](https://github.com/42dx/edgio-go-sdk/releases)
![GitHub Downloads](https://img.shields.io/github/downloads/42dx/edgio-go-sdk/total?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAACXBIWXMAAAsTAAALEwEAmpwYAAABvUlEQVR4nO3Zu48OURjA4dcqVuJWIpG4lKhcCkJJoiIaGxUtJSqXP4FGQiEioaG00SoEJVFJBKFDIrRIPDKdfI7Z2Z0Z+ebMecqZ5GTeX3Nm5kQURVEUxX+Ftf62JsYC2xMBtsVY4FAiwMEYC5xKBDgZY4FLiQAXYyxwIxHgeowF5hMBHsRY4EUiwPMYC3xOBPgUY4BZ/EoEqK7NRu6w1b9tidzhQE2A/ZE7zNUEOB65w7maAGcjd7haE+BK5A73awLci9zhWU2Ap5E7fKgJ8D5yhhn8qAlQ3ZuJXGGDha2PXGF3gwC7Ilc42iDAkcgVzjQIcDqGDHf1505MO6zC6x6GfzeYswPsWWC7W6yf2BcZ/Pldqssx0JeeR9p7guUxRNiILy2G/4bNMWQ41iLAicgBbi5h+NuRC6zEq0UM/3YwW15T2InvDbe8vZEjnG8Q4EJMExNarrUMD2uGf9x2y+vyeXtZEOvwMTH8V2zqYP3pDlDB4cTx2Fx0YBABKrj2x7K3oiNDCrACL/EGq0cXoIId1ZdjdGhQAfpQAkyYvqI9KwEmTF/RnpUAEzovOjQlQFsGrgQoiqIoiqKItN/2QY+dnO4r2wAAAABJRU5ErkJggg==&color=7777ff&label=Downloads)
[![Repo's Stars](https://img.shields.io/github/stars/42dx/edgio-go-sdk?style=flat&color=yellow&label=Repo%20Stars&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAACXBIWXMAAAsTAAALEwEAmpwYAAAEJ0lEQVR4nO2ZvW8bdRzGf1Q09r357MT3YikTZSgIwdAWCZA6MYBQ/oIKykCHDlUZujcUlQ2pVMAYJFQQUgckhKCxHb/cxU7sOunZvoiKdkCZUFGGVlS0gd/vQXcXv9upae4SG/mRnsny3ef5+vHd92xCJppoosAEW3sVdf0EGUfB0lTY2kPXlqaScRPq2jxsHag71ubJOAm/Ph+Crf+OuubAAzX1HoqzHBkXwU6cacHvuKp+QMZBAHkGNXWjCV5THXjAUm87r5FRF6r6221Tb8ADluL4LTLqQk1L9YW/pQDr8UUyyoKlvQRLZX3hvQDAuvIKGVXBUr/qqkwXfByozCyQkb1xWcpfPVNvh19zAkw/QiWeIKMmWMrH/SvTBr8243wCji+RURKKsxys+B994RvgLXjg5vQWLE04uOu8reioxY/B0uZgKWdhKdcGVqYd3IP3XJ6+hlLsLMrROazKx1BSdF/uE6grR7Chn4SdOIW6fgE1/Qpq6nVUtQKq6iaq6vbuV5kB8A1wDx4ox4CSZ7YadU1Xotu0KG/SFblAi5HrtBi5goJ0AcuRUyiIJ2HKz+0Ov5H4yF3AWktY9zrQfWMatjKd4H3g2UrDMlhxx4WI5+WGJcAQLw4OYOunUddZE757Heg39eEr0wHeCy+3wHvgJTBTAjUkBlN67wkV0t9FXdvumXoAlWH94Bvg3fCm9A8M8cxw34Oa+iaq2oNRqAxzJy8+RF56Zyj4Zoh1/QSqyr2DrAzz4LdgiG/8J/hWnZQjsOJ3DqIyzIHPi78hKx59KvhmCOf6vB6/5U9l5KEqwzx4G1ludk/wzRC2ImItntyPyjAXXsghFZN9ge94zq3MfBdkZZghOvDfo0i44NaIysynQVSGOfCG+DlADgUC3xGkHLvqZ2WYO3nxauDgzQCl2Cd+VYbtmOaFy/sXoBxb8qMyzHHeM83x6f2BBzmEUvT+XivTghfAcgJoln+Ai/vR/9XpF/dcmXwnPHOc5YGc+ELwAVbk9/2oTBM858G7ATL86cAD0FX5Sz8q44ELTXiW4UGXwl8EH6AYrfhVGbYD7pkDTYdvBgqPLAnTovx44NQ7dhlhC4ZwDlnhHM0JW/0qwxrwS5xrusQ9ds4RXICC9PoT19+8RGle/BoFofmnBkw5RrP8ZzQj/N0JzzXhWdpxGEiFXwsuQFH+cNf11xAzMISXB74/Kx6lGe6nXviw51QYSE6dDywAXZa+7QtuSpsw+d2fVduDpLk5mubudsM7pouhb4ILUJDudj6nin/SvDT/NL1FhRxGeuo8TYfve/AhsGQINDl1JxB4t8dmhDnw1JAoNcUF5Pk9/8aJn/kETU0t0GSIugFuhBh+JDF/qNtPtKgJ1BBvU1PKYTly3PfjJ589Tm9M5eji4V/wA+H9Pv5EE/1f9S9M4cKBWKiMUAAAAABJRU5ErkJggg==)](https://github.com/42dx/edgio-go-sdk)
[![All Contributors](https://img.shields.io/github/all-contributors/42dx/edgio-go-sdk/beta?color=ee8449&style=flat&label=Contributors)](https://github.com/42dx/edgio-go-sdk/blob/beta/README.md#contributors)

## Project Status

![GitHub milestone details](https://img.shields.io/github/milestones/progress-percent/42dx/edgio-go-sdk/1?label=MVP%20Milestone)
![GitHub milestone details](https://img.shields.io/github/milestones/progress-percent/42dx/edgio-go-sdk/2?label=V1%20Milestone)

## Table of Contents

- [Internal Packages](#internal-packages)
  - [Client](#client)
  - [Token](#token)
  - [Utils](#utils)
- [Public Packages](#public-packages)
  - [Edgio Common Lib](#edgio-common-lib)
  - [Edgio Organizations API](#edgio-organizations-api)
  - [Edgio Properties API](#edgio-properties-api)
  - [Edgio Environments API](#edgio-environments-api)
  - [Edgio Environment Variables API](#edgio-environment-variables-api)
  - [Edgio Cache API](#edgio-cache-api)
  - [Edgio CDN API](#edgio-cdn-api)
  - [Edgio Deployment API](#edgio-deployment-api)
  - [Edgio TSL Certificates API](#edgio-tsl-certificates-api)
  - [Edgio ACL API](#edgio-acl-api)
  - [Edgio Secure Ruleset API](#edgio-security-ruleset-api)
  - [Edgio Schemas API](#edgio-schemas-api)
  - [Rate Rules API](#rate-rules-api)
  - [Bot Config API](#bot-managers-config-api)
  - [Bot Ruleset API](#bot-ruleset-api)
  - [Known Bots API](#known-bots-api)
  - [Custom Rules API](#custom-rules-api)
  - [Managed Rules API](#managed-rules-api)
  - [Edgio Ruleset API](#edgio-ruleset-api)
  - [Security API](#security-apps-api)
- [Contributors](#contributors)
- [Changelog](#changelog)
- [Roadmap](#roadmap)

## Internal Packages

The internal package documentation is intended to potential contributors of the repository, since they are not exposed to be directly imported. If you are a user, you shoud use our [public api-specific packages](#public-packages) to develop your application.

### Client

This package provides a base client configuration and connection to Edgio's REST API, as well as configuration validation. The public packages under `edgio` namespace uses this client under the hood to perform their API calls.

#### `client.New(creds common.Creds, config common.ClientConfig) (client.Client, error)`

This constructor validates and assing default valued (if applicable) to the provided credentials and configurations and returns a client instance with a valid access token, or an error if anything goes wrong.

```go
client.New(
  common.Creds{Key: string, Secret: string, Scopes: string, AuthUrl: string},
  common.ClientConfig{Url: string, ApiVersion: string, Service: string, Scope: string, OrgId: string}
)
```

#### `GetServiceUrl(params common.UrlParams) string`

This function generates the fully formatted Edgio REST API's url for the desired resource, identified by its `service`, `scope` and `apiVersion`.

```go
edgioClient := client.New(
  common.Creds{Key: "your-api-key", Secret: "your-api-secret", Scopes: "scopes"},
  common.ClientConfig{ApiVersion: "v0.1", Service: "accounts", Scope: "properties", OrgId: "your-org-id"}
)

orgPropertyUrl := edgioClient.GetServiceUrl(common.UrlParams{Path: "your-property-id"}) // https://edgioapis.com/accounts/v0.1/properties/your-property-id
```

Check a more in-depth documentation of this package [here](internal/client/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Token

Check a more in-depth documentation of this package [here](internal/token/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Utils

Check a more in-depth documentation of this package [here](internal/utils/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

## Public Packages

### Edgio Common Lib

This package expose public interfaces, structs and other functions that could not be hosted on the internal packages, since in some cases they are required by both other internal packages and some public ones, leading to cyclic imports. Due to that, they were outsourced to their own package.

#### `common.ClientConfig.Merge(other common.ClientConfig{})`

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

Check a more in-depth documentation of this package [here](common/README.md).

[WIP]

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Organizations API

[Edgio Organizations REST API documentation reference](https://docs.edg.io/rest_api/#tag/organizations).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Properties API

[WIP]

[Edgio Properties REST API documentation reference](https://docs.edg.io/rest_api/#tag/properties).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Environments API

[WIP]

[Edgio Environments REST API documentation reference](https://docs.edg.io/rest_api/#tag/environments).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Environment Variables API

[WIP]

[Edgio Environment Variables REST API documentation reference](https://docs.edg.io/rest_api/#tag/environment-variables).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Cache API

[WIP]

[Edgio Cache REST API documentation reference](https://docs.edg.io/rest_api/#tag/purge-requests).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio CDN API

[WIP]

[Edgio CDN REST API documentation reference](https://docs.edg.io/rest_api/#tag/configs).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Deployment API

[WIP]

[Edgio Deployment REST API documentation reference](https://docs.edg.io/rest_api/#tag/deployments).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio TSL Certificates API

[WIP]

[Edgio TSL REST API documentation reference](https://docs.edg.io/rest_api/#tag/tls-certs).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio ACL API

[WIP]

[Edgio ACL REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Access-Control-List-(ACL)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Security Ruleset API

[WIP]

[Edgio Security RulesetREST API documentation reference](https://docs.edg.io/rest_api/#tag/API-Gateways).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Schemas API

[WIP]

[Edgio Schemas REST API documentation reference](https://docs.edg.io/rest_api/#tag/API-Schemas).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Rate Rules API

[WIP]

[Edgio Rate Rules REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Rate-Rules-(Limits)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Bot Managers Config API

[WIP]

[Edgio Managers Config REST API documentation reference](https://docs.edg.io/rest_api/#tag/Bot-Managers).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Bot Ruleset API

[WIP]

[Edgio Bot Ruleset REST API documentation reference](https://docs.edg.io/rest_api/#tag/Bot-Rules).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Known Bots API

[WIP]

[Edgio Known Bots REST API documentation reference](https://docs.edg.io/rest_api/#tag/Known-Bots).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Custom Rules API

[WIP]

[Edgio Custom Rules REST API documentation reference](https://docs.edg.io/rest_api/#tag/Custom-Rules).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Managed Rules API

[WIP]

[Edgio Managed Rules REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Managed-Rules-(Profiles)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Edgio Ruleset API

[WIP]

[Edgio Edgio Rulesets REST API documentation reference](https://docs.edg.io/rest_api/#tag/Edgio-Rulesets).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### Security Apps API

[WIP]

[Edgio Security Apps REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Security-Applications-(Scopes)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

## Contributors

Kudos to all our dear contributors. Without them, nothing would have been possible :heart:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->

<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/rafapaulin"><img src="https://avatars.githubusercontent.com/u/13452406?v=4?s=60" width="60px;" alt="Rafael Eduardo Paulin"/><br /><sub><b>Rafael Eduardo Paulin</b></sub></a><br /><a href="https://github.com/42dx/edgio-go-sdk/commits?author=rafapaulin" title="Code">💻</a> <a href="https://github.com/42dx/edgio-go-sdk/commits?author=rafapaulin" title="Documentation">📖</a> <a href="#ideas-rafapaulin" title="Ideas, Planning, & Feedback">🤔</a> <a href="#infra-rafapaulin" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#maintenance-rafapaulin" title="Maintenance">🚧</a> <a href="#projectManagement-rafapaulin" title="Project Management">📆</a> <a href="https://github.com/42dx/edgio-go-sdk/pulls?q=is%3Apr+reviewed-by%3Arafapaulin" title="Reviewed Pull Requests">👀</a> <a href="#tool-rafapaulin" title="Tools">🔧</a> <a href="https://github.com/42dx/edgio-go-sdk/commits?author=rafapaulin" title="Tests">⚠️</a> <a href="#tutorial-rafapaulin" title="Tutorials">✅</a></td>
    </tr>
  </tbody>
</table>
<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

Would you like to see your profile here? Take a look on our [Code of Conduct](https://github.com/42dx/.github/blob/main/CODE_OF_CONDUCT.md) and our [Contributing](https://github.com/42dx/.github/blob/main/CONTRIBUTING.md) docs, and start coding! We would be thrilled to review a PR of yours! :100:

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

## Changelog

All changes made to this module since the start of development can be found either on our release list or on the [changelog](CHANGELOG.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

## Roadmap

Any planned enhancement to the module will be described and tracked in our [project page](https://github.com/orgs/42dx/projects/1)

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>
