package router

import (
	"github.com/KeSilent/study-hub/server/router/edu_organization"
	"github.com/KeSilent/study-hub/server/router/edu_user_course"
	"github.com/KeSilent/study-hub/server/router/example"
	"github.com/KeSilent/study-hub/server/router/system"
)

type RouterGroup struct {
	System           system.RouterGroup
	Example          example.RouterGroup
	Edu_organization edu_organization.RouterGroup
	Edu_user_course  edu_user_course.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
