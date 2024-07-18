package convertHtmlToDb

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "modernc.org/sqlite"
	"strconv"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
	"ultimate-circassian-dictionary-helper/wordObject"
)

type DefinitionsInDifferentDictionaries struct {
	Title    string `json:"title"`
	Html     string `json:"html"`
	FromLang string `json:"fromLang"`
	ToLang   string `json:"toLang"`
	Spelling string `json:"spelling"`
}

type DictCounter struct {
	Count    int    `json:"count"`
	Title    string `json:"title"`
	FromLang string `json:"fromLang"`
	ToLang   string `json:"toLang"`
}

type WordToDiddMap map[string][]*DefinitionsInDifferentDictionaries

func (wtdm WordToDiddMap) ToSqliteDb(chunkSize int) error {
	fileName := "wordToDefsInDiffDicts.db"
	var db *sql.DB
	var err error

	// Check if the file exists and delete it if it does
	if utils.DoesFileExist(fileName) {
		err = utils.DeleteFile(fileName)
		if err != nil {
			return err
		}
	}

	// Open the SQLite database
	db, err = sql.Open("sqlite", fileName)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
			return
		}
	}(db)

	// Adjust the CREATE TABLE statement to set a limit based on the longest spelling
	maxLength := determineMaxLength(wtdm) // Call determineMaxLength to get the longest word length

	// Create the table with the original column names and additional boolean columns
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS keyvalue (
			key VARCHAR(`+strconv.Itoa(maxLength)+`) PRIMARY KEY,  -- Set VARCHAR length based on maxLength
			value TEXT,
			isFromKbd BOOLEAN,
			isToKbd BOOLEAN,
			isFromEn BOOLEAN,
			isToEn BOOLEAN,
			isFromAdy BOOLEAN,
			isToAdy BOOLEAN,
			isFromAr BOOLEAN,
			isToAr BOOLEAN,
			isFromTu BOOLEAN,
			isToTu BOOLEAN,
			isFromRu BOOLEAN,
			isToRu BOOLEAN,
			isFromHe BOOLEAN,
			isToHe BOOLEAN
		);`,
	)
	if err != nil {
		return err
	}

	// Create index on key column for faster lookup
	_, err = db.ExecContext(
		context.Background(),
		`CREATE INDEX IF NOT EXISTS idx_keyvalue_key ON keyvalue (key);`,
	)
	if err != nil {
		return err
	}

	// Prepare the insert statement
	stmt, err := db.PrepareContext(
		context.Background(),
		`INSERT INTO keyvalue (
			key, value, 
			isFromKbd, isToKbd, 
			isFromEn, isToEn, 
			isFromAdy, isToAdy, 
			isFromAr, isToAr, 
			isFromTu, isToTu, 
			isFromRu, isToRu, 
			isFromHe, isToHe
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error closing prepared statement:", err)
			return
		}
	}(stmt)

	// Begin a transaction for batch insertion
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			fmt.Println("Error rolling back transaction:", err)
			return
		}
	}(tx) // Rollback the transaction if any error occurs

	idx := 0
	total := len(wtdm)

	// Insert data in batches of size chunkSize
	for word, diddList := range wtdm {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err = enc.Encode(diddList)
		if err != nil {
			return err
		}
		valueStr := buf.String()
		valueStr = strings.Trim(valueStr, "\n")

		// Initialize the language flags
		isFromKbd, isToKbd := false, false
		isFromEn, isToEn := false, false
		isFromAdy, isToAdy := false, false
		isFromAr, isToAr := false, false
		isFromTu, isToTu := false, false
		isFromRu, isToRu := false, false
		isFromHe, isToHe := false, false

		// Check each dictionary entry for the languages
		for _, didd := range diddList {
			switch didd.FromLang {
			case "Kbd":
				isFromKbd = true
			case "En":
				isFromEn = true
			case "Ady":
				isFromAdy = true
			case "Ar":
				isFromAr = true
			case "Tu":
				isFromTu = true
			case "Ru":
				isFromRu = true
			case "He":
				isFromHe = true
			case "Ady/Kbd", "Kbd/Ady":
				isFromAdy = true
				isFromKbd = true
			}
			switch didd.ToLang {
			case "Kbd":
				isToKbd = true
			case "En":
				isToEn = true
			case "Ady":
				isToAdy = true
			case "Ar":
				isToAr = true
			case "Tu":
				isToTu = true
			case "Ru":
				isToRu = true
			case "He":
				isToHe = true
			case "Ady/Kbd", "Kbd/Ady":
				isToAdy = true
				isToKbd = true
			}
		}

		// Execute the prepared statement within the transaction
		_, err = tx.Stmt(stmt).Exec(
			word, valueStr,
			isFromKbd, isToKbd,
			isFromEn, isToEn,
			isFromAdy, isToAdy,
			isFromAr, isToAr,
			isFromTu, isToTu,
			isFromRu, isToRu,
			isFromHe, isToHe,
		)
		if err != nil {
			return err
		}

		idx++
		if idx%chunkSize == 0 || idx == total {
			// Commit the current transaction and start a new one for the next batch
			err = tx.Commit()
			if err != nil {
				return err
			}

			// Start a new transaction for the next batch if there are more records to insert
			if idx < total {
				tx, err = db.Begin()
				if err != nil {
					return err
				}
			}

			fmt.Printf("Inserted %d/%d records. Continuing...\n", idx, total)
		}
	}

	fmt.Println("Database and table created successfully!")
	return nil
}

func CallAllConverts() {
	wordToDefList := make(WordToDiddMap)
	// take all HTML files from the following directory
	distDictsDir := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDictsInHTML"
	filePathList := utils.GetAllFilesInDirectory(distDictsDir)
	dictWordCountMap := make(map[string]DictCounter)
	for _, filePath := range filePathList {
		fileToStr := utils.ReadEntireFile(filePath)
		fileStrToBytes := []byte(fileToStr)
		var dictObject wordObject.DictionaryObject
		if err := json.Unmarshal(fileStrToBytes, &dictObject); err != nil {
			fmt.Printf("Invalid json file: %s\n. Reason: %s", filePath, err.Error())
			return
		}

		for word, wordObj := range dictObject.Words {
			safeWord := regularWordToSafeWord(word)
			safeWord = strings.ToLower(safeWord)
			if _, ok := wordToDefList[safeWord]; !ok {
				wordToDefList[safeWord] = make([]*DefinitionsInDifferentDictionaries, 0)
			}
			wordToDefList[safeWord] = append(wordToDefList[safeWord], &DefinitionsInDifferentDictionaries{
				Title:    dictObject.Title,
				Html:     wordObj.FullDefinitionInHtml,
				FromLang: dictObject.FromLang,
				ToLang:   dictObject.ToLang,
				Spelling: safeWord,
			})

			if _, ok := dictWordCountMap[dictObject.Title]; !ok {
				dictWordCountMap[dictObject.Title] = DictCounter{
					Count:    0,
					FromLang: dictObject.FromLang,
					ToLang:   dictObject.ToLang,
					Title:    dictObject.Title,
				}
			}
			dictWordCountMap[dictObject.Title] = DictCounter{
				Count:    dictWordCountMap[dictObject.Title].Count + 1,
				FromLang: dictObject.FromLang,
				ToLang:   dictObject.ToLang,
				Title:    dictObject.Title,
			}
		}
	}

	dictWordCountMapToArray := make([]DictCounter, 0)
	for _, dictCounter := range dictWordCountMap {
		dictWordCountMapToArray = append(dictWordCountMapToArray, dictCounter)
	}
	dictWordCountMapJson, err := json.Marshal(dictWordCountMapToArray)
	if err != nil {
		panic(err)
	}

	utils.CreateFileWithString("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\dictTitles.txt", string(dictWordCountMapJson))

	err = wordToDefList.ToSqliteDb(10000)
	if err != nil {
		panic(err)
	}
}

func regularWordToSafeWord(word string) string {
	word = strings.ReplaceAll(word, " ", "_")
	word = strings.ReplaceAll(word, "/", "&")
	return word
}

// Helper function to determine the longest spelling for VARCHAR size
func determineMaxLength(wtdm WordToDiddMap) int {
	maxLength := 0
	for word := range wtdm {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}
	return maxLength
}
