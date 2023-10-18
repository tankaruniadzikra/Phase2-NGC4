package entity

type Character struct {
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"image_url"`
	Type     string `json:"type"`
}
