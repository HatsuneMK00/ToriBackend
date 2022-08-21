package apis

import (
	"ToriBackend/entity/response"
	"ToriBackend/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RecordApi interface {
	FindAllRecords(c *gin.Context)
	FindRecordsOfMonth(c *gin.Context)
	AddRecord(c *gin.Context)
	UpdateRecord(c *gin.Context)
}

type recordApi struct{}

func (r recordApi) FindAllRecords(c *gin.Context) {
	id := c.Query("user_id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	records, ok := recordService.FindAllRecords(uint(userId))
	if ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: records,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    500,
			Message: "records not found",
		})
	}
}

func (r recordApi) FindRecordsOfMonth(c *gin.Context) {
	id := c.Query("user_id")
	year, err1 := strconv.Atoi(c.Query("year"))
	month, err2 := strconv.Atoi(c.Query("month"))
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	if err1 != nil {
		global.Logger.Warnf("year is not a number")
	}
	if err2 != nil {
		global.Logger.Warnf("month is not a number")
	}
	records, ok := recordService.FindRecordsOfMonth(uint(userId), year, month)
	if ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: records,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    500,
			Message: "records not found",
		})
	}
}

func (r recordApi) AddRecord(c *gin.Context) {
	type param struct {
		UserId      uint   `json:"user_id"`
		CardId      uint   `json:"card_id"`
		Description string `json:"description"`
	}
	var p param
	err := c.ShouldBindJSON(&p)
	if err != nil {
		global.Logger.Errorf("error: %v", err)
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "params error",
		})
		return
	}
	record, rowAffected := recordService.AddRecord(p.UserId, p.CardId, p.Description)
	if rowAffected > 0 {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: record,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    500,
			Message: "record not added",
		})
	}
}

func (r recordApi) UpdateRecord(c *gin.Context) {
	type param struct {
		RecordId    uint   `json:"record_id"`
		CardId      uint   `json:"card_id"`
		Description string `json:"description"`
	}
	var p param
	err := c.ShouldBind(&p)
	if err != nil {
		global.Logger.Errorf("error: %v", err)
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "params error",
		})
		return
	}
	rowAffected := recordService.UpdateRecord(p.RecordId, p.CardId, p.Description)
	if rowAffected > 0 {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: "record updated",
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    500,
			Message: "record not updated",
		})
	}
}
