package http_handler

import (
	"go-base-cleancode/models"
	"go-base-cleancode/service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetRoomRef(c *gin.Context) {

	var response models.ResponseRoomref

	//begin bussiness logic
	res, err := usecase.GetRoomrefUsecase(idb.DB)
	if err == nil {
		response.Code = http.StatusOK
		response.Message = "Success"
		response.Status = "Success"
		response.Data = res
		c.JSON(http.StatusOK, response)
		return
	}

	response.Code = http.StatusInternalServerError
	response.Message = err.Error()
	response.Status = "Failed"
	response.Data = res
	c.JSON(http.StatusInternalServerError, response)
	return
}
