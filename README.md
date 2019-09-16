# msgraph.go

|v1.0|beta|
|---|---|
|[![GoDoc](https://godoc.org/github.com/yaegashi/msgraph.go/v1.0?status.svg)](https://godoc.org/github.com/yaegashi/msgraph.go/v1.0)|[![GoDoc](https://godoc.org/github.com/yaegashi/msgraph.go/beta?status.svg)](https://godoc.org/github.com/yaegashi/msgraph.go/beta)|

## Introduction 

[Microsoft Graph] client library for Go.  It's still in PoC or pre-alpha stage.
Don't use in production.

The library code is auto-generated from the REST API specification
available at https://graph.microsoft.com/v1.0/$metadata.

The code generator is written in pure Go,
in contrast to [the official code generator][Microsoft Graph SDK Code Generator]
heavily relying on C# and non-portable .NET Framework.

## Usage

[cmd/msgraph-me/main.go](cmd/msgraph-me/main.go) shows simple usage example.
It contains [Azure AD v2 device code flow](https://docs.microsoft.com/ja-jp/azure/active-directory/develop/v2-oauth2-device-code)
for the user authentication.
You can authenticate yourself with both pernsonal (Microsoft) account and organizational account.

```console
$ git clone https://github.com/yaegashi/msgraph.go
$ cd msgraph.go
$ go get ./cmd/msgraph-me
$ msgraph-me
To sign in, use a web browser to open the page https://microsoft.com/devicelogin and enter the code HN9KG98HS to authenticate.
2019/09/16 14:02:14 GET https://graph.microsoft.com/v1.0/me
{
  "id": "126871095554adb4",
  "displayName": "八重樫 剛史",
  "givenName": "剛史",
  "surname": "八重樫",
  "userPrincipalName": "yaegashi@live.jp"
}
2019/09/16 14:02:14 GET https://graph.microsoft.com/v1.0/me/drive/root/children?%24filter=file+ne+null&%24select=name%2Cfile%2Csize%2CwebUrl
[
  {
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
```

## Todo

- [ ] Process Action elements in metadata
- [ ] Split output into multiple files
- [ ] Generate camel cases in golang manner (`IpAddress` -> `IPAddress`)
- [ ] Improve API documentation in godoc.org
- [ ] Support more OAuth2 flows
- [ ] Persist OAuth2 tokens in file

## References

- [Microsoft Graph]
- [Microsoft Graph REST API reference]
- [Microsoft Graph SDKs - Requirements and Design]
- [Microsoft Graph SDK Code Generator]

[Microsoft Graph]: https://developer.microsoft.com/en-us/graph
[Microsoft Graph REST API reference]: https://docs.microsoft.com/en-us/graph/api/overview
[Microsoft Graph SDKs - Requirements and Design]: https://microsoftgraph.github.io/msgraph-sdk-design/
[Microsoft Graph SDK Code Generator]: https://github.com/microsoftgraph/MSGraph-SDK-Code-Generator
