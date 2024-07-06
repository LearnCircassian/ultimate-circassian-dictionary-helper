package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
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

	dictToJsonBytes, err := json.MarshalIndent(dict, "", "\t")
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
		return
	}

	dictToJsonStr := string(dictToJsonBytes)
	dictToJsonStr = strings.ReplaceAll(dictToJsonStr, "\\u003c", "<")
	dictToJsonStr = strings.ReplaceAll(dictToJsonStr, "\\u003e", ">")
	dictToJsonStr = strings.ReplaceAll(dictToJsonStr, "\\\"", "'")
	_, err = f.WriteString(dictToJsonStr)
	if err != nil {
		log.Fatalf("write file error: %v", err)
		return
	}

	fmt.Printf("\nFile created: %s\n", filePath)
}
