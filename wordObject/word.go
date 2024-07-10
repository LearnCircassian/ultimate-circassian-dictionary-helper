package wordObject

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
	Meaning             string    `json:"meaning"`
	MeaningWithExamples string    `json:"fullMeaning"`
	Examples            []Example `json:"examples"`
}

type WordObject struct {
	Spelling    string       `json:"spelling"`
	Cognates    []Cognate    `json:"cognates"`
	Redirect    string       `json:"redirect"`
	Definitions []Definition `json:"definitions"`
	Derivation  string       `json:"derivation"`
	Type        string       `json:"type"`
	Synonyms    []Synonym    `json:"synonyms"`
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
		Spelling:    spelling,
		Cognates:    make([]Cognate, 0),
		Redirect:    "",
		Definitions: make([]Definition, 0),
		Derivation:  "",
		Type:        wordType,
	}
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

func (w *WordObject) AddOneSimpleDefinition(meaningWithExamples string) {
	w.Definitions = append(w.Definitions, Definition{
		Meaning:             "",
		MeaningWithExamples: meaningWithExamples,
		Examples:            make([]Example, 0),
	})
}

func (w *WordObject) AddOneAdvancedDefinition(meaning string, examples []Example) {
	w.Definitions = append(w.Definitions, Definition{
		Meaning:             meaning,
		MeaningWithExamples: "",
		Examples:            examples,
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
