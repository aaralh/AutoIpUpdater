# autoipupdater
Automated IP updater for servers with dynamic IPs. Script checks if IP has been changed and will perform update to dynamic DNS provider.

## How to use
1. Compile `main.go`file with `go build main.go` command.
2. Create `config.json` file in the same folder.

Structure for config:
```JSON
{
  "username": "Username for dynamic DNS service.",
  "password": "Password for dynamic DNS service.",
  "wgetUrl": "Url to update the changed IP to dynamic DNS provider.",
  "checkIpUrl": "Url which returns the current ip."
}
```
Example provider for current ip https://ipify.org.

3. Run compiled code.
