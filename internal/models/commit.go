package models

type Commit struct {
	BaseModel
	Desc       string `json:"desc"`
	Content    string `json:"Content"`
	UserId     uint   `json:"UserId"`
	DocumentId uint   `json:"documentId"`
}
