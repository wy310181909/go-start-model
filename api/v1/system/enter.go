package system

import "start-model/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
)
