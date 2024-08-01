package wordObject

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

type Cognate struct {
	Language string `json:"language"`
	Spelling string `json:"spelling"`
}

type Synonym struct {
	Spelling    string `json:"spelling"`
	Explanation string `json:"explanation"`
}

type Antonym struct {
	Spelling string `json:"spelling"`
	Meaning  string `json:"meaning"`
}

type Example struct {
	Sentence    string `json:"sentence"`
	Translation string `json:"translation"`
}

type Definition struct {
	Meaning  string    `json:"meaning"`
	Examples []Example `json:"examples"`
}

type WordObject struct {
	Spelling             string       `json:"spelling"`
	Cognates             []Cognate    `json:"cognates"`
	Redirect             string       `json:"redirect"`
	FullDefinitionInHtml string       `json:"full_definition_in_html"`
	Definitions          []Definition `json:"definitions"`
	Derivation           string       `json:"derivation"`
	Type                 string       `json:"type"`
	Synonyms             []Synonym    `json:"synonyms"`
}

type DictionaryObject struct {
	Title    string                 `json:"title"`
	Id       int                    `json:"id"`
	Words    map[string]*WordObject `json:"words"`
	FromLang string                 `json:"fromLang"`
	ToLang   string                 `json:"toLang"`
	Format   string                 `json:"format"`
}

func NewDictionaryObject(title string, id int, fromLang string, toLang string, format string) *DictionaryObject {
	return &DictionaryObject{
		Title:    title,
		Id:       id,
		Words:    make(map[string]*WordObject),
		FromLang: fromLang,
		ToLang:   toLang,
		Format:   format,
	}
}

func NewWordObject(spelling string, wordType string) *WordObject {
	return &WordObject{
		Spelling:             spelling,
		Cognates:             make([]Cognate, 0),
		Redirect:             "",
		FullDefinitionInHtml: "",
		Definitions:          make([]Definition, 0),
		Derivation:           "",
		Type:                 wordType,
		Synonyms:             make([]Synonym, 0),
	}
}

func (w *WordObject) AddType(wordType string) {
	w.Type = wordType
}

func (w *WordObject) AddFullDefinitionInHtml(fullDefinitionInHtml string) {
	w.FullDefinitionInHtml = fullDefinitionInHtml
}

func (w *WordObject) AddCognate(language, spelling string) {
	w.Cognates = append(w.Cognates, Cognate{
		Spelling: spelling,
		Language: language,
	})
}

func (w *WordObject) AddSynonym(spelling, explanation string) {
	w.Synonyms = append(w.Synonyms, Synonym{
		Spelling:    spelling,
		Explanation: explanation,
	})
}

func (w *WordObject) AddOneDefinition(meaning string, examples []Example) {
	w.Definitions = append(w.Definitions, Definition{
		Meaning:  meaning,
		Examples: examples,
	})
}

func (w *WordObject) AddRedirect(redirect string) {
	w.Redirect = redirect
}

func (w *WordObject) AddDerivation(derivation string) {
	w.Derivation = derivation
}

func (ad Definition) AddExample(sentence, translation string) {
	ad.Examples = append(ad.Examples, Example{
		Sentence:    sentence,
		Translation: translation,
	})
}

func formatText(text string) string {
	re := regexp.MustCompile(`\|([^|]+)\|`)
	return re.ReplaceAllStringFunc(text, func(s string) string {
		innerText := s[1 : len(s)-1]
		return "<span style='font-weight:bold'>" + html.EscapeString(innerText) + "</span>"
	})
}

func formatMeaningText(meaning string) string {
	lines := strings.Split(meaning, "\n")
	var sb strings.Builder
	for _, line := range lines {
		tabs := strings.Count(line, "\t")
		line = strings.Trim(line, "\t")
		sb.WriteString(fmt.Sprintf("<div style='margin-left:%dem'>%s</div>", tabs, formatText(line)))
	}
	return sb.String()
}

func (w *WordObject) ToFullDefinitionHtml() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("<div><h2>%s</h2>", html.EscapeString(w.Spelling)))

	sb.WriteString(fmt.Sprintf("<p>Type: %s</p>", html.EscapeString(w.Type)))

	if len(w.Cognates) > 0 {
		sb.WriteString("<h3>Cognates:</h3>")
		for _, cognate := range w.Cognates {
			sb.WriteString(fmt.Sprintf("<div style='margin-left:1em'>%s: %s</div>", html.EscapeString(cognate.Language), html.EscapeString(cognate.Spelling)))
		}
	}

	if w.Redirect != "" {
		sb.WriteString(fmt.Sprintf("<p>Redirect: %s</p>", html.EscapeString(w.Redirect)))
	}

	sb.WriteString("<h3>Definitions:</h3>")
	for i, definition := range w.Definitions {
		sb.WriteString(fmt.Sprintf("<div style='margin-left:1em'><font color='darkblue'><span style='font-weight:bold'>%d.</span></font> %s</div>", i+1, formatMeaningText(definition.Meaning)))
		if len(definition.Examples) > 0 {
			for _, example := range definition.Examples {
				sb.WriteString(fmt.Sprintf("<div style='margin-left:3em'>%s — %s</div>", formatText(example.Sentence), formatText(example.Translation)))
			}
		}
	}

	if w.Derivation != "" {
		sb.WriteString(fmt.Sprintf("<p>Derivation: %s</p>", html.EscapeString(w.Derivation)))
	}

	if len(w.Synonyms) > 0 {
		sb.WriteString("<h3>Synonyms:</h3>")
		for _, synonym := range w.Synonyms {
			sb.WriteString(fmt.Sprintf("<div style='margin-left:1em'>%s: %s</div>", html.EscapeString(synonym.Spelling), html.EscapeString(synonym.Explanation)))
		}
	}

	sb.WriteString("</div>")

	return sb.String()
}

func (do *DictionaryObject) ConvertJsonToHtml() {
	for _, word := range do.Words {
		word.FullDefinitionInHtml = word.ToFullDefinitionHtml()
	}
}
