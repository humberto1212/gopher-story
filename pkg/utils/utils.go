package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/humberto1212/gopher-story/pkg/models"
)

func ParseJSON() map[string]models.StoryArc {

	// Open jsonFile
	jsonFile, err := os.Open("")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize histories array
	//var stories models.Histories

	var result = map[string]models.StoryArc{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
