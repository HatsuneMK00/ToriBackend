package services

type ServiceGroup struct {
	User        UserService
	Card        CardService
	Record      RecordService
	Course      CourseService
	Achievement AchievementService
}

var ServiceGroupApp = ServiceGroup{
	User:        userService{},
	Card:        cardService{},
	Record:      recordService{},
	Course:      courseService{},
	Achievement: achievementService{},
}
