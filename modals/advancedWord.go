package modals

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
	Meaning  string    `json:"meaning"`
	Examples []Example `json:"examples"`
	Type     string    `json:"type"`
}

type AdvancedWord struct {
	Spelling string    `json:"spelling"`
	Cognates []Cognate `json:"cognates"`
	Redirect string    `json:"redirect"`
}
