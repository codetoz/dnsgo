package version

import (
	"fmt"
	"os"
)

var Version = "0.1"

func HandleVersion() {
	fmt.Println("Current Version: DNS-GO:", Version)

	os.Exit(0)
}
