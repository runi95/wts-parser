package parser

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/runi95/wts-parser/models"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"math"
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
var NewLineRegex = regexp.MustCompile(`\r?\n`)

/*************************

	   WTS PARSERS

*************************/

func ReadWtsFile(input []byte) map[string]string {
	str := string(input)

	findAllStrings := stringRegex.FindAllString(str, -1)
	wtsMap := make(map[string]string)

	for _, stringObject := range findAllStrings {
		findContentContainer := contentContainerRegex.FindString(stringObject)
		idStr := idRegex.FindString(stringObject)

		removeBrackets := findContentContainer[1 : len(findContentContainer)-1]
		length := len(contentStartRegex.FindString(removeBrackets))
		parsedBeginning := removeBrackets[length:]
		lengthToEnd := len(parsedBeginning) - len(contextEndRegex.FindString(parsedBeginning))
		parsedContainer := parsedBeginning[:lengthToEnd]
		parsedContainer = NewLineRegex.ReplaceAllString(parsedContainer, "|n")

		wtsMap[idStr] = parsedContainer
	}

	return wtsMap
}

func ReadWtsFileAndUseIntAsKey(input []byte) map[int]string {
	str := string(input)

	findAllStrings := stringRegex.FindAllString(str, -1)
	wtsMap := make(map[int]string)

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
		strings.Replace(parsedContainer, "\n", "|n", -1)

		wtsMap[id] = parsedContainer
	}

	return wtsMap
}

func WtsToJson(input []byte) []byte {
	wtsMap := ReadWtsFile(input)

	jsonObject, err := json.Marshal(wtsMap)
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

func ReadW3uFile(input []byte) map[string]*models.W3uData {
	var lastUnitId string
	unitMap := make(map[string]*models.W3uData)

	for i := 12; i < len(input); i++ {
		baseUnitId := [4]byte{input[i-8], input[i-7], input[i-6], input[i-5]}
		if baseUnitId[0] > 96 && baseUnitId[0] < 123 && baseUnitId[1] > 96 && baseUnitId[1] < 123 && baseUnitId[2] > 96 && baseUnitId[2] < 123 && ((baseUnitId[3] > 96 && baseUnitId[3] < 123) || (baseUnitId[3] > 47 && baseUnitId[3] < 58)) {
			customUnitId := [4]byte{input[i-4], input[i-3], input[i-2], input[i-1]}
			if customUnitId[0] > 96 && customUnitId[0] < 123 && (customUnitId[1] == 48 || customUnitId[1] == 67) && ((customUnitId[2] > 47 && customUnitId[2] < 58) || (customUnitId[2] > 64 && customUnitId[2] < 91)) && ((customUnitId[3] > 47 && customUnitId[3] < 58) || (customUnitId[3] > 64 && customUnitId[3] < 91)) {
				unitId := string(customUnitId[0]) + string(customUnitId[1]) + string(customUnitId[2]) + string(customUnitId[3])
				lastUnitId = unitId
				newW3uUnit := new(models.W3uData)
				newW3uUnit.BaseUnitId = string(baseUnitId[0]) + string(baseUnitId[1]) + string(baseUnitId[2]) + string(baseUnitId[3])
				newW3uUnit.CustomUnitId = string(customUnitId[0]) + string(customUnitId[1]) + string(customUnitId[2]) + string(customUnitId[3])
				unitMap[unitId] = newW3uUnit
			}
		}

		if input[i] > 31 && input[i] < 127 && input[i-1] > 31 && input[i-1] < 127 && input[i-2] > 31 && input[i-2] < 127 && input[i-3] > 31 && input[i-3] < 127 {
			str := strings.Title(string(input[i-3]) + string(input[i-2]) + string(input[i-1]) + string(input[i]))
			currentUnit := unitMap[lastUnitId]
			if currentUnit != nil && reflectIsFieldNamePartOfStruct(currentUnit, str) {
				metaByte := input[i+1]
				i += 4
				switch metaByte {
				case 3:
					stringValue := ""
					for input[i+1] > 31 && input[i+1] < 127 {
						i++
						stringValue += string(input[i])
					}

					if stringValue == "" {
						stringValue = "-"
					}

					nullString := new(null.String)
					nullString.SetValid(stringValue)

					err := reflectUpdateValueOnFieldNullStruct(currentUnit, *nullString, str)
					if err != nil {
						log.Println(err)
					}
				case 2: // Don't know the difference between 1 and 2
					bytes := make([]byte, 4)
					i++
					bytes[0] = input[i]
					i++
					bytes[1] = input[i]
					i++
					bytes[2] = input[i]
					i++
					bytes[3] = input[i]

					bits := binary.LittleEndian.Uint32(bytes)
					float := math.Float32frombits(bits)

					nullString := new(null.String)
					nullString.SetValid(fmt.Sprint(float))

					err := reflectUpdateValueOnFieldNullStruct(currentUnit, *nullString, str)
					if err != nil {
						log.Println(err)
					}
				case 1: // Don't know the difference between 1 and 2
					bytes := make([]byte, 4)
					i++
					bytes[0] = input[i]
					i++
					bytes[1] = input[i]
					i++
					bytes[2] = input[i]
					i++
					bytes[3] = input[i]

					bits := binary.LittleEndian.Uint32(bytes)
					float := math.Float32frombits(bits)

					nullString := new(null.String)
					nullString.SetValid(fmt.Sprint(float))

					err := reflectUpdateValueOnFieldNullStruct(currentUnit, *nullString, str)
					if err != nil {
						log.Println(err)
					}
				case 0:
					bytes := make([]byte, 4)
					i++
					bytes[0] = input[i]
					i++
					bytes[1] = input[i]
					i++
					bytes[2] = input[i]
					i++
					bytes[3] = input[i]

					uint := binary.LittleEndian.Uint32(bytes)

					nullString := new(null.String)
					nullString.SetValid(fmt.Sprint(uint))

					err := reflectUpdateValueOnFieldNullStruct(currentUnit, *nullString, str)
					if err != nil {
						log.Println(err)
					}
				default:
					log.Println(fmt.Sprint(metaByte) + " is an unknown byte type for " + str)
				}
			}
		}
	}

	return unitMap
}

func W3uToJson(input []byte) []byte {
	unitMap := ReadW3uFile(input)

	mapValues := make([]*models.W3uData, len(unitMap))
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

func W3uToSLKUnits(input []byte) []*models.SLKUnit {
	unitMap := ReadW3uFile(input)

	slkUnits := make([]*models.SLKUnit, len(unitMap))
	index := 0
	for _, value := range unitMap {
		slkUnit := new(models.SLKUnit)
		value.TransformToSLKUnit(slkUnit)

		slkUnits[index] = slkUnit
		index++
	}

	return slkUnits
}

func W3uToSlkUnitsWithBaseSlk(baseSLKUnits map[string]*models.SLKUnit, unitMap map[string]*models.W3uData) []*models.SLKUnit {
	slkUnits := make([]*models.SLKUnit, len(unitMap))
	index := 0
	for _, value := range unitMap {
		slkUnit := new(models.SLKUnit)
		baseSLKUnit := *baseSLKUnits["\""+value.BaseUnitId+"\""]
		var unitUI models.UnitUI
		var unitData models.UnitData
		var unitWeapons models.UnitWeapons
		var unitBalance models.UnitBalance
		var unitAbilities models.UnitAbilities
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

func W3uToTxtUnitFuncsWithBaseTxtAndBaseWts(baseTxtUnitFuncs map[string]*models.UnitFunc, baseSLKUnits map[string]*models.SLKUnit, unitMap map[string]*models.W3uData, wtsMap map[string]string) *models.UnitFuncs {
	unitFuncs := new(models.UnitFuncs)

	for _, value := range unitMap {
		var unitRace string
		baseSlkUnit := baseSLKUnits["\""+value.BaseUnitId+"\""]
		if baseSlkUnit.UnitUI.Campaign.Valid && baseSlkUnit.UnitUI.Campaign.String == "1" {
			unitRace = "campaign"
		} else if value.Ucam.Valid && value.Ucam.String == "1" {
			unitRace = "campaign"
		} else {
			if value.Urac.Valid {
				unitRace = value.Urac.String
			} else {
				rawUnitRace := baseSlkUnit.UnitData.Race.String
				unitRace = strings.Replace(rawUnitRace, "\"", "", -1)
			}
		}

		//  We just want everything in the campaign file for now...
		unitRace = "campaign"

		unitFunc := new(models.UnitFunc)
		baseTXTUnitFunc := *baseTxtUnitFuncs[value.BaseUnitId]

		var baseUnitFunc models.UnitFunc
		baseUnitFunc = baseTXTUnitFunc

		*unitFunc = baseUnitFunc

		value.TransformToUnitFunc(unitFunc)

		switch unitRace {
		case "campaign":
			unitFuncs.CampaignUnitFuncs = append(unitFuncs.CampaignUnitFuncs, unitFunc)
		case "human":
			unitFuncs.HumanUnitFuncs = append(unitFuncs.HumanUnitFuncs, unitFunc)
		case "nightelf":
			unitFuncs.NightElfUnitFuncs = append(unitFuncs.NightElfUnitFuncs, unitFunc)
		case "orc":
			unitFuncs.OrcUnitFuncs = append(unitFuncs.OrcUnitFuncs, unitFunc)
		case "undead":
			unitFuncs.UndeadUnitFuncs = append(unitFuncs.UndeadUnitFuncs, unitFunc)
		case "naga":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		case "creeps":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		case "other":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		case "unknown":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		default:
			log.Println(fmt.Sprintf("Can't find the race %s", unitRace))
		}
	}

	return unitFuncs
}

func W3uToTxtUnitFuncsWithBaseTxt(baseTXTUnitfuncs map[string]*models.UnitFunc, baseSLKUnits map[string]*models.SLKUnit, unitMap map[string]*models.W3uData) *models.UnitFuncs {
	unitFuncs := new(models.UnitFuncs)

	for _, value := range unitMap {
		var unitRace string
		baseSlkUnit := baseSLKUnits["\""+value.BaseUnitId+"\""]
		if baseSlkUnit.UnitUI.Campaign.Valid && baseSlkUnit.UnitUI.Campaign.String == "1" {
			unitRace = "campaign"
		} else if value.Ucam.Valid && value.Ucam.String == "1" {
			unitRace = "campaign"
		} else {
			if value.Urac.Valid {
				unitRace = value.Urac.String
			} else {
				rawUnitRace := baseSlkUnit.UnitData.Race.String
				unitRace = strings.Replace(rawUnitRace, "\"", "", -1)
			}
		}

		unitFunc := new(models.UnitFunc)
		baseTXTUnitFunc := *baseTXTUnitfuncs[value.BaseUnitId]

		var baseUnitFunc models.UnitFunc
		baseUnitFunc = baseTXTUnitFunc

		*unitFunc = baseUnitFunc

		value.TransformToUnitFunc(unitFunc)

		switch unitRace {
		case "campaign":
			unitFuncs.CampaignUnitFuncs = append(unitFuncs.CampaignUnitFuncs, unitFunc)
		case "human":
			unitFuncs.HumanUnitFuncs = append(unitFuncs.HumanUnitFuncs, unitFunc)
		case "nightelf":
			unitFuncs.NightElfUnitFuncs = append(unitFuncs.NightElfUnitFuncs, unitFunc)
		case "orc":
			unitFuncs.OrcUnitFuncs = append(unitFuncs.OrcUnitFuncs, unitFunc)
		case "undead":
			unitFuncs.UndeadUnitFuncs = append(unitFuncs.UndeadUnitFuncs, unitFunc)
		case "naga":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		case "creeps":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		case "other":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		case "unknown":
			unitFuncs.NeutralUnitFuncs = append(unitFuncs.NeutralUnitFuncs, unitFunc)
		default:
			log.Println(fmt.Sprintf("Can't find the race %s", unitRace))
		}
	}

	return unitFuncs
}

func WtsToTxtUnitStringsWithBaseTxt(baseUnitStringMap map[string]*models.UnitString, baseSLKUnits map[string]*models.SLKUnit, unitMap map[string]*models.W3uData) *models.UnitStrings {
	unitStrings := new(models.UnitStrings)

	for _, value := range unitMap {
		var unitRace string
		baseSlkUnit := baseSLKUnits["\""+value.BaseUnitId+"\""]
		if baseSlkUnit.UnitUI.Campaign.Valid && baseSlkUnit.UnitUI.Campaign.String == "1" {
			unitRace = "campaign"
		} else {
			if value.Urac.Valid {
				unitRace = value.Urac.String
			} else {
				rawUnitRace := baseSlkUnit.UnitData.Race.String
				unitRace = strings.Replace(rawUnitRace, "\"", "", -1)
			}
		}

		unitString := new(models.UnitString)
		baseTxtUnitString := *baseUnitStringMap[value.BaseUnitId]

		var baseUnitString models.UnitString
		baseUnitString = baseTxtUnitString

		*unitString = baseUnitString

		value.TransformToUnitString(unitString)

		switch unitRace {
		case "campaign":
			unitStrings.CampaignUnitStrings = append(unitStrings.CampaignUnitStrings, unitString)
		case "human":
			unitStrings.HumanUnitStrings = append(unitStrings.HumanUnitStrings, unitString)
		case "nightelf":
			unitStrings.NightElfUnitStrings = append(unitStrings.NightElfUnitStrings, unitString)
		case "orc":
			unitStrings.OrcUnitStrings = append(unitStrings.OrcUnitStrings, unitString)
		case "undead":
			unitStrings.UndeadUnitStrings = append(unitStrings.UndeadUnitStrings, unitString)
		case "naga":
			unitStrings.NeutralUnitStrings = append(unitStrings.NeutralUnitStrings, unitString)
		case "creeps":
			unitStrings.NeutralUnitStrings = append(unitStrings.NeutralUnitStrings, unitString)
		case "other":
			unitStrings.NeutralUnitStrings = append(unitStrings.NeutralUnitStrings, unitString)
		case "unknown":
			unitStrings.NeutralUnitStrings = append(unitStrings.NeutralUnitStrings, unitString)
		default:
			log.Println(fmt.Sprintf("Can't find the race %s", unitRace))
		}
	}

	return unitStrings
}

/*************************

	PRIVATE FUNCTIONS

*************************/

func reflectUpdateValueOnFieldNullStruct(iface interface{}, fieldValue interface{}, fieldName string) error {
	if iface == nil {
		return fmt.Errorf("iface cannot be nil")
	}

	if fieldValue == nil {
		return fmt.Errorf("fieldValue cannot be nil")
	}

	if fieldName == "" {
		return fmt.Errorf("fieldName cannot be empty")
	}

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

func reflectIsFieldNamePartOfStruct(iface interface{}, fieldName string) bool {
	valueIface := reflect.ValueOf(iface)

	// Check if the passed interface is a pointer
	if valueIface.Type().Kind() != reflect.Ptr {
		return false
	}

	// 'dereference' with Elem() and get the field by name
	field := valueIface.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return false
	}

	return true
}

func reflectTypeFromFieldName(iface interface{}, fieldName string) reflect.Type {
	valueIface := reflect.ValueOf(iface)

	// Check if the passed interface is a pointer
	if valueIface.Type().Kind() != reflect.Ptr {
		return nil
	}

	// 'dereference' with Elem() and get the field by name
	field := valueIface.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}

	return field.Type()
}
