package repository

import (
	"go_api/model"

	"gorm.io/gorm"
)

type IMypageRepository interface {
	GetUser(user *model.User, userId uint) error
}

type mypageRepository struct {
	db *gorm.DB
}

func NewMypageRepository(db *gorm.DB) IMypageRepository {
	return &mypageRepository{db}
}

func (mr *mypageRepository) GetUser(user *model.User, userId uint) error {
	// userId に一致するユーザー情報を取得し、結果を user オブジェクトに格納
	if err := mr.db.Where("id = ?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}
