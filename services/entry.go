package services

type ServiceGroup struct {
	User   UserService
	Card   CardService
	Record RecordService
}

var ServiceGroupApp = ServiceGroup{
	User:   userService{},
	Card:   cardService{},
	Record: recordService{},
}
