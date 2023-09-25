package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var FilePath = "./"

const Filename = "prepush"

var files = []string{"prepush.yaml", "prepush.yml", "prepush.json"}

func GetConfigFile() (bool, os.FileInfo) {
	for _, file := range files {
		fileInfo, _ := os.Stat(fmt.Sprintf("%s", filepath.Join(FilePath, file)))
		if fileInfo != nil {
			return true, fileInfo
		}
	}

	return false, nil
}

func MakeConfigFile(ex string) {
	// Create a new file
	f := filepath.Join(FilePath, fmt.Sprintf("%s.%s", Filename, ex))

	file, err := os.Create(f)
	if err != nil {
		PrintError(err.Error())

		return
	}

	defer file.Close()

	// Write content to the file
	var content string

	if ex == "json" {
		content = getJsonContent()
	} else {
		content = getYamlContent()
	}

	_, err = io.WriteString(file, content)
	if err != nil {
		PrintError(err.Error())

		return
	}

	// Ensure the content is written to the disk
	err = file.Sync()
	if err != nil {
		PrintError(err.Error())

		return
	}

	PrintSuccess("Config file created successfully!")
}

func getYamlContent() string {
	return `config:
  name: value

commands:
  Hello PrePush!: echo "Hello PrePush"
  Listing current directory: ls -l`
}

func getJsonContent() string {
	return `{
    "config" : {
      "name" : "value"
    },

    "commands": {
      "Hello PrePush!" :  "echo \"Hello PrePush\"",
      "Listing current directory" :  "ls -l"
    }
}`
}
