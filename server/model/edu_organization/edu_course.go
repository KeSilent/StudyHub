/*
 * @Author: Yang
 * @Date: 2023-04-10 18:33:49
 * @Description: 请填写简介
 */
// 自动生成模板EduCourse
package edu_organization

import (
	"github.com/KeSilent/study-hub/server/global"
)

// EduCourse 结构体
type EduCourse struct {
	global.GVA_MODEL
	OrganizationId  *int            `json:"organizationId" form:"organizationId" gorm:"column:organization_id;comment:组织ID;size:10;"`
	AdminId         *int            `json:"adminId" form:"adminId" gorm:"column:admin_id;comment:管理员ID;size:10;"`
	CourseName      string          `json:"courseName" form:"courseName" gorm:"column:course_name;comment:课程名称;size:255;"`
	Description     string          `json:"description" form:"description" gorm:"column:description;comment:课程描述;size:255;"`
	ImageUrl        string          `json:"imageUrl" form:"imageUrl" gorm:"column:image_url;comment:课程展示图URL;size:255;"`
	EduOrganization EduOrganization `json:"eduOrganization" gorm:"foreignKey:OrganizationId"`
}

// TableName EduCourse 表名
func (EduCourse) TableName() string {
	return "edu_course"
}
