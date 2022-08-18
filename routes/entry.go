package routes

import (
	"ToriBackend/routes/api"
)

type routerGroup struct {
	api.RouterGroup
	AuthRouter
}

var RouterGroupApp = new(routerGroup)
