package dns

import (
	"os"
	"strings"

	"dnsgo/internal/constants"
	"dnsgo/internal/helpers"
	"dnsgo/internal/helpers/file"
)

func GetDnsServersUnix(servers *[]string) error {
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

func GetDnsServersWindows() {}

func SetDnsServersLinux(ips []string) error {
	var resolvConfContent string

	if ips == nil {
		existsDefaultFile, err := file.FileExists(constants.DefaultFilePath)
		if err != nil {
			return err
		}

		if existsDefaultFile {
			file.ReplaceFile(constants.DefaultFilePath, constants.UnixResolvFilePath)
			return nil
		} else {
			// copy current resolv.conf file to another file (resolve.conf.default)
			file.CopyFile(constants.UnixResolvFilePath, constants.DefaultFilePath)

			// google DNS
			googleDnsIPs := GetDnsIPsByName("google")
			resolvConfContent = GenerateResolvFileContent(googleDnsIPs)
		}
	} else {
		resolvConfContent = GenerateResolvFileContent(ips)
	}

	err := file.WriteStringFile(constants.UnixResolvFilePath, resolvConfContent)
	if err != nil {
		return err
	}

	return nil
}

func SetDnsServersMac(ips []string) error {
	if ips == nil {
		_, err := helpers.ExecuteCommand("networksetup -setdnsservers Wi-Fi empty")
		if err != nil {
			return err
		}
		return nil
	}

	// TODO handle set dns
	dnsIpsString := strings.Join(ips, " ")
	_, err := helpers.ExecuteCommand("networksetup -setdnsservers Wi-Fi " + dnsIpsString)
	if err != nil {
		return err
	}

	return nil
}

func GenerateResolvFileContent(dnsIps []string) string {
	fileContent := ""

	for _, option := range dnsIps {
		fileContent += "nameserver " + option + "\n"
	}

	return fileContent
}

func GetDnsIPsByName(name string) []string {
	for _, obj := range DnsOptions {
		if obj.Name == name {
			return obj.IPs
		}
	}
	return nil // DNS Option not found!
}
