package main

import (
	"context"
	"net"

	"github.com/miekg/dns"
	"github.com/semihalev/log"
	"github.com/semihalev/sdns/config"
	"github.com/semihalev/sdns/ctx"
)

// Example type
type Example struct {
	cfg  map[string]interface{}
	plog log.Logger
}

// New return example
func New(cfg *config.Config) ctx.Handler {
	plog := log.New("plugin", name)

	plog.Info("It's example plugin, loaded success")

	return &Example{
		cfg:  cfg.Plugins[name].Config,
		plog: plog,
	}
}

// Name return middleware name
func (e *Example) Name() string { return name }

// ServeDNS implements the Handle interface.
func (e *Example) ServeDNS(ctx context.Context, dc *ctx.Context) {
	/*
		$ dig value_1 @127.0.0.1

		;; Got answer:
		;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 4718
		;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

		;; OPT PSEUDOSECTION:
		; EDNS: version: 0, flags:; udp: 1232
		;; QUESTION SECTION:
		;value_1.			IN	A

		;; ANSWER SECTION:
		value_1.		3600	IN	A	0.0.0.0
	*/

	req, w := dc.DNSRequest, dc.DNSWriter

	testValue := dns.Fqdn(e.cfg["key_1"].(string))

	q := req.Question[0]

	if q.Name == testValue && q.Qtype == dns.TypeA {
		msg := new(dns.Msg)
		msg.SetReply(req)
		msg.Authoritative, msg.RecursionAvailable = true, true

		rrHeader := dns.RR_Header{
			Name:   testValue,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    3600,
		}
		a := &dns.A{Hdr: rrHeader, A: net.IPv4zero}
		msg.Answer = append(msg.Answer, a)

		w.WriteMsg(msg)
		dc.Abort()

		e.plog.Info("Example request received")
	}

	dc.NextDNS(ctx)
}

const name = "example"
