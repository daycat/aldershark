package utils

import (
	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
	"golang.org/x/net/context"
	"net"
	"sync"
	"time"
)

func ResolveDomain(domain string) string {
	IP, _ := net.LookupIP(domain)
	return IP[0].String()
}

func ResolveDomainSecure(domain string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := doh.New(doh.CloudflareProvider)
	rsp, err := client.Query(ctx, dns.Domain(domain), dns.TypeA)
	if err != nil {
		println("Error: ", err)
		return ""
	}
	answer := rsp.Answer
	return answer[0].Data
}

func Resolve(c struct {
	Country        string `json:"country"`
	Region         string `json:"region"`
	Location       string `json:"location"`
	ConnectionName string `json:"connectionName"`
	IP             []string
}, secure bool, wg sync.WaitGroup) {
	if secure {
		c.IP = append(c.IP, ResolveDomainSecure(c.ConnectionName))
	} else {
		c.IP = append(c.IP, ResolveDomain(c.ConnectionName))
	}
	wg.Done()
}
