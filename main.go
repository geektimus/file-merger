package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

// MergeCommand is a representation of the intention to merge the file contents
type MergeCommand struct {
	Input    flags.Filename `short:"i" long:"inputFile" description:"Input JSON file containing the descriptor" required:"true"`
	Output   flags.Filename `short:"o" long:"outputFile" description:"Output file name / location (local folder by default)" default:"output.txt"`
	BasePath string         `short:"b" long:"basePath" description:"Folder containing the files referenced by the descriptor"`
}

// Wrapper contains the definitions of the fields that describe a transformation
type Wrapper struct {
	Jars           []string `json:"jars"`
	Schema         []string `json:"schema"`
	Transformation []string `json:"transformation"`
}

func (w Wrapper) flat() []string {
	return append(w.Schema, w.Transformation...)
}

// Descriptor is just a container for the wrapper object
type Descriptor struct {
	Wrapper Wrapper `json:"transform"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	var command MergeCommand

	if _, err := flags.ParseArgs(&command, os.Args); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	descriptor, err := parseJSONToDescriptor(string(command.Input))
	if err != nil {
		log.Fatalln("Failed to create a descriptor from the json file", err)
	}

	if err := concatenateFiles(descriptor, command.BasePath, string(command.Output)); err != nil {
		log.Fatalf("failed to concatenate the files = %v", err)
	}
}

// parseJSONToDescriptor receives a Json file with some configuration and
// it returns a descriptor containing all the information required to
// concatenate the files (in order)
func parseJSONToDescriptor(jsonFile string) (Descriptor, error) {
	var descriptor Descriptor

	// Open our inputJSONFile
	inputJSONFile, err := os.Open(jsonFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		return descriptor, fmt.Errorf("failed to open file for reading %v", err)
	}
	// defer the closing of our inputJSONFile so that we can parse it later on
	defer func() {
		err := inputJSONFile.Close()
		if err != nil {
			log.Errorf("failed to close the file %v", err)
		}
	}()

	// bytes
	byteValue, err := ioutil.ReadAll(inputJSONFile)
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
func concatenateFiles(descriptor Descriptor, basePathForFiles string, outputFileName string) error {

	out, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for writing %v", err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Errorf("failed to close the file %v", err)
		}
	}()

	wrapper := descriptor.Wrapper

	for _, e := range wrapper.flat() {
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
