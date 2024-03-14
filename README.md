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
[![GH Issues](https://img.shields.io/github/issues/42dx/edgio-go-sdk?label=Issues)](https://github.com/42dx/edgio-go-sdk/issues)
![GitHub Sponsors](https://img.shields.io/github/sponsors/42dx?label=42dx%20Sponsors)

## Project Status

[![GitHub milestone details](https://img.shields.io/github/milestones/progress-percent/42dx/edgio-go-sdk/1?label=MVP%20Milestone)](https://github.com/42dx/edgio-go-sdk/milestone/1)
[![GitHub milestone details](https://img.shields.io/github/milestones/progress-percent/42dx/edgio-go-sdk/2?label=V1%20Milestone)](https://github.com/42dx/edgio-go-sdk/milestone/2)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=coverage)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=bugs)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=42dx_edgio-go-sdk&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=42dx_edgio-go-sdk)

## Table of Contents

- [Internal Packages](#internal-packages)
  - [internal/client](#internalclient)
  - [internal/token](#internaltoken)
  - [internal/utils](#internalutils)
- [Public Packages](#public-packages)
  - [edgio/common](#edgiocommon)
  - Cache APIs
    - [edgio/cache](#edgiocache)
  - Deployments APIs
    - [edgio/cdn](#edgiocdn)
    - [edgio/variables](#edgiovariables)
    - [edgio/deployment](#edgiodeployment)
  - Organizations APIs
    - [edgio/org](#edgioorg)
  - Properties APIs
    - [edgio/property](#edgioproperty)
    - [edgio/environment](#edgioenvironment)
  - TSL Certificates APIs
    - [edgio/tsl](#edgiotsl)
  - WAF (Security) APIs
    - [edgio/acl](#edgioacl)
    - [edgio/security_ruleset](#edgiosecurity_ruleset)
    - [edgio/schemas](#edgioschemas)
    - [edgio/rate](#edgiorate)
    - [edgio/bot_manager](#edgiobot_manager)
    - [edgio/bot_ruleset](#edgiobot_ruleset)
    - [edgio/bot_known](#edgiobot_known)
    - [edgio/custom_rule](#edgiocustom_rule)
    - [edgio/managed_rule](#edgiomanaged_rule)
    - [edgio/ruleset](#edgioruleset)
    - [edgio/security_app](#edgiosecurity_app)
- [Tooling](#tooling)
- [Contributors](#contributors)
- [Changelog](#changelog)
- [Roadmap](#roadmap)

## Internal Packages

The internal package documentation is intended to potential contributors of the repository, since they are not exposed to be directly imported. If you are a user, you shoud use our [public api-specific packages](#public-packages) to develop your application.

### internal/client

This package provides a base client configuration and connection to Edgio's REST API, as well as configuration validation. The public packages under `edgio` namespace uses this client under the hood to perform their API calls.

#### `client.New(creds common.Creds, config common.ClientConfig) (client.Client, error)`

This constructor validates and assing default valued (if applicable) to the provided credentials and configurations and returns a client instance with a valid access token, or an error if anything goes wrong.

#### `utils.GetServiceURL(params common.URLParams) string`

This function generates the fully formatted Edgio REST API's url for the desired resource, identified by its `service`, `scope` and `apiVersion`.

Check a more in-depth documentation of thw `internal/client` package [here](internal/client/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### internal/token

The main goal of this package is to hold any logic related to the aquisition/refreshing/invalidation generated by Edgio REST API client.

#### `token.GetAccessToken(credentials common.Creds) (string, error)`

This func outsources the process of getting an auth token from Edgio's Auth service. It assumes Edgio's standard auth endpoint as a default URL, but that can be overwritten by your own, in the event of an enterprise/self-hosted app.

Check a more in-depth documentation of the `internal/token` package [here](internal/token/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### internal/utils

This package package holds some utility functions that are used to outsource some common logic from other packages to avoid repetition/ease testing.

#### `utils.GetHttpJsonResult(httpClient *http.Client, request *http.Request, token string, model interface{}) (interface{}, error)`

This function has mainly three related goals: 1- Process http requests; 2- treat HTTP errors in a standardized way, and; 3- Process and decode returned json data from the endpoints.

Check a more in-depth documentation of the `internal/utils` package [here](internal/utils/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

## Public Packages

The public packages are the parts of the SDK that are actually intended to be used. You should be able to import just those you need for your project.

### edgio/common

This package expose public interfaces, structs and other functions that could not be hosted on the internal packages, since in some cases they are required by both other internal packages and some public ones, leading to cyclic imports. Due to that, they were outsourced to their own package.

#### `common.ClientConfig.Merge(other common.ClientConfig{})`

Merges the `other` ClientConfig into the default one, overwriting the current values with the other's if they are not empty.

Check a more in-depth documentation of the edgio/common package [here](common/README.md).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/org

[WIP]

This package groups Edgio Organization specific funcs.

#### `org.NewClient(params ClientParams) (ClientStruct, error)`

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's orgs.

#### `org.Get(params common.URLParams) (getResultType, error)`

This func returns the relevant organization details (name and id).

Check a more in-depth documentation of the `edgio/org` package [here](org/README.md).

**Reference**: [Edgio Organizations REST API documentation reference](https://docs.edg.io/rest_api/#tag/organizations).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/property

[WIP]

This package groups Edgio Property specific funcs.

#### `property.NewClient(params ClientParams) (ClientStruct, error)`

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's properties.

#### `property.List() (ListResultType, error)`

This func lists properties for a given Edgio Organization. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all properties until actual pagination is implemented. Returns a list of properties for a given Organization or an error if anything goes wrong.

Check a more in-depth documentation of the `edgio/property` package [here](property/README.md).

**Reference**: [Edgio Properties REST API documentation reference](https://docs.edg.io/rest_api/#tag/properties).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/environment

[WIP]

This package groups Edgio Environment specific funcs.

#### `env.NewClient(params common.ClientParams) (ClientStruct, error)`

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's environments.

#### `env.List(propertyID string) (ListResultType, error)`

This func list environments for a given Edgio Property. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all environments until actual pagination is implemented. Returns a list of environments for a given Property or an error if anything goes wrong.

Check a more in-depth documentation of the `edgio/property` package [here](env/README.md).

**Reference**: [Edgio Environments REST API documentation reference](https://docs.edg.io/rest_api/#tag/environments).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/variables

[WIP]

This package groups Edgio Environment Variables specific funcs.

#### `variable.NewClient(params common.ClientParams) (ClientStruct, error)`

This func returns a new client with the provided parameters, with an access token already set, and the default edgio params values for service, scope and API version (which can be overwritten if needed) to interact with your application's environments.

#### `variable.List(environmentID string) (ListResultType, error)`

This func list environment variables for a given Edgio Environment. Edgio's list page size was defaulted to 100 for now, which is the highest value. The idea is to return all environment variables until actual pagination is implemented. Returns a list of environment variables for a given Property or an error if anything goes wrong.

Check a more in-depth documentation of the `edgio/variables` package [here](variables/README.md).

**Reference**: [Edgio Environment Variables REST API documentation reference](https://docs.edg.io/rest_api/#tag/environment-variables).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/cache

[WIP]

[Edgio Cache REST API documentation reference](https://docs.edg.io/rest_api/#tag/purge-requests).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/cdn

[WIP]

[Edgio CDN REST API documentation reference](https://docs.edg.io/rest_api/#tag/configs).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/deployment

[WIP]

[Edgio Deployment REST API documentation reference](https://docs.edg.io/rest_api/#tag/deployments).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/tsl

[WIP]

[Edgio TSL REST API documentation reference](https://docs.edg.io/rest_api/#tag/tls-certs).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/acl

[WIP]

[Edgio ACL REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Access-Control-List-(ACL)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/security_ruleset

[WIP]

[Edgio Security Ruleset REST API documentation reference](https://docs.edg.io/rest_api/#tag/API-Gateways).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/schemas

[WIP]

[Edgio Schemas REST API documentation reference](https://docs.edg.io/rest_api/#tag/API-Schemas).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/rate

[WIP]

[Edgio Rate Rules REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Rate-Rules-(Limits)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/bot_manager

[WIP]

[Edgio Managers Config REST API documentation reference](https://docs.edg.io/rest_api/#tag/Bot-Managers).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/bot_ruleset

[WIP]

[Edgio Bot Ruleset REST API documentation reference](https://docs.edg.io/rest_api/#tag/Bot-Rules).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/bot_known

[WIP]

[Edgio Known Bots REST API documentation reference](https://docs.edg.io/rest_api/#tag/Known-Bots).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/custom_rule

[WIP]

[Edgio Custom Rules REST API documentation reference](https://docs.edg.io/rest_api/#tag/Custom-Rules).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/managed_rule

[WIP]

[Edgio Managed Rules REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Managed-Rules-(Profiles)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/ruleset

[WIP]

[Edgio Edgio Rulesets REST API documentation reference](https://docs.edg.io/rest_api/#tag/Edgio-Rulesets).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

### edgio/security_app

[WIP]

[Edgio Security Apps REST API documentation reference](<https://docs.edg.io/rest_api/#tag/Security-Applications-(Scopes)>).

<p align="right"><em><a href="#table-of-contents">back to top</a></em></p>

## Tooling

There are a few tools we provide alongside with the source code to ease a little bit the burden of following a bunch of patterns and standards we set up, as well as automate some boring processes.

- `comitizen`: This CLI helps with writting commit messages in a meaningful and standardized way, so that our automation process can use them to properly write our software, changelog. To use it, you just need to run `./tools/commitizen-go install` from the repository's root folder. After that, you just need to use `git cz` command instead of the standard `git commit` and follow the cli interactive steps :)

## Contributors

Kudos to all our dear contributors. Without them, nothing would have been possible :heart:

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/rafapaulin"><img src="https://avatars.githubusercontent.com/u/13452406?v=4?s=60" width="60px;" alt="Rafael Eduardo Paulin"/><br /><sub><b>Rafael Eduardo Paulin</b></sub></a><br /><a href="https://github.com/42dx/edgio-go-sdk/commits?author=rafapaulin" title="Code">💻</a> <a href="#design-rafapaulin" title="Design">🎨</a> <a href="https://github.com/42dx/edgio-go-sdk/commits?author=rafapaulin" title="Documentation">📖</a> <a href="#ideas-rafapaulin" title="Ideas, Planning, & Feedback">🤔</a> <a href="#infra-rafapaulin" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="#maintenance-rafapaulin" title="Maintenance">🚧</a> <a href="#projectManagement-rafapaulin" title="Project Management">📆</a> <a href="https://github.com/42dx/edgio-go-sdk/pulls?q=is%3Apr+reviewed-by%3Arafapaulin" title="Reviewed Pull Requests">👀</a> <a href="https://github.com/42dx/edgio-go-sdk/commits?author=rafapaulin" title="Tests">⚠️</a> <a href="#tool-rafapaulin" title="Tools">🔧</a> <a href="#tutorial-rafapaulin" title="Tutorials">✅</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/rafaelassumpcao"><img src="https://avatars.githubusercontent.com/u/7454306?v=4?s=60" width="60px;" alt="Rafael A"/><br /><sub><b>Rafael A</b></sub></a><br /><a href="https://github.com/42dx/edgio-go-sdk/pulls?q=is%3Apr+reviewed-by%3Arafaelassumpcao" title="Reviewed Pull Requests">👀</a></td>
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
