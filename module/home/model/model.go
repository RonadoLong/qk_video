package model

type HomeBanner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Type int    `json:"type"`
}
