package common

import (
	"fmt"
	"io/ioutil"
)

func Print_file_to_screen(path string) {
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
