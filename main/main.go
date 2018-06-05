package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"regexp"
	"WTSParser/parsers"
)

var fileNameRegex = regexp.MustCompile("[^.]*")
var fileExtensionRegex = regexp.MustCompile("\\.[^.]+$") // Find file extension

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Error: Expected 1 argument, but got %d\n", len(os.Args) - 1)
		os.Exit(10)
	}

	fileExtension := fileExtensionRegex.FindString(os.Args[1])
	fileName := fileNameRegex.FindString(os.Args[1])

	inputFilePath := os.Args[1]
	var outputFilePath string

	b, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Print(err)
		os.Exit(10)
	}

	if fileExtension == ".wts" {
		outputFilePath = fileName + ".json"

		jsonObject := parsers.WtsToJson(b)

		err = ioutil.WriteFile(outputFilePath, jsonObject, 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else if fileExtension == ".json" {
		outputFilePath = fileName + ".wts"

		wtsObject := parsers.JsonToWts(b)

		err = ioutil.WriteFile(outputFilePath, wtsObject, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
