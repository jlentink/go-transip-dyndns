# Go-transip-dyndns
[![](https://goreportcard.com/badge/github.com/jlentink/go-transip-dyndns)](https://goreportcard.com/report/github.com/jlentink/go-transip-dyndns)

Is a small little executable that will update a domain record of your choice that is hosted by [TransIP](https://www.transip.nl/). Run it via cron to keep your DNS record up to date.

## Setup
1. Make sure you have a domain at Transip
2. Create an access key for the API. `(Control panel > My Account > API)`
3. Add a label and press create.
4. Save Key to a file. `e.g. privatekey.pem`
5. Create config file. run `go-transip-dyndns init` to create a configuration file. [Example config file](#example-config-file)
6. Add it to the crontab `0 * * * *  /usr/local/bin/go-transip-dyndns update` 
7. run `go-transip-dyndns create` to create the record
8. run `go-transip-dyndns validate` to see if all is 👍
9. You are set. Well done!


## The configuration file
### Key file
The key file can be loaded via two options:
1. Load key from external file.
2. Inject keyfile into the configuration file.

#### Loading key via external key file

Simply put the path to the file into the private-key field. 
Mind the single double quote at the beginning and end of the path.<br />

    private-key = "/some/path/to/key.file"

#### Loading key directly into the configuration file
When you want to store the keyfile in the configuration file for ease to only have to maintain one file
put the file directly behind the private-key value enclosed by three double quotes at the beginning and the
end of the certificate. Ensure you don't put any spaces in the key structure to not break the key. 

    private-key = """-----BEGIN PRIVATE KEY-----
    ...Your certificate data...
    ...Your certificate data...
    ...Your certificate data...
    ...Your certificate data...
    ...Your certificate data...
    ...Your certificate data...
    ...Your certificate data...
    ...Your certificate data...
    -----END PRIVATE KEY-----"""

#### TXT, MX etc records.
These records are handy and can have more markup than just an IP. For this purpose you have the ability to
completly format the record and just inject the IP addresses. This can be done to inject a ip via the tags:

1. {{.IPv4}} for the IPv4 address.
2. {{.IPv6}} for the IPv6 address.

Example SPF record:

    v=spf1 ip4:{{.IPv4}} ip6:{{.IPv6}} include:thirdpartydomain.com -all

### Example config file
name: go-transip-dyndns.toml

place the config file at /etc/go-transip-dyndns.toml or in the directory where you execute the command.


    [general]
    #
    # Enable verbose mode (debugging information).
    # Disabled by default.
    #
    verbose = false
    
    #
    # Pull in your public IPv4 address.
    #
    IPv4 = true
    
    #
    # Pull in your public IPv6 address.
    # Only use when you have an IPv6 address.
    #
    IPv6 = false
    
    #
    # Update in keep running mode every x (in minutes)
    #
    update-frequency = 1
    
    [account]
    #
    # Your account name on transip.
    #
    username = "your-account name"
    #
    # Private key to get access the API.
    # Create your own key here: https://www.transip.nl/cp/account/api/.
    #
    # You have two options here.
    # Include the private key in the configuration file.
    #
    # Example:
    # private-key = """-----BEGIN PRIVATE KEY-----
    #...Your certificate data...
    #-----END PRIVATE KEY-----"""
    #
    # or
    #
    # provide the path to the file that contains the private key.
    #
    # Example:
    # private-key = "/path/to/key.pem"
    #
    # Mind the """content""" (3x) quote for including the key in the config and the "path" (1) for the path...
    #
    private-key = """-----BEGIN PRIVATE KEY-----
    ...Your certificate data...
    -----END PRIVATE KEY-----"""
    
    #
    # The DNS record you want to update.
    # You can have as many as you want.
    #
    #[[record]]
    #
    # the domain name where the record should be updated.
    #
    #hostname = "example.com"
    #
    # The entry key for the domain
    # in this example my-home.example.com is the full dns entry we are creating here.
    #
    # use @ if you want to redirect the root domain.
    #
    #entry = "my-home"
    #
    # The caching time in seconds.
    #
    #ttl = 60
    #
    # The record type.
    # A for IPv4
    # AAAA for IPv6
    # but can also be MX TXT SRV
    #
    #type = "A"
    #
    # content that will be pushed into the record.
    # this value is ignored for A and AAAA records.
    # for other records you can use the placeholders {{.IPv4}} and {{.IPv6}}
    # to inject the IP's
    #
    # content = ""
    
    [[record]]
    hostname = "example.com"
    entry = "my-home"
    ttl = 60
    type = "A"
    content = ""

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

## Todo:
1. Code clean-up
2. DNS based change check (not only API)
