package models

type User struct {
	BaseModel
	Username     string   `json:"username"`
	PasswordHash string   `json:"-"`
	Email        string   `json:",omitempty"`
	Role         string   `json:"role"`
	Commits      []Commit `json:",omitempty" gorm:"foreignKey:UserId"`
}
