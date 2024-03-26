package log

import (
	"fmt"
	"strings"
)

func DEBUG(msg ...string) {
	fmt.Println(">>>> DEBUG <<<<")
	fmt.Println(strings.Join(msg, " "))
	fmt.Println()
}
