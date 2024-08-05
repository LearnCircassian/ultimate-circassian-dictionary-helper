package main

import (
	"ultimate-circassian-dictionary-helper/convertHtmlToDb"
	"ultimate-circassian-dictionary-helper/convertJsonToHtml"
	"ultimate-circassian-dictionary-helper/convertOriginalFilesToJson"
)

func main() {
	convertOriginalFilesToJson.CallAllConverts()
	convertJsonToHtml.CallAllConverts()
	convertHtmlToDb.CallAllConverts()
}
