package apis

import "ToriBackend/services"

type ApiGroup struct {
	User   UserApi
	Card   CardApi
	Record RecordApi
}

var (
	userService   = services.ServiceGroupApp.User
	cardService   = services.ServiceGroupApp.Card
	recordService = services.ServiceGroupApp.Record
)

var ApiGroupApp = ApiGroup{
	User:   userApi{},
	Card:   cardApi{},
	Record: recordApi{},
}
