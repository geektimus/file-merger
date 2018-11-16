package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	inputPtr := flag.String("input", "", "Input JSON file containing the descriptor")
	outputPtr := flag.String("output", "output.txt", "Output file name / location (local folder by default)")
	inBasePathPtr := flag.String("basePath", "", "Folder containing the files referenced by the descriptor")
	flag.Parse()

	if *inputPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	descriptor, err := parseJsonToDescriptor(*inputPtr)
	if err != nil {
		log.Fatalln("Failed to create a descriptor from the json file", err)
	}

	if err := concatenateFiles(descriptor, *inBasePathPtr, *outputPtr); err != nil {
		log.Fatalf("failed to concatenate the files = %v", err)
	}
}

// parseJsonToDescriptor receives a Json file with some configuration and
// it returns a descriptor containing all the information required to
// concatenate the files (in order)
func parseJsonToDescriptor(jsonFile string) (HLTPDescriptor, error) {
	var descriptor HLTPDescriptor

	// Open our inputJsonFile
	inputJsonFile, err := os.Open(jsonFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		return descriptor, fmt.Errorf("failed to open file for reading %v", err)
	}
	// defer the closing of our inputJsonFile so that we can parse it later on
	defer inputJsonFile.Close()

	// bytes
	byteValue, err := ioutil.ReadAll(inputJsonFile)
	if err != nil {
		return descriptor, fmt.Errorf("failed to read file before parsing %v", err)
	}

	err = json.Unmarshal([]byte(byteValue), &descriptor)
	if err != nil {
		return descriptor, fmt.Errorf("failed to parse the descriptor from JSON data %v", err)
	}
	return descriptor, nil
}

// concatenateFiles receives a descriptor with the location of the files and
// it concatenates all of them on the outputFileName
func concatenateFiles(descriptor HLTPDescriptor, basePathForFiles string, outputFileName string) error {

	out, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing %v", err)
	}
	defer out.Close()

	wrapper := descriptor.Wrapper
	paths := append(wrapper.Schema, wrapper.Transformation...)

	for _, e := range paths {
		err := concatenateFile(basePathForFiles+"/"+e, out)
		if err != nil {
			return err
		}
	}

	return nil
}

// concatenateFile reads the inFile and write its contests to the outFile (pointer)
func concatenateFile(inFile string, outFile *os.File) error {
	f, err := os.Open(inFile)
	if err != nil {
		return err
	}
	n, err := io.Copy(outFile, f)
	if err != nil {
		return err
	}
	log.Printf("Wrote %d bytes of %s to the end of %s\n", n, inFile, outFile.Name())
	return nil
}
