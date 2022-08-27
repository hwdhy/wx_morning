package models

type Words struct {
	Data WordsData `json:"data"`
}

type WordsData struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
