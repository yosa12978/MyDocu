package models

type Document struct {
	BaseModel
	CurrentCommit Commit   `json:"currentCommit"`
	Commits       []Commit `json:"commits,omitempty" gorm:"foreignKey:DocumentId"`
}
