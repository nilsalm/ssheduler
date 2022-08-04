package pawn

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	charmfs "github.com/charmbracelet/charm/fs"
)

type FileSystem struct {
	Files *charmfs.FS
}

func (cfs *FileSystem) download_file_from_charm(local_path string, charm_path string) {

	fmt.Printf("Downloading file from %s to %s ... ", charm_path, local_path)
	// Get a file from the DB
	file, err := cfs.Files.Open(charm_path)
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

func mark_job_as_done(charm_path string) {
	// TODO
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
