# Cloudflare Dynamic DNS Tool
cfdns is a small tool used to update dynamic DNS entries on Cloudflare.

## Instructions

### Standalone
 1. Download the latest release binary from https://github.com/someone-stole-my-name/cfdns/releases
 2. Create a json config file with your Cloudflare credentials, zones and names that looks like this (you can add as many records as you need):

 ```
 {
    "IPEndpoint": "https://ipinfo.io/ip",
    "Sleep": 60,
    "Records":
    [
        {
            "Username": "myaccount@someone.com",
            "API-Key": "88b2b8e3d2b68b9cc4b945d81516v91d77k6g",
            "Zone": "myzone.xyz",
            "Entry": "myentry.xyz"
        },
        {
            "Username": "myaccount_1@someone.com",
            "API-Key": "55b2b8e3d2b68b9cc4b945d81516v91d77k6g",
            "Zone": "anotherzone.xyz",
            "Entry": "anotherentry.xyz"
        }
    ]
}
 ```

 3. Run the program `cfdns --config config.json`

### Container

TODO