package parser

import (
	"fmt"
	"github.com/runi95/wts-parser/models"
	"github.com/runi95/wts-parser/templates"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

/*************************

	   VARIABLES

*************************/

var xReg = regexp.MustCompile(`^X([0-9]+)$`)
var yReg = regexp.MustCompile(`^Y([0-9]+)$`)
var kReg = regexp.MustCompile(`^K.*`)
var unitIdReg = regexp.MustCompile(`^\[(\w+)]`)
var keyNameReg = regexp.MustCompile(`^(.+)=`)

/*************************

	   STRUCTURES

*************************/

type SlkInformation struct {
	split          []string
	headerMap      map[string]string
	columnSize     int
	headerEndIndex int
}

/*************************

	   SLK PARSERS

*************************/

func WriteToFiles(unitList []*models.SLKUnit) {
	WriteToFilesAndSaveToFolder(unitList, "out", true)
}

func WriteToFilesAndSaveToFolder(unitList []*models.SLKUnit, outputFolder string, sortBeforeSave bool) {
	if len(unitList) < 1 {
		return
	}

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 2
		},
	}

	if sortBeforeSave {
		log.Println("Sorting units according to unit ID...")

		sort.Slice(unitList, func(i, j int) bool {
			return unitList[i].UnitID.String < unitList[j].UnitID.String
		})
	}

	var writeToFileWaitGroup sync.WaitGroup
	log.Println("Writing to CampaignUnitFunc...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		campaignUnitFuncFile, err := os.Create(outputFolder + "/CampaignUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				return isCampaign == "1"
			},
		}
		campaignUnitFuncTemplate := template.New("UnitFunc").Funcs(customFuncMap)
		campaignUnitFuncTemplate, err = campaignUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = campaignUnitFuncTemplate.ExecuteTemplate(campaignUnitFuncFile, "UnitFunc", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to CampaignUnitStrings...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		campaignUnitStringsFile, err := os.Create(outputFolder + "/CampaignUnitStrings.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				return isCampaign == "1"
			},
		}
		campaignUnitStringsTemplate := template.New("UnitStrings").Funcs(customFuncMap)
		campaignUnitStringsTemplate, err = campaignUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = campaignUnitStringsTemplate.ExecuteTemplate(campaignUnitStringsFile, "UnitStrings", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to HumanUnitFunc...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		humanUnitFuncFile, err := os.Create(outputFolder + "/HumanUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "human"
			},
		}
		humanUnitFuncTemplate := template.New("UnitFunc").Funcs(customFuncMap)
		humanUnitFuncTemplate, err = humanUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = humanUnitFuncTemplate.ExecuteTemplate(humanUnitFuncFile, "UnitFunc", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to HumanUnitStrings...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		humanUnitStringsFile, err := os.Create(outputFolder + "/HumanUnitStrings.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "human"
			},
		}
		humanUnitStringsTemplate := template.New("UnitStrings").Funcs(customFuncMap)
		humanUnitStringsTemplate, err = humanUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = humanUnitStringsTemplate.ExecuteTemplate(humanUnitStringsFile, "UnitStrings", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to NightElfUnitFunc...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		nightElfUnitFuncFile, err := os.Create(outputFolder + "/NightElfUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "nightelf"
			},
		}
		nightElfUnitFuncTemplate := template.New("UnitFunc").Funcs(customFuncMap)
		nightElfUnitFuncTemplate, err = nightElfUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = nightElfUnitFuncTemplate.ExecuteTemplate(nightElfUnitFuncFile, "UnitFunc", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to NightElfUnitStrings...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		nightElfUnitStringsFile, err := os.Create(outputFolder + "/NightElfUnitStrings.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "nightelf"
			},
		}
		nightElfUnitStringsTemplate := template.New("UnitStrings").Funcs(customFuncMap)
		nightElfUnitStringsTemplate, err = nightElfUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = nightElfUnitStringsTemplate.ExecuteTemplate(nightElfUnitStringsFile, "UnitStrings", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to OrcUnitFunc...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		orcUnitFuncFile, err := os.Create(outputFolder + "/OrcUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "orc"
			},
		}
		orcUnitFuncTemplate := template.New("UnitFunc").Funcs(customFuncMap)
		orcUnitFuncTemplate, err = orcUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = orcUnitFuncTemplate.ExecuteTemplate(orcUnitFuncFile, "UnitFunc", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to OrcUnitStrings...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		orcUnitStringsFile, err := os.Create(outputFolder + "/OrcUnitStrings.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "orc"
			},
		}
		orcUnitStringsTemplate := template.New("UnitStrings").Funcs(customFuncMap)
		orcUnitStringsTemplate, err = orcUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = orcUnitStringsTemplate.ExecuteTemplate(orcUnitStringsFile, "UnitStrings", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UndeadUnitFunc...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		undeadUnitFuncFile, err := os.Create(outputFolder + "/UndeadUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "undead"
			},
		}
		undeadUnitFuncTemplate := template.New("UnitFunc").Funcs(customFuncMap)
		undeadUnitFuncTemplate, err = undeadUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = undeadUnitFuncTemplate.ExecuteTemplate(undeadUnitFuncFile, "UnitFunc", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UndeadUnitStrings...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		undeadUnitStringsFile, err := os.Create(outputFolder + "/UndeadUnitStrings.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				return isCampaign != "1" && strings.ToLower(trimmedRace) == "undead"
			},
		}
		undeadUnitStringsTemplate := template.New("UnitStrings").Funcs(customFuncMap)
		undeadUnitStringsTemplate, err = undeadUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = undeadUnitStringsTemplate.ExecuteTemplate(undeadUnitStringsFile, "UnitStrings", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to NeutralUnitFunc...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		neutralUnitFuncFile, err := os.Create(outputFolder + "/NeutralUnitFunc.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				lowercaseRace := strings.ToLower(trimmedRace)
				return isCampaign != "1" && lowercaseRace != "human" && lowercaseRace != "nightelf" && lowercaseRace != "orc" && lowercaseRace != "undead"
			},
		}
		neutralUnitFuncTemplate := template.New("UnitFunc").Funcs(customFuncMap)
		neutralUnitFuncTemplate, err = neutralUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = neutralUnitFuncTemplate.ExecuteTemplate(neutralUnitFuncFile, "UnitFunc", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to NeutralUnitStrings...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
		neutralUnitStringsFile, err := os.Create(outputFolder + "/NeutralUnitStrings.txt")
		if err != nil {
			log.Println(err)
			return
		}

		customFuncMap := template.FuncMap{
			"campaignAndRaceCheck": func(race string, isCampaign string) bool {
				trimmedRace := strings.Replace(race, "\"", "", -1)
				lowercaseRace := strings.ToLower(trimmedRace)
				return isCampaign != "1" && lowercaseRace != "human" && lowercaseRace != "nightelf" && lowercaseRace != "orc" && lowercaseRace != "undead"
			},
		}
		neutralUnitStringsTemplate := template.New("UnitStrings").Funcs(customFuncMap)
		neutralUnitStringsTemplate, err = neutralUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = neutralUnitStringsTemplate.ExecuteTemplate(neutralUnitStringsFile, "UnitStrings", unitList)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UnitAbilities.slk...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
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

		err = unitAbilitiesTemplate.ExecuteTemplate(unitAbilitiesFile, "UnitAbilities", models.RowCountTemplate{RowCount: len(unitList) + 1, Unit: unitList})
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UnitData.slk...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
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

		err = unitDataTemplate.ExecuteTemplate(unitDataFile, "UnitData", models.RowCountTemplate{RowCount: len(unitList) + 1, Unit: unitList})
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UnitBalance.slk...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
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

		err = unitBalanceTemplate.ExecuteTemplate(unitBalanceFile, "UnitBalance", models.RowCountTemplate{RowCount: len(unitList) + 1, Unit: unitList})
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UnitUI.slk...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
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

		err = unitUITemplate.ExecuteTemplate(unitUIFile, "UnitUI", models.RowCountTemplate{RowCount: len(unitList) + 1, Unit: unitList})
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Writing to UnitWeapons.slk...")
	writeToFileWaitGroup.Add(1)

	go func() {
		defer writeToFileWaitGroup.Done()
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

		err = unitWeaponsTemplate.ExecuteTemplate(unitWeaponsFile, "UnitWeapons", models.RowCountTemplate{RowCount: len(unitList) + 1, Unit: unitList})
		if err != nil {
			log.Println(err)
			return
		}
	}()

	writeToFileWaitGroup.Wait()
}

/*
bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		var x *string
		var y *string
		var k *string

		cleanBodyLine := strings.Replace(strings.Replace(bodyLine, "\r", "", -1), "\n", "", -1)
		bodySplit := strings.Split(cleanBodyLine, ";")
		for _, s := range bodySplit {
			if xReg.MatchString(s) {
				newX := s[1:]
				x = &newX
			} else if kReg.MatchString(s) {
				newK := s[1:]
				k = &newK
			} else if yReg.MatchString(s) {
				newY := s[1:]
				y = &newY
			}
		}

		if x != nil && k != nil {
			if y != nil { // y != nil means we've reached the beginning of a new unit
				trimmedId := strings.Replace(*k, "\"", "", -1)
				currentUnitId = &trimmedId

				// Check if unitMap does not already have this key
				if _, ok := unitMap[trimmedId]; !ok {
					// New unit
				}
			}

			if currentUnitId != nil {
				// Do stuff
			}
		}
	}
*/

func GenericSlkReader(input []byte) (*SlkInformation, error) {
	slkInformation := new(SlkInformation)
	str := string(input)
	split := strings.Split(str, "\n")
	metaSplit := strings.Split(split[1], ";")

	var metaY *string
	for _, s := range metaSplit {
		cleanString := strings.Replace(strings.Replace(s, "\r", "", -1), "\n", "", -1)
		if yReg.MatchString(cleanString) {
			newY := cleanString[1:]
			metaY = &newY
		}
	}

	if metaY == nil {
		return nil, fmt.Errorf("meta match is nil")
	}

	columnSize, err := strconv.Atoi(*metaY)
	if err != nil {
		return nil, err
	}

	headerMap := make(map[string]string)
	headerLineIndex := 2
	isNotAtEndOfHeader := true
	for isNotAtEndOfHeader {
		headerSplit := strings.Split(split[headerLineIndex], ";")
		var x *string
		var y *string
		var k *string

		for _, s := range headerSplit {
			cleanString := strings.Replace(strings.Replace(s, "\r", "", -1), "\n", "", -1)
			if xReg.MatchString(cleanString) {
				newX := cleanString[1:]
				x = &newX
			} else if kReg.MatchString(cleanString) {
				newK := cleanString[1:]
				k = &newK
			} else if yReg.MatchString(cleanString) {
				newY := cleanString[1:]
				y = &newY
			}
		}

		if y != nil && *y != "1" {
			isNotAtEndOfHeader = false
		} else {
			if *k == "comment(s)" {
				headerMap[*x] = "comment"
			} else {
				headerMap[*x] = *k
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

/**
 * A function that populates a map of SLKUnit structures with data from any SLK file like the UnitData.slk, UnitBalance.slk and UnitUI.slk files.
 * If you want to add file data from a TXT file like CampaignUnitFunc.txt you need to run PopulateUnitMapWithTxtFileData.
 * To populate a map with values from multiple files you need to call this function several times like in the code example below
 *
 * func populateMapWithDataFromMultipleFiles() error {
 *   unitDataBytes, err := ioutil.ReadFile("./UnitData.slk")
 *   if err != nil {
 *     return err
 *   }
 *
 *   unitBalanceBytes, err := ioutil.ReadFile("./UnitBalance.slk")
 *   if err != nil {
 *     return err
 *   }
 *
 *   var unitMap = make(map[string]*models.SLKUnit)
 *   parser.PopulateUnitMapWithSlkFileData(unitDataBytes, unitMap)
 *   parser.PopulateUnitMapWithSlkFileData(unitBalanceBytes, unitMap)
 *
 *   // You can also add data from txt files like seen below
 *   campaignUnitFuncBytes, err := ioutil.ReadFile("./CampaignUnitFunc.txt")
 *   if err != nil {
 *     return err
 *   }
 *
 *   // Note that this is not the same function as used above!
 *   parser.PopulateUnitMapWithTxtFileData(campaignUnitFuncBytes, unitMap)
 *
 *   return nil
 * }
 *
 */
func PopulateUnitMapWithSlkFileData(inputFileData []byte, unitMap map[string]*models.SLKUnit) {
	var currentUnitId *string
	slkInformation, err := GenericSlkReader(inputFileData)
	if err != nil {
		log.Println(err)
		return
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		var x *string
		var y *string
		var k *string

		cleanBodyLine := strings.Replace(strings.Replace(bodyLine, "\r", "", -1), "\n", "", -1)
		bodySplit := strings.Split(cleanBodyLine, ";")
		for _, s := range bodySplit {
			if xReg.MatchString(s) {
				newX := s[1:]
				x = &newX
			} else if kReg.MatchString(s) {
				newK := s[1:]
				k = &newK
			} else if yReg.MatchString(s) {
				newY := s[1:]
				y = &newY
			}
		}

		if x != nil && k != nil {
			if y != nil { // y != nil means we've reached the beginning of a new unit
				trimmedId := strings.Replace(*k, "\"", "", -1)
				currentUnitId = &trimmedId

				// Check if unitMap does not already have this key
				if _, ok := unitMap[trimmedId]; !ok {
					newUnit := new(models.SLKUnit)
					newUnit.UnitUI = new(models.UnitUI)
					newUnit.UnitData = new(models.UnitData)
					newUnit.UnitBalance = new(models.UnitBalance)
					newUnit.UnitWeapons = new(models.UnitWeapons)
					newUnit.UnitAbilities = new(models.UnitAbilities)
					newUnit.UnitFunc = new(models.UnitFunc)
					newUnit.UnitString = new(models.UnitString)

					unitMap[trimmedId] = newUnit
				}
			}

			if currentUnitId != nil {
				nullString := new(null.String)
				nullString.SetValid(*k)

				if unit, ok := unitMap[*currentUnitId]; ok {
					if header, headerMapOk := slkInformation.headerMap[*x]; headerMapOk {
						err = reflectUpdateValueOnFieldNullStruct(unit, *nullString, strings.Title(header))
						if err != nil {
							log.Println(err)
						}
					}
				}
			}
		}
	}
}

/**
 * A function that populates a map of SLKUnit structures with data from any TXT file like the CampaignUnitFunc.txt, CampaignUnitStrings.txt and HumanUnitFunc.txt files.
 * If you want to add file data from an SLK file like UnitData.slk you need to run PopulateUnitMapWithSlkFileData.
 * To populate a map with values from multiple files you need to call this function several times like in the code example below
 *
 * func populateMapWithDataFromMultipleFiles() error {
 *   campaignUnitFuncBytes, err := ioutil.ReadFile("./CampaignUnitFunc.txt")
 *   if err != nil {
 *     return err
 *   }
 *
 *   campaignUnitStringsBytes, err := ioutil.ReadFile("./CampaignUnitStrings.txt")
 *   if err != nil {
 *     return err
 *   }
 *
 *   var unitMap = make(map[string]*models.SLKUnit)
 *   parser.PopulateUnitMapWithTxtFileData(campaignUnitFuncBytes, unitMap)
 *   parser.PopulateUnitMapWithTxtFileData(campaignUnitStringsBytes, unitMap)
 *
 *   // You can also add data from slk files like seen below
 *   unitDataBytes, err := ioutil.ReadFile("./UnitData.slk")
 *   if err != nil {
 *     return err
 *   }
 *
 *   // Note that this is not the same function as used above!
 *   parser.PopulateUnitMapWithSlkFileData(unitDataBytes, unitMap)
 *
 *   return nil
 * }
 *
 */
func PopulateUnitMapWithTxtFileData(inputFileData []byte, unitMap map[string]*models.SLKUnit) {
	var currentUnitId *string

	str := string(inputFileData)
	split := strings.Split(str, "\n")

	for _, line := range split {
		var unitId *string
		var keyName *string
		var value *string
		cleanLine := strings.Replace(strings.Replace(line, "\r", "", -1), "\n", "", -1)

		if keyNameReg.MatchString(cleanLine) {
			lineSubmatch := keyNameReg.FindStringSubmatch(cleanLine)
			keyName = &lineSubmatch[1]

			newValue := cleanLine[len(*keyName) + 1:]
			value = &newValue
		} else if unitIdReg.MatchString(cleanLine) {
			lineSubmatch := unitIdReg.FindStringSubmatch(cleanLine)
			unitId = &lineSubmatch[1]
		}

		if unitId != nil {
			currentUnitId = unitId

			// Check if unitMap does not already have this key
			if _, ok := unitMap[*unitId]; ok {
				unitMap[*unitId].UnitFuncId.SetValid(*unitId)
				unitMap[*unitId].UnitStringId.SetValid(*unitId)
			} else {
				newUnit := new(models.SLKUnit)
				newUnit.UnitUI = new(models.UnitUI)
				newUnit.UnitData = new(models.UnitData)
				newUnit.UnitBalance = new(models.UnitBalance)
				newUnit.UnitWeapons = new(models.UnitWeapons)
				newUnit.UnitAbilities = new(models.UnitAbilities)
				newUnit.UnitFunc = new(models.UnitFunc)
				newUnit.UnitString = new(models.UnitString)
				newUnit.UnitFuncId.SetValid(*unitId)
				newUnit.UnitStringId.SetValid(*unitId)

				unitMap[*unitId] = newUnit
			}
		}

		if currentUnitId != nil && keyName != nil && value != nil {
			nullString := new(null.String)
			nullString.SetValid(*value)

			if _, ok := unitMap[*currentUnitId]; ok {
				err := reflectUpdateValueOnFieldNullStruct(unitMap[*currentUnitId], *nullString, strings.Title(strings.ToLower(*keyName)))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

/**
 * A function that populates a map of SLKItem structures with data from the ItemData.slk file.
 * If you want to add file data from a TXT file like ItemFunc.txt you need to run PopulateItemMapWithTxtFileData.
 * You can take a look at the example below on how to use the function
 *
 * func populateMapWithDataFromMultipleFiles() error {
 *   itemDataBytes, err := ioutil.ReadFile("./ItemData.slk")
 *   if err != nil {
 *     return err
 *   }
 *
 *   var itemMap = make(map[string]*models.SLKItem)
 *   parser.PopulateItemMapWithSlkFileData(itemDataBytes, itemMap)
 *
 *   // You can also add data from txt files like seen below
 *   itemFuncBytes, err := ioutil.ReadFile("./ItemFunc.txt")
 *   if err != nil {
 *     return err
 *   }
 *
 *   // Note that this is not the same function as used above!
 *   parser.PopulateItemMapWithTxtFileData(itemFuncBytes, itemMap)
 *
 *   return nil
 * }
 *
 */
func PopulateItemMapWithSlkFileData(inputFileData []byte, itemMap map[string]*models.SLKItem) {
	var currentItemId *string
	slkInformation, err := GenericSlkReader(inputFileData)
	if err != nil {
		log.Println(err)
		return
	}

	bodyLines := slkInformation.split[slkInformation.headerEndIndex:]
	for _, bodyLine := range bodyLines {
		var x *string
		var y *string
		var k *string

		cleanBodyLine := strings.Replace(strings.Replace(bodyLine, "\r", "", -1), "\n", "", -1)
		bodySplit := strings.Split(cleanBodyLine, ";")
		for _, s := range bodySplit {
			if xReg.MatchString(s) {
				newX := s[1:]
				x = &newX
			} else if kReg.MatchString(s) {
				newK := s[1:]
				k = &newK
			} else if yReg.MatchString(s) {
				newY := s[1:]
				y = &newY
			}
		}

		if x != nil && k != nil {
			if y != nil { // y != nil means we've reached the beginning of a new unit
				trimmedId := strings.Replace(*k, "\"", "", -1)
				currentItemId = &trimmedId

				// Check if itemMap does not already have this key
				if _, ok := itemMap[trimmedId]; !ok {
					newItem := new(models.SLKItem)
					newItem.ItemData = new(models.ItemData)
					newItem.ItemFunc = new(models.ItemFunc)
					newItem.ItemString = new(models.ItemString)

					itemMap[trimmedId] = newItem
				}
			}

			if currentItemId != nil {
				nullString := new(null.String)
				nullString.SetValid(*k)

				if unit, ok := itemMap[*currentItemId]; ok {
					if header, headerMapOk := slkInformation.headerMap[*x]; headerMapOk {
						err = reflectUpdateValueOnFieldNullStruct(unit, *nullString, strings.Title(header))
						if err != nil {
							log.Println(err)
						}
					}
				}
			}
		}
	}
}

/**
 * A function that populates a map of SLKItem structures with data from any TXT file like the ItemFunc.txt and ItemStrings.txt files.
 * If you want to add file data from an SLK file like ItemData.slk you need to run PopulateItemMapWithSlkFileData.
 * To populate a map with values from multiple files you need to call this function several times like in the code example below
 *
 * func populateMapWithDataFromMultipleFiles() error {
 *   itemFuncBytes, err := ioutil.ReadFile("./ItemFunc.txt")
 *   if err != nil {
 *     return err
 *   }
 *
 *   itemStringBytes, err := ioutil.ReadFile("./ItemStrings.txt")
 *   if err != nil {
 *     return err
 *   }
 *
 *   var itemMap = make(map[string]*models.SLKItem)
 *   parser.PopulateItemMapWithTxtFileData(itemFuncBytes, itemMap)
 *   parser.PopulateItemMapWithTxtFileData(itemStringBytes, itemMap)
 *
 *   // You can also add data from slk files like seen below
 *   itemDataBytes, err := ioutil.ReadFile("./ItemData.slk")
 *   if err != nil {
 *     return err
 *   }
 *
 *   // Note that this is not the same function as used above!
 *   parser.PopulateItemMapWithSlkFileData(itemDataBytes, itemMap)
 *
 *   return nil
 * }
 *
 */
func PopulateItemMapWithTxtFileData(inputFileData []byte, itemMap map[string]*models.SLKItem) {
	var currentItemId *string

	str := string(inputFileData)
	split := strings.Split(str, "\n")

	for _, line := range split {
		var itemId *string
		var keyName *string
		var value *string
		cleanLine := strings.Replace(strings.Replace(line, "\r", "", -1), "\n", "", -1)

		if keyNameReg.MatchString(cleanLine) {
			lineSubmatch := keyNameReg.FindStringSubmatch(cleanLine)
			keyName = &lineSubmatch[1]

			newValue := cleanLine[len(*keyName) + 1:]
			value = &newValue
		} else if unitIdReg.MatchString(cleanLine) {
			lineSubmatch := unitIdReg.FindStringSubmatch(cleanLine)
			itemId = &lineSubmatch[1]
		}

		if itemId != nil {
			currentItemId = itemId

			// Check if itemMap does not already have this key
			if _, ok := itemMap[*itemId]; ok {
				itemMap[*itemId].ItemFuncId.SetValid(*itemId)
				itemMap[*itemId].ItemStringId.SetValid(*itemId)
			} else {
				newItem := new(models.SLKItem)
				newItem.ItemData = new(models.ItemData)
				newItem.ItemFunc = new(models.ItemFunc)
				newItem.ItemString = new(models.ItemString)
				newItem.ItemID.SetValid(*itemId)
				newItem.ItemFuncId.SetValid(*itemId)
				newItem.ItemStringId.SetValid(*itemId)

				itemMap[*itemId] = newItem
			}
		}

		if currentItemId != nil && keyName != nil && value != nil {
			nullString := new(null.String)
			nullString.SetValid(*value)

			if _, ok := itemMap[*currentItemId]; ok {
				err := reflectUpdateValueOnFieldNullStruct(itemMap[*currentItemId], *nullString, strings.Title(strings.ToLower(*keyName)))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}