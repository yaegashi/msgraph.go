# msgraph-me

This example shows user's profile of signed in user, then downloads files in the root folder of user's drive.

```console
$ msgraph-me
To sign in, use a web browser to open the page https://microsoft.com/devicelogin and enter the code H5N62P7KQ to authenticate.
2019/12/07 02:39:21 Get current logged in user information
2019/12/07 02:39:21 GET https://graph.microsoft.com/v1.0/me
{
  "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users/$entity",
  "id": "6764eb11-841c-444e-b770-0e0d8748ea0a",
  "displayName": "八重樫 剛史",
  "givenName": "剛史",
  "mail": "yaegashi@l0wdev.onmicrosoft.com",
  "surname": "八重樫",
  "userPrincipalName": "yaegashi@l0wdev.onmicrosoft.com"
}
2019/12/07 02:39:21 Get files in the root folder of user's drive
2019/12/07 02:39:21 GET https://graph.microsoft.com/v1.0/me/drive/root/children
2019/12/07 02:39:21   FILE 2019-10-27T07:57:28Z       7972 Book.xlsx
2019/12/07 02:39:21   FILE 2019-10-27T07:57:19Z      11395 ドキュメント.docx
2019/12/07 02:39:21   FILE 2019-10-27T07:57:36Z      29014 プレゼンテーション.pptx
Press ENTER to download files or Ctrl-C to abort: 
2019/12/07 02:39:25 Download "Book.xlsx" (7972 bytes)
2019/12/07 02:39:25 Download "ドキュメント.docx" (11395 bytes)
2019/12/07 02:39:25 Download "プレゼンテーション.pptx" (29014 bytes)
```

It uses [Azure AD v2 device code flow] for the user authentication.
You can authenticate yourself with your personal (Microsoft) or organizational (Azure AD) account.

It saves auth tokens in `token_cache.json` in the current directory.
You won't be asked for authentication again until tokens in this file expires.

## Reference

- [Azure AD v2 device code flow]

[Azure AD v2 device code flow]: https://docs.microsoft.com/ja-jp/azure/active-directory/develop/v2-oauth2-device-code