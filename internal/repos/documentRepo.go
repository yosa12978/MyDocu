package repos

import (
	"github.com/yosa12978/MyDocu/internal/db"
	"github.com/yosa12978/MyDocu/internal/models"
	"gorm.io/gorm"
)

type DocumentRepo interface {
	GetDocuments() []models.Document
	GetDocument(id uint) (models.Document, error)
	SearchDocuments(q string) []models.Document
	CreateDocument(doc models.Document) (models.Document, error)
	UpdateDocument(doc models.Document) (models.Document, error)
	DeleteDocument(id uint, doc models.Document) (models.Document, error)
}

type documentRepo struct {
	db *gorm.DB
}

func NewDocumentRepo() DocumentRepo {
	repo := new(documentRepo)
	repo.db = db.GetDB()
	return repo
}

func (repo *documentRepo) GetDocuments() []models.Document {
	var docs []models.Document
	repo.db.Find(&docs)
	return docs
}

func (repo *documentRepo) GetDocument(id uint) (models.Document, error) {
	var doc models.Document
	err := repo.db.First(&doc, "id = ?", id).Error
	return doc, err
}

func (repo *documentRepo) SearchDocuments(q string) []models.Document {
	var docs []models.Document
	repo.db.Model(&models.Document{}).Where("title LIKE %?%", q).Find(&docs)
	return docs
}

func (repo *documentRepo) CreateDocument(doc models.Document) (models.Document, error) {
	err := repo.db.Create(&doc).Error
	return doc, err
}

func (repo *documentRepo) UpdateDocument(doc models.Document) (models.Document, error) {
	docu, err := repo.GetDocument(doc.ID)
	if err != nil {
		return docu, err
	}
	err = repo.db.Save(&doc).Error
	return doc, err
}

func (repo *documentRepo) DeleteDocument(id uint, doc models.Document) (models.Document, error) {
	docu, err := repo.GetDocument(doc.ID)
	if err != nil {
		return docu, err
	}
	err = repo.db.Save(&doc).Error
	return doc, err
}
