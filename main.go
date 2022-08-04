package main

import (
	queen "github.com/nilsalm/ssheduler/queen"
)

func main() {
	cmd_filepath_synced := "/cmd"
	cmd_filepath := "ssheduler_cmd.sh"
	queen.Upload_file_to_charm(cmd_filepath, cmd_filepath_synced)
}
