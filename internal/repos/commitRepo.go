package repos

import (
	"github.com/yosa12978/MyDocu/internal/db"
	"github.com/yosa12978/MyDocu/internal/models"
	"gorm.io/gorm"
)

type CommitRepo interface {
	GetCommits(docId uint) []models.Commit
	GetCommit(id uint) (models.Commit, error)
	CreateCommit(commit models.Commit) (models.Commit, error)
	DeleteCommit(id uint) (models.Commit, error)
}

type commitRepo struct {
	db *gorm.DB
}

func NewCommitRepo() CommitRepo {
	repo := new(commitRepo)
	repo.db = db.GetDB()
	return repo
}

func (repo *commitRepo) GetCommits(docId uint) []models.Commit {
	var commits []models.Commit
	repo.db.Where("DocumentId = ?", docId).Find(&commits)
	return commits
}

func (repo *commitRepo) GetCommit(id uint) (models.Commit, error) {
	var commit models.Commit
	err := repo.db.First(&commit, "id = ?", id).Error
	return commit, err
}

func (repo *commitRepo) CreateCommit(commit models.Commit) (models.Commit, error) {
	err := repo.db.Create(&commit).Error
	return commit, err
}

func (repo *commitRepo) DeleteCommit(id uint) (models.Commit, error) {
	commit, err := repo.GetCommit(id)
	if err != nil {
		return commit, err
	}
	repo.db.Delete(commit)
	return commit, err
}
