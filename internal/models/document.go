package models

type Document struct {
	BaseModel
	Title         string   `json:"title"`
	CurrentCommit Commit   `json:"currentCommit"`
	Commits       []Commit `json:"commits,omitempty" gorm:"foreignKey:DocumentId"`
	Protected     bool     `json:"protected"`
}
