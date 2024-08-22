// 自动生成模板EduOrganization
package edu_organization

import (
	"github.com/KeSilent/study-hub/server/global"
)

// EduOrganization 结构体
type EduOrganization struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:组织名称;size:255;"`
	Brief       string `json:"brief" form:"brief" gorm:"column:brief;comment:组织简介;size:255;"`
	Details     string `json:"details" form:"details" gorm:"column:details;comment:组织详情（HTML字符串）;size:4294967295;"`
	ContactInfo string `json:"contactInfo" form:"contactInfo" gorm:"column:contact_info;comment:联系方式;size:255;"`
	Address     string `json:"address" form:"address" gorm:"column:address;comment:地址;size:255;"`
}

// TableName EduOrganization 表名
func (EduOrganization) TableName() string {
	return "edu_organization"
}
