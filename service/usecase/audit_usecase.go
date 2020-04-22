package usecase

import (
	"go-base-cleancode/models"
	"go-base-cleancode/service/repository"
	"math"

	"github.com/jinzhu/gorm"
)

func GetBuildingUsecase(db *gorm.DB, limit int, page int, filter string) (data models.Data, err error) {
	// get total
	res, err := repository.GetCountBuildList(db)
	if err != nil {
		return data, err
	}

	//calculate audit status
	belumAudit := 0
	auditUlang := 0
	auditComplete := 0
	for _, item := range res {
		if item.BuildAudit == 0 {
			belumAudit += 1
		}
		if item.BuildAudit == 1 {
			auditUlang += 1
		}
		if item.BuildAudit == 2 {
			auditComplete += 1
		}
	}

	//get list
	building, err := repository.GetBuildList(db, limit, page, filter)
	if err != nil {
		return data, err
	}

	data.Listresult = building
	data.TotalData = len(res)
	data.BelumAudit = belumAudit
	data.AuditUlang = auditUlang
	data.AuditComplete = auditComplete
	data.Page = page
	data.TotalPage = int(math.Ceil(float64(len(res)) / float64(10)))

	return data, err

}

func GetBuildingDetailUsecase(db *gorm.DB, auditID int) (data models.BuildingDetail, err error) {

	//get list
	data, err = repository.GetBuildingDetail(db, auditID)
	if err != nil {
		return data, err
	}

	//get room
	rooms, err := repository.GetRoomType(db, data.AuditBuildID)
	if err != nil {
		return data, err
	}

	//get document
	docs, err := repository.GetDocument(db, data.AuditBuildID)
	if err != nil {
		return data, err
	}

	data.Room = rooms
	data.Document = docs
	data.RoomType = len(rooms)

	return data, err

}

func UpdateAuditUsecase(db *gorm.DB, auditID int, req models.AuditDetailRequest) (err error) {

	tx := db.Begin()

	//get list
	err = repository.UpdateAudit(tx, auditID, req.UserFirstname, req.AuditDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	// update room type
	for _, item := range req.RoomType {
		err := repository.UpdateRoomType(tx, item.RoomtypeID, item.RoomtypeName, item.RoomtypeBuildID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return err

}
