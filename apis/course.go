package apis

import (
	"ToriBackend/entity/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CourseApi interface {
	GetAllCourses(c *gin.Context)
	GetCourse(c *gin.Context)
	EnrollCourse(c *gin.Context)
	GetEnrolledCourses(c *gin.Context)
}

type courseApi struct {
}

func (a courseApi) GetEnrolledCourses(c *gin.Context) {
	userId := c.Query("user_id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	if courses, ok := courseService.GetEnrolledCourses(uint(id)); ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: courses,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    500,
			Message: "fail to get courses",
		})
	}
}

func (a courseApi) GetAllCourses(c *gin.Context) {
	courses, ok := courseService.GetAllCourses()
	if ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: courses,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "failed to get courses",
		})
	}
}

func (a courseApi) GetCourse(c *gin.Context) {
	courseId := c.Param("course_id")
	id, err := strconv.Atoi(courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "course id needs to be uint",
		})
		return
	}
	if course, ok := courseService.GetCourse(uint(id)); ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: course,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    404,
			Message: "course not found",
		})
	}
}

func (a courseApi) EnrollCourse(c *gin.Context) {
	type param struct {
		CourseId uint `json:"course_id"`
		UserId   uint `json:"user_id"`
	}
	var p param
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid params",
		})
		return
	}
	if ok := courseService.EnrollCourse(p.UserId, p.CourseId); ok {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "success to enroll course",
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusInternalServerError,
			Message: "failed to enroll course",
		})
	}
}
