package dtos

type CreateDocumentDTO struct {
	Title     string `json:"title"`
	Content   string `json:"Content"`
	Protected bool   `json:"protected,omitempty"`
}

type DocumentDTO struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	CurrentCommit CommitDTO `json:"currentCommit"`
	Protected     bool      `json:"protected"`
}
