package version

import (
	"fmt"
	"os"
)

var Version = "0.1"

func Handle() {
	fmt.Println("Current version of DNS-GO:", Version)

	os.Exit(0)
}
