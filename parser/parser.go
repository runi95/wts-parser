package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/runi95/wts-parser/models"
	"regexp"
	"strconv"
)

var idRegex = regexp.MustCompile(`[0-9]+`)                  // Find string id
var stringRegex = regexp.MustCompile(`STRING [0-9]*[^}]*}`) // Regex to find each string object
var contentContainerRegex = regexp.MustCompile(`{[^}]*}$`)  // Regex to find the content container of each string object
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

		removeBrackets := findContentContainer[1 : len(findContentContainer)-1]
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

func W3uToJson(input []byte) []byte {
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
			switch str {
			case "unam":
				unitUnam := ""
				for input[i+1] > 31 {
					i++
					unitUnam += string(input[i])
				}
				currentUnit.Unam = unitUnam
				break
			case "umvr":
				unitUmvr := ""
				for input[i+1] > 31 {
					i++
					unitUmvr += string(input[i])
				}
				currentUnit.Umvr = unitUmvr
				break
			case "umas":
				unitUmas := ""
				for input[i+1] > 31 {
					i++
					unitUmas += string(input[i])
				}
				currentUnit.Umas = unitUmas
				break
			case "umis":
				unitUmis := ""
				for input[i+1] > 31 {
					i++
					unitUmis += string(input[i])
				}
				currentUnit.Umis = unitUmis
				break
			case "util":
				unitUtil := ""
				for input[i+1] > 31 {
					i++
					unitUtil += string(input[i])
				}
				currentUnit.Util = unitUtil
				break
			case "uine":
				unitUine := ""
				for input[i+1] > 31 {
					i++
					unitUine += string(input[i])
				}
				currentUnit.Uine = unitUine
				break
			case "udro":
				unitUdro := ""
				for input[i+1] > 31 {
					i++
					unitUdro += string(input[i])
				}
				currentUnit.Udro = unitUdro
				break
			case "uspa":
				unitUspa := ""
				for input[i+1] > 31 {
					i++
					unitUspa += string(input[i])
				}
				currentUnit.Uspa = unitUspa
				break
			case "ua1z":
				unitUa1z := ""
				for input[i+1] > 31 {
					i++
					unitUa1z += string(input[i])
				}
				currentUnit.Ua1z = unitUa1z
				break
			case "utc1":
				unitUtc1 := ""
				for input[i+1] > 31 {
					i++
					unitUtc1 += string(input[i])
				}
				currentUnit.Utc1 = unitUtc1
				break
			case "urb1":
				unitUrb1 := ""
				for input[i+1] > 31 {
					i++
					unitUrb1 += string(input[i])
				}
				currentUnit.Urb1 = unitUrb1
				break
			case "ucs1":
				unitUcs1 := ""
				for input[i+1] > 31 {
					i++
					unitUcs1 += string(input[i])
				}
				currentUnit.Ucs1 = unitUcs1
				break
			case "ua1w":
				unitUa1w := ""
				for input[i+1] > 31 {
					i++
					unitUa1w += string(input[i])
				}
				currentUnit.Ua1w = unitUa1w
				break
			case "umh1":
				unitUmh1 := ""
				for input[i+1] > 31 {
					i++
					unitUmh1 += string(input[i])
				}
				currentUnit.Umh1 = unitUmh1
				break
			case "udef":
				unitUdef := ""
				for input[i+1] > 31 {
					i++
					unitUdef += string(input[i])
				}
				currentUnit.Udef = unitUdef
				break
			case "uarm":
				unitUarm := ""
				for input[i+1] > 31 {
					i++
					unitUarm += string(input[i])
				}
				currentUnit.Uarm = unitUarm
				break
			case "udty":
				unitUdty := ""
				for input[i+1] > 31 {
					i++
					unitUdty += string(input[i])
				}
				currentUnit.Udty = unitUdty
				break
			case "ides":
				unitIdes := ""
				for input[i+1] > 31 {
					i++
					unitIdes += string(input[i])
				}
				currentUnit.Ides = unitIdes
				break
			case "unsf":
				unitUnsf := ""
				for input[i+1] > 31 {
					i++
					unitUnsf += string(input[i])
				}
				currentUnit.Unsf = unitUnsf
				break
			case "utyp":
				unitUtyp := ""
				for input[i+1] > 31 {
					i++
					unitUtyp += string(input[i])
				}
				currentUnit.Utyp = unitUtyp
				break
			case "ufor":
				unitUfor := ""
				for input[i+1] > 31 {
					i++
					unitUfor += string(input[i])
				}
				currentUnit.Ufor = unitUfor
				break
			case "ulbd":
				unitUlbd := ""
				for input[i+1] > 31 {
					i++
					unitUlbd += string(input[i])
				}
				currentUnit.Ulbd = unitUlbd
				break
			case "ulba":
				unitUlba := ""
				for input[i+1] > 31 {
					i++
					unitUlba += string(input[i])
				}
				currentUnit.Ulba = unitUlba
				break
			case "ulbs":
				unitUlbs := ""
				for input[i+1] > 31 {
					i++
					unitUlbs += string(input[i])
				}
				currentUnit.Ulbs = unitUlbs
				break
			case "ubba":
				unitUbba := ""
				for input[i+1] > 31 {
					i++
					unitUbba += string(input[i])
				}
				currentUnit.Ubba = unitUbba
				break
			case "ubsi":
				unitUbsi := ""
				for input[i+1] > 31 {
					i++
					unitUbsi += string(input[i])
				}
				currentUnit.Ubsi = unitUbsi
				break
			case "ubdi":
				unitUbdi := ""
				for input[i+1] > 31 {
					i++
					unitUbdi += string(input[i])
				}
				currentUnit.Ubdi = unitUbdi
				break
			case "ugor":
				unitUgor := ""
				for input[i+1] > 31 {
					i++
					unitUgor += string(input[i])
				}
				currentUnit.Ugor = unitUgor
				break
			case "ulur":
				unitUlur := ""
				for input[i+1] > 31 {
					i++
					unitUlur += string(input[i])
				}
				currentUnit.Ulur = unitUlur
				break
			case "umpm":
				unitUmpm := ""
				for input[i+1] > 31 {
					i++
					unitUmpm += string(input[i])
				}
				currentUnit.Umpm = unitUmpm
				break
			case "umpr":
				unitUmpr := ""
				for input[i+1] > 31 {
					i++
					unitUmpr += string(input[i])
				}
				currentUnit.Umpr = unitUmpr
				break
			case "upri":
				unitUpri := ""
				for input[i+1] > 31 {
					i++
					unitUpri += string(input[i])
				}
				currentUnit.Upri = unitUpri
				break
			case "urtm":
				unitUrtm := ""
				for input[i+1] > 31 {
					i++
					unitUrtm += string(input[i])
				}
				currentUnit.Urtm = unitUrtm
				break
			case "usle":
				unitUsle := ""
				for input[i+1] > 31 {
					i++
					unitUsle += string(input[i])
				}
				currentUnit.Usle = unitUsle
				break
			case "usin":
				unitUsin := ""
				for input[i+1] > 31 {
					i++
					unitUsin += string(input[i])
				}
				currentUnit.Usin = unitUsin
				break
			case "usid":
				unitUsid := ""
				for input[i+1] > 31 {
					i++
					unitUsid += string(input[i])
				}
				currentUnit.Usid = unitUsid
				break
			case "ulev":
				unitUlev := ""
				for input[i+1] > 31 {
					i++
					unitUlev += string(input[i])
				}
				currentUnit.Ulev = unitUlev
				break
			case "ucar":
				unitUcar := ""
				for input[i+1] > 31 {
					i++
					unitUcar += string(input[i])
				}
				currentUnit.Ucar = unitUcar
				break
			case "uhpr":
				unitUhpr := ""
				for input[i+1] > 31 {
					i++
					unitUhpr += string(input[i])
				}
				currentUnit.Uhpr = unitUhpr
				break
			case "uhrt":
				unitUhrt := ""
				for input[i+1] > 31 {
					i++
					unitUhrt += string(input[i])
				}
				currentUnit.Uhrt = unitUhrt
				break
			case "umpi":
				unitUmpi := ""
				for input[i+1] > 31 {
					i++
					unitUmpi += string(input[i])
				}
				currentUnit.Umpi = unitUmpi
				break
			case "usst":
				unitUsst := ""
				for input[i+1] > 31 {
					i++
					unitUsst += string(input[i])
				}
				currentUnit.Usst = unitUsst
				break
			case "usrg":
				unitUsrg := ""
				for input[i+1] > 31 {
					i++
					unitUsrg += string(input[i])
				}
				currentUnit.Usrg = unitUsrg
				break
			case "usma":
				unitUsma := ""
				for input[i+1] > 31 {
					i++
					unitUsma += string(input[i])
				}
				currentUnit.Usma = unitUsma
				break
			case "upgr":
				unitUpgrade := ""
				for input[i+1] > 31 {
					i++
					unitUpgrade += string(input[i])
				}
				currentUnit.Upgr = unitUpgrade
				break
			case "uabi":
				unitUabi := ""
				for input[i+1] > 31 {
					i++
					unitUabi += string(input[i])
				}
				currentUnit.Uabi = unitUabi
				break
			case "utip":
				unitUtip := ""
				for input[i+1] > 31 {
					i++
					unitUtip += string(input[i])
				}
				currentUnit.Utip = unitUtip
				break
			case "utub":
				unitUtub := ""
				for input[i+1] > 31 {
					i++
					unitUtub += string(input[i])
				}
				currentUnit.Utub = unitUtub
				break
			case "ugol":
				unitUgol := "" // ?
				for input[i+1] > 31 {
					i++
					unitUgol += string(input[i])
				}
				currentUnit.Ugol = unitUgol
				break
			case "utar":
				unitUtar := ""
				for input[i+1] > 31 {
					i++
					unitUtar += string(input[i])
				}
				currentUnit.Utar = unitUtar
				break
			case "umvt":
				unitUmvt := ""
				for input[i+1] > 31 {
					i++
					unitUmvt += string(input[i])
				}
				currentUnit.Umvt = unitUmvt
				break
			case "umvh":
				unitUmvh := "" // Movement height?
				for input[i+1] > 31 {
					i++
					unitUmvh += string(input[i])
				}
				currentUnit.Umvh = unitUmvh
				break
			case "ubld":
				unitUbld := "" // ? It's not buildables as that's ubui
				for input[i+1] > 31 {
					i++
					unitUbld += string(input[i])
				}
				currentUnit.Ubld = unitUbld
				break
			case "uhpm":
				unitUhpm := "" // ?
				for input[i+1] > 31 {
					i++
					unitUhpm += string(input[i])
				}
				currentUnit.Uhpm = unitUhpm
				break
			case "umvs":
				unitUmvs := "" // ?
				for input[i+1] > 31 {
					i++
					unitUmvs += string(input[i])
				}
				currentUnit.Umvs = unitUmvs
				break
			case "umdl":
				unitUmdl := ""
				for input[i+1] > 31 {
					i++
					unitUmdl += string(input[i])
				}
				currentUnit.Umdl = unitUmdl
				break
			case "uico":
				unitUico := ""
				for input[i+1] > 31 {
					i++
					unitUico += string(input[i])
				}
				currentUnit.Uico = unitUico
				break
			case "urac":
				unitUrac := ""
				for input[i+1] > 31 {
					i++
					unitUrac += string(input[i])
				}
				currentUnit.Urac = unitUrac
				break
			case "usnd":
				unitUsnd := "" // ?
				for input[i+1] > 31 {
					i++
					unitUsnd += string(input[i])
				}
				currentUnit.Usnd = unitUsnd
				break
			case "usca":
				unitUsca := "" // ?
				for input[i+1] > 31 {
					i++
					unitUsca += string(input[i])
				}
				currentUnit.Usca = unitUsca
				break
			case "uhot":
				unitUhot := "" // Hotkey?
				for input[i+1] > 31 {
					i++
					unitUhot += string(input[i])
				}
				currentUnit.Uhot = unitUhot
				break
			case "ubui":
				unitUbui := ""
				for input[i+1] > 31 {
					i++
					unitUbui += string(input[i])
				}
				currentUnit.Ubui = unitUbui
				break
			case "uclb":
				unitUclb := "" // ?
				for input[i+1] > 31 {
					i++
					unitUclb += string(input[i])
				}
				currentUnit.Uclb = unitUclb
				break
			case "ufoo":
				unitUfoo := "" // ?
				for input[i+1] > 31 {
					i++
					unitUfoo += string(input[i])
				}
				currentUnit.Ufoo = unitUfoo
				break
			case "ua1c":
				unitUa1c := "" // ? Has something to do with attack 1
				for input[i+1] > 31 {
					i++
					unitUa1c += string(input[i])
				}
				currentUnit.Ua1c = unitUa1c
				break
			case "ua1b":
				unitUa1b := "" // ? Has something to do with attack 1
				for input[i+1] > 31 {
					i++
					unitUa1b += string(input[i])
				}
				currentUnit.Ua1b = unitUa1b
				break
			case "ua1d":
				unitUa1d := "" // ? Has something to do with attack 1
				for input[i+1] > 31 {
					i++
					unitUa1d += string(input[i])
				}
				currentUnit.Ua1d = unitUa1d
				break
			case "ua1s":
				unitUa1s := "" // ? Has something to do with attack 1
				for input[i+1] > 31 {
					i++
					unitUa1s += string(input[i])
				}
				currentUnit.Ua1s = unitUa1s
				break
			case "ua1g":
				unitUa1g := "" // ? Has something to do with attack 1
				for input[i+1] > 31 {
					i++
					unitUa1g += string(input[i])
				}
				currentUnit.Ua1g = unitUa1g
				break
			case "ulum":
				unitUlum := "" // ?
				for input[i+1] > 31 {
					i++
					unitUlum += string(input[i])
				}
				currentUnit.Ulum = unitUlum
				break
			case "upoi":
				unitUpoi := "" // ?
				for input[i+1] > 31 {
					i++
					unitUpoi += string(input[i])
				}
				currentUnit.Upoi = unitUpoi
				break
			case "utra":
				unitUtra := "" // Unit trains?
				for input[i+1] > 31 {
					i++
					unitUtra += string(input[i])
				}
				currentUnit.Utra = unitUtra
				break
			case "upap":
				unitUpap := "" // ? A unit had unbuilable as value here
				for input[i+1] > 31 {
					i++
					unitUpap += string(input[i])
				}
				currentUnit.Upap = unitUpap
				break
			case "ubs1":
				unitUbs1 := "" // ?
				for input[i+1] > 31 {
					i++
					unitUbs1 += string(input[i])
				}
				currentUnit.Ubs1 = unitUbs1
				break
			case "udp1":
				unitUdp1 := "" // ?
				for input[i+1] > 31 {
					i++
					unitUdp1 += string(input[i])
				}
				currentUnit.Udp1 = unitUdp1
				break
			case "ubpy":
				unitUbpy := "" // ?
				for input[i+1] > 31 {
					i++
					unitUbpy += string(input[i])
				}
				currentUnit.Ubpy = unitUbpy
				break
			case "ua1t":
				unitUa1t := ""
				for input[i+1] > 31 {
					i++
					unitUa1t += string(input[i])
				}
				currentUnit.Ua1t = unitUa1t
				break
			case "utc2":
				unitUtc2 := ""
				for input[i+1] > 31 {
					i++
					unitUtc2 += string(input[i])
				}
				currentUnit.Utc2 = unitUtc2
				break
			case "ua1p":
				unitUa1p := ""
				for input[i+1] > 31 {
					i++
					unitUa1p += string(input[i])
				}
				currentUnit.Ua1p = unitUa1p
				break
			case "ubpx":
				unitUbpx := ""
				for input[i+1] > 31 {
					i++
					unitUbpx += string(input[i])
				}
				currentUnit.Ubpx = unitUbpx
				break
			case "uacq":
				unitUacq := "" // Unit acquisition?
				for input[i+1] > 31 {
					i++
					unitUacq += string(input[i])
				}
				currentUnit.Uacq = unitUacq
				break
			case "ua1r":
				unitUa1r := ""
				for input[i+1] > 31 {
					i++
					unitUa1r += string(input[i])
				}
				currentUnit.Ua1r = unitUa1r
				break
			case "ua1m":
				unitUa1m := ""
				for input[i+1] > 31 {
					i++
					unitUa1m += string(input[i])
				}
				currentUnit.Ua1m = unitUa1m
				break
			}
		}
	}

	mapValues := make([]*W3uData, len(unitMap))
	index := 0
	for _, value := range unitMap {
		mapValues[index] = value
		index++
	}

	jsonObject, err := json.Marshal(mapValues)
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
