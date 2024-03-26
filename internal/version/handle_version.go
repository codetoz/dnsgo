package version

import (
	"fmt"

	"dnsgo/internal/helpers"
)

var Version = "0.1"

func Handle() {
	fmt.Println("Current version of DNS-GO:", Version)

	helpers.Cleanup()
}
