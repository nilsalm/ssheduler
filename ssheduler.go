package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	charmfs "github.com/charmbracelet/charm/fs"
)

func main() {
	// Open the file system
	cfs, err := charmfs.NewFS()
	if err != nil {
		panic(err)
	}

	// Load the prepared file with commands and magic
	file, err := os.Open("r1_cmd.sh")
	if err != nil {
		panic(err)
	}
	// Write the prepared file to the DB
	err = cfs.WriteFile("/testcmd", file)
	if err != nil {
		panic(err)
	}

	// Get a file from the DB
	f, err := cfs.Open("/testcmd")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes()))

	// Save to local executable temporary file
	err = os.WriteFile("/tmp/ssheduler_cmd.sh", buf.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Execute the file
	out, err := exec.Command("/tmp/ssheduler_cmd.sh").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

}
