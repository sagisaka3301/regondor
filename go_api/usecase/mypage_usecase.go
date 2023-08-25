package usecase

import (
	"go_api/model"
	"go_api/repository"
)

type IMypageUsecase interface {
	GetUser(userId uint) (model.MypageResponse, error)
}

type mypageUsecase struct {
	mr repository.IMypageRepository
	// mv validator.IMypageRepository
}

func NewMypageUsecase(mr repository.IMypageRepository) IMypageUsecase {
	return &mypageUsecase{mr}
}

func (mu *mypageUsecase) GetUser(userId uint) (model.MypageResponse, error) {
	user := model.User{}
	if err := mu.mr.GetUser(&user, userId); err != nil {
		return model.MypageResponse{}, err
	}
	resUser := model.MypageResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt,
	}
	return resUser, nil
}
