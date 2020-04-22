package models

// import(
// 	"github.com/jinzhu/gorm"
// )

type Building struct {
	AuditID        int    `json:"audit_id" form:"audit_id"`
	AuditBuildName string `json:"audit_build_name" form:"audit_build_name"`
	OwnerName      string `json:"owner_name" form:"owner_name"`
	OwnerPhone     string `json:"owner_phone" form:"owner_phone"`
	BuildKabupaten string `json:"build_kabupaten" form:"build_kabupaten"`
	BuildTotalRoom int    `json:"build_total_room" form:"build_total_room"`
	BuildAudit     int    `json:"build_audit" form:"build_audit"`
}

type ResponseBuilding struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type BuildingTotalDetail struct {
	AuditID    int `json:"audit_id" form:"audit_id"`
	BuildAudit int `json:"build_audit" form:"build_audit"`
}

type Data struct {
	Listresult    []Building `json:"listresult"`
	Page          int        `json:"page"`
	TotalData     int        `json:"totaldata"`
	TotalPage     int        `json:"totalpage"`
	BelumAudit    int        `json:"belum_audit"`
	AuditUlang    int        `json:"audit_ulang"`
	AuditComplete int        `json:"audit_complete"`
}

type ResponseAuditUpdated struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    int    `json:"auditID"`
}

type ResponseBuildingDetail struct {
	Code    int            `json:"code"`
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    BuildingDetail `json:"data"`
}

type BuildingDetail struct {
	AuditID          int        `json:"audit_id" form:"audit_id"`
	AuditBuildID     int        `json:"audit_build_id" form:"audit_build_id"`
	AuditDate        string     `json:"audit_date" form:"audit_date"`
	AuditAuditorName string     `json:"audit_auditor_name" form:"audit_auditor_name"`
	AuditBuildName   string     `json:"audit_build_name" form:"audit_build_name"`
	OwnerName        string     `json:"owner_name" form:"owner_name"`
	OwnerPhone       string     `json:"owner_phone" form:"owner_phone"`
	BuildTotalRoom   int        `json:"build_total_room" form:"build_total_room"`
	BuildAudit       int        `json:"build_audit" form:"build_audit"`
	BuildStaffName   string     `json:"build_staff_name" form:"build_staff_name"`
	BuildStaffPhone  string     `json:"build_staff_phone" form:"build_staff_phone"`
	Address          string     `json:"address" form:"address"`
	Document         []Document `json:"document"`
	Room             []Room     `json:"room"`
	RoomType         int        `json:"room_type"`
}

type Document struct {
	DocID           int    `json:"doc_id" form:"doc_id"`
	DocOriginalName string `json:"doc_original_name" form:"doc_original_name"`
	DocFileName     string `json:"doc_file_name" form:"doc_file_name"`
	DocPath         string `json:"doc_path" form:"doc_path"`
}

type Room struct {
	RoomtypeName    string `json:"roomtype_name" form:"roomtype_name"`
	RoomtypeID      int    `json:"roomtype_id" form:"roomtype_id"`
	RoomtypeBuildID int    `json:"roomtype_build_id" form:"roomtype_build_id"`
}

type AuditDetailRequest struct {
	UserFirstname string `json:"user_firstname"`
	AuditDate     string `json:"audit_date"`
	RoomType      []Room `json:"room_type"`
}
