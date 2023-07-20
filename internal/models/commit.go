package models

type Commit struct {
	BaseModel
	Name       string `json:"name"`
	Body       string `json:"body"`
	UserId     uint   `json:"UserId"`
	DocumentId uint   `json:"documentId"`
}
