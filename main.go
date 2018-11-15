package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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
	inputFileName := "testdata/workflow.json"
	outputFileName := "output.sql"
	basePathForFiles := "testdata"
	parseTransformationFiles(inputFileName, basePathForFiles, outputFileName)
}

func parseTransformationFiles(inputFileName string, basePathForFiles string, outputFileName string) {
	// Open our inputJsonFile
	inputJsonFile, err := os.Open(inputFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln("Failed to open file for reading:", err)
	}
	log.Println("Successfully Opened " + inputFileName)
	// defer the closing of our inputJsonFile so that we can parse it later on
	defer inputJsonFile.Close()

	// bytes
	byteValue, err := ioutil.ReadAll(inputJsonFile)
	if err != nil {
		log.Fatalln("Failed to read file before parsing:", err)
	}

	var descriptor HLTPDescriptor
	err = json.Unmarshal([]byte(byteValue), &descriptor)
	if err != nil {
		log.Fatalln("Failed to parse the descriptor from JSON data", err)
	}

	out, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Failed to open file for writing:", err)
	}
	defer out.Close()

	for _, e := range descriptor.Wrapper.Schema {
		concatenateFiles(basePathForFiles + "/" + e, out)
	}

	for _, e := range descriptor.Wrapper.Transformation {
		concatenateFiles(basePathForFiles + "/" + e, out)
	}
}

// concatenateFiles reads the inFile and write its contests to the outFile (pointer)
func concatenateFiles(inFile string, outFile *os.File) {
	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalln("Failed to open the file:", err)
	}
	n, err := io.Copy(outFile, f)
	if err != nil {
		log.Fatalln("Failed to append the files:", err)
	}
	log.Printf("Wrote %d bytes of %s to the end of %s\n", n, inFile, outFile.Name())
}
