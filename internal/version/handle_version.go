package version

import (
	"fmt"
	"os"
)

var Version = "0.1"

func HandleVersion() {
	fmt.Println("Current version of DNS-GO:", Version)

	os.Exit(0)
}
