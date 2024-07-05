package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"ultimate-circassian-dictionary-helper/wordObject"
)

func ReadEntireFile(filePath string) string {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return ""
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file error: %v", err)
		}
	}(f)

	// Create a scanner to read the file
	sc := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
	// Initialize a strings.Builder for efficient string concatenation
	var builder strings.Builder

	// Scan the file line by line
	for sc.Scan() {
		builder.WriteString(sc.Text())
		builder.WriteString("\n")
	}

	// Check for scanning errors
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return ""
	}

	return builder.String()
}

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

func CreateFileWithDictionaryObject(filePath string, dict *wordObject.DictionaryObject) {
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

func CreateFileWithString(filePath string, content string) {
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

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatalf("write file error: %v", err)
		return
	}

	fmt.Printf("\nFile created: %s\n", filePath)
}

func GetAllFilesInDirectory(dirPath string) []string {
	var filePathList []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("walk error: %v", err)
			return err
		}
		if !info.IsDir() {
			filePathList = append(filePathList, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("walk error: %v", err)
		return nil
	}
	return filePathList
}

func DoesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
