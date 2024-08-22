/*
 * @Author: Yang
 * @Date: 2023-04-10 18:35:56
 * @Description: 请填写简介
 */
// 自动生成模板EduEnrollment
package edu_user_course

import (
	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/edu_organization"
)

// EduEnrollment 结构体
type EduEnrollment struct {
	global.GVA_MODEL
	UserId            *int                       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:10;"`
	CourseId          *int                       `json:"courseId" form:"courseId" gorm:"column:course_id;comment:课程ID;size:10;"`
	TotalSessions     *int                       `json:"totalSessions" form:"totalSessions" gorm:"column:total_sessions;comment:总课时数;size:10;"`
	RemainingSessions *int                       `json:"remainingSessions" form:"remainingSessions" gorm:"column:remaining_sessions;comment:剩余课时数;size:10;"`
	EduCourse         edu_organization.EduCourse `json:"eduCourse" gorm:"foreignKey:CourseId"`
}

// TableName EduEnrollment 表名
func (EduEnrollment) TableName() string {
	return "edu_enrollment"
}

// ConsumptionClassReq 消耗课时请求参数
type ConsumptionClassResp struct {
	UserId            int    `json:"userId"`            // 用户ID
	CourseId          int    `json:"courseId"`          // 课程ID
	SessionsToConsume int    `json:"sessionsToConsume"` // 消耗课时数
	Reason            string `json:"reason"`            // 消耗原因
	UseDate           string `json:"useDate"`
}

// AddSessionReq 增加课时请求参数
type AddSessionResp struct {
	UserId        int    `json:"userId"`        // 用户ID
	CourseId      int    `json:"courseId"`      // 课程ID
	SessionsToAdd int    `json:"sessionsToAdd"` // 增加课时数
	Reason        string `json:"reason"`        // 消耗原因
	UseDate       string `json:"useDate"`
}
