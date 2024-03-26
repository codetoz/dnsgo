package current

import (
	"fmt"
	"strings"

	"dnsgo/internal/constants"
	"dnsgo/internal/helpers"
	"dnsgo/internal/helpers/dns"
)

func Handle() {
	goos := helpers.Goos()

	var dnsServers []string

	switch goos {
	case "linux", "darwin":
		err := dns.GetDnsServersUnix(&dnsServers)
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

	if len(dnsServers) < 1 || helpers.IsLocalDns(dnsServers[0]) {
		println(constants.Strings["no_dns"])
		helpers.Cleanup()
	}

	var currentOption *dns.DnsOption
	for _, option := range dns.DnsOptions {
		if strings.Contains(dnsServers[0], option.IPs[0]) {
			currentOption = &option
			break
		}
	}
	if currentOption != nil {
		println(constants.Strings["current_dns"], currentOption.Name)
	} else {
		println(constants.Strings["unknown_dns"])
	}

	helpers.Cleanup()
}
