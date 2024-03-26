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

func SetDnsServersLinux(servers []string) error {
	if servers == nil {
		existsBackup, err := file.FileExists(constants.BackupFilePath)
		if err != nil {
			return err
		}

		if existsBackup {
			file.ReplaceFile(constants.UnixResolvFilePath, constants.BackupFilePath)
		} else {
			// set google DNS

			googleDnsIPs := GetDnsIPsByName("google")
			resolvConfContent := GenerateResolvFileContent(googleDnsIPs)
			err := file.WriteStringFile(constants.UnixResolvFilePath, resolvConfContent)
			if err != nil {
				return err
			}
		}

		return nil
	}

	// TODO handle set dns
	return nil
}

func SetDnsServersMac(servers []string) error {
	if servers == nil {
		_, err := helpers.ExecuteCommand("networksetup -setdnsservers Wi-Fi empty")
		if err != nil {
			return err
		}
		return nil
	}

	// TODO handle set dns
	return nil
}

func GenerateResolvFileContent(dnsServers []string) string {
	fileContent := ""

	for _, option := range dnsServers {
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
