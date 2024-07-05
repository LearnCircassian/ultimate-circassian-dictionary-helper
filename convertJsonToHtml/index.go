package convertJsonToHtml

import (
	"encoding/json"
	"fmt"
	"strings"
	"ultimate-circassian-dictionary-helper/utils"
	"ultimate-circassian-dictionary-helper/wordObject"
)

func CallAllConverts() {
	// take all JSON files from the following directory
	distDictsDir := "D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDicts"
	filePathList := utils.GetAllFilesInDirectory(distDictsDir)
	for _, filePath := range filePathList {
		fileToStr := utils.ReadEntireFile(filePath)
		fileStrToBytes := []byte(fileToStr)
		fileName := strings.Split(filePath, "\\")[len(strings.Split(filePath, "\\"))-1]
		var dictObject wordObject.DictionaryObject
		if err := json.Unmarshal(fileStrToBytes, &dictObject); err != nil {
			fmt.Printf("Invalid json file: %s\n. Reason: %s", filePath, err.Error())
			return
		}

		if dictObject.Format == "JSON" {
			dictObject.ConvertJsonToHtml()
		}

		targetFilePath := fmt.Sprintf("D:\\Github\\Cir\\ultimate-circassian-dictionary-helper\\distDictsInHTML\\%s", fileName)
		utils.CreateFileWithDictionaryObject(targetFilePath, &dictObject)
	}
}
