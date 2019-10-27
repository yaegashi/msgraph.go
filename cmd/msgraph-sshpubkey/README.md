# msgraph-sshpubkey

## Introduction

msgraph-sshpubkey is a CLI utility to manage OpenSSH public keys
in the custom property ([open extensions]) of user resources in Azure Active Directory.

On SSH server hosts, it acquires SSH public keys via `AuthorizedKeysCommand`
in /etc/ssh/sshd_config when users want to log in.

It's still a PoC implementation, never use it in production.
There're many security consideration to solve as well as feature and imporvement ideas (see todos below).

## Manage your keys in the directory

By default msgraph-sshpubkey authenticates users using [Azure AD v2 device code flow].
It will always ask you to open https://microsoft.com/devicelogin to enter the code then sign in with your Azure AD account.
You can bypass sign-ins by specifying `-token-store /path/to/token.json`.

Set SSH public keys with a file:

```console
$ msgraph-sshpubkey -op set -in ~/.ssh/id_rsa.pub
```

Set SSH public keys with keys kept in the SSH agent:

```console
$ ssh-add -L | msgraph-sshpubkey -op set -in -
```

Get SSH public keys (you can omit `-op get`):

```console
$ msgraph-sshpubkey
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDTv2zRefcFMXafSRneDlULwCPh0v7SM9rJPIlySgd8WEJwk3/bY4B6j6hMPk3xS/JAqvQG0hc5cRSSmo4tG9H7TDjmGKBptIsGr5skTx181nbv/qRLYrej80KFrKyt2yHxg7BFOMGDSG1RnRVDUQJxlYxluavky0dv3KGRt6TtDuzuLGi6flHcqJymlZleqprEEwZwc0ju/ZNBfpEW2A+e69nJkudgT8jsO3a61iQ9myf7Jdk/0dxHPoHhu2VWEv/YcFPr0OX5fp7OHVL56vYb6yQVSVp1MtqjqSLpSK+O1eEGnwLsI9/93DXUj3gFncqjddgD75SQ1N9e1DPYK9sz /Users/yaegashi/.ssh/id_rsa
```

Delete SSH public keys:

```console
$ msgraph-sshpubkey -op delete
```

## Authenticate users on SSH server hosts

First, you have to register the application with client credentials grant on Azure Portal.
It will need `User.Read.All` permission to access exntension properties of users in the directory.

On SSH server hosts, prepare /etc/msgraph-sshpubkey.json with something like the following in it:

```json
{
  "tenant_id": "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX",
  "client_id": "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY",
  "client_secret": "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ",
  "login_map": {
    "admin": "admin@l0wdev.onmicrosoft.com",
    "yaegashi": "yaegashi@l0wdev.onmicrosoft.com",
    "takeshi": "takeshi@l0wdev.onmicrosoft.com"
  }
}
```

Run msgraph-sshpubkey on the shell to see it can certainly retrieve SSH public keys that user has registered:

```console
$ msgraph-sshpubkey -config /etc/msgraph-sshpubkey.json -login yaegashi
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDTv2zRefcFMXafSRneDlULwCPh0v7SM9rJPIlySgd8WEJwk3/bY4B6j6hMPk3xS/JAqvQG0hc5cRSSmo4tG9H7TDjmGKBptIsGr5skTx181nbv/qRLYrej80KFrKyt2yHxg7BFOMGDSG1RnRVDUQJxlYxluavky0dv3KGRt6TtDuzuLGi6flHcqJymlZleqprEEwZwc0ju/ZNBfpEW2A+e69nJkudgT8jsO3a61iQ9myf7Jdk/0dxHPoHhu2VWEv/YcFPr0OX5fp7OHVL56vYb6yQVSVp1MtqjqSLpSK+O1eEGnwLsI9/93DXUj3gFncqjddgD75SQ1N9e1DPYK9sz /Users/yaegashi/.ssh/id_rsa
```

Put the following lines in /etc/ssh/sshd_config then reload sshd:

```
AuthorizedKeysCommand /usr/bin/msgraph-sshpubkey -config /etc/msgraph-sshpubkey.json -login %u
AuthorizedKeysCommandUser root
```

## Todo

- [ ] Provide a reasonable way to translate from SSH login names to Azure AD user principal names (`login_map` is cumbersome)
- [ ] Azure AD group authorization
- [ ] SSH public key validation
- [ ] File permission sanity check (config, token store)
- [ ] Web app to manage keys
- [ ] Caching to improve performance
- [ ] Logging
- [ ] [Managed identity][Managed identities] integration on Azure VMs

## Referenecs

- [Open extensions]
- [Managed identities]
- [Azure Active Directory authentication for Linux]

[Open extensions]: https://docs.microsoft.com/en-us/graph/extensibility-open-users
[Managed identities]: https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview
[Azure AD v2 device code flow]: https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-device-code
[Azure Active Directory authentication for Linux]: https://docs.microsoft.com/en-us/azure/virtual-machines/linux/login-using-aad