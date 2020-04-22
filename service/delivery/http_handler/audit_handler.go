package http_handler

import (
	"encoding/json"
	"go-base-cleancode/models"
	"go-base-cleancode/service/usecase"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetList(c *gin.Context) {

	var response models.ResponseBuilding

	//prevent sql injection with queryEscape
	filter := url.QueryEscape(c.DefaultQuery("filter", ""))
	pageParam := url.QueryEscape(c.DefaultQuery("page", "1"))
	limitParam := url.QueryEscape(c.DefaultQuery("limit", "10"))

	//parse page and limit into int using uint to check value whether string or negative number
	page, err := strconv.ParseUint(pageParam, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Failed"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Failed"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	res, err := usecase.GetBuildingUsecase(idb.DB, int(limit), int(page), filter)
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

func (idb *InDB) GetListDetail(c *gin.Context) {

	var response models.ResponseBuildingDetail
	auditParamID := c.Param("auditID")

	//parse params id into int using uint to check value whether string or negative number
	auditID, err := strconv.ParseUint(auditParamID, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Failed"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	res, err := usecase.GetBuildingDetailUsecase(idb.DB, int(auditID))
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

func (idb *InDB) UpdateAudit(c *gin.Context) {

	var response models.ResponseAuditUpdated
	var req models.AuditDetailRequest

	auditParamID := c.Param("auditID")
	body, _ := ioutil.ReadAll(c.Request.Body)

	//parse params id into int using uint to check value whether string or negative number
	auditID, err := strconv.ParseUint(auditParamID, 10, 64)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Failed"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "Failed"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	err = usecase.UpdateAuditUsecase(idb.DB, int(auditID), req)
	if err == nil {
		response.Code = http.StatusOK
		response.Message = "Success"
		response.Status = "Success"
		response.Data = int(auditID)
		c.JSON(http.StatusOK, response)
		return
	}

	response.Code = http.StatusInternalServerError
	response.Message = err.Error()
	response.Status = "Failed"
	c.JSON(http.StatusInternalServerError, response)
	return
}
