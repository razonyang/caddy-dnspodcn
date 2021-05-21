# Caddy2 DNSPod.cn DNS Provider Module
[![Build Status](https://img.shields.io/travis/clevergo/caddy-dnspodcn?style=flat-square)](https://travis-ci.org/clevergo/caddy-dnspodcn)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/caddy-dnspodcn?style=flat-square)](https://coveralls.io/github/clevergo/caddy-dnspodcn)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/clevergo.tech/caddy-dnspodcn?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/caddy-dnspodcn?style=flat-square)](https://goreportcard.com/report/github.com/clevergo/caddy-dnspodcn)
[![Release](https://img.shields.io/github/release/clevergo/caddy-dnspodcn.svg?style=flat-square)](https://github.com/clevergo/caddy-dnspodcn/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/caddy-dnspodcn?style=flat-square)](https://pkg.clevergo.tech/clevergo.tech/caddy-dnspodcn)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat-square)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat-square&color=orange)](https://forum.clevergo.tech)


## Usage

Install [xcaddy](https://github.com/caddyserver/xcaddy)

```
$ go get -u github.com/caddyserver/xcaddy/cmd/xcaddy
```

Rebuild caddy with `dnspodcn` module

```shell
$ xcaddy build --with clevergo.tech/caddy-dnspodcn
```

## Configuration

`Caddyfile` example:

```
tls {
  dns dnspodcn APP_ID APP_TOKEN
}
```

`json` example:

```
  "challenges": {
    "dns": {
      "provider": {
        "Language": "en",
        "BaseURL": "https://dnsapi.cn",
        "name": "dnspodcn",
        "AppID": "APP_ID",
        "AppToken": "APP_TOKEN"
      }
    }
  }
```
- `APP_ID`: API ID.
- `APP_TOKEN`: API TOKEN.
