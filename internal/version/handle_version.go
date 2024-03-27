package version

import (
	"fmt"

	"dnsgo/internal/constants"
	"dnsgo/internal/helpers"
)

func Handle() {
	fmt.Println("Current version of DNS-GO:", constants.Version)

	helpers.Cleanup()
}
