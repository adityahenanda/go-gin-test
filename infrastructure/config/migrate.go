package config

import (
	. "go-base-cleancode/models"

	"github.com/jinzhu/gorm"
)

func Migrate(idb *gorm.DB) {

	idb.Debug().AutoMigrate(
		&Building{},
	)

}
