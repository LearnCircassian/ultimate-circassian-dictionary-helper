package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
)

type JontyWord struct {
	Type        string `json:"type"`
	Definitions []struct {
		Meaning string `json:"meaning"`
	} `json:"definitions"`
}

func main() {
	pathFile := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\originDicts\\Kbd-En-Jonty.json"
	utils.ReadFileLineByLine(pathFile, func(line string, idx int) {
		line = strings.Trim(line, " ")
		line = strings.Trim(line, "\t")
		split := strings.SplitN(line, ":", 2)
		if len(split) < 2 {
			fmt.Printf("Invalid line: %s", line)
			return
		}

		var spelling = split[0]
		if strings.Count(spelling, "\"") != 2 {
			fmt.Printf("Invalid line: %s\n", line)
			log.Fatal("Invalid spelling")
		}
		spelling = strings.Trim(spelling, "\"")
		spelling = strings.Trim(spelling, " ")
		spelling = strings.Trim(spelling, "\t")

		var obj = split[1]
		obj = strings.TrimSuffix(obj, ",")
		var objMap JontyWord
		if err := json.Unmarshal([]byte(obj), &objMap); err != nil {
			fmt.Printf("Invalid line: %s\n", line)
			log.Fatal(err)
		}
	})
}
