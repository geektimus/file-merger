package main

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestValidJsonToDescriptor(t *testing.T) {
	descriptor, err := parseJSONToDescriptor("testdata/workflow.json")
	if err != nil {
		t.Fatal("Could not parse the json configuration required for the descriptor")
	}

	wrapper := descriptor.Wrapper
	expectedJarName := "${project.build.finalName}.jar"
	expectedSampleSchema := "schema/schema-01.sql"
	expectedSampleTransformation := "transformation/transformation-02.sql"

	if len(wrapper.Jars) != 1 {
		t.Fatalf("Expecting one jar, got %d", len(wrapper.Jars))
	} else {
		first := wrapper.Jars[0]
		if first != expectedJarName {
			t.Fatal("The jar element should be " + expectedJarName)
		}
	}

	if len(wrapper.Schema) != 2 {
		t.Fatalf("Expecting two schemas, got %d", len(wrapper.Schema))
	} else {
		first := wrapper.Schema[0]
		if first != expectedSampleSchema {
			t.Fatalf("Expecting fist schema to be %s, got %s", expectedSampleSchema, first)
		}
	}

	if len(wrapper.Transformation) != 5 {
		t.Fatalf("Expecting five transformations, got %d", len(wrapper.Transformation))
	} else {
		first := wrapper.Transformation[3]
		if first != expectedSampleTransformation {
			t.Fatalf("Expecting third transformation to be %s, got %s", expectedSampleTransformation, first)
		}
	}
}

func TestInvalidJsonToDescriptor(t *testing.T) {
	_, err := parseJSONToDescriptor("testdata/non-existent.json")
	if err == nil {
		t.Fatal("It should fail because the json file doesn't exist")
	}
}

func TestConcatenateFiles(t *testing.T) {
	logrus.SetLevel(logrus.ErrorLevel)
	outputFileName := "test-output.txt"
	descriptor, _ := parseJSONToDescriptor("testdata/workflow.json")
	err := concatenateFiles(descriptor, "testdata", outputFileName)
	if err != nil {
		t.Fatal("Could not concatenate the files")
	}
	if err := os.Remove(outputFileName); err != nil {
		t.Fatal("Could not remove the test output file")
	}
}
