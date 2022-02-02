package service

import (
	"start-model/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
