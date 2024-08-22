/*
 * @Author: Yang
 * @Date: 2023-04-10 18:33:49
 * @Description: 组织
 */
package edu_organization

import (
	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/edu_organization"
	edu_organizationReq "github.com/KeSilent/study-hub/server/model/edu_organization/request"
)

type EduCourseService struct {
}

// CreateEduCourse 创建EduCourse记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduCourseService *EduCourseService) CreateEduCourse(eduCourse *edu_organization.EduCourse) (err error) {
	err = global.GVA_DB.Create(eduCourse).Error
	return err
}

// DeleteEduCourse 删除EduCourse记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduCourseService *EduCourseService) DeleteEduCourse(eduCourse edu_organization.EduCourse) (err error) {
	err = global.GVA_DB.Delete(&eduCourse).Error
	return err
}

// DeleteEduCourseByIds 批量删除EduCourse记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduCourseService *EduCourseService) DeleteEduCourseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]edu_organization.EduCourse{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEduCourse 更新EduCourse记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduCourseService *EduCourseService) UpdateEduCourse(eduCourse edu_organization.EduCourse) (err error) {
	err = global.GVA_DB.Save(&eduCourse).Error
	return err
}

// GetEduCourse 根据id获取EduCourse记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduCourseService *EduCourseService) GetEduCourse(id uint) (eduCourse edu_organization.EduCourse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eduCourse).Error
	return
}

// GetEduCourseInfoList 分页获取EduCourse记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduCourseService *EduCourseService) GetEduCourseInfoList(info edu_organizationReq.EduCourseSearch) (list []edu_organization.EduCourse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&edu_organization.EduCourse{})
	var eduCourses []edu_organization.EduCourse
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrganizationId != nil {
		db = db.Where("organization_id = ? ", info.OrganizationId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Preload("EduOrganization").Limit(limit).Offset(offset).Find(&eduCourses).Error
	return eduCourses, total, err
}

/**
 * @Description: 通过组织ID获取当前组织下的课程
 * @param {uint} oId
 * @return {*}
 */
func (eduCourseService *EduCourseService) GetEduCourseInfoListByOrganId(oId uint) (list []edu_organization.EduCourse, err error) {

	db := global.GVA_DB.Model(&edu_organization.EduCourse{})
	var eduCourses []edu_organization.EduCourse
	db = db.Where("organization_id = ?", oId)

	err = db.Preload("EduOrganization").Find(&eduCourses).Error

	return eduCourses, err
}
