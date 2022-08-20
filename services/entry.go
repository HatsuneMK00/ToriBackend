package services

type ServiceGroup struct {
	User  UserService
	Login LoginService
	Card  CardService
}

var ServiceGroupApp = ServiceGroup{
	User:  userService{},
	Login: loginService{},
	Card:  cardService{},
}
