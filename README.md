# msgraph.go

[![](https://github.com/yaegashi/msgraph.go/workflows/go%20generate%20test/badge.svg?branch=master)](https://github.com/yaegashi/msgraph.go/actions)

|v1.0|beta|
|---|---|
|[![go.dev](https://img.shields.io/badge/go.dev-reference-000000?logo=go)](https://pkg.go.dev/github.com/yaegashi/msgraph.go/v1.0)|[![go.dev](https://img.shields.io/badge/go.dev-reference_(missing)-000000?logo=go)](https://pkg.go.dev/github.com/yaegashi/msgraph.go/beta)|

(Online reference generation is broken due to huge number of files in the package)

## Introduction 

[Microsoft Graph] client library for Go.  Still in PoC or pre-alpha stage.
Don't use in production.

The library code is auto-generated from the REST API specification
available at https://graph.microsoft.com/v1.0/$metadata.

The code generator is written in pure Go,
in contrast to [the official code generator][Microsoft Graph SDK Code Generator]
heavily relying on C# and non-portable .NET Framework.

## v0.x.x releases

See [GitHub releases](https://github.com/yaegashi/msgraph.go/releases)
for all release tags and release notes,
and [pkg.go.dev](https://pkg.go.dev/mod/github.com/yaegashi/msgraph.go)
for all Go module versions available for your applications.

Until v1.0.0, all types of changes might be included in every release:
bug fixes, new features, even incompatible API updates.

## Usage

You can choose API version when importing `msgraph` package:

```go
import msgraph "github.com/yaegashi/msgraph.go/v1.0"
```

```go
import msgraph "github.com/yaegashi/msgraph.go/beta"
```

You could benefit from better IDE assisted coding experience
because Graph API specs are completely translated to Go codes by msgraph.go.

![](assets/msgraph.go-vscode2.gif)

Code examples in the repository:

- [cmd/msgraph-me](cmd/msgraph-me): Show the profile of signed in user (me) and download files in the root folder of their drive
- [cmd/msgraph-usergroup](cmd/msgraph-usergroup): Graph user/group manipulation example
- [cmd/msgraph-sshpubkey](cmd/msgraph-sshpubkey): Manage SSH public keys in the open extension of graph user resources
- [cmd/msgraph-spoget](cmd/msgraph-spoget): Download a file with SharePoint Online URL

## Hacking

Run `go generate ./gen` to download the metadata and generate library code from it.

You need `goimports` to format outputs and fix imports of them.

```console
$ go get golang.org/x/tools/cmd/goimports
$ rm gen/*.xml
$ go generate ./gen
2019/10/23 03:56:06 Downloading https://graph.microsoft.com/v1.0/$metadata to metadata-v1.0.xml
2019/10/23 03:56:07 Downloading https://graph.microsoft.com/beta/$metadata to metadata-beta.xml
2019/10/23 03:56:07 Creating directory ../v1.0
2019/10/23 03:56:07 Removing ../v1.0/InstallIntentEnum.go
2019/10/23 03:56:07 Removing ../v1.0/EditionUpgradeConfigurationModel.go
2019/10/23 03:56:07 Removing ../v1.0/AssignedPlanModel.go
...
2019/10/23 03:56:07 Creating ../v1.0/msgraph.go
2019/10/23 03:56:07 Creating ../v1.0/ActionStateEnum.go
2019/10/23 03:56:07 Creating ../v1.0/ActivityDomainEnum.go
...
2019/10/23 03:56:10 Formatting ../v1.0/ContentTypeModel.go ../v1.0/AuditLogRootRequest.go ../v1.0/DeviceComplianceScheduledActionForRuleRequest.go ...
```

## Todo

- [x] Save indented metadata.xml
- [x] Support Action elements in metadata
- [ ] Support Function elements in metadata
- [ ] Support batch requests
- [x] Access to additional properties like `@odata.type` `@odata.id`
- [x] Split output into multiple files
- [x] Generate camel cases in golang manner (`IpAddress` -> `IPAddress`)
- [x] Provide easy way to generate pointers to literals
- [x] Provide easy way to generate pointers to constants
- [x] Provide easy way to add queries like `$expand` `$select` `$filter`
- [x] Every request method should take a ctx as the first arg for better control
- [ ] Online API docs (the output is too big for pkg.go.dev to handle)
- [ ] Unit tests
- [x] CI
- [x] Persist OAuth2 tokens in file
- [ ] Persist OAuth2 tokens in object storage like Azure Blob
- [x] OAuth2 device auth grant
- [x] OAuth2 client credentials grant
- [x] Use string for EnumType (pointed out in #6)
- [x] Reduce number of generated files (#11)
- [x] Provide easy way to add HTTP headers like `Prefer: outlook.timezone="Tokyo Standard Time"`
- [ ] Support max number of pages to retrieve from a collection
- [ ] Support Windows time zone names used in DateTimeTimeZone (incorporate [genzabbrs.go](https://golang.org/src/time/genzabbrs.go))

## References

- [Microsoft Graph]
- [Microsoft Graph REST API reference]
- [Microsoft Graph SDKs - Requirements and Design]
- [Microsoft Graph SDK Code Generator]
- [GitHub repository search for msgraph in Go]
- [Microsoft Graph API Library for Go] (presentation in Japanese)
- [msgraph.go demo - SharePoint Online + Microsoft Flow + GitLab CI] (screencast in YouTube)

[Microsoft Graph]: https://developer.microsoft.com/en-us/graph
[Microsoft Graph REST API reference]: https://docs.microsoft.com/en-us/graph/api/overview
[Microsoft Graph SDKs - Requirements and Design]: https://microsoftgraph.github.io/msgraph-sdk-design/
[Microsoft Graph SDK Code Generator]: https://github.com/microsoftgraph/MSGraph-SDK-Code-Generator
[GitHub repository search for msgraph in Go]: https://github.com/search?l=Go&q=msgraph&type=Repositories
[Microsoft Graph API Library for Go]: https://www.slideshare.net/yaegashi/microsoft-graph-api-library-for-go
[msgraph.go demo - SharePoint Online + Microsoft Flow + GitLab CI]: https://www.youtube.com/watch?v=DwKk405XyF4

## Applications

- [Terraform Provider for Microsoft Graph](https://github.com/yaegashi/terraform-provider-msgraph)
