package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmdFlags := os.Args[2:]

	inputReader := bufio.NewScanner(os.Stdin)
	var cmdParams []string
	for inputReader.Scan() {
		cmdParams = append(cmdParams, inputReader.Text())
	}

	cmdArgs := append(cmdFlags, cmdParams...)

	cmd := exec.Command(os.Args[1], cmdArgs...)

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
