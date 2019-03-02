package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/runi95/wts-parser/models"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var idRegex = regexp.MustCompile(`[0-9]+`)                  // Find string id
var stringRegex = regexp.MustCompile(`STRING [0-9]*[^}]*}`) // Regex to find each string object
var contentContainerRegex = regexp.MustCompile(`{[^}]*}$`)  // Regex to find the content container of each string object
var contentStartRegex = regexp.MustCompile(`^[\r]*[\n]`)
var contextEndRegex = regexp.MustCompile(`[\r]*[\n]$`)
var SLKRegex = regexp.MustCompile(`C;X([0-9]+)(?:;Y([0-9]+))?;K([-"\\\w]*)`)
var SLKMetaRegex = regexp.MustCompile(`B;X([0-9]+);(?:Y([0-9]+);)D([-"\w]*)`)

/*************************

       WTS PARSERS

*************************/

func WtsToJson(input []byte) []byte {
	str := string(input)

	findAllStrings := stringRegex.FindAllString(str, -1)
	m := make(map[int]string)

	for _, stringObject := range findAllStrings {
		findContentContainer := contentContainerRegex.FindString(stringObject)
		idStr := idRegex.FindString(stringObject)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println(err)
		}

		removeBrackets := findContentContainer[1 : len(findContentContainer)-1]
		length := len(contentStartRegex.FindString(removeBrackets))
		parsedBeginning := removeBrackets[length:]
		lengthToEnd := len(parsedBeginning) - len(contextEndRegex.FindString(parsedBeginning))
		parsedContainer := parsedBeginning[:lengthToEnd]

		m[id] = parsedContainer
	}

	jsonObject, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return jsonObject
}

/*************************

      JSON PARSERS

*************************/

func JsonToWts(input []byte) []byte {
	var m map[string]string

	err := json.Unmarshal(input, &m)
	if err != nil {
		log.Println(err)
	}

	var buffer bytes.Buffer

	for key := range m {
		buffer.WriteString("STRING " + key + "\r\n{\r\n" + m[key] + "\r\n}\r\n\r\n")
	}

	return buffer.Bytes()
}

/*************************

       W3U PARSERS

*************************/

func ReadW3uFile(input []byte) map[string]*W3uData {
	var lastUnitId string
	unitMap := make(map[string]*W3uData)

	for i := 12; i < len(input); i++ {
		baseUnitId := [4]byte{input[i-8], input[i-7], input[i-6], input[i-5]}
		if baseUnitId[0] > 96 && baseUnitId[0] < 123 && baseUnitId[1] > 96 && baseUnitId[1] < 123 && baseUnitId[2] > 96 && baseUnitId[2] < 123 && ((baseUnitId[3] > 96 && baseUnitId[3] < 123) || (baseUnitId[3] > 47 && baseUnitId[3] < 58)) {
			customUnitId := [4]byte{input[i-4], input[i-3], input[i-2], input[i-1]}
			if customUnitId[0] > 96 && customUnitId[0] < 123 && (customUnitId[1] == 48 || customUnitId[1] == 67) && ((customUnitId[2] > 47 && customUnitId[2] < 58) || (customUnitId[2] > 64 && customUnitId[2] < 91)) && ((customUnitId[3] > 47 && customUnitId[3] < 58) || (customUnitId[3] > 64 && customUnitId[3] < 91)) {
				unitId := string(customUnitId[0]) + string(customUnitId[1]) + string(customUnitId[2]) + string(customUnitId[3])
				lastUnitId = unitId
				newW3uUnit := new(W3uData)
				newW3uUnit.BaseUnitId = string(baseUnitId[0]) + string(baseUnitId[1]) + string(baseUnitId[2]) + string(baseUnitId[3])
				newW3uUnit.CustomUnitId = string(customUnitId[0]) + string(customUnitId[1]) + string(customUnitId[2]) + string(customUnitId[3])
				unitMap[unitId] = newW3uUnit
			}
		}

		if input[i] > 31 && input[i-1] > 31 && input[i-2] > 31 && input[i-3] > 31 && input[i+1] == 3 {
			str := string(input[i-3]) + string(input[i-2]) + string(input[i-1]) + string(input[i])
			currentUnit := unitMap[lastUnitId]

			i += 4
			value := ""
			for input[i+1] > 31 {
				i++
				value += string(input[i])
			}

			if len(value) > 0 {
				nullString := new(null.String)
				nullString.SetValid(value)

				err := reflectUpdateValueOnFieldNullString(currentUnit, *nullString, strings.Title(str))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

	return unitMap
}

func W3uToJson(input []byte) []byte {
	unitMap := ReadW3uFile(input)

	mapValues := make([]*W3uData, len(unitMap))
	index := 0
	for _, value := range unitMap {
		mapValues[index] = value
		index++
	}

	jsonObject, err := json.Marshal(mapValues)
	if err != nil {
		log.Println(err)
	}

	return jsonObject
}

func W3uToSLKUnits(input []byte) []*SLKUnit {
	unitMap := ReadW3uFile(input)

	slkUnits := make([]*SLKUnit, len(unitMap))
	index := 0
	for _, value := range unitMap {
		slkUnit := new(SLKUnit)
		value.TransformToSLKUnit(slkUnit)

		slkUnits[index] = slkUnit
		index++
	}

	return slkUnits
}

func W3uToSLKUnitsWithBaseSLK(baseSLKUnits map[string]*SLKUnit, input []byte) []*SLKUnit {
	unitMap := ReadW3uFile(input)

	slkUnits := make([]*SLKUnit, len(unitMap))
	index := 0
	for _, value := range unitMap {
		slkUnit := new(SLKUnit)
		baseSLKUnit := *baseSLKUnits["\"" + value.BaseUnitId + "\""]
		var unitUI UnitUI
		var unitData UnitData
		var unitWeapons UnitWeapons
		var unitBalance UnitBalance
		var unitAbilities UnitAbilities
		unitUI = *baseSLKUnit.UnitUI
		unitData = *baseSLKUnit.UnitData
		unitWeapons = *baseSLKUnit.UnitWeapons
		unitBalance = *baseSLKUnit.UnitBalance
		unitAbilities = *baseSLKUnit.UnitAbilities

		slkUnit.UnitUI = &unitUI
		slkUnit.UnitData = &unitData
		slkUnit.UnitWeapons = &unitWeapons
		slkUnit.UnitBalance = &unitBalance
		slkUnit.UnitAbilities = &unitAbilities

		value.TransformToSLKUnit(slkUnit)
		slkUnits[index] = slkUnit

		index++
	}

	return slkUnits
}

/*************************

       SLK PARSERS

*************************/

type SLKInformation struct {
	split          []string
	headerMap      map[string]string
	columnSize     int
	headerEndIndex int
}

func GenericSLKReader(input []byte) (*SLKInformation, error) {
	slkInformation := new(SLKInformation)
	str := string(input)
	split := strings.Split(str, "\n")

	meta := split[1]
	metaSubmatch := SLKMetaRegex.FindStringSubmatch(meta)
	if !(len(metaSubmatch) > 0) {
		return nil, fmt.Errorf("meta submatch is empty")
	}

	columnSize, err := strconv.Atoi(metaSubmatch[1])
	if err != nil {
		return nil, err
	}

	headerMap := make(map[string]string)

	headerLineIndex := 2
	isNotAtEndOfHeader := true
	for isNotAtEndOfHeader {
		headerLineSubmatch := SLKRegex.FindStringSubmatch(split[headerLineIndex])
		if len(headerLineSubmatch[2]) > 0 && headerLineSubmatch[2] != "1" {
			isNotAtEndOfHeader = false
		} else {
			if len(headerLineSubmatch) > 1 {
				headerMap[headerLineSubmatch[1]] = headerLineSubmatch[3]
			}

			headerLineIndex++
		}
	}

	slkInformation.split = split
	slkInformation.headerMap = headerMap
	slkInformation.columnSize = columnSize
	slkInformation.headerEndIndex = headerLineIndex

	return slkInformation, nil
}

func SLKToUnitData(input []byte) map[string]*UnitData {
	unitDataMap := make(map[string]*UnitData)
	var currentUnitData *UnitData
	slkInformation, err := GenericSLKReader(input)
	if err != nil {
		log.Println(err)

		return nil
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		bodyLineSubmatch := SLKRegex.FindStringSubmatch(bodyLine)
		if bodyLineSubmatch != nil {
			if len(bodyLineSubmatch[2]) > 0 {
				if currentUnitData != nil && currentUnitData.UnitID.Valid {
					unitDataMap[currentUnitData.UnitID.String] = currentUnitData
				}

				currentUnitData = new(UnitData)
			}

			nullString := new(null.String)
			nullString.SetValid(bodyLineSubmatch[3])

			err = reflectUpdateValueOnFieldNullString(currentUnitData, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	return unitDataMap
}

func SLKToUnitUI(input []byte) map[string]*UnitUI {
	unitUIMap := make(map[string]*UnitUI)
	var currentUnitUI *UnitUI
	slkInformation, err := GenericSLKReader(input)
	if err != nil {
		log.Println(err)

		return nil
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		bodyLineSubmatch := SLKRegex.FindStringSubmatch(bodyLine)
		if bodyLineSubmatch != nil {
			if len(bodyLineSubmatch[2]) > 0 {
				if currentUnitUI != nil && currentUnitUI.UnitUIID.Valid {
					unitUIMap[currentUnitUI.UnitUIID.String] = currentUnitUI
				}

				currentUnitUI = new(UnitUI)
			}

			nullString := new(null.String)
			nullString.SetValid(bodyLineSubmatch[3])

			err = reflectUpdateValueOnFieldNullString(currentUnitUI, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	return unitUIMap
}

func SLKToUnitWeapons(input []byte) map[string]*UnitWeapons {
	unitWeaponsMap := make(map[string]*UnitWeapons)
	var currentUnitWeapons *UnitWeapons
	slkInformation, err := GenericSLKReader(input)
	if err != nil {
		log.Println(err)

		return nil
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		bodyLineSubmatch := SLKRegex.FindStringSubmatch(bodyLine)
		if bodyLineSubmatch != nil {
			if len(bodyLineSubmatch[2]) > 0 {
				if currentUnitWeapons != nil && currentUnitWeapons.UnitWeapID.Valid {
					unitWeaponsMap[currentUnitWeapons.UnitWeapID.String] = currentUnitWeapons
				}

				currentUnitWeapons = new(UnitWeapons)
			}

			nullString := new(null.String)
			nullString.SetValid(bodyLineSubmatch[3])

			err = reflectUpdateValueOnFieldNullString(currentUnitWeapons, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	return unitWeaponsMap
}

func SLKToUnitBalance(input []byte) map[string]*UnitBalance {
	unitBalanceMap := make(map[string]*UnitBalance)
	var currentUnitBalance *UnitBalance
	slkInformation, err := GenericSLKReader(input)
	if err != nil {
		log.Println(err)

		return nil
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		bodyLineSubmatch := SLKRegex.FindStringSubmatch(bodyLine)
		if bodyLineSubmatch != nil && bodyLineSubmatch[1] != "6" {
			if len(bodyLineSubmatch[2]) > 0 {
				if currentUnitBalance != nil && currentUnitBalance.UnitBalanceID.Valid {
					unitBalanceMap[currentUnitBalance.UnitBalanceID.String] = currentUnitBalance
				}

				currentUnitBalance = new(UnitBalance)
			}

			nullString := new(null.String)
			nullString.SetValid(bodyLineSubmatch[3])

			err = reflectUpdateValueOnFieldNullString(currentUnitBalance, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	return unitBalanceMap
}

/*************************

    PRIVATE FUNCTIONS

*************************/

func reflectUpdateValueOnFieldNullString(iface interface{}, fieldValue null.String, fieldName string) error {
	fieldName = strings.Replace(fieldName, "\"", "", -1)
	valueIface := reflect.ValueOf(iface)

	// Check if the passed interface is a pointer
	if valueIface.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("can't swap values if the reflected interface is not a pointer")
	}

	// 'dereference' with Elem() and get the field by name
	field := valueIface.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("interface `%s` does not have the field `%s`", valueIface.Type(), fieldName)
	}

	// A Value can be changed only if it is
	// addressable and was not obtained by
	// the use of unexported struct fields.
	if field.CanSet() {
		// change value of the field with name fieldName
		if field.Kind() == reflect.Struct {
			field.Set(reflect.ValueOf(fieldValue))
		}
	}

	return nil
}

func reflectUpdateValueOnField(iface interface{}, fieldValue string, fieldName string) error {
	valueIface := reflect.ValueOf(iface)

	// Check if the passed interface is a pointer
	if valueIface.Type().Kind() != reflect.Ptr {
		return fmt.Errorf("can't swap values if the reflected interface is not a pointer")
	}

	// 'dereference' with Elem() and get the field by name
	field := valueIface.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("interface `%s` does not have the field `%s`", valueIface.Type(), fieldName)
	}

	// A Value can be changed only if it is
	// addressable and was not obtained by
	// the use of unexported struct fields.
	if field.CanSet() {
		// change value of the field with name fieldName
		if field.Kind() == reflect.String {
			field.SetString(fieldValue)
		}
	}

	return nil
}
