package helpers

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCommand(commandString string) string {
	commandArray := strings.Split(commandString, " ")

	cmd := exec.Command(commandArray[0], commandArray[1:]...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running command:", err)
		Cleanup()
	}

	return string(output)
}
