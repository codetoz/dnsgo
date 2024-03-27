package change

import (
	"math/rand"
	"time"

	"dnsgo/internal/helpers/dns"
	"dnsgo/internal/set"
)

func Handle() {
	randomDnsOption := chooseRandomDnsServer()

	set.Handle(randomDnsOption.Name)
}

func chooseRandomDnsServer() dns.DnsOption {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(dns.DnsOptions))

	randomDnsServer := dns.DnsOptions[randomIndex]

	return randomDnsServer
}
