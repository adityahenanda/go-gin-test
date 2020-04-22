package http_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type InDB struct {
	DB *gorm.DB
}

func NewHttpHandler(route *gin.Engine, db *gorm.DB) {

	handler := &InDB{DB: db}

	v1 := route.Group("/v1")
	{
		api := v1.Group("/api")
		{
			audit := api.Group("/list")
			{
				audit.GET("", handler.GetList)
				audit.GET("/:auditID", handler.GetListDetail)
				audit.PUT("/:auditID", handler.UpdateAudit)
				// users.PUT("/:id", handler.Update)
				// users.DELETE("/:id", handler.Delete)
			}
			user := api.Group("/user")
			{
				user.GET("", handler.GetUser)
			}
			roomref := api.Group("/roomref")
			{
				roomref.GET("", handler.GetRoomRef)
			}
		}
	}
}
