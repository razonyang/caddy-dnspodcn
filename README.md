# Caddy2 DNSPod.cn DNS Provider Module
[![Build Status](https://img.shields.io/travis/clevergo/caddy-dnspodcn?style=for-the-badge)](https://travis-ci.org/clevergo/caddy-dnspodcn)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/caddy-dnspodcn?style=for-the-badge)](https://coveralls.io/github/clevergo/caddy-dnspodcn)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/caddy-dnspodcn?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/caddy-dnspodcn?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/caddy-dnspodcn)
[![Release](https://img.shields.io/github/release/clevergo/caddy-dnspodcn.svg?style=for-the-badge)](https://github.com/clevergo/caddy-dnspodcn/releases)


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

```nginx
domain.tld {
    tls {
        dns dnspodcn <APP_ID> <APP_TOKEN>
    }
}
```

- `APP_ID`: API ID.
- `APP_TOKEN`: API TOKEN.
