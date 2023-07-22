# cloudflare-dynamic-ip-fixer-upper

simple program to update ips on your CF zones

should be run periodically (e.g. crontab)

## Build

`go build` shall output ./cloudflare-dynamic-ip-fixer-upper

## Usage

this program will go through every zone allocated to whatever api token you give it, so generally you should use a scoped token instead of a global one.


Create an API Token

![https://imgur.com/5l47vfs.png](https://imgur.com/5l47vfs.png)

Edit zone DNS template works fine.

![https://imgur.com/GRKgoHq.png](https://imgur.com/GRKgoHq.png)

Add whatever zones you want this to affect to zone resources

You could do client IP address filtering too, but something tells me if you're using this script you probably don't have the liberty of a static IP :p

Set the TTL for however long you wish.

![https://imgur.com/wJoo4Ip.png](https://imgur.com/wJoo4Ip.png)

Create your token and you're good to go :)

![https://imgur.com/hWFwgPo.png](https://imgur.com/hWFwgPo.png)




copy .env.example and change values
`cp .env.example .env`

`CLOUDFLARE_API_TOKEN` is the token you generated earlier

`DOMAINS` is a comma deliminated list of DNS records you want to update

