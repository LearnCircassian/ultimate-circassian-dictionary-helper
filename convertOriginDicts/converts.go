package converts

import (
	"encoding/json"
	"fmt"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
	"ultimate-circassian-dictionary-helper/wordObject"
)

/*
0, "Адыгабзэм изэхэф гущыIалъ",
1, "Адыгэбзэ псалъалъэ",
2, "Адыгэ-араб гущыIалъ",
3, "Къардэн",
4, "Яхуэмыфащэу лъэныкъуэ едгъэза псалъэхэр",
5, "Тхьаркъуахъо",
6, "Хъуажь",
7, "Блэгъожъ",
8, "Одэжьдэкъо",
9, "Урыс-адыгэ школ псалъалъэ",
10, "Абазэ",
11, "Хъуажь",
12, "ТIэшъу",
13, "Adyghe-English Dictionary",
14, "Adam Adyghe-English Dictionary",
15, "English-Adyghe Dictionary",
16, "Adam English-Adyghe Dictionary",
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

type JontyWord struct {
	Type        string `json:"type"`
	Definitions []struct {
		Meaning string `json:"meaning"`
	} `json:"definitions"`
}

func CallAllConverts() {
	convert1()
	// convert22()
}

// convert1() Адыгэбзэ псалъалъэ
func convert1() {
	dictObj := wordObject.NewDictionaryObject("Адыгэбзэ псалъалъэ", 1, "Kbd", "Ru")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Ady-Ady_AIG.json"
	distFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\Ady-Ady_AIG.json"
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
		spelling = strings.ReplaceAll(spelling, "I", "1")

		// second part
		var obj = split[1]
		obj = strings.Trim(obj, " ")
		obj = strings.Trim(obj, "\t")
		obj = strings.Trim(obj, "\"")
		obj = strings.ReplaceAll(obj, "\\\"", "\u0022")

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

// convert22() Kbd-En-Jonty.json
func convert22() {
	dictObj := wordObject.NewDictionaryObject("Jonty Kabardian-English Dictionary", 22, "Kbd", "En")
	srcFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\srcDicts\\Kbd-En-Jonty.json"
	distFilePath := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts\\Kbd-En-Jonty.json"
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

		// second part
		var obj = split[1]
		obj = strings.TrimSuffix(obj, ",")
		var objMap JontyWord
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
