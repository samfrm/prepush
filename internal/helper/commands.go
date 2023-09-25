package helper

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"path/filepath"
)

var result map[string]map[string]string

func getCommands() {
	e, file := GetConfigFile()
	if !e {
		PrintWarning("Config file not found!")
		PrintInfo("You can run prepush init to make new one.")

		return
	}

	if file.Name() == filepath.Join(FilePath, fmt.Sprintf("%s.%s", Filename, "json")) {
		readFromJsonFile(file.Name())
	} else {
		readFromYamlFile(file.Name())
	}
}

func RunCommands() {
	getCommands()

	for job, command := range result["commands"] {
		PrintSuccess(fmt.Sprintf("Executing: %s...", job))
		cmd := exec.Command("bash", "-c", command)

		output, err := cmd.Output()

		if err != nil {
			PrintError(err.Error())
			return
		}

		fmt.Println(string(output))
	}
}

func readFromYamlFile(fileName string) {
	// Read YAML file
	data, err := os.ReadFile(fileName)
	if err != nil {
		PrintError(err.Error())
		return
	}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		PrintError(err.Error())
		return
	}
}

func readFromJsonFile(fileName string) {
	content, err := os.ReadFile(fileName)

	err = json.Unmarshal(content, &result)
	if err != nil {
		PrintError(err.Error())
		return
	}
}
