package converts

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
	"ultimate-circassian-dictionary-helper/wordObject"
)

/*
0, "Адыгабзэм изэхэф гущыIалъ" - Ady-Ady_AIG.json
1, "Адыгэбзэ псалъалъэ" - Ady-Ady_AP.json
2, "Адыгэ-араб гущыIалъ" - Ady-Ara.json
3, "адыгэбзэ-инджылыбзэ гущы1алъЭ - Ady-En.json
4, "Адэм Шъэджашъ и адыгэбзэ-инджылыбзэ гущы1алъЭ - Ady-En_Adam.json
5 "Къардэн" - Ady-Rus_Qarden.json
6, "Яхуэмыфащэу лъэныкъуэ едгъэза псалъэхэр" - Ady-Rus_Sherdjes.json,
7, "Тхьаркъуахъо" - Ady-Rus_Tharkaho.json,
8, "Хъуажь" - Ady-Tur_Huvaj.json,
8,

8, "Блэгъожъ",
9, "Одэжьдэкъо",
10, "Урыс-адыгэ школ псалъалъэ",
11, "Абазэ",
12, "Хъуажь",
13, "ТIэшъу",
14, "Adyghe-English Dictionary",
15, "Adam Adyghe-English Dictionary",
16, "English-Adyghe Dictionary",
17, "Adam English-Adyghe Dictionary",
17, "Jonty English-Kabardian Dictionary",
18, "Ziwar English-Kabardian Dictionary",
19, "Jonty Kabardian-Arabic Dictionary",
20, "Amjad Kabardian-English Dictionary",
21, "Jonty Kabardian-English Dictionary with Examples",
22, "Jonty Kabardian-English Dictionary",
23, "Ziwar Kabardian-English Dictionary",
24, "Ziwar Kabardian-English Dictionary",
25, "Kabardian-Russian Dictionary",
26, "Jonty Kabardian-Russian Dictionary",
27, "Jonty Kabardian-Turkish Dictionary",
28, "Jonty Russian-Kabardian Dictionary",
29, "Jonty Turkish-Kabardian Dictionary",
*/

func CallAllConverts() {
	// convert0()
	// convert1()
	//convert2()
	//convert3()
	convert4()
	// convert22()
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
	dictObj := wordObject.NewDictionaryObject("Адыгэ-араб гущыIалъ", 2, "Ady", "Ar", "JSON")
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

// convert22() Kbd-En-Jonty.json
func convert22() {
	dictObj := wordObject.NewDictionaryObject("Jonty Kabardian-English Dictionary", 22, "Kbd", "En")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En-Jonty.json"
	distFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\%d.Kbd-En-Jonty.json", dictObj.Id)
	invalidLinesList := make([]string, 0)

	utils.ReadFileLineByLine(srcFilePath, func(line string, idx int) {
		var errMsg string
		// split
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
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
		obj = strings.TrimSuffix(obj, ",")
		obj = convertIToCirStick(obj)
		var objMap struct {
			Type        string `json:"type"`
			Definitions []struct {
				Meaning string `json:"meaning"`
			} `json:"definitions"`
		}
		if err := json.Unmarshal([]byte(obj), &objMap); err != nil {
			errMsg = fmt.Sprintf("Invalid line %d: %s", idx, line)
			invalidLinesList = append(invalidLinesList, errMsg)
			fmt.Printf("%s\n", errMsg)
			return
		}

		// Check if word exist, if it does, add definition, otherwise create new word
		if _, ok := dictObj.Words[spelling]; !ok {
			dictObj.Words[spelling] = wordObject.NewWordObject(spelling, objMap.Type)
		}
		for _, definition := range objMap.Definitions {
			dictObj.Words[spelling].AddOneSimpleDefinition(definition.Meaning)
		}
		fmt.Printf("line %d: %s\n", idx, line)
	})

	// print invalid lines
	fmt.Printf("\n--Invalid lines:--\n")
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
