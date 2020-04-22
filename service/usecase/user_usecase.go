package usecase

import (
	"go-base-cleancode/models"
	"go-base-cleancode/service/repository"

	"github.com/jinzhu/gorm"
)

func GetUserUsecase(db *gorm.DB) (data []models.User, err error) {

	//get list
	data, err = repository.GetUser(db)
	if err != nil {
		return data, err
	}

	return data, err

}
