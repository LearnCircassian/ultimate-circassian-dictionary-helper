package wordObject

type Cognate struct {
	Language string `json:"language"`
	Spelling string `json:"spelling"`
}

type Synonym struct {
	Spelling string `json:"spelling"`
	Meaning  string `json:"meaning"`
}

type Antonym struct {
	Spelling string `json:"spelling"`
	Meaning  string `json:"meaning"`
}

type Example struct {
	Sentence    string `json:"sentence"`
	Translation string `json:"translation"`
}

type AllDefinitions struct {
	Meaning             string    `json:"meaning"`
	MeaningWithExamples string    `json:"fullMeaning"`
	Examples            []Example `json:"examples"`
}

type WordObject struct {
	Spelling       string           `json:"spelling"`
	Cognates       []Cognate        `json:"cognates"`
	Redirect       string           `json:"redirect"`
	AllDefinitions []AllDefinitions `json:"allDefinitions"`
	Derivation     string           `json:"derivation"`
	Type           string           `json:"type"`
}

type DictionaryObject struct {
	Title    string                 `json:"title"`
	Id       int                    `json:"id"`
	Words    map[string]*WordObject `json:"words"`
	FromLang string                 `json:"fromLang"`
	ToLang   string                 `json:"toLang"`
}

func NewDictionaryObject(title string, id int, fromLang string, toLang string) *DictionaryObject {
	return &DictionaryObject{
		Title:    title,
		Id:       id,
		Words:    make(map[string]*WordObject),
		FromLang: fromLang,
		ToLang:   toLang,
	}
}

func NewWordObject(spelling string, wordType string) *WordObject {
	return &WordObject{
		Spelling:       spelling,
		Cognates:       make([]Cognate, 0),
		Redirect:       "",
		AllDefinitions: make([]AllDefinitions, 0),
		Derivation:     "",
		Type:           wordType,
	}
}

func (w *WordObject) AddCognate(language, spelling string) {
	w.Cognates = append(w.Cognates, Cognate{
		Spelling: spelling,
		Language: language,
	})
}

func (w *WordObject) AddOneSimpleDefinition(meaningWithExamples string) {
	w.AllDefinitions = append(w.AllDefinitions, AllDefinitions{
		Meaning:             "",
		MeaningWithExamples: meaningWithExamples,
		Examples:            make([]Example, 0),
	})
}

func (w *WordObject) AddOneAdvancedDefinition(meaning string, examples []Example) {
	w.AllDefinitions = append(w.AllDefinitions, AllDefinitions{
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

func (ad AllDefinitions) AddExample(sentence, translation string) {
	ad.Examples = append(ad.Examples, Example{
		Sentence:    sentence,
		Translation: translation,
	})
}
