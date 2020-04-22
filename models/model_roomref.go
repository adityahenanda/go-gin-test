package models

type ResponseRoomref struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Roomref `json:"data"`
}

type Roomref struct {
	RoomtypeRefID   int    `json:"roomtype_ref_id" form:"roomtype_ref_id"`
	RoomtypeRefName string `json:"roomtype_ref_name" form:"roomtype_ref_name"`
	RoomtypeRefBed  string `json:"roomtype_ref_bed" form:"roomtype_ref_bed"`
}
