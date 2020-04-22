package repository

import (
	"fmt"
	"go-base-cleancode/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetBuildList(db *gorm.DB, limit int, page int, filter string) (building []models.Building, err error) {

	query := `SELECT audit.audit_id, audit.audit_build_name , owner.owner_name, owner.owner_phone,building.build_kabupaten, building.build_total_room, building.build_audit
	FROM audit audit
	join owner owner on audit.audit_build_id = owner.owner_building_id
	join building building on audit.audit_build_id = building.build_id `

	//search filter by name or building name or district
	if filter != "" {
		query += "where audit.audit_build_name = '" + filter + "' or owner.owner_name = '" + filter + "' or building.build_kabupaten = '" + filter + "' "
	}

	//default limit, offset
	offset := 0
	limit = 10
	if limit > 0 {
		limit = limit
	}
	if page > 0 {
		offset = (page - 1) * limit
	}

	fmt.Println(query)
	rows, err := db.Raw(query+"order by audit.audit_id asc LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		return building, err
	}
	defer rows.Close()
	for rows.Next() {
		var temp models.Building
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return building, err
		}
		building = append(building, temp)

	}

	return building, err
}

func GetCountBuildList(db *gorm.DB) (buildingTotalDetail []models.BuildingTotalDetail, err error) {

	rows, err := db.Raw(`SELECT audit.audit_id, building.build_audit
	FROM audit audit
	join building building on audit.audit_build_id = building.build_id`).Rows()
	if err != nil {
		return buildingTotalDetail, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.BuildingTotalDetail
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return buildingTotalDetail, err
		}

		buildingTotalDetail = append(buildingTotalDetail, temp)

	}
	return buildingTotalDetail, err
}

func GetBuildingDetail(db *gorm.DB, auditID int) (buildingDetail models.BuildingDetail, err error) {

	rows, err := db.Raw(`
	SELECT audit.audit_auditor_name,audit.audit_date, audit.audit_id,audit.audit_build_id,audit.audit_build_name ,owner.owner_name, owner.owner_phone ,
	CONCAT(building.build_address,' RT ',building.build_address_rt,' RW ',building.build_address_rw, ' Kel. ',building.build_kelurahan,' Kec. ',building.build_kecamatan,' ',building.build_kabupaten) as address,building.build_total_room, building.build_audit, bs.build_staff_name as nama_pengurus,  
	bs.build_staff_phone as kontak_pengurus
	FROM audit audit 
	join owner owner on audit.audit_build_id = owner.owner_building_id 
	join building building on audit.audit_build_id = building.build_id
	join build_staff bs on building.build_id = bs.build_staff_build_id
	where audit.audit_id = ?`, auditID).Rows()
	if err != nil {
		return buildingDetail, err
	}

	defer rows.Close()
	for rows.Next() {
		err = db.ScanRows(rows, &buildingDetail)
		if err != nil {
			return buildingDetail, err
		}
	}
	return buildingDetail, err
}

func GetRoomType(db *gorm.DB, buildID int) (room []models.Room, err error) {

	rows, err := db.Raw(`SELECT roomtype_build_id,roomtype_name,roomtype_id 
	FROM operation.room_type where roomtype_build_id = ?`, buildID).Rows()
	if err != nil {
		return room, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Room
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return room, err
		}

		room = append(room, temp)

	}
	return room, err
}

func GetDocument(db *gorm.DB, buildID int) (document []models.Document, err error) {

	rows, err := db.Raw(`SELECT doc_id, doc_original_name,doc_file_name,doc_path 
	FROM operation.document where doc_build_id = ?`, buildID).Rows()
	if err != nil {
		return document, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.Document
		err = db.ScanRows(rows, &temp)
		if err != nil {
			return document, err
		}

		document = append(document, temp)

	}
	return document, err
}

func UpdateAudit(db *gorm.DB, auditID int, name string, date string) (err error) {

	//convert string into time
	layoutDate := "2006-01-02T15:04:05Z"
	updatedDate, err := time.Parse(layoutDate, date)
	if err != nil {
		return err
	}

	_, err = db.Raw(`update audit
	set audit_auditor_name = ? , audit_update_date = ?
	where audit_id = ?`, name, updatedDate, auditID).Rows()
	if err != nil {
		return err
	}

	return err
}

func UpdateRoomType(db *gorm.DB, roomtypeID int, name string, buildID int) (err error) {

	_, err = db.Raw(`update room_type
	set roomtype_name = ? , roomtype_update_date = now()
	where roomtype_id = ? and roomtype_build_id = ?`, name, roomtypeID, buildID).Rows()
	if err != nil {
		return err
	}
	return err
}
