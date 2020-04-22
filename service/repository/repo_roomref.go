package repository

import (
	"go-base-cleancode/models"

	"github.com/jinzhu/gorm"
)

func GetRoomref(db *gorm.DB) (roomref []models.Roomref, err error) {

	rows, err := db.Raw(`SELECT roomtype_ref_id,roomtype_ref_name,roomtype_ref_bed FROM roomtype_ref`).Rows()
	if err != nil {
		return roomref, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Roomref
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return roomref, err
		}

		roomref = append(roomref, temp)

	}
	return roomref, err
}
