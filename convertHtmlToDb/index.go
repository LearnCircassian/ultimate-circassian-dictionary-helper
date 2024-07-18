package convertHtmlToDb

import (
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

var db *sql.DB // Global variable for SQLite database connection

// InitializeDB initializes the SQLite database connection
func InitializeDB(fileName string) error {
	var err error
	db, err = sql.Open("sqlite", fileName)
	if err != nil {
		return err
	}
	return nil
}

// CloseDB closes the SQLite database connection
func CloseDB() error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// ToSqliteDb inserts data into SQLite database tables
func (wtdm WordToDiddMap) ToSqliteDb(tableName string, chunkSize int) error {
	// Ensure db is initialized
	if db == nil {
		return fmt.Errorf("database connection is nil, call InitializeDB with the database filename first")
	}

	// Adjust the CREATE TABLE statement to set a limit based on the longest spelling
	// Call determineMaxLength to get the longest word length
	maxLength, longestWord := determineMaxLength(wtdm)
	fmt.Printf("Longest word: %s (%d characters)\n", longestWord, maxLength)

	// Create the table with the original column names and additional boolean columns
	_, err := db.ExecContext(
		context.Background(),
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
            key VARCHAR(%d) PRIMARY KEY,  -- Set VARCHAR length based on maxLength
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
        );`, tableName, maxLength),
	)
	if err != nil {
		return err
	}

	// Create index on key column for faster lookup
	_, err = db.ExecContext(
		context.Background(),
		fmt.Sprintf(`CREATE INDEX IF NOT EXISTS idx_%s_key ON %s (key);`, tableName, tableName),
	)
	if err != nil {
		return err
	}

	// Prepare the insert statement
	stmt, err := db.PrepareContext(
		context.Background(),
		fmt.Sprintf(`INSERT INTO %s (
            key, value, 
            isFromKbd, isToKbd, 
            isFromEn, isToEn, 
            isFromAdy, isToAdy, 
            isFromAr, isToAr, 
            isFromTu, isToTu, 
            isFromRu, isToRu, 
            isFromHe, isToHe
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`, tableName),
	)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Error closing prepared statement:", err)
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
		}
	}(tx) // Rollback the transaction if any error occurs

	idx := 0
	total := len(wtdm)

	// Insert data in batches of size chunkSize
	for word, diddList := range wtdm {
		// Encode diddList to JSON string
		valueStr, err := json.Marshal(diddList)
		if err != nil {
			return err
		}

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
			word, string(valueStr),
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

	fmt.Println("Data insertion completed successfully!")
	return nil
}

// CallAllConverts processes HTML files and stores data into SQLite database
func CallAllConverts() {
	latinWordToDefList := make(WordToDiddMap)
	cyrllicWordToDefList := make(WordToDiddMap)
	dictWordCountMap := make(map[string]*DictCounter) // Change to store pointers

	// Take all HTML files from the following directory
	distDictsDir := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDictsInHTML"
	filePathList := utils.GetAllFilesInDirectory(distDictsDir)

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

			// Determine if the word is Latin or Cyrillic
			isLatin := utils.IsLatin(safeWord)

			if isLatin {
				if _, ok := latinWordToDefList[safeWord]; !ok {
					latinWordToDefList[safeWord] = make([]*DefinitionsInDifferentDictionaries, 0)
				}
				latinWordToDefList[safeWord] = append(latinWordToDefList[safeWord], &DefinitionsInDifferentDictionaries{
					Title:    dictObject.Title,
					Html:     wordObj.FullDefinitionInHtml,
					FromLang: dictObject.FromLang,
					ToLang:   dictObject.ToLang,
					Spelling: safeWord,
				})
			} else {
				if _, ok := cyrllicWordToDefList[safeWord]; !ok {
					cyrllicWordToDefList[safeWord] = make([]*DefinitionsInDifferentDictionaries, 0)
				}
				cyrllicWordToDefList[safeWord] = append(cyrllicWordToDefList[safeWord], &DefinitionsInDifferentDictionaries{
					Title:    dictObject.Title,
					Html:     wordObj.FullDefinitionInHtml,
					FromLang: dictObject.FromLang,
					ToLang:   dictObject.ToLang,
					Spelling: safeWord,
				})
			}

			// Update dictionary word count map using pointers
			if _, ok := dictWordCountMap[dictObject.Title]; !ok {
				dictWordCountMap[dictObject.Title] = &DictCounter{
					Count:    0,
					FromLang: dictObject.FromLang,
					ToLang:   dictObject.ToLang,
					Title:    dictObject.Title,
				}
			}
			dictWordCountMap[dictObject.Title].Count++
		}
	}

	// Convert dictionary word count map to array and save to file
	dictWordCountMapToArray := make([]DictCounter, 0, len(dictWordCountMap))
	for _, dictCounter := range dictWordCountMap {
		dictWordCountMapToArray = append(dictWordCountMapToArray, *dictCounter) // Dereference pointer
	}
	dictWordCountMapJson, err := json.Marshal(dictWordCountMapToArray)
	if err != nil {
		panic(err)
	}
	utils.CreateFileWithString("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\dictTitles.txt", string(dictWordCountMapJson))

	// Insert Latin and Cyrillic words into separate tables
	err = latinWordToDefList.ToSqliteDb("latinWordsTable", 10000)
	if err != nil {
		panic(err)
	}

	err = cyrllicWordToDefList.ToSqliteDb("cyrillicWordsTable", 10000)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data processing completed successfully!")
}

func regularWordToSafeWord(word string) string {
	word = strings.ReplaceAll(word, " ", "_")
	word = strings.ReplaceAll(word, "/", "&")
	return word
}

// Helper function to determine the longest spelling for VARCHAR size
func determineMaxLength(wtdm WordToDiddMap) (maxLength int, longestWord string) {
	maxLength = 0
	longestWord = ""

	for word := range wtdm {
		if len(word) > maxLength {
			maxLength = len(word)
			longestWord = word
		}
	}
	return maxLength, longestWord
}
