# msgraph.go

(Online reference generation is broken due to huge number of files in the package)

|v1.0|beta|
|---|---|
|[![GoDoc](https://godoc.org/github.com/yaegashi/msgraph.go/v1.0?status.svg)](https://godoc.org/github.com/yaegashi/msgraph.go/v1.0)|[![GoDoc](https://godoc.org/github.com/yaegashi/msgraph.go/beta?status.svg)](https://godoc.org/github.com/yaegashi/msgraph.go/beta)|

[![](https://github.com/yaegashi/msgraph.go/workflows/go%20generate%20test/badge.svg?branch=master)](https://github.com/yaegashi/msgraph.go/actions)

## Introduction 

[Microsoft Graph] client library for Go.  Still in PoC or pre-alpha stage.
Don't use in production.

The library code is auto-generated from the REST API specification
available at https://graph.microsoft.com/v1.0/$metadata.

The code generator is written in pure Go,
in contrast to [the official code generator][Microsoft Graph SDK Code Generator]
heavily relying on C# and non-portable .NET Framework.

## Usage

You can choose API version when importing `msgraph` package:

```go
import msgraph "github.com/yaegashi/msgraph.go/v1.0"
```

```go
import msgraph "github.com/yaegashi/msgraph.go/beta"
```

## Examples

- [cmd/msgraph-me](cmd/msgraph-me): Show the profile of signed in user (me) and download files in the root folder of their drive
- [cmd/msgraph-usergroup](cmd/msgraph-usergroup): Graph user/group manipulation example
- [cmd/msgraph-sshpubkey](cmd/msgraph-sshpubkey): Manage SSH public keys in the open extension of graph user resources

## Hacks

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
- [ ] Online API docs (the output is too big for godoc.org to handle)
- [ ] Unit tests
- [x] CI
- [x] Persist OAuth2 tokens in file
- [x] OAuth2 device auth grant
- [x] OAuth2 client credentials grant
- [ ] OAuth2 authorization code grant

## References

- [Microsoft Graph]
- [Microsoft Graph REST API reference]
- [Microsoft Graph SDKs - Requirements and Design]
- [Microsoft Graph SDK Code Generator]
- [GitHub repository search for msgraph in Go]

[Microsoft Graph]: https://developer.microsoft.com/en-us/graph
[Microsoft Graph REST API reference]: https://docs.microsoft.com/en-us/graph/api/overview
[Microsoft Graph SDKs - Requirements and Design]: https://microsoftgraph.github.io/msgraph-sdk-design/
[Microsoft Graph SDK Code Generator]: https://github.com/microsoftgraph/MSGraph-SDK-Code-Generator
[GitHub repository search for msgraph in Go]: https://github.com/search?l=Go&q=msgraph&type=Repositories