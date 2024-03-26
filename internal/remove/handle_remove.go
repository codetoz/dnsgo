package remove

import (
	"fmt"

	"dnsgo/internal/constants"
	"dnsgo/internal/helpers"
	"dnsgo/internal/helpers/dns"
)

func Handle() {
	goos := helpers.Goos()
	switch goos {
	case "linux":
		err := dns.SetDnsServersLinux(nil)
		if err != nil {
			fmt.Println(err)
			helpers.Cleanup()
		}
	case "darwin":
		err := dns.SetDnsServersMac(nil)
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

	fmt.Println(constants.Strings["dns_removed"])

	helpers.Cleanup()
}
