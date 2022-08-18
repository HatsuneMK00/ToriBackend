package routes

import (
	"WebAppStructure/routes/api"
)

type routerGroup struct {
	api.RouterGroup
	AuthRouter
}

var RouterGroupApp = new(routerGroup)
