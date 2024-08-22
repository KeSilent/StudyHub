package edu_organization

import (
	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/edu_organization"
	edu_organizationReq "github.com/KeSilent/study-hub/server/model/edu_organization/request"
)

type EduEnrollmentService struct {
}

// CreateEduEnrollment 创建EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) CreateEduEnrollment(eduEnrollment *edu_organization.EduEnrollment) (err error) {
	err = global.GVA_DB.Create(eduEnrollment).Error
	return err
}

// DeleteEduEnrollment 删除EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) DeleteEduEnrollment(eduEnrollment edu_organization.EduEnrollment) (err error) {
	err = global.GVA_DB.Delete(&eduEnrollment).Error
	return err
}

// DeleteEduEnrollmentByIds 批量删除EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) DeleteEduEnrollmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]edu_organization.EduEnrollment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEduEnrollment 更新EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) UpdateEduEnrollment(eduEnrollment edu_organization.EduEnrollment) (err error) {
	err = global.GVA_DB.Save(&eduEnrollment).Error
	return err
}

// GetEduEnrollment 根据id获取EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) GetEduEnrollment(id uint) (eduEnrollment edu_organization.EduEnrollment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eduEnrollment).Error
	return
}

// GetEduEnrollmentInfoList 分页获取EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) GetEduEnrollmentInfoList(info edu_organizationReq.EduEnrollmentSearch) (list []edu_organization.EduEnrollment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&edu_organization.EduEnrollment{})
	var eduEnrollments []edu_organization.EduEnrollment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&eduEnrollments).Error
	return eduEnrollments, total, err
}
