package convertHtmlToDb

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "modernc.org/sqlite"
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

func (wtdm WordToDiddMap) ToSqliteDb(n int) error {
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
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS keyvalue (key TEXT PRIMARY KEY, value TEXT);`,
	)
	if err != nil {
		return err
	}

	// Begin a transaction for batch insertion
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback the transaction if any error occurs

	idx := 0
	total := len(wtdm)

	// Insert data in batches of size n
	for word, diddList := range wtdm {
		buf := &bytes.Buffer{} // or &strings.Builder{} as from the example of @mkopriva
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err = enc.Encode(diddList)
		if err != nil {
			return err
		}
		valueStr := buf.String()
		valueStr = strings.Trim(valueStr, "\n")

		// Execute the insertion statement within the transaction
		_, err = tx.Exec(`INSERT INTO keyvalue (key, value) VALUES (?, ?)`, word, valueStr)
		if err != nil {
			return err
		}

		idx++
		if idx%n == 0 || idx == total {
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
