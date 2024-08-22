// 自动生成模板EduClassSession
package edu_user_course

import (
	"time"

	"github.com/KeSilent/study-hub/server/global"
)

// EduClassSession 结构体
type EduClassSession struct {
	global.GVA_MODEL
	EnrollmentId  *int          `json:"enrollmentId" form:"enrollmentId" gorm:"column:enrollment_id;comment:报名ID;size:10;"`
	Action        string        `json:"action" form:"action" gorm:"column:action;type:enum('add', 'subtract');comment:操作类型（增加或扣除）;"`
	Reason        string        `json:"reason" form:"reason" gorm:"column:reason;comment:操作原因;size:255;"`
	NumSessions   *int          `json:"numSessions" form:"numSessions" gorm:"column:num_sessions;comment:课时数量;size:10;"`
	CourseName    string        `json:"courseName" form:"courseName" gorm:"column:course_name;comment:课程名称;size:255;"`
	EduEnrollment EduEnrollment `json:"eduEnrollment" gorm:"foreignKey:EnrollmentId"` // 用户科目
	UseDate       time.Time     `json:"useDate" form:"useDate" gorm:"column:use_date;comment:使用日期;size:255;"`
}

// TableName EduClassSession 表名
func (EduClassSession) TableName() string {
	return "edu_class_session"
}
