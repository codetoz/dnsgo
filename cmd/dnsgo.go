package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	// args := os.Args
	// for i, v := range args {
	// 	fmt.Printf("The %dth arg: %s\n", i, v)
	// }

	cmd := exec.Command("echo", "hi")
	cmd.Stdout = &bytes.Buffer{}
	output := cmd.Stdout

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(output)
}
