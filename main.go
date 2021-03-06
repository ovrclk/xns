package main

import (
	"log"

	"github.com/miekg/dns"
	"github.com/ovrclk/xns/handler"
	"github.com/spf13/cobra"
)

var addr, root, ip string
var ns []string
var cmd = &cobra.Command{
	Use:   "xns",
	Short: "xns is a simple and fast wildcard DNS service",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		doStart(addr, root)
	},
}

func init() {
	cmd.PersistentFlags().StringVarP(&ip, "ip", "i", "127.0.0.1", "public IP to resolve root domain and NS records")
	cmd.PersistentFlags().StringVarP(&addr, "bind", "b", "0.0.0.0:53", "address to bind (udp)")
	cmd.PersistentFlags().StringVarP(&root, "root", "r", "aksh.io", "root domain (example: aksh.io)")
	cmd.PersistentFlags().StringSliceVarP(&ns, "ns", "n", []string{"ns1.aksh.io", "ns2.aksh.io"}, "ns servers")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func doStart(addr, root string) {
	log.SetPrefix("[xns]")
	log.Println("starting DNS server, binding to:", addr)
	log.Println("resolving wildcard addresses for domain:", root)
	srv := &dns.Server{Addr: addr, Net: "udp"}
	srv.Handler = handler.New(root, ip, ns)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}

var desc = `xns is a simple and fast wildcard DNS service for any IP addressm,
written in go and inspired by xip.io. xns powers aksh.io.
`
