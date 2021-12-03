package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func speedTest() {
	fmt.Println("Speedtest routine started....")

	cmd := exec.Command("./speedtest-cli", "--simple")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print("The result of speedtest is:\n", string(cmdOutput.Bytes()))
}

func main() {

	speedTest()
}
