# autoipupdater
Automated IP updater for servers with dynamic IPs. Script checks if IP has been changed and will perform update to dynamic DNS provider. Works out of the box with [dy.fi](https://www.dy.fi) DNS.

## How to use
1. Compile `main.go`file with `go build main.go` command.
2. Create `config.json` file in the same folder.

Structure for config:
```JSON
{
  "username": "Username for dynamic DNS service.",
  "password": "Password for dynamic DNS service.",
  "wgetUrl": ["Urls in list which to be updated with the changed IP."],
  "checkIpUrl": "Url which returns the current ip."
}
```
Example provider for current ip https://www.ipify.org.

3. Run compiled code.
