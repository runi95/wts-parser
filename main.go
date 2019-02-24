package main

import (
	"fmt"
	"github.com/runi95/wts-parser/parser"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var fileNameRegex = regexp.MustCompile("[^.]*")
var fileExtensionRegex = regexp.MustCompile("\\.[^.]+$") // Find file extension

func main() {
	if len(os.Args) != 2 {
		log.Printf("Error: Expected 1 argument, but got %d\n", len(os.Args)-1)
		os.Exit(10)
	}

	fileExtension := fileExtensionRegex.FindString(os.Args[1])
	fileName := fileNameRegex.FindString(os.Args[1])

	inputFilePath := os.Args[1]
	var outputFilePath string

	b, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Println(err)
		os.Exit(10)
	}

	if fileExtension == ".wts" {
		outputFilePath = fileName + ".json"

		log.Println("Reading wts file...")
		jsonObject := parser.WtsToJson(b)

		log.Println("Writing json file...")
		err = ioutil.WriteFile(outputFilePath, jsonObject, 0644)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Successfully created a file called " + outputFilePath)
		}
	} else if fileExtension == ".json" {
		outputFilePath = fileName + ".wts"

		log.Println("Reading json file...")
		wtsObject := parser.JsonToWts(b)

		log.Println("Writing wts file...")
		err = ioutil.WriteFile(outputFilePath, wtsObject, 0644)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Successfully created a file called " + outputFilePath)
		}
	} else if fileExtension == ".w3u" {
		outputFilePath = fileName + ".json"

		log.Println("Reading .w3u file...")
		jsonObject := parser.W3uToJson(b)

		log.Println("Writing json file...")
		err = ioutil.WriteFile(outputFilePath, jsonObject, 0644)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Successfully created a file called " + outputFilePath)
		}
	} else {
		log.Fatal(fmt.Errorf("expected .wts or .json, but got " + fileExtension))
	}
}
