package utils

import (
	"bufio"
	"log"
	"os"
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
