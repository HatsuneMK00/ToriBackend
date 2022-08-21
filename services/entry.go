package services

type ServiceGroup struct {
	User   UserService
	Card   CardService
	Record RecordService
	Course CourseService
}

var ServiceGroupApp = ServiceGroup{
	User:   userService{},
	Card:   cardService{},
	Record: recordService{},
	Course: courseService{},
}
