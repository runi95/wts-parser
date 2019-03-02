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
		log.Println("Reading UnitData.slk...")

		unitDataBytes, err := ioutil.ReadFile("resources/UnitData.slk")
		if err != nil {
			log.Println(err)
			os.Exit(10)
		}

		unitDataMap := parser.SLKToUnitData(unitDataBytes)

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
			slkUnit.UnitData = unitDataMap[k]
			slkUnit.UnitUI = unitUIMap[k]
			slkUnit.UnitWeapons = unitWeaponsMap[k]
			slkUnit.UnitBalance = unitBalanceMap[k]
			slkUnit.UnitAbilities = new(models.UnitAbilities)

			baseUnitMap[k] = slkUnit
		}

		log.Println("Reading .w3u file...")

		parsedSLKUnits := parser.W3uToSLKUnitsWithBaseSLK(baseUnitMap, b)

		parsedSLKUnitsData := make([]*models.UnitData, len(parsedSLKUnits))
		parsedSLKUnitsUI := make([]*models.UnitUI, len(parsedSLKUnits))
		parsedSLKUnitsWeapons := make([]*models.UnitWeapons, len(parsedSLKUnits))
		parsedSLKUnitsBalance := make([]*models.UnitBalance, len(parsedSLKUnits))

		for i, parsedSLKUnit := range parsedSLKUnits {
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
