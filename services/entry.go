package services

type ServiceGroup struct {
	User  UserService
	Login LoginService
}

var ServiceGroupApp = ServiceGroup{
	User:  userService{},
	Login: loginService{},
}
