// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package caddydnspodcn

import (
	"clevergo.tech/dnspodcn"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

var (
	_ caddyfile.Unmarshaler = (*Module)(nil)
)

func init() {
	caddy.RegisterModule(&Module{})
}

// Module wraps dnspodcn.Provider.
type Module struct {
	*dnspodcn.Provider
}

// CaddyModule returns the Caddy module information.
func (m *Module) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.dnspodcn",
		New: func() caddy.Module {
			return &Module{
				Provider: &dnspodcn.Provider{
					Client: dnspodcn.NewClient("", ""),
				},
			}
		},
	}
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// dnspodcn [<app_id> <app_token>] {
//     app_id <app_id>
//     app_token <app_token>
// }
//
func (m *Module) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	repl := caddy.NewReplacer()
	for d.Next() {
		if d.NextArg() {
			m.AppID = repl.ReplaceAll(d.Val(), "")
		}
		if d.NextArg() {
			m.AppToken = repl.ReplaceAll(d.Val(), "")
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "app_id":
				if m.AppID != "" {
					return d.Err("APP ID already set")
				}
				m.AppID = repl.ReplaceAll(d.Val(), "")
				if d.NextArg() {
					return d.ArgErr()
				}
			case "app_token":
				if m.AppToken != "" {
					return d.Err("APP token already set")
				}
				m.AppToken = repl.ReplaceAll(d.Val(), "")
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if m.AppID == "" || m.AppToken == "" {
		return d.Err("missing APP ID or Token")
	}
	return nil
}
