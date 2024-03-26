package version

import (
	"fmt"

	"dnsgo/internal/helpers"
)

var Version = "1.0.0"

func Handle() {
	fmt.Println("Current version of DNS-GO:", Version)

	helpers.Cleanup()
}
