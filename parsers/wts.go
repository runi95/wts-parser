package parsers

import (
	"regexp"
	"fmt"
	"strconv"
	"encoding/json"
	"bytes"
)

var idRegex = regexp.MustCompile(`[0-9]+`) // Find string id
var stringRegex = regexp.MustCompile(`STRING [0-9]*[^}]*}`) // Regex to find each string object
var contentContainerRegex = regexp.MustCompile(`{[^}]*}$`) // Regex to find the content container of each string object
var contentStartRegex = regexp.MustCompile(`^[\r]*[\n]`)
var contextEndRegex = regexp.MustCompile(`[\r]*[\n]$`)

func WtsToJson(input []byte) []byte {
	str := string(input)

	findAllStrings := stringRegex.FindAllString(str, -1)
	m := make(map[int]string)

	for _, stringObject := range findAllStrings {
		findContentContainer := contentContainerRegex.FindString(stringObject)
		idStr := idRegex.FindString(stringObject)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Print(err)
		}

		removeBrackets := findContentContainer[1:len(findContentContainer) - 1]
		length := len(contentStartRegex.FindString(removeBrackets))
		parsedBeginning := removeBrackets[length:]
		lengthToEnd := len(parsedBeginning) - len(contextEndRegex.FindString(parsedBeginning))
		parsedContainer := parsedBeginning[:lengthToEnd]

		m[id] = parsedContainer
	}

	jsonObject, err := json.Marshal(m)
	if err != nil {
		fmt.Print(err)
	}

	return jsonObject
}

func JsonToWts(input []byte) []byte {
	var m map[string]string

	err := json.Unmarshal(input, &m)
	if err != nil {
		fmt.Print(err)
	}

	var buffer bytes.Buffer

	for key := range m {
		buffer.WriteString("STRING " + key + "\r\n{\r\n" + m[key] + "\r\n}\r\n\r\n")
	}

	return buffer.Bytes()
}