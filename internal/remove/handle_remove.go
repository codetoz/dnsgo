package remove

import (
	"dnsgo/internal/helpers"
	"dnsgo/internal/references"
	"fmt"
)

func Handle() {
	goos := helpers.Goos()
	switch goos {
	case "linux":
		err := removeDnsLinux()
		if err != nil {
			fmt.Println(err)
			helpers.Cleanup()
		}
	case "darwin":
		err := removeDnsMac()
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

	fmt.Println(references.Strings["dns_removed"])

	helpers.Cleanup()
}

func removeDnsLinux() error {
	//if backup file is exist
	// replace /etc/resolve.conf with backup file

	//if backup file is not exist
	// set google DNS
	return nil
}

func removeDnsMac() error {
	_, err := helpers.ExecuteCommand("networksetup -setdnsservers Wi-Fi empty")
	if err != nil {
		return err
	}

	return nil
}
