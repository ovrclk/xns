package handler

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/miekg/dns"
)

type Handler struct {
	RootDomain, PublicIP string
	Nameservers          []string

	dotRegexp, dashRegexp *regexp.Regexp
}

func New(rootDomain, publicIP string, nameServers []string) *Handler {
	dotted, _ := regexp.Compile(fmt.Sprintf(`(\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b).%s.`, rootDomain))
	dashed, _ := regexp.Compile(fmt.Sprintf(`(\b\d{1,3}\-\d{1,3}\-\d{1,3}\-\d{1,3}\b).%s.`, rootDomain))
	return &Handler{
		RootDomain:  rootDomain,
		PublicIP:    publicIP,
		Nameservers: nameServers,
		dotRegexp:   dotted,
		dashRegexp:  dashed,
	}
}

func (this *Handler) Resolve(name string) string {
	// resolve to public ip if root domain
	if name == this.RootDomain+"." {
		return this.PublicIP
	}
	// resolve to public ip if ns domain
	for _, ns := range this.Nameservers {
		if name == ns+"." {
			return this.PublicIP
		}
	}
	if s := this.dotRegexp.FindStringSubmatch(name); len(s) > 0 {
		return s[len(s)-1]
	}
	if s := this.dashRegexp.FindStringSubmatch(name); len(s) > 0 {
		return strings.Replace(s[len(s)-1], "-", ".", -1)
	}
	return ""
}

func (this *Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	start := time.Now()
	log.Printf("[%s] QUESTION: %s", w.RemoteAddr(), r.Question[0].String())
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address := this.Resolve(domain)
		if address != "" {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(address),
			})
		}
		log.Printf("[%s] DOMAIN: %s ANSWER: %s. elapsed: %s", w.RemoteAddr(), domain, address, time.Since(start))
	}
	w.WriteMsg(&msg)
}
