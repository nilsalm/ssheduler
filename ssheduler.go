package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	charmfs "github.com/charmbracelet/charm/fs"
)

func main() {

	running_remote := true

	if runtime.GOOS == "darwin" {
		running_remote = false
	}

	cmd_filepath_synced := "/cmd"

	if running_remote == false {
		cmd_filepath := "ssheduler_cmd.sh"
		upload_file_to_charm(cmd_filepath, cmd_filepath_synced)
		print_file_to_screen(cmd_filepath)
	} else {
		cmd_filepath := "/tmp/ssheduler_cmd_" + fmt.Sprintf("%d", time.Now().Unix()) + ".sh"
		download_file_from_charm(cmd_filepath, cmd_filepath_synced)
		print_file_to_screen(cmd_filepath)
		execute_cmd_file(cmd_filepath)
	}
}

func upload_file_to_charm(local_path string, charm_path string) {
	// Open the file system
	cfs, err := charmfs.NewFS()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uploading file %s to %s ... ", local_path, charm_path)

	// Load the prepared file with commands and magic
	file, err := os.Open(local_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Write the prepared file to the DB
	err = cfs.WriteFile(charm_path, file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success!")
}

func download_file_from_charm(local_path string, charm_path string) {
	// Open the file system
	cfs, err := charmfs.NewFS()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Downloading file from %s to %s ... ", charm_path, local_path)
	// Get a file from the DB
	file, err := cfs.Open(charm_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	// Save to local executable temporary file
	err = os.WriteFile(local_path, buf.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success!")
}

func print_file_to_screen(path string) {
	fmt.Println("Reading ", path)

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	fmt.Println(text)

}

func execute_cmd_file(path string) []byte {
	fmt.Println("Executing ", path)
	// Execute the file
	out, err := exec.Command(path).Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	return out
}
