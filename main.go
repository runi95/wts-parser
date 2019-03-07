package main

import (
	"fmt"
	"github.com/runi95/wts-parser/models"
	"github.com/runi95/wts-parser/parser"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"text/template"
)

var fileNameRegex = regexp.MustCompile("[^.]*")
var fileExtensionRegex = regexp.MustCompile("\\.[^.]+$") // Find file extension

type UnitBalanceTemplate struct {
	RowCount    int
	UnitBalance []*models.UnitBalance
}

type UnitUITemplate struct {
	RowCount int
	UnitUI   []*models.UnitUI
}

type UnitAbilitiesTemplate struct {
	RowCount int
	UnitAbilities []*models.UnitAbilities
}

type UnitDataTemplate struct {
	RowCount int
	UnitData []*models.UnitData
}

type UnitWeaponsTemplate struct {
	RowCount    int
	UnitWeapons []*models.UnitWeapons
}

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
		log.Println("Reading UnitAbilities.slk...")

		unitAbilitiesBytes, err := ioutil.ReadFile("resources/UnitAbilities.slk")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		unitAbilitiesMap := parser.SlkToUnitAbilities(unitAbilitiesBytes)

		log.Println("Reading UnitData.slk...")

		unitDataBytes, err := ioutil.ReadFile("resources/UnitData.slk")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		unitDataMap := parser.SlkToUnitData(unitDataBytes)

		log.Println("Reading UnitUI.slk...")

		unitUIBytes, err := ioutil.ReadFile("resources/UnitUI.slk")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		unitUIMap := parser.SLKToUnitUI(unitUIBytes)

		log.Println("Reading UnitWeapons.slk...")

		unitWeaponsBytes, err := ioutil.ReadFile("resources/UnitWeapons.slk")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		unitWeaponsMap := parser.SLKToUnitWeapons(unitWeaponsBytes)

		log.Println("Reading UnitBalance.slk...")

		unitBalanceBytes, err := ioutil.ReadFile("resources/UnitBalance.slk")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		unitBalanceMap := parser.SLKToUnitBalance(unitBalanceBytes)

		baseUnitMap := make(map[string]*models.SLKUnit)
		for k := range unitDataMap {
			slkUnit := new(models.SLKUnit)
			slkUnit.UnitAbilities = unitAbilitiesMap[k]
			slkUnit.UnitData = unitDataMap[k]
			slkUnit.UnitUI = unitUIMap[k]
			slkUnit.UnitWeapons = unitWeaponsMap[k]
			slkUnit.UnitBalance = unitBalanceMap[k]

			baseUnitMap[k] = slkUnit
		}

		log.Println("Reading .w3u file...")

		unitMap := parser.ReadW3uFile(b)

		parsedSLKUnits := parser.W3uToSlkUnitsWithBaseSlk(baseUnitMap, unitMap)

		parsedSLKUnitsAbilities := make([]*models.UnitAbilities, len(parsedSLKUnits))
		parsedSLKUnitsData := make([]*models.UnitData, len(parsedSLKUnits))
		parsedSLKUnitsUI := make([]*models.UnitUI, len(parsedSLKUnits))
		parsedSLKUnitsWeapons := make([]*models.UnitWeapons, len(parsedSLKUnits))
		parsedSLKUnitsBalance := make([]*models.UnitBalance, len(parsedSLKUnits))

		for i, parsedSLKUnit := range parsedSLKUnits {
			parsedSLKUnitsAbilities[i] = parsedSLKUnit.UnitAbilities
			parsedSLKUnitsData[i] = parsedSLKUnit.UnitData
			parsedSLKUnitsUI[i] = parsedSLKUnit.UnitUI
			parsedSLKUnitsWeapons[i] = parsedSLKUnit.UnitWeapons
			parsedSLKUnitsBalance[i] = parsedSLKUnit.UnitBalance
		}

		funcMap := template.FuncMap{
			"inc": func(i int) int {
				return i + 2
			},
		}

		log.Println("Reading .wts file...")

		wtsInputFilePath := "/home/runi95/Downloads/war3map.wts"

		wtsBytes, err := ioutil.ReadFile(wtsInputFilePath)
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		wtsMap := parser.ReadWtsFile(wtsBytes)

		log.Println("Reading CampaignUnitFunc.txt...")

		campaignUnitFuncBytes, err := ioutil.ReadFile("resources/CampaignUnitFunc.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedCampaignUnitFunc := parser.TxtToUnitFunc(campaignUnitFuncBytes)

		log.Println("Reading HumanUnitFunc.txt...")

		humanUnitFuncBytes, err := ioutil.ReadFile("resources/HumanUnitFunc.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedHumanUnitFunc := parser.TxtToUnitFunc(humanUnitFuncBytes)

		log.Println("Reading NeutralUnitFunc.txt...")

		neutralUnitFuncBytes, err := ioutil.ReadFile("resources/NeutralUnitFunc.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedNeutralUnitFunc := parser.TxtToUnitFunc(neutralUnitFuncBytes)

		log.Println("Reading NightElfUnitFunc.txt...")

		nightElfUnitFuncBytes, err := ioutil.ReadFile("resources/NightElfUnitFunc.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedNightElfUnitFunc := parser.TxtToUnitFunc(nightElfUnitFuncBytes)

		log.Println("Reading OrcUnitFunc.txt...")

		orcUnitFuncBytes, err := ioutil.ReadFile("resources/OrcUnitFunc.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedOrcUnitFunc := parser.TxtToUnitFunc(orcUnitFuncBytes)

		log.Println("Reading UndeadUnitFunc.txt...")

		undeadUnitFuncBytes, err := ioutil.ReadFile("resources/UndeadUnitFunc.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedUndeadUnitFunc := parser.TxtToUnitFunc(undeadUnitFuncBytes)

		log.Println("Building baseUnitFunc map...")

		baseUnitFuncMap := make(map[string]*models.UnitFunc)
		for k, v := range parsedCampaignUnitFunc {
			baseUnitFuncMap[k] = v
		}

		for k, v := range parsedHumanUnitFunc {
			baseUnitFuncMap[k] = v
		}

		for k, v := range parsedNeutralUnitFunc {
			baseUnitFuncMap[k] = v
		}

		for k, v := range parsedNightElfUnitFunc {
			baseUnitFuncMap[k] = v
		}

		for k, v := range parsedOrcUnitFunc {
			baseUnitFuncMap[k] = v
		}

		for k, v := range parsedUndeadUnitFunc {
			baseUnitFuncMap[k] = v
		}

		customUnitFuncs := parser.W3uToTxtUnitFuncsWithBaseTxtAndBaseWts(baseUnitFuncMap, baseUnitMap, unitMap, wtsMap)

		/*
		log.Println("Reading CampaignUnitStrings.txt...")

		campaignUnitStringsBytes, err := ioutil.ReadFile("resources/CampaignUnitStrings.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedCampaignUnitStrings := parser.TxtToUnitStrings(campaignUnitStringsBytes)

		log.Println("Reading HumanUnitStrings.txt...")

		humanUnitStringsBytes, err := ioutil.ReadFile("resources/HumanUnitStrings.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedHumanUnitStrings := parser.TxtToUnitStrings(humanUnitStringsBytes)

		log.Println("Reading NeutralUnitStrings.txt...")

		neutralUnitStringsBytes, err := ioutil.ReadFile("resources/NeutralUnitStrings.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedNeutralUnitStrings := parser.TxtToUnitStrings(neutralUnitStringsBytes)

		log.Println("Reading NightElfUnitStrings.txt...")

		nightElfUnitStringsBytes, err := ioutil.ReadFile("resources/NightElfUnitStrings.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedNightElfUnitStrings := parser.TxtToUnitStrings(nightElfUnitStringsBytes)

		log.Println("Reading OrcUnitStrings.txt...")

		orcUnitStringsBytes, err := ioutil.ReadFile("resources/OrcUnitStrings.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedOrcUnitStrings := parser.TxtToUnitStrings(orcUnitStringsBytes)

		log.Println("Reading UndeadUnitStrings.txt...")

		undeadUnitStringsBytes, err := ioutil.ReadFile("resources/UndeadUnitStrings.txt")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		parsedUndeadUnitStrings := parser.TxtToUnitStrings(undeadUnitStringsBytes)

		log.Println("Building baseUnitString map...")

		baseUnitStringMap := make(map[string]*models.UnitString)
		for k, v := range parsedCampaignUnitStrings {
			baseUnitStringMap[k] = v
		}

		for k, v := range parsedHumanUnitStrings {
			baseUnitStringMap[k] = v
		}

		for k, v := range parsedNeutralUnitStrings {
			baseUnitStringMap[k] = v
		}

		for k, v := range parsedNightElfUnitStrings {
			baseUnitStringMap[k] = v
		}

		for k, v := range parsedOrcUnitStrings {
			baseUnitStringMap[k] = v
		}

		for k, v := range parsedUndeadUnitStrings {
			baseUnitStringMap[k] = v
		}

		customUnitStrings := parser.WtsToTxtUnitStringsWithBaseTxt(baseUnitStringMap, baseUnitMap, unitMap)

		log.Println("Writing to CampaignUnitString...")

		campaignUnitStringFile, err := os.Create("out/CampaignUnitString.txt")
		if err != nil {
			log.Println(err)
		}

		campaignUnitStringTemplate := template.New("UnitStrings")
		campaignUnitStringTemplate, err = campaignUnitStringTemplate.ParseFiles("templates/UnitStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = campaignUnitStringTemplate.ExecuteTemplate(campaignUnitStringFile, "UnitStrings", customUnitStrings.CampaignUnitFuncs)
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to HumanUnitString...")

		humanUnitStringFile, err := os.Create("out/HumanUnitString.txt")
		if err != nil {
			log.Println(err)
		}

		humanUnitStringTemplate := template.New("UnitStrings")
		humanUnitStringTemplate, err = humanUnitStringTemplate.ParseFiles("templates/UnitStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = humanUnitStringTemplate.ExecuteTemplate(humanUnitStringFile, "UnitStrings", customUnitStrings.HumanUnitStrings)
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to NeutralUnitString...")

		neutralUnitStringFile, err := os.Create("out/NeutralUnitString.txt")
		if err != nil {
			log.Println(err)
		}

		neutralUnitStringTemplate := template.New("UnitStrings")
		neutralUnitStringTemplate, err = neutralUnitStringTemplate.ParseFiles("templates/UnitStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = neutralUnitStringTemplate.ExecuteTemplate(neutralUnitStringFile, "UnitStrings", customUnitStrings.NeutralUnitStrings)
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to NightElfUnitString...")

		nightElfUnitStringFile, err := os.Create("out/NightElfUnitString.txt")
		if err != nil {
			log.Println(err)
		}

		nightElfUnitStringTemplate := template.New("UnitStrings")
		nightElfUnitStringTemplate, err = nightElfUnitStringTemplate.ParseFiles("templates/UnitStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = nightElfUnitStringTemplate.ExecuteTemplate(nightElfUnitStringFile, "UnitStrings", customUnitStrings.NightElfUnitStrings)
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to OrcUnitString...")

		orcUnitStringFile, err := os.Create("out/OrcUnitString.txt")
		if err != nil {
			log.Println(err)
		}

		orcUnitStringTemplate := template.New("UnitStrings")
		orcUnitStringTemplate, err = orcUnitStringTemplate.ParseFiles("templates/UnitStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = orcUnitStringTemplate.ExecuteTemplate(orcUnitStringFile, "UnitStrings", customUnitStrings.OrcUnitStrings)
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to UndeadUnitString...")

		undeadUnitStringFile, err := os.Create("out/UndeadUnitString.txt")
		if err != nil {
			log.Println(err)
		}

		undeadUnitStringTemplate := template.New("UnitStrings")
		undeadUnitStringTemplate, err = undeadUnitStringTemplate.ParseFiles("templates/UnitStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = undeadUnitStringTemplate.ExecuteTemplate(undeadUnitStringFile, "UnitStrings", customUnitStrings.UndeadUnitStrings)
		if err != nil {
			log.Println(err)
		}
		*/

		log.Println("Writing to CampaignUnitFunc...")

		campaignUnitFuncFile, err := os.Create("out/CampaignUnitFunc.txt")
		if err != nil {
			log.Println(err)
		}

		campaignUnitFuncTemplate := template.New("UnitFunc")
		campaignUnitFuncTemplate, err = campaignUnitFuncTemplate.ParseFiles("templates/UnitFuncAndStrings.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = campaignUnitFuncTemplate.ExecuteTemplate(campaignUnitFuncFile, "UnitFunc", customUnitFuncs.CampaignUnitFuncs)
		if err != nil {
			log.Println(err)
		}

		if len(customUnitFuncs.HumanUnitFuncs) > 0 {
			log.Println("Writing to HumanUnitFunc...")

			humanUnitFuncFile, err := os.Create("out/HumanUnitFunc.txt")
			if err != nil {
				log.Println(err)
			}

			humanUnitFuncTemplate := template.New("UnitFunc")
			humanUnitFuncTemplate, err = humanUnitFuncTemplate.ParseFiles("templates/UnitFunc.tmpl")
			if err != nil {
				log.Println(err)
			}

			err = humanUnitFuncTemplate.ExecuteTemplate(humanUnitFuncFile, "UnitFunc", customUnitFuncs.HumanUnitFuncs)
			if err != nil {
				log.Println(err)
			}
		}

		if len(customUnitFuncs.NightElfUnitFuncs) > 0 {
			log.Println("Writing to NightElfUnitFunc...")

			nightElfUnitFuncFile, err := os.Create("out/NightElfUnitFunc.txt")
			if err != nil {
				log.Println(err)
			}

			nightElfUnitFuncTemplate := template.New("UnitFunc")
			nightElfUnitFuncTemplate, err = nightElfUnitFuncTemplate.ParseFiles("templates/UnitFunc.tmpl")
			if err != nil {
				log.Println(err)
			}

			err = nightElfUnitFuncTemplate.ExecuteTemplate(nightElfUnitFuncFile, "UnitFunc", customUnitFuncs.NightElfUnitFuncs)
			if err != nil {
				log.Println(err)
			}
		}

		if len(customUnitFuncs.OrcUnitFuncs) > 0 {
			log.Println("Writing to OrcUnitFunc...")

			orcUnitFuncFile, err := os.Create("out/OrcUnitFunc.txt")
			if err != nil {
				log.Println(err)
			}

			orcUnitFuncTemplate := template.New("UnitFunc")
			orcUnitFuncTemplate, err = orcUnitFuncTemplate.ParseFiles("templates/UnitFunc.tmpl")
			if err != nil {
				log.Println(err)
			}

			err = orcUnitFuncTemplate.ExecuteTemplate(orcUnitFuncFile, "UnitFunc", customUnitFuncs.OrcUnitFuncs)
			if err != nil {
				log.Println(err)
			}
		}

		if len(customUnitFuncs.UndeadUnitFuncs) > 0 {
			log.Println("Writing to UndeadUnitFunc...")

			undeadUnitFuncFile, err := os.Create("out/UndeadUnitFunc.txt")
			if err != nil {
				log.Println(err)
			}

			undeadUnitFuncTemplate := template.New("UnitFunc")
			undeadUnitFuncTemplate, err = undeadUnitFuncTemplate.ParseFiles("templates/UnitFunc.tmpl")
			if err != nil {
				log.Println(err)
			}

			err = undeadUnitFuncTemplate.ExecuteTemplate(undeadUnitFuncFile, "UnitFunc", customUnitFuncs.UndeadUnitFuncs)
			if err != nil {
				log.Println(err)
			}
		}

		if len(customUnitFuncs.NeutralUnitFuncs) > 0 {
			log.Println("Writing to NeutralUnitFunc...")

			neutralUnitFuncFile, err := os.Create("out/NeutralUnitFunc.txt")
			if err != nil {
				log.Println(err)
			}

			neutralUnitFuncTemplate := template.New("UnitFunc")
			neutralUnitFuncTemplate, err = neutralUnitFuncTemplate.ParseFiles("templates/UnitFunc.tmpl")
			if err != nil {
				log.Println(err)
			}

			err = neutralUnitFuncTemplate.ExecuteTemplate(neutralUnitFuncFile, "UnitFunc", customUnitFuncs.NeutralUnitFuncs)
			if err != nil {
				log.Println(err)
			}
		}

		log.Println("Writing to UnitAbilities.slk...")

		unitAbilitiesFile, err := os.Create("out/UnitAbilities.slk")
		if err != nil {
			log.Println(err)
		}

		unitAbilitiesTemplate := template.New("UnitAbilities").Funcs(funcMap)
		unitAbilitiesTemplate, err = unitAbilitiesTemplate.ParseFiles("templates/UnitAbilities.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = unitAbilitiesTemplate.ExecuteTemplate(unitAbilitiesFile, "UnitAbilities", UnitAbilitiesTemplate{RowCount: len(parsedSLKUnitsAbilities) + 1, UnitAbilities: parsedSLKUnitsAbilities})
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to UnitData.slk...")

		unitDataFile, err := os.Create("out/UnitData.slk")
		if err != nil {
			log.Println(err)
		}

		unitDataTemplate := template.New("UnitData").Funcs(funcMap)
		unitDataTemplate, err = unitDataTemplate.ParseFiles("templates/UnitData.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = unitDataTemplate.ExecuteTemplate(unitDataFile, "UnitData", UnitDataTemplate{RowCount: len(parsedSLKUnitsData) + 1, UnitData: parsedSLKUnitsData})
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to UnitBalance.slk...")

		unitBalanceFile, err := os.Create("out/UnitBalance.slk")
		if err != nil {
			log.Println(err)
		}

		unitBalanceTemplate := template.New("UnitBalance").Funcs(funcMap)
		unitBalanceTemplate, err = unitBalanceTemplate.ParseFiles("templates/UnitBalance.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = unitBalanceTemplate.ExecuteTemplate(unitBalanceFile, "UnitBalance", UnitBalanceTemplate{RowCount: len(parsedSLKUnitsBalance) + 1, UnitBalance: parsedSLKUnitsBalance})
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to UnitUI.slk...")

		unitUIFile, err := os.Create("out/UnitUI.slk")
		if err != nil {
			log.Println(err)
		}

		unitUITemplate := template.New("UnitUI").Funcs(funcMap)
		unitUITemplate, err = unitUITemplate.ParseFiles("templates/UnitUI.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = unitUITemplate.ExecuteTemplate(unitUIFile, "UnitUI", UnitUITemplate{RowCount: len(parsedSLKUnitsUI) + 1, UnitUI: parsedSLKUnitsUI})
		if err != nil {
			log.Println(err)
		}

		log.Println("Writing to UnitWeapons.slk...")

		unitWeaponsFile, err := os.Create("out/UnitWeapons.slk")
		if err != nil {
			log.Println(err)
		}

		unitWeaponsTemplate := template.New("UnitWeapons").Funcs(funcMap)
		unitWeaponsTemplate, err = unitWeaponsTemplate.ParseFiles("templates/UnitWeapons.tmpl")
		if err != nil {
			log.Println(err)
		}

		err = unitWeaponsTemplate.ExecuteTemplate(unitWeaponsFile, "UnitWeapons", UnitWeaponsTemplate{RowCount: len(parsedSLKUnitsWeapons) + 1, UnitWeapons: parsedSLKUnitsWeapons})
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Fatal(fmt.Errorf("expected .wts or .json, but got " + fileExtension))
	}
}
