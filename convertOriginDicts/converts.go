package converts

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
	"ultimate-circassian-dictionary-helper/wordObject"
)

func CallAllConverts() {
	convert0()
	convert1()
	convert2()
	convert3()
	convert4()
	convert5()
	convert6()
	convert7()
	convert8()
	convert9()
	convert10()
	convert11()
	convert12()
	convert13()
	convert14()
	convert15()
	convert16()
	convert17()
	convert18()
	convert19()
	convert20()
	convert21()
	convert22()
	convert23()
	convert24()
	convert25()
	convert26()
	convert27()
	convert28()
	convert29()
}

// convert0() Адыгабзэм изэхэф гущыIалъ - Ady-Ady_AIG.json
func convert0() {
	dictObj := wordObject.NewDictionaryObject("Адыгабзэм изэхэф гущы1алъ", 0, "Ady", "Ru", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Ady_AIG.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Ady_AIG.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert1() Адыгэбзэ псалъалъэ - Ady-Ady_AP.json
func convert1() {
	dictObj := wordObject.NewDictionaryObject("Адыгэбзэ псалъалъэ", 1, "Kbd", "Ru", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Ady_AP.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Ady_AP.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert2() Адыгэ-араб гущыIалъ - Ady-Ara.json
func convert2() {
	dictObj := wordObject.NewDictionaryObject("Адыгэ-араб гущыIалъ", 2, "Ady", "Ar", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Ara.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Ara.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert3() Адыгэбзэ-инджылыбзэ гущы1алъэ - Ady-En.json
func convert3() {
	dictObj := wordObject.NewDictionaryObject("Адыгэбзэ-инджылыбзэ гущы1алъэ", 3, "Ady", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-En.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-En.json", dictObj.Id)

	// covert file to json
	var originalDict map[string]struct {
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert4() Адэм Шъэджашъ и адыгэбзэ-инджылыбзэ гущы1алъэ - Ady-En_Adam.json
func convert4() {
	dictObj := wordObject.NewDictionaryObject("Адыгэбзэ-инджылыбзэ гущы1алъэ", 4, "Ady", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-En_Adam.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-En_Adam.json", dictObj.Id)

	// covert file to json
	var originalDict map[string]struct {
		Word     string `json:"word"`
		Type     string `json:"type"`
		Synonyms []struct {
			Word string `json:"word"`
		} `json:"synonyms"`
		Shapsug     string `json:"shapsug"`
		Kabardian   string `json:"kabardian"`
		Definitions []struct {
			Meaning  string               `json:"meaning"`
			Examples []wordObject.Example `json:"examples"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n. Reason: %s", srcFilePath, err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Definitions {
			if strings.Contains(definition.Meaning, "alternative form of") {
				redirectWord, ok := extractRedirectWord(definition.Meaning)
				if ok {
					dictObj.Words[key].Redirect = redirectWord
				}
			} else {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, definition.Examples)
			}
		}
		if value.Shapsug != "" {
			dictObj.Words[key].AddCognate("Shapsug", value.Shapsug)
		}
		if value.Kabardian != "" {
			dictObj.Words[key].AddCognate("Kabardian", value.Kabardian)
		}
		if value.Synonyms != nil && len(value.Synonyms) > 0 {
			for _, synonym := range value.Synonyms {
				dictObj.Words[key].AddSynonym(synonym.Word, "")
			}
		}
		if value.Type != "" {
			dictObj.Words[key].Type = value.Type
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert5() Къардэн - Ady-Rus_Qarden.json
func convert5() {
	dictObj := wordObject.NewDictionaryObject("Къардэн", 5, "Ady", "Ru", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Rus_Qarden.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Rus_Qarden.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert6() Ady-Rus_Sherdjes.json
func convert6() {
	dictObj := wordObject.NewDictionaryObject("Яхуэмыфащэу лъэныкъуэ едгъэза псалъэхэр", 6, "Kbd", "Ru", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Rus_Sherdjes.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Rus_Sherdjes.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert7() Ady-Rus_Tharkaho.json
func convert7() {
	dictObj := wordObject.NewDictionaryObject("Тхьаркъуахъо", 7, "Ady", "Ru", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Rus_Tharkaho.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Rus_Tharkaho.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert8() Ady-Tur_Huvaj.json
func convert8() {
	dictObj := wordObject.NewDictionaryObject("Хъуажь", 8, "Ady/Kbd", "Tu", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Tur_Huvaj.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Tur_Huvaj.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s.\n(Reason: contains 0 or 1 ':')", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s.\n(Reason: Does not contain 2 \")", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert9() En-Ady.json
func convert9() {
	dictObj := wordObject.NewDictionaryObject("Simple English to Adyghe dictionary", 9, "En", "Ady", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\En-Ady.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.En-Ady.json", dictObj.Id)

	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word string `json:"word"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert10() En-Ady_Adam.json
func convert10() {
	dictObj := wordObject.NewDictionaryObject("Адэм Шъэджашъ и инджылыбзэ-адыгэбзэ гущы1алъэ", 10, "En", "Ady", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\En-Ady_Adam.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.En-Ady_Adam.json", dictObj.Id)

	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word    string  `json:"word"`
			Clarity *string `json:"clarity"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			if definition.Clarity == nil {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
			} else {
				meaningWithClarity := fmt.Sprintf("%s (%s)", definition.Word, *definition.Clarity)
				dictObj.Words[key].AddOneAdvancedDefinition(meaningWithClarity, []wordObject.Example{})
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert11() En-Kbd-Jonty.json
func convert11() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's English to Kabardian dictionary", 11, "En", "Kbd", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\En-Kbd-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.En-Kbd-Jonty.json", dictObj.Id)

	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word string `json:"word"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert12() En-Kbd-Ziwar.json
func convert12() {
	dictObj := wordObject.NewDictionaryObject("Ziwar Gish's English to Kabardian dictionary", 12, "En", "Kbd", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\En-Kbd-Ziwar.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.En-Kbd-Ziwar.json", dictObj.Id)

	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word string `json:"word"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert13() Kbd-Ar-Jonty.json
func convert13() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to Arabic dictionary", 13, "Kbd", "Ar", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-Ar-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-Ar-Jonty.json", dictObj.Id)

	// {"type":"verb","definitions":[{"meaning":"يستلم"}]}
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert14() Kbd-En-2-Jonty.json
func convert14() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to English dictionary 2", 14, "Kbd", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En-2-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-En-2-Jonty.json", dictObj.Id)

	// {"type":"adverb","definitions":[{"meaning":"very"},{"examples":[{"sentence":"he behaved badly.","translation":"ар щiыкiэншэу щытащ."}],"meaning":"badly"},{"meaning":"poorly"},{"meaning":"badly"},{"meaning":"badly"},{"meaning":"badly"},{"meaning":"wrongly"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning  string                `json:"meaning"`
			Examples *[]wordObject.Example `json:"examples"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			if definition.Examples == nil {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
			} else {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, *definition.Examples)
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert15() Kbd-En-Jonty.json
func convert15() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to English dictionary", 15, "Kbd", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-En-Jonty.json", dictObj.Id)

	// {"type":"adjective","definitions":[{"meaning":"delicate"},{"meaning":"meek"},{"meaning":"courteous"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert16() Kbd-En-Ziwar.json
func convert16() {
	dictObj := wordObject.NewDictionaryObject("Ziwar Gish's Kabardian to English dictionary", 16, "Kbd", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En-Ziwar.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-En-Ziwar.json", dictObj.Id)

	// {"links":[{"clarity":"bad","word":"awful"},{"word":"bad"},{"word":"dashing"},{"word":"dashing"},{"word":"nasty"}]},
	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word    string  `json:"word"`
			Clarity *string `json:"clarity"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			if definition.Clarity == nil {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
			} else {
				meaningWithClarity := fmt.Sprintf("%s (%s)", definition.Word, *definition.Clarity)
				dictObj.Words[key].AddOneAdvancedDefinition(meaningWithClarity, []wordObject.Example{})
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert17() Kbd-En_Amjad.json
func convert17() {
	dictObj := wordObject.NewDictionaryObject("Amjad Jaimoukha's Kabardian to English dictionary", 17, "Kbd", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En_Amjad.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-En_Amjad.json", dictObj.Id)

	// "1уеин": {"type":"intransitive verb","definitions":[{"meaning":"to bleat (of goats)"},{"meaning":"to stick with putty, to do up (ex.: window)"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert18() Kbd-Ru&En.json
func convert18() {
	dictObj := wordObject.NewDictionaryObject("Kabardian to Russian & English", 18, "Kbd", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-Ru&En.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-Ru&En.json", dictObj.Id)

	// "1уеин": {"type":"intransitive verb","definitions":[{"meaning":"to bleat (of goats)"},{"meaning":"to stick with putty, to do up (ex.: window)"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert19() Kbd-Ru-2-Jonty.json
func convert19() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to Russian dictionary 2", 19, "Kbd", "Ru", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-Ru-2-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-Ru-2-Jonty.json", dictObj.Id)

	// {"type":"noun","definitions":[{"examples":[{"sentence":"мы легко поднялись на холм.","translation":"дэ тыншу iуащхьэм дыдэкiуеящ."}],"meaning":"холм"},{"meaning":"курган"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning  string                `json:"meaning"`
			Examples *[]wordObject.Example `json:"examples"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			if definition.Examples == nil {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
			} else {
				dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, *definition.Examples)
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert20() Kbd-Ru-Jonty.json
func convert20() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to Russian dictionary", 20, "Kbd", "Ru", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-Ru-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-Ru-Jonty.json", dictObj.Id)

	// {"type":"adjective","definitions":[{"meaning":"деликатный"},{"meaning":"кроткий"},{"meaning":"обходительный"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert21() Kbd-Tu-Jonty.json
func convert21() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to Turkish dictionary", 21, "Kbd", "Tu", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-Tu-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-Tu-Jonty.json", dictObj.Id)

	// {"type":"adjective","definitions":[{"meaning":"деликатный"},{"meaning":"кроткий"},{"meaning":"обходительный"}]},
	// covert file to json
	var originalDict map[string]struct {
		Type        string `json:"type"`
		Definitions []struct {
			Meaning string `json:"meaning"`
		} `json:"definitions"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		fmt.Printf("Reason: %s\n", err.Error())
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
			dictObj.Words[key].Type = value.Type
		}
		for _, definition := range value.Definitions {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert22() Ru-Kbd-Jonty.json
func convert22() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Russian to Kabardian dictionary", 22, "Ru", "Kbd", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ru-Kbd-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ru-Kbd-Jonty.json", dictObj.Id)

	// {"type":"adjective","definitions":[{"meaning":"деликатный"},{"meaning":"кроткий"},{"meaning":"обходительный"}]},
	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word string `json:"word"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert22() Rus-Ady_Blaghoj.json
func convert23() {
	dictObj := wordObject.NewDictionaryObject("Блэгъожъ", 23, "Ru", "Ady", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Rus-Ady_Blaghoj.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Rus-Ady_Blaghoj.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert24() Rus-Ady_UAG.json
func convert24() {
	dictObj := wordObject.NewDictionaryObject("Одэжьдэкъо", 24, "Ru", "Ady", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Rus-Ady_UAG.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Rus-Ady_UAG.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert25() Rus-Ady_UAG.json
func convert25() {
	dictObj := wordObject.NewDictionaryObject("Урыс-адыгэ школ псалъалъэ", 25, "Ru", "Kbd", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Rus-Ady_UASP.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Rus-Ady_UASP.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert26() Tu-Kbd-Jonty.json
func convert26() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Turkish to Kabardian dictionary", 26, "Tu", "Kbd", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Tu-Kbd-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Tu-Kbd-Jonty.json", dictObj.Id)

	// {"type":"adjective","definitions":[{"meaning":"деликатный"},{"meaning":"кроткий"},{"meaning":"обходительный"}]},
	// covert file to json
	var originalDict map[string]struct {
		Links []struct {
			Word string `json:"word"`
		} `json:"links"`
	}

	fileToStr := utils.ReadEntireFile(srcFilePath)
	fileToStr = strings.ReplaceAll(fileToStr, "\\\"", "'")
	fileToStr = convertIToCirStick(fileToStr)
	fileStrToBytes := []byte(fileToStr)
	if err := json.Unmarshal(fileStrToBytes, &originalDict); err != nil {
		fmt.Printf("Invalid json file: %s\n", srcFilePath)
		return
	}

	// Check if word exist, if it does, add definition, otherwise create new word
	idx := 0
	for key, value := range originalDict {
		if _, ok := dictObj.Words[key]; !ok {
			dictObj.Words[key] = wordObject.NewWordObject(key, "")
		}
		for _, definition := range value.Links {
			dictObj.Words[key].AddOneAdvancedDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert27() Tur-Ady_Abaze.json
func convert27() {
	dictObj := wordObject.NewDictionaryObject("Абазэ", 27, "Tu", "Kbd", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Tur-Ady_Abaze.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Tur-Ady_Abaze.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert28() Tur-Ady_Huvaj.json
func convert28() {
	dictObj := wordObject.NewDictionaryObject("Абазэ", 28, "Tu", "Ady/Kbd", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Tur-Ady_Huvaj.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Tur-Ady_Huvaj.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

// convert29() Tur-Ady_Teshu.json
func convert29() {
	dictObj := wordObject.NewDictionaryObject("ТIэшъу", 29, "Tu", "Ady", "HTML")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Tur-Ady_Teshu.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Tur-Ady_Teshu.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line (0) %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// first part
		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")
		spelling = strings.ToLower(spelling)
		spelling = convertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = convertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneSimpleDefinition(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithContent(distFilePath, dictObj)
}

func convertIToCirStick(str string) string {
	str = strings.ReplaceAll(str, "ӏ", "1")
	// small i
	str = strings.ReplaceAll(str, "iа", "1а")
	str = strings.ReplaceAll(str, "iо", "1о")
	str = strings.ReplaceAll(str, "iе", "1е")
	str = strings.ReplaceAll(str, "iя", "1я")
	str = strings.ReplaceAll(str, "iы", "1ы")
	str = strings.ReplaceAll(str, "iэ", "1э")
	str = strings.ReplaceAll(str, "iи", "1и")
	str = strings.ReplaceAll(str, "iу", "1у")
	str = strings.ReplaceAll(str, "кi", "к1")
	str = strings.ReplaceAll(str, "шi", "ш1")
	str = strings.ReplaceAll(str, "пi", "п1")
	str = strings.ReplaceAll(str, "цi", "ц1")
	str = strings.ReplaceAll(str, "сi", "с1")
	str = strings.ReplaceAll(str, "чi", "ч1")
	str = strings.ReplaceAll(str, "тi", "т1")
	str = strings.ReplaceAll(str, "лi", "л1")
	str = strings.ReplaceAll(str, "фi", "ф1")
	str = strings.ReplaceAll(str, "щi", "щ1")

	// upper I
	str = strings.ReplaceAll(str, "Iа", "1а")
	str = strings.ReplaceAll(str, "Iо", "1о")
	str = strings.ReplaceAll(str, "Iе", "1е")
	str = strings.ReplaceAll(str, "Iя", "1я")
	str = strings.ReplaceAll(str, "Iы", "1ы")
	str = strings.ReplaceAll(str, "Iэ", "1э")
	str = strings.ReplaceAll(str, "Iи", "1и")
	str = strings.ReplaceAll(str, "Iу", "1у")
	str = strings.ReplaceAll(str, "кI", "к1")
	str = strings.ReplaceAll(str, "шI", "ш1")
	str = strings.ReplaceAll(str, "пI", "п1")
	str = strings.ReplaceAll(str, "цI", "ц1")
	str = strings.ReplaceAll(str, "сI", "с1")
	str = strings.ReplaceAll(str, "чI", "ч1")
	str = strings.ReplaceAll(str, "тI", "т1")
	str = strings.ReplaceAll(str, "лI", "л1")
	str = strings.ReplaceAll(str, "фI", "ф1")
	str = strings.ReplaceAll(str, "щI", "щ1")
	return str
}

// extractRedirectWord checks if a string matches the pattern "alternative form of \"кiэхьан\""
// and extracts the text between the escaped double quotes.
func extractRedirectWord(input string) (string, bool) {
	// Define the regular expression pattern
	pattern := `alternative form of "(.*?)"`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the substring that matches the pattern
	match := re.FindStringSubmatch(input)

	// Check if there was a match and extract the text between the quotes
	if len(match) > 1 {
		return match[1], true
	}
	return "", false
}
