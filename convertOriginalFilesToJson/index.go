package convertOriginalFilesToJson

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
	"ultimate-circassian-dictionary-helper/wordObject"
)

var (
	ruWordsMap  = make(map[string]bool)
	cirWordsMap = make(map[string]bool)
)

func CallAllConverts() {
	convert0()  // Ady-Ady_AIG.json
	convert1()  // Ady-Ady_AP.json
	convert2()  // Ady-Ara.json
	convert3()  // Ady-En.json
	convert4()  // Ady-En_Adam.json
	convert5()  // Ady-Rus_Qarden.json
	convert6()  // Ady-Rus_Sherdjes.json
	convert7()  // Ady-Rus_Tharkaho.json
	convert8()  // Ady-Tur_Huvaj.json
	convert9()  // En-Ady.json
	convert10() // En-Ady_Adam.json
	convert11() // En-Kbd-Jonty.json
	convert12() // En-Kbd-Ziwar.json
	convert13() // Kbd-Ar-Jonty.json
	convert14() // Kbd-En-2-Jonty.json
	convert15() // Kbd-En-Jonty.json
	convert16() // Kbd-En-Ziwar.json
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
	convert30()
	convert31()
	convert32()
	convert33()

	// print all the words in Circassian
	fmt.Printf("\n--All the words in Circassian--\n")
	fmt.Printf("Total number of words: %d\n", len(cirWordsMap))
	for key := range cirWordsMap {
		fmt.Printf("%s\n", key)
	}

	// print all the words in Russian
	fmt.Printf("\n--All the words in Russian--\n")
	fmt.Printf("Total number of words: %d\n", len(ruWordsMap))
	for key := range ruWordsMap {
		fmt.Printf("%s\n", key)
	}
}

// convert0() Адыгабзэм изэхэф гущыIалъ - Ady-Ady_AIG.json
func convert0() {
	dictObj := wordObject.NewDictionaryObject("Адыгабзэм изэхэф гущы1алъ (2006)", 0, "Ady", "Ru", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to cirWordsMap
		if _, ok := cirWordsMap[spelling]; !ok {
			cirWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert1() Адыгэбзэ псалъалъэ - Ady-Ady_AP.json
func convert1() {
	dictObj := wordObject.NewDictionaryObject("Адыгэ-урыс псалъалъэ (2008)", 1, "Kbd", "Ru", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to cirWordsMap
		if _, ok := cirWordsMap[spelling]; !ok {
			cirWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert2() Адыгэ-араб гущыIалъ - Ady-Ara.json
func convert2() {
	dictObj := wordObject.NewDictionaryObject("Adel abdulsalam lash", 2, "Ady", "Ar", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to cirWordsMap
		if _, ok := cirWordsMap[spelling]; !ok {
			cirWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++

		// Add to cirWordsMap
		if _, ok := cirWordsMap[key]; !ok {
			cirWordsMap[key] = true
		}
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert4() Адэм Шъэджашъ и адыгэбзэ-инджылыбзэ гущы1алъэ - Ady-En_Adam.json
func convert4() {
	dictObj := wordObject.NewDictionaryObject("Adam Shagash's Adyghe to English Dictionary (2020)", 4, "Ady", "En", "JSON")
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
				dictObj.Words[key].AddOneDefinition(definition.Meaning, definition.Examples)
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

		// Add to cirWordsMap
		if _, ok := cirWordsMap[key]; !ok {
			cirWordsMap[key] = true
		}
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert5() Къардэн - Ady-Rus_Qarden.json
func convert5() {
	dictObj := wordObject.NewDictionaryObject("Къардэн (1957)", 5, "Kbd", "Ru", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to cirWordsMap
		if _, ok := cirWordsMap[spelling]; !ok {
			cirWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert6() Ady-Rus_Sherdjes.json
func convert6() {
	dictObj := wordObject.NewDictionaryObject("Шэрджэс Алий - Яхуэмыфащэу лъэныкъуэ едгъэза псалъэхэр (2009)", 6, "Kbd", "Ru", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to cirWordsMap
		if _, ok := cirWordsMap[spelling]; !ok {
			cirWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert7() Ady-Rus_Tharkaho.json
func convert7() {
	dictObj := wordObject.NewDictionaryObject("Тхьаркъуахъо (1991)", 7, "Ady", "Ru", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to cirWordsMap
		if _, ok := cirWordsMap[spelling]; !ok {
			cirWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert8() Ady-Tur_Huvaj.json
func convert8() {
	dictObj := wordObject.NewDictionaryObject("Хъуажь - Circassian to Turkish (2007)", 8, "Ady/Kbd", "Tr", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)
		splitSpelling := strings.Split(spelling, "/")

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		for _, s := range splitSpelling {
			s = strings.Trim(s, " ")
			s = strings.Trim(s, ",")
			// Check if word exist, if it does, add definition, otherwise create new word
			if _, ok := dictObj.Words[s]; !ok {
				dictObj.Words[s] = wordObject.NewWordObject(spelling, "")
			}
			dictObj.Words[s].AddFullDefinitionInHtml(obj)
		}
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert10() En-Ady_Adam.json
func convert10() {
	dictObj := wordObject.NewDictionaryObject("Adam Shagash's English to Adyghe Dictionary (2020)", 10, "En", "Ady", "JSON")
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
				dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
			} else {
				meaningWithClarity := fmt.Sprintf("%s (%s)", definition.Word, *definition.Clarity)
				dictObj.Words[key].AddOneDefinition(meaningWithClarity, []wordObject.Example{})
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
				dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
			} else {
				dictObj.Words[key].AddOneDefinition(definition.Meaning, *definition.Examples)
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
				dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
			} else {
				meaningWithClarity := fmt.Sprintf("%s (%s)", definition.Word, *definition.Clarity)
				dictObj.Words[key].AddOneDefinition(meaningWithClarity, []wordObject.Example{})
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert17() Kbd-En_Amjad.json
func convert17() {
	dictObj := wordObject.NewDictionaryObject("Amjad Jaimoukha's Kabardian to English dictionary", 17, "Kbd", "En", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En_Amjad.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-En_Amjad.json", dictObj.Id)

	addNewLineAndTabToNumberDot := func(s string) string {
		s = strings.ReplaceAll(s, " 1.", "\n1.")
		s = strings.ReplaceAll(s, " 2.", "\n2.")
		s = strings.ReplaceAll(s, " 3.", "\n3.")
		s = strings.ReplaceAll(s, " 4.", "\n4.")
		s = strings.ReplaceAll(s, " 5.", "\n5.")
		s = strings.ReplaceAll(s, " 6.", "\n6.")
		s = strings.ReplaceAll(s, " 7.", "\n7.")
		s = strings.ReplaceAll(s, " 8.", "\n8.")
		s = strings.ReplaceAll(s, " 9.", "\n9.")
		s = strings.ReplaceAll(s, " 1)", "\n1)")
		s = strings.ReplaceAll(s, " 2)", "\n2)")
		s = strings.ReplaceAll(s, " 3)", "\n3)")
		s = strings.ReplaceAll(s, " 4)", "\n4)")
		s = strings.ReplaceAll(s, " 5)", "\n5)")
		s = strings.ReplaceAll(s, " 6)", "\n6)")
		s = strings.ReplaceAll(s, " 7)", "\n7)")
		s = strings.ReplaceAll(s, " 8)", "\n8)")
		s = strings.ReplaceAll(s, " 9)", "\n9)")
		return s
	}

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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			meaningAdjusted := addNewLineAndTabToNumberDot(definition.Meaning)
			dictObj.Words[key].AddOneDefinition(meaningAdjusted, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++

		// Add to cirWordsMap
		if _, ok := cirWordsMap[key]; !ok {
			cirWordsMap[key] = true
		}
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
				dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
			} else {
				dictObj.Words[key].AddOneDefinition(definition.Meaning, *definition.Examples)
			}
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++

		// Add to cirWordsMap
		if _, ok := cirWordsMap[key]; !ok {
			cirWordsMap[key] = true
		}
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert21() Kbd-Tu-Jonty.json
func convert21() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Kabardian to Turkish dictionary", 21, "Kbd", "Tr", "JSON")
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Meaning, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++

		// Add to ruWordsMap
		if _, ok := ruWordsMap[key]; !ok {
			ruWordsMap[key] = true
		}
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert22() Rus-Ady_Blaghoj.json
func convert23() {
	dictObj := wordObject.NewDictionaryObject("Блэгъожъ (1991)", 23, "Ru", "Ady", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to ruWordsMap
		if _, ok := ruWordsMap[spelling]; !ok {
			ruWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert24() Rus-Ady_UAG.json
func convert24() {
	dictObj := wordObject.NewDictionaryObject("Одэжьдэкъо (1960)", 24, "Ru", "Ady", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to ruWordsMap
		if _, ok := ruWordsMap[spelling]; !ok {
			ruWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert25() Rus-Ady_UAG.json
func convert25() {
	dictObj := wordObject.NewDictionaryObject("Урыс-адыгэ школ псалъалъэ (1991)", 25, "Ru", "Kbd", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)

		// Add to ruWordsMap
		if _, ok := ruWordsMap[spelling]; !ok {
			ruWordsMap[spelling] = true
		}
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert26() Tu-Kbd-Jonty.json
func convert26() {
	dictObj := wordObject.NewDictionaryObject("Jonty Yamisha's Turkish to Kabardian dictionary", 26, "Tr", "Kbd", "JSON")
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
	fileToStr = utils.ConvertIToCirStick(fileToStr)
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
			dictObj.Words[key].AddOneDefinition(definition.Word, []wordObject.Example{})
		}
		fmt.Printf("Index %d key %s\n", idx, key)
		idx++
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert27() Tur-Ady_Abaze.json
func convert27() {
	dictObj := wordObject.NewDictionaryObject("Ибрагим Алхаз Абазэ (2005)", 27, "Tr", "Kbd", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert28() Tur-Ady_Huvaj.json
func convert28() {
	dictObj := wordObject.NewDictionaryObject("Хъуажь - Turkish to Circassian (2007)", 28, "Tr", "Ady/Kbd", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert29() Tur-Ady_Teshu.json
func convert29() {
	dictObj := wordObject.NewDictionaryObject("Т1эшъу (1991)", 29, "Tr", "Ady", "HTML")
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
		spelling = utils.ConvertIToCirStick(spelling)

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, ",")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")
		obj = utils.ConvertIToCirStick(obj)

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddFullDefinitionInHtml(obj)
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert30() Ady-Rus_ThreeVolumes.txt
func convert30() {
	dictObj := wordObject.NewDictionaryObject("Адыгабзэм изэхэф гущы1алъ томищ мэхъу (2011)", 30, "Ady", "Ru", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Rus_ThreeVolumes.txt"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Rus_ThreeVolumes.json", dictObj.Id)
	invalidLinesList := make([]string, 0)
	entireFileStr := utils.ReadEntireFile(srcFilePath)
	lines := strings.Split(entireFileStr, "\n")
	mapRes := make(map[string][]string)
	var currentKey string
	var currentValue strings.Builder

	addNewLineAndTabToNumberDot := func(s string) string {
		s = strings.ReplaceAll(s, " 1.", "\n\t1.")
		s = strings.ReplaceAll(s, " 2.", "\n\t2.")
		s = strings.ReplaceAll(s, " 3.", "\n\t3.")
		s = strings.ReplaceAll(s, " 4.", "\n\t4.")
		s = strings.ReplaceAll(s, " 5.", "\n\t5.")
		s = strings.ReplaceAll(s, " 6.", "\n\t6.")
		s = strings.ReplaceAll(s, " 7.", "\n\t7.")
		s = strings.ReplaceAll(s, " 8.", "\n\t8.")
		s = strings.ReplaceAll(s, " 9.", "\n\t9.")

		// if string starts with a number
		if utils.StartsWithNumber(s) {
			s = strings.ReplaceAll(s, "1.", "\n\t1.")
			s = strings.ReplaceAll(s, "2.", "\n\t2.")
			s = strings.ReplaceAll(s, "3.", "\n\t3.")
			s = strings.ReplaceAll(s, "4.", "\n\t4.")
			s = strings.ReplaceAll(s, "5.", "\n\t5.")
			s = strings.ReplaceAll(s, "6.", "\n\t6.")
			s = strings.ReplaceAll(s, "7.", "\n\t7.")
			s = strings.ReplaceAll(s, "8.", "\n\t8.")
			s = strings.ReplaceAll(s, "9.", "\n\t9.")
		}
		return s
	}

	for idx, line := range lines {
		trimmedLine := utils.TrimSlashes(line)

		// Split the line into words
		words := strings.Fields(trimmedLine)

		// Add new line and tab to number dot
		trimmedLine = addNewLineAndTabToNumberDot(trimmedLine)

		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, trimmedLine)
		} else if len(words) > 0 && utils.IsFullyCapitalized(words[0]) && !utils.StartsWithNumber(words[0]) {
			// New key detected
			if currentKey != "" {
				keyWithoutSuffix := utils.RemoveSuffixes(currentKey)
				mapRes[keyWithoutSuffix] = append(mapRes[keyWithoutSuffix], currentValue.String())
			}
			currentKey = utils.RemoveSuffixes(words[0])
			currentValue.Reset()
			currentValue.WriteString(trimmedLine)
		} else {
			// Continuation of the current value
			if currentValue.Len() > 0 {
				currentValue.WriteString(" ")
			}
			currentValue.WriteString(trimmedLine)
		}

		fmt.Printf("line %d: %s\n", idx, line)
	}

	// Add the last key-value pair
	if currentKey != "" {
		keyWithoutSuffix := utils.RemoveSuffixes(currentKey)
		mapRes[keyWithoutSuffix] = append(mapRes[keyWithoutSuffix], currentValue.String())
	}

	for key, values := range mapRes {
		// Check if word exist, if it does, add definition, otherwise create new word
		spelling := strings.ToLower(key)
		spelling = utils.ConvertIToCirStick(spelling)

		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		for _, value := range values {
			value = utils.ConvertIToCirStick(value)
			dictObj.Words[spelling].AddOneDefinition(value, []wordObject.Example{})
		}
	}

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert31() Tu-Ady_Hilmi.txt
func convert31() {
	dictObj := wordObject.NewDictionaryObject("Ацумыжъ Хилми (2013)", 31, "Tr", "Ady", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Tu-Ady_Hilmi.txt"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Tu-Ady_Hilmi.json", dictObj.Id)
	invalidLinesList := make([]string, 0)
	entireFileStr := utils.ReadEntireFile(srcFilePath)
	lines := strings.Split(entireFileStr, "\n")
	mapRes := make(map[string][]string)
	var currentKey string
	var currentValue strings.Builder

	addNewLineAndTabToNumberDot := func(s string) string {
		s = strings.ReplaceAll(s, " 1.", "\n\t1.")
		s = strings.ReplaceAll(s, " 2.", "\n\t2.")
		s = strings.ReplaceAll(s, " 3.", "\n\t3.")
		s = strings.ReplaceAll(s, " 4.", "\n\t4.")
		s = strings.ReplaceAll(s, " 5.", "\n\t5.")
		s = strings.ReplaceAll(s, " 6.", "\n\t6.")
		s = strings.ReplaceAll(s, " 7.", "\n\t7.")
		s = strings.ReplaceAll(s, " 8.", "\n\t8.")
		s = strings.ReplaceAll(s, " 9.", "\n\t9.")
		s = strings.ReplaceAll(s, " 1)", "\n\t\t1)")
		s = strings.ReplaceAll(s, " 2)", "\n\t\t2)")
		s = strings.ReplaceAll(s, " 3)", "\n\t\t3)")
		s = strings.ReplaceAll(s, " 4)", "\n\t\t4)")
		s = strings.ReplaceAll(s, " 5)", "\n\t\t5)")
		s = strings.ReplaceAll(s, " 6)", "\n\t\t6)")
		s = strings.ReplaceAll(s, " 7)", "\n\t\t7)")
		s = strings.ReplaceAll(s, " 8)", "\n\t\t8)")
		s = strings.ReplaceAll(s, " 9)", "\n\t\t9)")
		return s
	}

	for idx, line := range lines {
		trimmedLine := utils.TrimSlashes(line)

		// Split the line into words
		words := strings.Fields(trimmedLine)

		// Add new line and tab to number dot
		trimmedLine = addNewLineAndTabToNumberDot(trimmedLine)

		if strings.Trim(trimmedLine, " ") == "" {
			// do nothing
		} else if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, trimmedLine)
		} else if len(words) == 1 && len(words[0]) == 3 && strings.Contains(words[0], "-") {
			invalidLinesList = append(invalidLinesList, trimmedLine)
		} else if len(words) > 0 && utils.IsFullyCapitalized(words[0]) && !utils.StartsWithNumber(words[0]) && !utils.StartsWithSpecialCharacter(words[0]) {
			// New key detected
			if currentKey != "" {
				keyWithoutSuffix := utils.RemoveSuffixes(currentKey)
				mapRes[keyWithoutSuffix] = append(mapRes[keyWithoutSuffix], currentValue.String())
			}
			currentKey = utils.RemoveSuffixes(words[0])
			currentValue.Reset()
			currentValue.WriteString(trimmedLine)
		} else {
			// Continuation of the current value
			if currentValue.Len() > 0 {
				currentValue.WriteString(" ")
			}
			currentValue.WriteString(trimmedLine)
		}

		fmt.Printf("line %d: %s\n", idx, line)
	}

	// Add the last key-value pair
	if currentKey != "" {
		keyWithoutSuffix := utils.RemoveSuffixes(currentKey)
		mapRes[keyWithoutSuffix] = append(mapRes[keyWithoutSuffix], currentValue.String())
	}

	for key, values := range mapRes {
		// Check if word exist, if it does, add definition, otherwise create new word
		spelling := strings.ToLower(key)
		spelling = strings.Trim(spelling, ":")

		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		for _, value := range values {
			value = utils.ConvertIToCirStick(value)
			dictObj.Words[spelling].AddOneDefinition(value, []wordObject.Example{})
		}
	}

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert32() Rus-Kbd_Nalchik_2013.txt
func convert32() {
	dictObj := wordObject.NewDictionaryObject("Еджап1эм папщ1э урыс-адыгэ псалъалъэ (2013)", 32, "Ru", "Kbd", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Rus-Kbd_Nalchik_2013.txt"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Rus-Kbd_Nalchik_2013.json", dictObj.Id)
	invalidLinesList := make([]string, 0)
	entireFileStr := utils.ReadEntireFile(srcFilePath)
	lines := strings.Split(entireFileStr, "\n")

	addNewLineAndTabToNumberDot := func(s string) string {
		s = strings.ReplaceAll(s, " 1.", "\n\t1.")
		s = strings.ReplaceAll(s, " 2.", "\n\t2.")
		s = strings.ReplaceAll(s, " 3.", "\n\t3.")
		s = strings.ReplaceAll(s, " 4.", "\n\t4.")
		s = strings.ReplaceAll(s, " 5.", "\n\t5.")
		s = strings.ReplaceAll(s, " 6.", "\n\t6.")
		s = strings.ReplaceAll(s, " 7.", "\n\t7.")
		s = strings.ReplaceAll(s, " 8.", "\n\t8.")
		s = strings.ReplaceAll(s, " 9.", "\n\t9.")
		s = strings.ReplaceAll(s, " 1)", "\n\t\t1)")
		s = strings.ReplaceAll(s, " 2)", "\n\t\t2)")
		s = strings.ReplaceAll(s, " 3)", "\n\t\t3)")
		s = strings.ReplaceAll(s, " 4)", "\n\t\t4)")
		s = strings.ReplaceAll(s, " 5)", "\n\t\t5)")
		s = strings.ReplaceAll(s, " 6)", "\n\t\t6)")
		s = strings.ReplaceAll(s, " 7)", "\n\t\t7)")
		s = strings.ReplaceAll(s, " 8)", "\n\t\t8)")
		s = strings.ReplaceAll(s, " 9)", "\n\t\t9)")
		return s
	}

	for idx, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "\u200B")
		line = strings.Trim(line, "\uFEFF")
		line = strings.Trim(line, "\u200D")
		line = strings.Trim(line, "\u200C")
		line = strings.Trim(line, "")

		// Split the line into words
		words := strings.Fields(line)

		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, line)
			continue
		}

		spelling := strings.ToLower(words[0])
		value := addNewLineAndTabToNumberDot(line)
		value = utils.ConvertIToCirStick(value)

		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		dictObj.Words[spelling].AddOneDefinition(value, []wordObject.Example{})

		fmt.Printf("line %d: %s\n", idx, line)
	}

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
}

// convert33() Ady-Rus-1960.txt
func convert33() {
	dictObj := wordObject.NewDictionaryObject("Адыгабзэм изэхэф гущы1алъ жъы (1960)", 33, "Ady", "Ru", "JSON")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Rus-1960.txt"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Ady-Rus-1960.json", dictObj.Id)
	invalidLinesList := make([]string, 0)
	entireFileStr := utils.ReadEntireFile(srcFilePath)
	lines := strings.Split(entireFileStr, "\n")
	mapRes := make(map[string][]string)
	var currentKey string
	var currentValue strings.Builder

	addNewLineAndTabToNumberDot := func(s string) string {
		s = strings.ReplaceAll(s, " 1.", "\n\t1.")
		s = strings.ReplaceAll(s, " 2.", "\n\t2.")
		s = strings.ReplaceAll(s, " 3.", "\n\t3.")
		s = strings.ReplaceAll(s, " 4.", "\n\t4.")
		s = strings.ReplaceAll(s, " 5.", "\n\t5.")
		s = strings.ReplaceAll(s, " 6.", "\n\t6.")
		s = strings.ReplaceAll(s, " 7.", "\n\t7.")
		s = strings.ReplaceAll(s, " 8.", "\n\t8.")
		s = strings.ReplaceAll(s, " 9.", "\n\t9.")

		// if string starts with a number
		if utils.StartsWithNumber(s) {
			s = strings.ReplaceAll(s, "1.", "\n\t1.")
			s = strings.ReplaceAll(s, "2.", "\n\t2.")
			s = strings.ReplaceAll(s, "3.", "\n\t3.")
			s = strings.ReplaceAll(s, "4.", "\n\t4.")
			s = strings.ReplaceAll(s, "5.", "\n\t5.")
			s = strings.ReplaceAll(s, "6.", "\n\t6.")
			s = strings.ReplaceAll(s, "7.", "\n\t7.")
			s = strings.ReplaceAll(s, "8.", "\n\t8.")
			s = strings.ReplaceAll(s, "9.", "\n\t9.")
		}
		return s
	}

	removeFirstWordSpaces := func(line string) string {
		// Regular expression to match spaces between two capitalized Cyrillic letters or a letter and the letter 'I'
		re := regexp.MustCompile(`([А-ЯЁI])\s+([А-ЯЁI])`)
		// Replace the matched spaces with an empty string
		modifiedLine := re.ReplaceAllString(line, "$1$2")
		return modifiedLine
	}

	for idx, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, "\u200B")
		line = strings.Trim(line, "\uFEFF")
		line = strings.Trim(line, "\u200D")
		line = strings.Trim(line, "\u200C")
		line = strings.Trim(line, "")
		line = removeFirstWordSpaces(line)
		trimmedLine := utils.TrimSlashes(line)

		// Split the line into words
		words := strings.Fields(trimmedLine)

		// Add new line and tab to number dot
		trimmedLine = addNewLineAndTabToNumberDot(trimmedLine)

		if len(words) == 0 {
			invalidLinesList = append(invalidLinesList, trimmedLine)
		} else if len(words) > 0 && utils.IsFullyCapitalized(words[0]) && !utils.StartsWithNumber(words[0]) && !utils.StartsWithSpecialCharacter(words[0]) {
			// New key detected
			if currentKey != "" {
				keyWithoutSuffix := utils.RemoveSuffixes(currentKey)
				mapRes[keyWithoutSuffix] = append(mapRes[keyWithoutSuffix], currentValue.String())
			}
			currentKey = utils.RemoveSuffixes(words[0])
			currentValue.Reset()
			currentValue.WriteString(trimmedLine)
		} else {
			// Continuation of the current value
			if currentValue.Len() > 0 {
				currentValue.WriteString(" ")
			}
			currentValue.WriteString(trimmedLine)
		}

		fmt.Printf("line %d: %s\n", idx, line)
	}

	// Add the last key-value pair
	if currentKey != "" {
		keyWithoutSuffix := utils.RemoveSuffixes(currentKey)
		mapRes[keyWithoutSuffix] = append(mapRes[keyWithoutSuffix], currentValue.String())
	}

	for key, values := range mapRes {
		// Check if word exist, if it does, add definition, otherwise create new word
		spelling := strings.ToLower(key)
		spelling = utils.ConvertIToCirStick(spelling)

		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, "")
		}
		for _, value := range values {
			value = utils.ConvertIToCirStick(value)
			dictObj.Words[spelling].AddOneDefinition(value, []wordObject.Example{})
		}
	}

	// print invalid lines
	fmt.Printf("\n--Invalid lines in %s:--\n", srcFilePath)
	for idx, line := range invalidLinesList {
		fmt.Printf("%d. %s\n", idx, line)
	}

	utils.CreateFileWithDictionaryObject(distFilePath, dictObj)
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
