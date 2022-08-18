package apis

import "WebAppStructure/services"

type ApiGroup struct {
	Login LoginApi
	User  UserApi
}

var (
	loginService = services.ServiceGroupApp.Login
	userService  = services.ServiceGroupApp.User
)

var ApiGroupApp = ApiGroup{
	Login: loginApi{},
	User:  userApi{},
}
