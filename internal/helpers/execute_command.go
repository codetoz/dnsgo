package helpers

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(commandString string) (string, error) {
	commandArray := strings.Split(commandString, " ")

	cmd := exec.Command(commandArray[0], commandArray[1:]...)
	output, err := cmd.Output()
	return string(output), err
}
