package common

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func PrintFileToScreen(path string) {
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

type Mode int

const (
	Init Mode = iota
	Queen
	Pawn
	Error
)

type Configuration struct {
	Mode Mode
}

const (
	confPath string = "config.yaml"
)

var config Configuration

func SaveConfig() {
	data, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile(confPath, data, 0664)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func ReadConfig() {
	yfile, err := ioutil.ReadFile(confPath)

	if err != nil {
		config.Mode = Init
		return
	}

	err2 := yaml.Unmarshal(yfile, &config)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func GetMode() Mode {
	return config.Mode
}

func SetMode(m Mode) {
	config.Mode = m
}
