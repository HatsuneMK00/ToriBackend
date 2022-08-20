package apis

import (
	"ToriBackend/entity"
	"ToriBackend/entity/response"
	"ToriBackend/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserApi interface {
	FindUser(c *gin.Context)
	FindUsersWithOffset(c *gin.Context)
	AddUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userApi struct{}

func (api userApi) FindUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	if user, ok := userService.FindUser(uint(id)); ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: user,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    404,
			Message: "user not found",
		})
	}
}

func (api userApi) FindUsersWithOffset(c *gin.Context) {
	param := c.Param("offset")
	offset, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	if users, ok := userService.FindUsersWithOffset(offset); ok {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Data: users,
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    404,
			Message: "user not found",
		})
	}
}

func (api userApi) AddUser(c *gin.Context) {
	var user entity.User
	var result *entity.User
	var rowAffected int64
	if err := c.ShouldBindJSON(&user); err == nil {
		result, rowAffected = userService.AddUser(&user)
	}
	if rowAffected > 0 {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "",
			Data:    result,
		})
		global.Logger.Infof("add user: %v", result.Username)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "unable to add user"})
	}
}

func (api userApi) UpdateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (api userApi) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: "user id needs to be uint",
		})
		return
	}
	if rowAffected := userService.DeleteUser(uint(id)); rowAffected > 0 {
		c.JSON(http.StatusOK, response.Response{
			Code:    http.StatusOK,
			Message: "",
		})
		global.Logger.Infof("delete user: %v", id)
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code:    500,
			Message: "unable to delete user",
		})
	}
}
