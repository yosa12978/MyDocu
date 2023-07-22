package repos

import (
	"github.com/yosa12978/MyDocu/internal/db"
	"github.com/yosa12978/MyDocu/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUser(id uint) (models.User, error)
	GetUsers() []models.User
	Search(q string) []models.User
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id uint) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	repo := new(userRepo)
	repo.db = db.GetDB()
	return repo
}

func (repo *userRepo) GetUser(id uint) (models.User, error) {
	var user models.User
	err := repo.db.First(&user, "id = ?", id).Error
	return user, err
}

func (repo *userRepo) GetUsers() []models.User {
	var users []models.User
	repo.db.Find(&users)
	return users
}

func (repo *userRepo) Search(q string) []models.User {
	var users []models.User
	repo.db.Where("username LIKE %?%", q).Find(&users)
	return users
}

func (repo *userRepo) CreateUser(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

func (repo *userRepo) DeleteUser(id uint) (models.User, error) {
	user, err := repo.GetUser(id)
	if err != nil {
		return user, err
	}
	repo.db.Delete(user)
	return user, err
}

func (repo *userRepo) UpdateUser(user models.User) (models.User, error) {
	_, err := repo.GetUser(user.ID)
	if err != nil {
		return user, err
	}
	err = repo.db.Save(&user).Error
	return user, err
}
