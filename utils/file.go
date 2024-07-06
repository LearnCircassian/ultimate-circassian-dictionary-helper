package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"ultimate-circassian-dictionary-helper/wordObject"
)

func ReadFileLineByLine(filePath string, callback func(string, int)) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file error: %v", err)
		}
	}(f)

	sc := bufio.NewScanner(f)
	idx := 0
	for sc.Scan() {
		line := sc.Text() // GET the line string
		callback(line, idx)
		// log.Printf("line %d: %s", idx, line)
		idx++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}

func CreateFileWithContent(filePath string, dict *wordObject.DictionaryObject) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("create file error: %v", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file error: %v", err)
		}
	}(f)

	dictToJson, err := json.MarshalIndent(dict, "", " ")
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
		return
	}

	_, err = f.Write(dictToJson)
	if err != nil {
		log.Fatalf("write file error: %v", err)
		return
	}

	fmt.Printf("\nFile created: %s\n", filePath)
}
