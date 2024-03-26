package current

import (
	"fmt"
	"os"
	"strings"

	"dnsgo/internal/helpers"
	"dnsgo/internal/list"
	"dnsgo/internal/references"
)

func Handle() {
	goos := helpers.Goos()

	var dnsServers []string

	switch goos {
	case "linux", "darwin":
		err := getDnsServersUnix(&dnsServers)
		if err != nil {
			fmt.Println(err)
			helpers.Cleanup()
		}
	case "windows":
		fmt.Println("windows")
		helpers.Cleanup()
	default:
		fmt.Println(references.Strings["unsupported_os"])
		helpers.Cleanup()
	}

	if len(dnsServers) < 1 || helpers.IsLocalDns(dnsServers[0]) {
		println(references.Strings["no_dns"])
		helpers.Cleanup()
	}

	var currentOption *list.DnsOption
	for _, option := range list.DnsOptions {
		if strings.Contains(dnsServers[0], option.IPs[0]) {
			currentOption = &option
			break
		}
	}
	if currentOption != nil {
		println(references.Strings["current_dns"], currentOption.Name)
	} else {
		println(references.Strings["unknown_dns"])
	}

	helpers.Cleanup()
}

func getDnsServersWindows() {}

func getDnsServersUnix(servers *[]string) error {
	data, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") {
			parts := strings.Split(line, " ")
			if len(parts) > 1 {
				*servers = append(*servers, parts[1])
			}
		}
	}
	return nil
}
