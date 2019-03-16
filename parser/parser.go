package parser

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	. "github.com/runi95/wts-parser/models"
	"github.com/runi95/wts-parser/templates"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var idRegex = regexp.MustCompile(`[0-9]+`)                  // Find string id
var stringRegex = regexp.MustCompile(`STRING [0-9]*[^}]*}`) // Regex to find each string object
var contentContainerRegex = regexp.MustCompile(`{[^}]*}$`)  // Regex to find the content container of each string object
var contentStartRegex = regexp.MustCompile(`^[\r]*[\n]`)
var contextEndRegex = regexp.MustCompile(`[\r]*[\n]$`)
var SLKRegex = regexp.MustCompile(`C;X([0-9]+)(?:;Y([0-9]+))?;K([0-9.]+|"[\w-.,()\\]+")`)
var SLKHeadRegex = regexp.MustCompile(`C;X([0-9]+)(?:;Y([0-9]+))?;K"([\w-()]+)"`)
var SLKMetaRegex = regexp.MustCompile(`B;X([0-9]+);(?:Y([0-9]+);)D([-"\w]*)`)
var TXTHeadRegex = regexp.MustCompile(`^\[(\w+)]`)
var TXTRegex = regexp.MustCompile(`([A-Z]\w+)=([^\r\n]*)`)
var TrigstrRegex = regexp.MustCompile(`TRIGSTR_([0-9]+)`)
var NewLineRegex = regexp.MustCompile(`\r?\n`)

func WriteToFiles(customUnitFuncs *UnitFuncs, parsedSLKUnitsAbilities []*UnitAbilities, parsedSLKUnitsData []*UnitData, parsedSLKUnitsUI []*UnitUI, parsedSLKUnitsWeapons []*UnitWeapons, parsedSLKUnitsBalance []*UnitBalance) {
	WriteToFilesAndSaveToFolder(customUnitFuncs, parsedSLKUnitsAbilities, parsedSLKUnitsData, parsedSLKUnitsUI, parsedSLKUnitsWeapons, parsedSLKUnitsBalance, "out")
}

func WriteToFilesAndSaveToFolder(customUnitFuncs *UnitFuncs, parsedSLKUnitsAbilities []*UnitAbilities, parsedSLKUnitsData []*UnitData, parsedSLKUnitsUI []*UnitUI, parsedSLKUnitsWeapons []*UnitWeapons, parsedSLKUnitsBalance []*UnitBalance, outputFolder string) {
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 2
		},
	}

	log.Println("Writing to CampaignUnitFunc...")

	campaignUnitFuncFile, err := os.Create(outputFolder + "/CampaignUnitFunc.txt")
	if err != nil {
		log.Println(err)
		return
	}

	campaignUnitFuncTemplate := template.New("UnitFunc")
	campaignUnitFuncTemplate, err = campaignUnitFuncTemplate.Parse(templates.GetUnitFuncAndStringsTemplate())
	if err != nil {
		log.Println(err)
		return
	}

	err = campaignUnitFuncTemplate.ExecuteTemplate(campaignUnitFuncFile, "UnitFunc", customUnitFuncs.CampaignUnitFuncs)
	if err != nil {
		log.Println(err)
		return
	}

	if len(customUnitFuncs.HumanUnitFuncs) > 0 {
		log.Println("Writing to HumanUnitFunc...")

		humanUnitFuncFile, err := os.Create(outputFolder + "/HumanUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		humanUnitFuncTemplate := template.New("UnitFunc")
		humanUnitFuncTemplate, err = humanUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = humanUnitFuncTemplate.ExecuteTemplate(humanUnitFuncFile, "UnitFunc", customUnitFuncs.HumanUnitFuncs)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if len(customUnitFuncs.NightElfUnitFuncs) > 0 {
		log.Println("Writing to NightElfUnitFunc...")

		nightElfUnitFuncFile, err := os.Create(outputFolder + "/NightElfUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		nightElfUnitFuncTemplate := template.New("UnitFunc")
		nightElfUnitFuncTemplate, err = nightElfUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = nightElfUnitFuncTemplate.ExecuteTemplate(nightElfUnitFuncFile, "UnitFunc", customUnitFuncs.NightElfUnitFuncs)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if len(customUnitFuncs.OrcUnitFuncs) > 0 {
		log.Println("Writing to OrcUnitFunc...")

		orcUnitFuncFile, err := os.Create(outputFolder + "/OrcUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		orcUnitFuncTemplate := template.New("UnitFunc")
		orcUnitFuncTemplate, err = orcUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = orcUnitFuncTemplate.ExecuteTemplate(orcUnitFuncFile, "UnitFunc", customUnitFuncs.OrcUnitFuncs)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if len(customUnitFuncs.UndeadUnitFuncs) > 0 {
		log.Println("Writing to UndeadUnitFunc...")

		undeadUnitFuncFile, err := os.Create(outputFolder + "/UndeadUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		undeadUnitFuncTemplate := template.New("UnitFunc")
		undeadUnitFuncTemplate, err = undeadUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = undeadUnitFuncTemplate.ExecuteTemplate(undeadUnitFuncFile, "UnitFunc", customUnitFuncs.UndeadUnitFuncs)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if len(customUnitFuncs.NeutralUnitFuncs) > 0 {
		log.Println("Writing to NeutralUnitFunc...")

		neutralUnitFuncFile, err := os.Create(outputFolder + "/NeutralUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		neutralUnitFuncTemplate := template.New("UnitFunc")
		neutralUnitFuncTemplate, err = neutralUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = neutralUnitFuncTemplate.ExecuteTemplate(neutralUnitFuncFile, "UnitFunc", customUnitFuncs.NeutralUnitFuncs)
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("Writing to UnitAbilities.slk...")

	unitAbilitiesFile, err := os.Create(outputFolder + "/UnitAbilities.slk")
	if err != nil {
		log.Println(err)
		return
	}

	unitAbilitiesTemplate := template.New("UnitAbilities").Funcs(funcMap)
	unitAbilitiesTemplate, err = unitAbilitiesTemplate.Parse(templates.GetUnitAbilitiesTemplate())
	if err != nil {
		log.Println(err)
		return
	}

	err = unitAbilitiesTemplate.ExecuteTemplate(unitAbilitiesFile, "UnitAbilities", UnitAbilitiesTemplate{RowCount: len(parsedSLKUnitsAbilities) + 1, UnitAbilities: parsedSLKUnitsAbilities})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Writing to UnitData.slk...")

	unitDataFile, err := os.Create(outputFolder + "/UnitData.slk")
	if err != nil {
		log.Println(err)
		return
	}

	unitDataTemplate := template.New("UnitData").Funcs(funcMap)
	unitDataTemplate, err = unitDataTemplate.Parse(templates.GetUnitDataTemplate())
	if err != nil {
		log.Println(err)
		return
	}

	err = unitDataTemplate.ExecuteTemplate(unitDataFile, "UnitData", UnitDataTemplate{RowCount: len(parsedSLKUnitsData) + 1, UnitData: parsedSLKUnitsData})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Writing to UnitBalance.slk...")

	unitBalanceFile, err := os.Create(outputFolder + "/UnitBalance.slk")
	if err != nil {
		log.Println(err)
		return
	}

	unitBalanceTemplate := template.New("UnitBalance").Funcs(funcMap)
	unitBalanceTemplate, err = unitBalanceTemplate.Parse(templates.GetUnitBalanceTemplate())
	if err != nil {
		log.Println(err)
		return
	}

	err = unitBalanceTemplate.ExecuteTemplate(unitBalanceFile, "UnitBalance", UnitBalanceTemplate{RowCount: len(parsedSLKUnitsBalance) + 1, UnitBalance: parsedSLKUnitsBalance})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Writing to UnitUI.slk...")

	unitUIFile, err := os.Create(outputFolder + "/UnitUI.slk")
	if err != nil {
		log.Println(err)
		return
	}

	unitUITemplate := template.New("UnitUI").Funcs(funcMap)
	unitUITemplate, err = unitUITemplate.Parse(templates.GetUnitUITemplate())
	if err != nil {
		log.Println(err)
		return
	}

	err = unitUITemplate.ExecuteTemplate(unitUIFile, "UnitUI", UnitUITemplate{RowCount: len(parsedSLKUnitsUI) + 1, UnitUI: parsedSLKUnitsUI})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Writing to UnitWeapons.slk...")

	unitWeaponsFile, err := os.Create(outputFolder + "/UnitWeapons.slk")
	if err != nil {
		log.Println(err)
		return
	}

	unitWeaponsTemplate := template.New("UnitWeapons").Funcs(funcMap)
	unitWeaponsTemplate, err = unitWeaponsTemplate.Parse(templates.GetUnitWeaponsTemplate())
	if err != nil {
		log.Println(err)
		return
	}

	err = unitWeaponsTemplate.ExecuteTemplate(unitWeaponsFile, "UnitWeapons", UnitWeaponsTemplate{RowCount: len(parsedSLKUnitsWeapons) + 1, UnitWeapons: parsedSLKUnitsWeapons})
	if err != nil {
		log.Println(err)
		return
	}
}

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

func W3uToSlkUnitsWithBaseSlk(baseSLKUnits map[string]*SLKUnit, unitMap map[string]*W3uData) []*SLKUnit {
	slkUnits := make([]*SLKUnit, len(unitMap))
	index := 0
	for _, value := range unitMap {
		slkUnit := new(SLKUnit)
		baseSLKUnit := *baseSLKUnits["\""+value.BaseUnitId+"\""]
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

func W3uToTxtUnitFuncsWithBaseTxtAndBaseWts(baseTxtUnitFuncs map[string]*UnitFunc, baseSLKUnits map[string]*SLKUnit, unitMap map[string]*W3uData, wtsMap map[string]string) *UnitFuncs {
	unitFuncs := new(UnitFuncs)

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

		unitFunc := new(UnitFunc)
		baseTXTUnitFunc := *baseTxtUnitFuncs[value.BaseUnitId]

		var baseUnitFunc UnitFunc
		baseUnitFunc = baseTXTUnitFunc

		*unitFunc = baseUnitFunc

		value.TransformToUnitFunc(unitFunc)

		nameSubmatch := TrigstrRegex.FindStringSubmatch(unitFunc.Name.String)
		if len(nameSubmatch) > 0 {
			unitFunc.Name.SetValid(wtsMap[nameSubmatch[1]])
		}
		tipSubmatch := TrigstrRegex.FindStringSubmatch(unitFunc.Tip.String)
		if len(tipSubmatch) > 0 {
			unitFunc.Tip.SetValid(wtsMap[tipSubmatch[1]])
		}
		uberTipSubmatch := TrigstrRegex.FindStringSubmatch(unitFunc.Ubertip.String)
		if len(uberTipSubmatch) > 0 {
			unitFunc.Ubertip.SetValid("\"" + wtsMap[uberTipSubmatch[1]] + "\"")
		}
		descriptionSubmatch := TrigstrRegex.FindStringSubmatch(unitFunc.Description.String)
		if len(descriptionSubmatch) > 0 {
			unitFunc.Description.SetValid(wtsMap[descriptionSubmatch[1]])
		}
		hotkeySubmatch := TrigstrRegex.FindStringSubmatch(unitFunc.Hotkey.String)
		if len(hotkeySubmatch) > 0 {
			unitFunc.Hotkey.SetValid(wtsMap[hotkeySubmatch[1]])
		}
		editorSuffixSubmatch := TrigstrRegex.FindStringSubmatch(unitFunc.Editorsuffix.String)
		if len(editorSuffixSubmatch) > 0 {
			unitFunc.Editorsuffix.SetValid(wtsMap[editorSuffixSubmatch[1]])
		}

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

func W3uToTxtUnitFuncsWithBaseTxt(baseTXTUnitfuncs map[string]*UnitFunc, baseSLKUnits map[string]*SLKUnit, unitMap map[string]*W3uData) *UnitFuncs {
	unitFuncs := new(UnitFuncs)

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

		unitFunc := new(UnitFunc)
		baseTXTUnitFunc := *baseTXTUnitfuncs[value.BaseUnitId]

		var baseUnitFunc UnitFunc
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

func WtsToTxtUnitStringsWithBaseTxt(baseUnitStringMap map[string]*UnitString, baseSLKUnits map[string]*SLKUnit, unitMap map[string]*W3uData) *UnitStrings {
	unitStrings := new(UnitStrings)

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

		unitString := new(UnitString)
		baseTxtUnitString := *baseUnitStringMap[value.BaseUnitId]

		var baseUnitString UnitString
		baseUnitString = baseTxtUnitString

		*unitString = baseUnitString

		value.TransformToUnitString(unitString)

		switch unitRace {
		case "campaign":
			unitStrings.CampaignUnitFuncs = append(unitStrings.CampaignUnitFuncs, unitString)
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

       SLK PARSERS

*************************/

type SlkInformation struct {
	split          []string
	headerMap      map[string]string
	columnSize     int
	headerEndIndex int
}

func GenericSlkReader(input []byte) (*SlkInformation, error) {
	slkInformation := new(SlkInformation)
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
		headerLineSubmatch := SLKHeadRegex.FindStringSubmatch(split[headerLineIndex])
		if len(headerLineSubmatch[2]) > 0 && headerLineSubmatch[2] != "1" {
			isNotAtEndOfHeader = false
		} else {
			if len(headerLineSubmatch) > 1 {
				val := headerLineSubmatch[3]
				if val == "comment(s)" {
					headerMap[headerLineSubmatch[1]] = "comment"
				} else {
					headerMap[headerLineSubmatch[1]] = val
				}
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

func SlkToUnitAbilities(input []byte) map[string]*UnitAbilities {
	unitAbilitiesMap := make(map[string]*UnitAbilities)
	var currentUnitAbilities *UnitAbilities
	slkInformation, err := GenericSlkReader(input)
	if err != nil {
		log.Println(err)

		return nil
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		bodyLineSubmatch := SLKRegex.FindStringSubmatch(bodyLine)
		if bodyLineSubmatch != nil {
			if len(bodyLineSubmatch[2]) > 0 {
				if currentUnitAbilities != nil && currentUnitAbilities.UnitAbilID.Valid {
					unitAbilitiesMap[currentUnitAbilities.UnitAbilID.String] = currentUnitAbilities
				}

				currentUnitAbilities = new(UnitAbilities)
			}

			nullString := new(null.String)
			nullString.SetValid(bodyLineSubmatch[3])

			err = reflectUpdateValueOnFieldNullStruct(currentUnitAbilities, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	if currentUnitAbilities != nil && currentUnitAbilities.UnitAbilID.Valid {
		unitAbilitiesMap[currentUnitAbilities.UnitAbilID.String] = currentUnitAbilities
	}

	return unitAbilitiesMap
}

func SlkToUnitData(input []byte) map[string]*UnitData {
	unitDataMap := make(map[string]*UnitData)
	var currentUnitData *UnitData
	slkInformation, err := GenericSlkReader(input)
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

			err = reflectUpdateValueOnFieldNullStruct(currentUnitData, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	if currentUnitData != nil && currentUnitData.UnitID.Valid {
		unitDataMap[currentUnitData.UnitID.String] = currentUnitData
	}

	return unitDataMap
}

func SLKToUnitUI(input []byte) map[string]*UnitUI {
	unitUIMap := make(map[string]*UnitUI)
	var currentUnitUI *UnitUI
	slkInformation, err := GenericSlkReader(input)
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

			err = reflectUpdateValueOnFieldNullStruct(currentUnitUI, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	if currentUnitUI != nil && currentUnitUI.UnitUIID.Valid {
		unitUIMap[currentUnitUI.UnitUIID.String] = currentUnitUI
	}

	return unitUIMap
}

func SLKToUnitWeapons(input []byte) map[string]*UnitWeapons {
	unitWeaponsMap := make(map[string]*UnitWeapons)
	var currentUnitWeapons *UnitWeapons
	slkInformation, err := GenericSlkReader(input)
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

			err = reflectUpdateValueOnFieldNullStruct(currentUnitWeapons, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	if currentUnitWeapons != nil && currentUnitWeapons.UnitWeapID.Valid {
		unitWeaponsMap[currentUnitWeapons.UnitWeapID.String] = currentUnitWeapons
	}

	return unitWeaponsMap
}

func SLKToUnitBalance(input []byte) map[string]*UnitBalance {
	unitBalanceMap := make(map[string]*UnitBalance)
	var currentUnitBalance *UnitBalance
	slkInformation, err := GenericSlkReader(input)
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

			err = reflectUpdateValueOnFieldNullStruct(currentUnitBalance, *nullString, strings.Title(slkInformation.headerMap[bodyLineSubmatch[1]]))
			if err != nil {
				log.Println(err)
			}
		}
	}

	if currentUnitBalance != nil && currentUnitBalance.UnitBalanceID.Valid {
		unitBalanceMap[currentUnitBalance.UnitBalanceID.String] = currentUnitBalance
	}

	return unitBalanceMap
}

/*************************

       TXT PARSERS

*************************/

func TxtToUnitFunc(input []byte) map[string]*UnitFunc {
	unitFuncMap := make(map[string]*UnitFunc)
	var currentUnitFunc *UnitFunc

	str := string(input)
	split := strings.Split(str, "\n")

	for _, line := range split {
		if TXTHeadRegex.MatchString(line) {
			if currentUnitFunc != nil {
				unitFuncMap[currentUnitFunc.UnitId] = currentUnitFunc
			}

			currentUnitFunc = new(UnitFunc)
			lineSubmatch := TXTHeadRegex.FindStringSubmatch(line)
			currentUnitFunc.UnitId = lineSubmatch[1]
		} else {
			lineSubmatch := TXTRegex.FindStringSubmatch(line)
			if lineSubmatch != nil {
				nullString := new(null.String)
				nullString.SetValid(lineSubmatch[2])

				err := reflectUpdateValueOnFieldNullStruct(currentUnitFunc, *nullString, strings.Title(strings.ToLower(lineSubmatch[1])))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

	if currentUnitFunc != nil {
		unitFuncMap[currentUnitFunc.UnitId] = currentUnitFunc
	}

	return unitFuncMap
}

func TxtToUnitStrings(input []byte) map[string]*UnitString {
	unitStringsMap := make(map[string]*UnitString)
	var currentUnitString *UnitString

	str := string(input)
	split := strings.Split(str, "\n")

	for _, line := range split {
		if TXTHeadRegex.MatchString(line) {
			if currentUnitString != nil {
				unitStringsMap[currentUnitString.UnitId] = currentUnitString
			}

			currentUnitString = new(UnitString)
			lineSubmatch := TXTHeadRegex.FindStringSubmatch(line)
			currentUnitString.UnitId = lineSubmatch[1]
		} else {
			lineSubmatch := TXTRegex.FindStringSubmatch(line)
			if lineSubmatch != nil {
				nullString := new(null.String)
				nullString.SetValid(lineSubmatch[2])

				err := reflectUpdateValueOnFieldNullStruct(currentUnitString, *nullString, strings.Title(strings.ToLower(lineSubmatch[1])))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

	if currentUnitString != nil {
		unitStringsMap[currentUnitString.UnitId] = currentUnitString
	}

	return unitStringsMap
}

/*************************

    PRIVATE FUNCTIONS

*************************/

func reflectUpdateValueOnFieldNullStruct(iface interface{}, fieldValue interface{}, fieldName string) error {
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
