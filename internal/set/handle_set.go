package set

import (
	"fmt"

	"dnsgo/internal/constants"
	"dnsgo/internal/helpers"
	"dnsgo/internal/helpers/dns"
	"dnsgo/internal/list"
)

func Handle(dnsName string) {
	ips := dns.GetDnsIPsByName(dnsName)
	if len(ips) < 1 {
		fmt.Println(constants.Strings["enter_valid_dns"])
		list.Handle()
	}

	goos := helpers.Goos()
	switch goos {
	case "linux":
		err := dns.SetDnsServersLinux(ips)
		if err != nil {
			fmt.Println(err)
			helpers.Cleanup()
		}
	case "darwin":
		err := dns.SetDnsServersMac(ips)
		if err != nil {
			fmt.Println(err)
			helpers.Cleanup()
		}
	case "windows":
		fmt.Println("windows")
		helpers.Cleanup()
	default:
		fmt.Println(constants.Strings["unsupported_os"])
		helpers.Cleanup()
	}

	fmt.Println(constants.Strings["dns_set"], dnsName)

	helpers.Cleanup()
}
