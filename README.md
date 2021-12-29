# Go-transip-dyndns
[![](https://goreportcard.com/badge/github.com/jlentink/go-transip-dyndns)](https://goreportcard.com/report/github.com/jlentink/go-transip-dyndns)

Is a small little executable that will update a domain record of your choice that is hosted by [TransIP](https://www.transip.nl/). Run it via cron to keep your DNS record up to date.

## Setup
1. Make sure you have a domain at Transip
2. Create an access key for the API. `(Control panel > My Account > API)`
3. Add a label and press create.
4. Save Key to a file. `e.g. privatekey.pem`
5. Create config file. [Example config file](#example-config-file)
6. Add it to the crontab `0 * * * *  /usr/local/bin/go-transip-dyndns`
7. run `go-transip-dyndns create` to create the record
8. run `go-transip-dyndns validate` to see if all is 👍
9. You are set. Well done!


## Example config file
name: go-transip-dyndns.toml

place the config file at /etc/go-transip-dyndns.toml or in the directory where you execute the command.

```
username = "transip-username"
private-key = "/path-to/private.key"

verbose = false

domain = "yourdomain.nl"
domain-entry = "subdomain"
domain-ttl = 60
```

## Binaries available for download
Binaries are available for download in multiple formats

* Windows (32/64 Bit)
* Linux Intel (32/64 Bit)
* Linux ARM (32/64 Bit) - Run directly on most routers
* Linux ARM for Raspberry PI
* Linux MIPS 64 - Unifi USG
* MacOS 64 Bit

Download them [here](https://github.com/jlentink/go-transip-dyndns/releases/latest).

## No association with Transip
This tool has been created for me own comfort. There is no association with Transip. But I would like to thank TransIP for there fine service!

## PHP version
[Previous version](https://github.com/jlentink/transip-dyndns) was build in PHP.
