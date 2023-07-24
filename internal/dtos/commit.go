package dtos

import "time"

type CreateCommitDTO struct {
	Desc    string `json:"desc"`
	Content string `json:"Content"`
}

type CommitDTO struct {
	ID         uint      `json:"id"`
	Desc       string    `json:"desc"`
	Content    string    `json:"Content"`
	Created    time.Time `json:"created"`
	UserId     uint      `json:"UserId"`
	DocumentId uint      `json:"documentId"`
}
