package parser

import (
	"fmt"
	"github.com/runi95/wts-parser/models"
	"github.com/runi95/wts-parser/templates"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

/*************************

	   STRUCTURES

*************************/

type SlkInformation struct {
	split          []string
	headerMap      map[string]string
	columnSize     int
	headerEndIndex int
}

type SlkSaveInformation struct {
	CampaignUnitFuncs []*models.UnitFunc
	NeutralUnitFuncs  []*models.UnitFunc
	HumanUnitFuncs    []*models.UnitFunc
	NightElfUnitFuncs []*models.UnitFunc
	OrcUnitFuncs      []*models.UnitFunc
	UndeadUnitFuncs   []*models.UnitFunc

	CampaignUnitStrings []*models.UnitString
	NeutralUnitStrings  []*models.UnitString
	HumanUnitStrings    []*models.UnitString
	NightElfUnitStrings []*models.UnitString
	OrcUnitStrings      []*models.UnitString
	UndeadUnitStrings   []*models.UnitString

	UnitAbilities []*models.UnitAbilities
	UnitBalance   []*models.UnitBalance
	UnitData      []*models.UnitData
	UnitUI        []*models.UnitUI
	UnitWeapons   []*models.UnitWeapons
}

/*************************

	   SLK PARSERS

*************************/

func WriteToFiles(slkSaveInfo *SlkSaveInformation) {
	WriteToFilesAndSaveToFolder(slkSaveInfo, "out", true)
}

func WriteToFilesAndSaveToFolder(slkSaveInfo *SlkSaveInformation, outputFolder string, sortBeforeSave bool) {
	if slkSaveInfo == nil {
		return
	}

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 2
		},
	}

	if sortBeforeSave {
		log.Println("Sorting units according to unit ID...")

		var sortWaitGroup sync.WaitGroup

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.CampaignUnitFuncs, func(i, j int) bool {
				return slkSaveInfo.CampaignUnitFuncs[i].UnitId < slkSaveInfo.CampaignUnitFuncs[j].UnitId
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.NightElfUnitFuncs, func(i, j int) bool {
				return slkSaveInfo.NightElfUnitFuncs[i].UnitId < slkSaveInfo.NightElfUnitFuncs[j].UnitId
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.OrcUnitFuncs, func(i, j int) bool {
				return slkSaveInfo.OrcUnitFuncs[i].UnitId < slkSaveInfo.OrcUnitFuncs[j].UnitId
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.UndeadUnitFuncs, func(i, j int) bool {
				return slkSaveInfo.UndeadUnitFuncs[i].UnitId < slkSaveInfo.UndeadUnitFuncs[j].UnitId
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.NeutralUnitFuncs, func(i, j int) bool {
				return slkSaveInfo.NeutralUnitFuncs[i].UnitId < slkSaveInfo.NeutralUnitFuncs[j].UnitId
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.HumanUnitFuncs, func(i, j int) bool {
				return slkSaveInfo.HumanUnitFuncs[i].UnitId < slkSaveInfo.HumanUnitFuncs[j].UnitId
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.UnitAbilities, func(i, j int) bool {
				return slkSaveInfo.UnitAbilities[i].UnitAbilID.String < slkSaveInfo.UnitAbilities[j].UnitAbilID.String
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.UnitData, func(i, j int) bool {
				return slkSaveInfo.UnitData[i].UnitID.String < slkSaveInfo.UnitData[j].UnitID.String
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.UnitUI, func(i, j int) bool {
				return slkSaveInfo.UnitUI[i].UnitUIID.String < slkSaveInfo.UnitUI[j].UnitUIID.String
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.UnitWeapons, func(i, j int) bool {
				return slkSaveInfo.UnitWeapons[i].UnitWeapID.String < slkSaveInfo.UnitWeapons[j].UnitWeapID.String
			})
		}()

		sortWaitGroup.Add(1)
		go func() {
			defer sortWaitGroup.Done()

			sort.Slice(slkSaveInfo.UnitBalance, func(i, j int) bool {
				return slkSaveInfo.UnitBalance[i].UnitBalanceID.String < slkSaveInfo.UnitBalance[j].UnitBalanceID.String
			})
		}()

		sortWaitGroup.Wait()
	}

	for _, campaignUnitFunc := range slkSaveInfo.CampaignUnitFuncs {
		if campaignUnitFunc.Missileart1.Valid || campaignUnitFunc.Missileart2.Valid {
			missileArt1 := "_"
			if campaignUnitFunc.Missileart1.Valid {
				missileArt1 = campaignUnitFunc.Missileart1.String
			}

			missileArt := missileArt1
			if campaignUnitFunc.Missileart2.Valid {
				missileArt += "," + campaignUnitFunc.Missileart2.String
			}

			campaignUnitFunc.Missileart.SetValid(missileArt)
		}

		if campaignUnitFunc.Missilearc1.Valid || campaignUnitFunc.Missilearc2.Valid {
			missilearc1 := "_"
			if campaignUnitFunc.Missilearc1.Valid {
				missilearc1 = campaignUnitFunc.Missilearc1.String
			}

			missilearc := missilearc1
			if campaignUnitFunc.Missilearc2.Valid {
				missilearc += "," + campaignUnitFunc.Missilearc2.String
			}

			campaignUnitFunc.Missilearc.SetValid(missilearc)
		}

		if campaignUnitFunc.Missilespeed1.Valid || campaignUnitFunc.Missilespeed2.Valid {
			missilespeed1 := "_"
			if campaignUnitFunc.Missilespeed1.Valid {
				missilespeed1 = campaignUnitFunc.Missilespeed1.String
			}

			missilespeed := missilespeed1
			if campaignUnitFunc.Missilespeed2.Valid {
				missilespeed += "," + campaignUnitFunc.Missilespeed2.String
			}

			campaignUnitFunc.Missilespeed.SetValid(missilespeed)
		}

		if campaignUnitFunc.Missilehoming1.Valid || campaignUnitFunc.Missilehoming2.Valid {
			missilehoming1 := "_"
			if campaignUnitFunc.Missilehoming1.Valid {
				missilehoming1 = campaignUnitFunc.Missilehoming1.String
			}

			missilehoming := missilehoming1
			if campaignUnitFunc.Missilehoming2.Valid {
				missilehoming += "," + campaignUnitFunc.Missilehoming2.String
			}

			campaignUnitFunc.Missilehoming.SetValid(missilehoming)
		}
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

		campaignUnitFuncTemplate := template.New("UnitFunc")
		campaignUnitFuncTemplate, err = campaignUnitFuncTemplate.Parse(templates.GetUnitFuncTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = campaignUnitFuncTemplate.ExecuteTemplate(campaignUnitFuncFile, "UnitFunc", slkSaveInfo.CampaignUnitFuncs)
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

		campaignUnitStringsTemplate := template.New("UnitStrings")
		campaignUnitStringsTemplate, err = campaignUnitStringsTemplate.Parse(templates.GetUnitStringsTemplate())
		if err != nil {
			log.Println(err)
			return
		}

		err = campaignUnitStringsTemplate.ExecuteTemplate(campaignUnitStringsFile, "UnitStrings", slkSaveInfo.CampaignUnitStrings)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	if len(slkSaveInfo.HumanUnitFuncs) > 0 {
		log.Println("Writing to HumanUnitFunc...")
		writeToFileWaitGroup.Add(1)

		go func() {
			defer writeToFileWaitGroup.Done()
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

			err = humanUnitFuncTemplate.ExecuteTemplate(humanUnitFuncFile, "UnitFunc", slkSaveInfo.HumanUnitFuncs)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

	if len(slkSaveInfo.NightElfUnitFuncs) > 0 {
		log.Println("Writing to NightElfUnitFunc...")
		writeToFileWaitGroup.Add(1)

		go func() {
			defer writeToFileWaitGroup.Done()
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

			err = nightElfUnitFuncTemplate.ExecuteTemplate(nightElfUnitFuncFile, "UnitFunc", slkSaveInfo.NightElfUnitFuncs)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

	if len(slkSaveInfo.OrcUnitFuncs) > 0 {
		log.Println("Writing to OrcUnitFunc...")
		writeToFileWaitGroup.Add(1)

		go func() {
			defer writeToFileWaitGroup.Done()
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

			err = orcUnitFuncTemplate.ExecuteTemplate(orcUnitFuncFile, "UnitFunc", slkSaveInfo.OrcUnitFuncs)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

	if len(slkSaveInfo.UndeadUnitFuncs) > 0 {
		log.Println("Writing to UndeadUnitFunc...")
		writeToFileWaitGroup.Add(1)

		go func() {
			defer writeToFileWaitGroup.Done()
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

			err = undeadUnitFuncTemplate.ExecuteTemplate(undeadUnitFuncFile, "UnitFunc", slkSaveInfo.UndeadUnitFuncs)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

	if len(slkSaveInfo.NeutralUnitFuncs) > 0 {
		log.Println("Writing to NeutralUnitFunc...")
		writeToFileWaitGroup.Add(1)

		go func() {
			defer writeToFileWaitGroup.Done()
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

			err = neutralUnitFuncTemplate.ExecuteTemplate(neutralUnitFuncFile, "UnitFunc", slkSaveInfo.NeutralUnitFuncs)
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

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

		err = unitAbilitiesTemplate.ExecuteTemplate(unitAbilitiesFile, "UnitAbilities", models.UnitAbilitiesTemplate{RowCount: len(slkSaveInfo.UnitAbilities) + 1, UnitAbilities: slkSaveInfo.UnitAbilities})
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

		err = unitDataTemplate.ExecuteTemplate(unitDataFile, "UnitData", models.UnitDataTemplate{RowCount: len(slkSaveInfo.UnitData) + 1, UnitData: slkSaveInfo.UnitData})
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

		err = unitBalanceTemplate.ExecuteTemplate(unitBalanceFile, "UnitBalance", models.UnitBalanceTemplate{RowCount: len(slkSaveInfo.UnitBalance) + 1, UnitBalance: slkSaveInfo.UnitBalance})
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

		err = unitUITemplate.ExecuteTemplate(unitUIFile, "UnitUI", models.UnitUITemplate{RowCount: len(slkSaveInfo.UnitUI) + 1, UnitUI: slkSaveInfo.UnitUI})
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

		err = unitWeaponsTemplate.ExecuteTemplate(unitWeaponsFile, "UnitWeapons", models.UnitWeaponsTemplate{RowCount: len(slkSaveInfo.UnitWeapons) + 1, UnitWeapons: slkSaveInfo.UnitWeapons})
		if err != nil {
			log.Println(err)
			return
		}
	}()

	writeToFileWaitGroup.Wait()
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

func SLKToUnitUI(input []byte) map[string]*models.UnitUI {
	unitUIMap := make(map[string]*models.UnitUI)
	var currentUnitUI *models.UnitUI
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

				currentUnitUI = new(models.UnitUI)
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

func SlkToUnitData(input []byte) map[string]*models.UnitData {
	unitDataMap := make(map[string]*models.UnitData)
	var currentUnitData *models.UnitData
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

				currentUnitData = new(models.UnitData)
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

func SlkToUnitAbilities(input []byte) map[string]*models.UnitAbilities {
	unitAbilitiesMap := make(map[string]*models.UnitAbilities)
	var currentUnitAbilities *models.UnitAbilities
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

				currentUnitAbilities = new(models.UnitAbilities)
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

func SLKToUnitWeapons(input []byte) map[string]*models.UnitWeapons {
	unitWeaponsMap := make(map[string]*models.UnitWeapons)
	var currentUnitWeapons *models.UnitWeapons
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

				currentUnitWeapons = new(models.UnitWeapons)
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

func SLKToUnitBalance(input []byte) map[string]*models.UnitBalance {
	unitBalanceMap := make(map[string]*models.UnitBalance)
	var currentUnitBalance *models.UnitBalance
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

				currentUnitBalance = new(models.UnitBalance)
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

func TxtToUnitFunc(input []byte) map[string]*models.UnitFunc {
	unitFuncMap := make(map[string]*models.UnitFunc)
	var currentUnitFunc *models.UnitFunc

	str := string(input)
	split := strings.Split(str, "\n")

	for _, line := range split {
		if TXTHeadRegex.MatchString(line) {
			if currentUnitFunc != nil {
				unitFuncMap[currentUnitFunc.UnitId] = currentUnitFunc
			}

			currentUnitFunc = new(models.UnitFunc)
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

func TxtToUnitStrings(input []byte) map[string]*models.UnitString {
	unitStringsMap := make(map[string]*models.UnitString)
	var currentUnitString *models.UnitString

	str := string(input)
	split := strings.Split(str, "\n")

	for _, line := range split {
		if TXTHeadRegex.MatchString(line) {
			if currentUnitString != nil {
				unitStringsMap[currentUnitString.UnitId] = currentUnitString
			}

			currentUnitString = new(models.UnitString)
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
