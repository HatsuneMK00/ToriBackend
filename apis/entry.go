package apis

import "ToriBackend/services"

type ApiGroup struct {
	Login LoginApi
	User  UserApi
	Card  CardApi
}

var (
	loginService = services.ServiceGroupApp.Login
	userService  = services.ServiceGroupApp.User
	cardService  = services.ServiceGroupApp.Card
)

var ApiGroupApp = ApiGroup{
	Login: loginApi{},
	User:  userApi{},
	Card:  cardApi{},
}
