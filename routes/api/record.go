package api

import (
	"ToriBackend/apis"
	"github.com/gin-gonic/gin"
)

type recordRouter struct{}

func (r recordRouter) AddRecordRoutes(rg *gin.RouterGroup) {
	record := rg.Group("/record")
	{
		record.GET("/all", apis.ApiGroupApp.Record.FindAllRecords)
		record.GET("/monthly", apis.ApiGroupApp.Record.FindRecordsOfMonth)
		record.POST("/", apis.ApiGroupApp.Record.AddRecord)
		record.PUT("/", apis.ApiGroupApp.Record.UpdateRecord)
	}
}
