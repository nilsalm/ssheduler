package queen

import (
	"fmt"
	"os"

	// "runtime"
	// "time"

	"ssheduler/common"

	charmfs "github.com/charmbracelet/charm/fs"
)

// func ssheduler() {

// 	running_remote := true

// 	if runtime.GOOS == "darwin" {
// 		running_remote = false
// 	}

// 	cmd_filepath_synced := "/cmd"

// 	if running_remote == false {
// 		cmd_filepath := "ssheduler_cmd.sh"
// 		upload_file_to_charm(cmd_filepath, cmd_filepath_synced)
// 		print_file_to_screen(cmd_filepath)
// 	} else {
// 		cmd_filepath := "/tmp/ssheduler_cmd_" + fmt.Sprintf("%d", time.Now().Unix()) + ".sh"
// 		download_file_from_charm(cmd_filepath, cmd_filepath_synced)
// 		print_file_to_screen(cmd_filepath)
// 		execute_cmd_file(cmd_filepath)
// 		mark_job_as_done(cmd_filepath_synced)
// 	}
// }

type FileSystem struct {
	Files *charmfs.FS
}

func GetFS() *charmfs.FS {
	cfs, err := charmfs.NewFS()
	if err != nil {
		panic(err)
	}

	return cfs
}

func (cfs *FileSystem) UploadFileToCharm(local_path string, charm_path string) {

	fmt.Printf("Uploading file %s to %s ... ", local_path, charm_path)

	// Load the prepared file with commands and magic
	file, err := os.Open(local_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Write the prepared file to the DB
	err = cfs.Files.WriteFile(charm_path, file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success!")

	common.PrintFileToScreen(local_path)
}
