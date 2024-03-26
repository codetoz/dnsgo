package helpers

import "strings"

var localDnsList = []string{"192.168", "127.0.0.53"}

func IsLocalDns(dns string) bool {
	for _, localDns := range localDnsList {
		if strings.Contains(dns, localDns) {
			return true
		}
	}

	return false
}
