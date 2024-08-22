package service

import (
	"github.com/KeSilent/study-hub/server/service/edu_organization"
	"github.com/KeSilent/study-hub/server/service/edu_user_course"
	"github.com/KeSilent/study-hub/server/service/example"
	"github.com/KeSilent/study-hub/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup           system.ServiceGroup
	ExampleServiceGroup          example.ServiceGroup
	Edu_organizationServiceGroup edu_organization.ServiceGroup
	Edu_user_courseServiceGroup  edu_user_course.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
