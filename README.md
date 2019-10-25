# msgraph.go

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

## Example

[cmd/msgraph-me/main.go](cmd/msgraph-me/main.go) shows a simple example.
It contains [Azure AD v2 device code flow](https://docs.microsoft.com/ja-jp/azure/active-directory/develop/v2-oauth2-device-code) for the user authentication.
You can authenticate yourself with your personal (Microsoft) or organizational (Azure AD) account.

```console
$ go get ./cmd/msgraph-me
$ msgraph-me
To sign in, use a web browser to open the page https://microsoft.com/devicelogin and enter the code GQYZ5KM5H to authenticate.
2019/10/22 01:17:51 Get current logged in user information
2019/10/22 01:17:51 GET https://graph.microsoft.com/v1.0/me
{
  "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users/$entity",
  "id": "126871095554adb4",
  "displayName": "八重樫 剛史",
  "givenName": "剛史",
  "surname": "八重樫",
  "userPrincipalName": "yaegashi@live.jp"
}
2019/10/22 01:17:52 Get files in the root folder of user's drive
2019/10/22 01:17:52 GET https://graph.microsoft.com/v1.0/me/drive/root/children?%24filter=file+ne+null&%24select=id%2Cname%2Cfile%2Csize%2CwebUrl
[
  {
    "id": "126871095554ADB4!1425",
    "name": "ブック.xlsx",
    "webUrl": "https://1drv.ms/x/s!ALStVFUJcWgSixE",
    "file": {
      "hashes": {
        "quickXorHash": "",
        "sha1Hash": "54665588E155AE30FB0D2CA4FC7201A39ACCBE5C"
      },
      "mimeType": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
    },
    "size": 7990
  },
  {
    "id": "126871095554ADB4!1428",
    "name": "プレゼンテーション.pptx",
    "webUrl": "https://1drv.ms/p/s!ALStVFUJcWgSixQ",
    "file": {
      "hashes": {
        "quickXorHash": "5GWMST7DLJqxYS28Avo794YgGOQ=",
        "sha1Hash": "B4E607C06DDEF71C3C523E9993DCA761C85AA5B6"
      },
      "mimeType": "application/vnd.openxmlformats-officedocument.presentationml.presentation"
    },
    "size": 29313
  },
  {
    "id": "126871095554ADB4!1333",
    "name": "文書.docx",
    "webUrl": "https://1drv.ms/w/s!ALStVFUJcWgSijU",
    "file": {
      "hashes": {
        "quickXorHash": "",
        "sha1Hash": "F1CFE569876E83122C54366F918500F6D778B777"
      },
      "mimeType": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    "size": 11393
  }
]
2019/10/22 01:17:52 GET https://graph.microsoft.com/v1.0/me/drive/items/126871095554ADB4!1425
2019/10/22 01:17:53 Download to "ブック.xlsx" (7990 bytes)
2019/10/22 01:17:54 GET https://graph.microsoft.com/v1.0/me/drive/items/126871095554ADB4!1428
2019/10/22 01:17:55 Download to "プレゼンテーション.pptx" (29313 bytes)
2019/10/22 01:17:55 GET https://graph.microsoft.com/v1.0/me/drive/items/126871095554ADB4!1333
2019/10/22 01:17:56 Download to "文書.docx" (11393 bytes)
```

It saves auth tokens in `token_store.json` in the current directory.
You won't be asked for authentication again until tokens in this file expires.

## Code generation

Use `go generate` to download the metadata and generate library code from it.
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