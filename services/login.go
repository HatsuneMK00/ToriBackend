package services

type LoginService interface {
	IsUserExist(username string) bool
}

type loginService struct{}

func (s loginService) IsUserExist(username string) bool {
	if username == "admin" {
		return true
	} else {
		return false
	}
}
