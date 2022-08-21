package apis

import "ToriBackend/services"

type ApiGroup struct {
	User        UserApi
	Card        CardApi
	Record      RecordApi
	Course      CourseApi
	Achievement AchievementApi
}

var (
	userService        = services.ServiceGroupApp.User
	cardService        = services.ServiceGroupApp.Card
	recordService      = services.ServiceGroupApp.Record
	courseService      = services.ServiceGroupApp.Course
	achievementService = services.ServiceGroupApp.Achievement
)

var ApiGroupApp = ApiGroup{
	User:        userApi{},
	Card:        cardApi{},
	Record:      recordApi{},
	Course:      courseApi{},
	Achievement: achievementApi{},
}
