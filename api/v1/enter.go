package v1

import (
	"start-model/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
