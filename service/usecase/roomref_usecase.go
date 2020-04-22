package usecase

import (
	"go-base-cleancode/models"
	"go-base-cleancode/service/repository"

	"github.com/jinzhu/gorm"
)

func GetRoomrefUsecase(db *gorm.DB) (data []models.Roomref, err error) {

	//get list
	data, err = repository.GetRoomref(db)
	if err != nil {
		return data, err
	}

	return data, err

}
