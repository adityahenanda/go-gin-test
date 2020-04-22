package repository

import (
	"go-base-cleancode/models"

	"github.com/jinzhu/gorm"
)

func GetUser(db *gorm.DB) (users []models.User, err error) {

	rows, err := db.Raw(`SELECT user_id, user_firstname FROM user`).Rows()
	if err != nil {
		return users, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.User
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return users, err
		}

		users = append(users, temp)

	}
	return users, err
}
