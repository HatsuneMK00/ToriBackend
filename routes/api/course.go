package api

import (
	"ToriBackend/apis"
	"github.com/gin-gonic/gin"
)

type courseRouter struct{}

func (r courseRouter) AddCourseRoutes(rg *gin.RouterGroup) {
	course := rg.Group("/course")
	{
		course.GET("/:course_id", apis.ApiGroupApp.Course.GetCourse)
		course.GET("/all", apis.ApiGroupApp.Course.GetAllCourses)
		course.GET("/enrolled", apis.ApiGroupApp.Course.GetEnrolledCourses)
		course.POST("/enroll", apis.ApiGroupApp.Course.EnrollCourse)
	}
}
