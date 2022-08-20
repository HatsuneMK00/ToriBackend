package services

type ServiceGroup struct {
	User   UserService
	Login  LoginService
	Card   CardService
	Record RecordService
}

var ServiceGroupApp = ServiceGroup{
	User:   userService{},
	Login:  loginService{},
	Card:   cardService{},
	Record: recordService{},
}
