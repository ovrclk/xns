package main

import (
	"flag"
	"log"

	"github.com/miekg/dns"
	"github.com/ovrclk/xns/handler"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:53", "address to bind (udp)")
	root := flag.String("root", "", "root domain (example: xns.akash.network)")
	flag.Parse()

	log.SetPrefix("[xns] ")
	log.Println("starting DNS server, binding to:", addr)
	log.Println("resolving wildcard addresses for domain:", *root)
	srv := &dns.Server{Addr: *addr, Net: "udp"}
	srv.Handler = &handler.Handler{RootDomain: *root}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
