// 自动生成模板EduEnrollment
package edu_organization

import (
	"github.com/KeSilent/study-hub/server/global"
)

// EduEnrollment 结构体
type EduEnrollment struct {
	global.GVA_MODEL
	UserId            *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:10;"`
	CourseId          *int `json:"courseId" form:"courseId" gorm:"column:course_id;comment:课程ID;size:10;"`
	TotalSessions     *int `json:"totalSessions" form:"totalSessions" gorm:"column:total_sessions;comment:总课时数;size:10;"`
	RemainingSessions *int `json:"remainingSessions" form:"remainingSessions" gorm:"column:remaining_sessions;comment:剩余课时数;size:10;"`
}

// TableName EduEnrollment 表名
func (EduEnrollment) TableName() string {
	return "edu_enrollment"
}
