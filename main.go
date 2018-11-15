package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type HLTPTransformationWrapper struct {
	Jars           []string `json:"jars"`
	Schema         []string `json:"schema"`
	Transformation []string `json:"transformation"`
}

type HLTPDescriptor struct {
	Wrapper HLTPTransformationWrapper `json:"transform"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("workflow.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened workflow.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// bytes
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var descriptor HLTPDescriptor
	err = json.Unmarshal([]byte(byteValue), &descriptor)
	if err != nil {
		fmt.Println(err)
	}

	//
	for _, e := range descriptor.Wrapper.Schema {
		fmt.Println(e)
	}

	for _, e := range descriptor.Wrapper.Transformation {
		fmt.Println(e)
	}

}
