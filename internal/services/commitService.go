package services

import (
	"github.com/yosa12978/MyDocu/internal/dtos"
)

type CommitService interface {
	AddCommit(commit dtos.CreateCommitDTO)
	DeleteCommit(id uint)
	GetCommits(docId uint)
	GetCommit(id uint)
}
