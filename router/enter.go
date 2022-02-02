package router

import (
	"start-model/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
