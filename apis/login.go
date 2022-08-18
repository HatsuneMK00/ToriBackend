package apis

import (
	"WebAppStructure/entity"
)

type LoginApi interface {
	LoginUser(username string, password string) (*entity.User, bool)
	LogoutUser(userId string)
}

type loginApi struct{}

func (c loginApi) LoginUser(username string, password string) (*entity.User, bool) {
	if loginService.IsUserExist(username) {
		return &entity.User{}, true
	} else {
		return nil, false
	}
}

func (c loginApi) LogoutUser(userId string) {
	//TODO implement me
	panic("implement me")
}
