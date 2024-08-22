/*
 * @Author: Yang
 * @Date: 2023-04-10 18:35:56
 * @Description: 学生绑定课程信息
 */
package edu_user_course

import (
	"errors"
	"time"

	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/edu_user_course"
	edu_user_courseReq "github.com/KeSilent/study-hub/server/model/edu_user_course/request"
)

type EduEnrollmentService struct {
}

var eduClassSessionService EduClassSessionService

// CreateEduEnrollment 创建EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) CreateEduEnrollment(eduEnrollment *edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Create(eduEnrollment).Error
	return err
}

// DeleteEduEnrollment 删除EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) DeleteEduEnrollment(eduEnrollment edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Delete(&eduEnrollment).Error
	return err
}

// DeleteEduEnrollmentByIds 批量删除EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) DeleteEduEnrollmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]edu_user_course.EduEnrollment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEduEnrollment 更新EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) UpdateEduEnrollment(eduEnrollment edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Save(&eduEnrollment).Error
	return err
}

// UpdateEduEnrollment 更新课程选择
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) UpdateCourseId(eduEnrollment edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Model(&edu_user_course.EduEnrollment{}).Where("ID = ?", eduEnrollment.ID).Update("course_id", eduEnrollment.CourseId).Error

	return err
}

// GetEduEnrollment 根据id获取EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) GetEduEnrollment(id uint) (eduEnrollment edu_user_course.EduEnrollment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eduEnrollment).Error
	return
}

// GetEduEnrollmentInfoList 分页获取EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) GetEduEnrollmentInfoList(info edu_user_courseReq.EduEnrollmentSearch) (list []edu_user_course.EduEnrollment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&edu_user_course.EduEnrollment{})
	var eduEnrollments []edu_user_course.EduEnrollment
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

// ConsumeSession 消耗课时
func (eduEnrollmentService *EduEnrollmentService) ConsumeSession(userID, courseID, sessionsToConsume int, reason string, useData string) error {
	var enrollment edu_user_course.EduEnrollment
	db := global.GVA_DB

	// 查询用户在指定课程中的报名信息
	result := db.Where("user_id = ? AND course_id = ?", userID, courseID).Preload("EduCourse").First(&enrollment)

	// 如果找不到报名信息或发生错误，则返回错误
	if result.Error != nil {
		return errors.New("报名信息未找到或查询错误")
	}

	// 检查剩余课时是否足够
	if enrollment.RemainingSessions == nil || *enrollment.RemainingSessions < sessionsToConsume {
		return errors.New("剩余课时不足")
	}

	// 扣除课时
	*enrollment.RemainingSessions -= sessionsToConsume

	enenrollmentId := int(enrollment.ID)
	t, err := time.Parse("2006-01-02", useData)
	if err != nil {
		return errors.New("日期转换失败:" + err.Error())
	}

	// 更新数据库中的报名信息
	result = db.Save(&enrollment)

	if result.Error != nil {
		return errors.New("更新报名信息失败")
	}
	// 记录课时操作
	classSession := edu_user_course.EduClassSession{
		EnrollmentId: &enenrollmentId,
		Action:       "subtract",
		Reason:       reason,
		NumSessions:  &sessionsToConsume,
		CourseName:   enrollment.EduCourse.CourseName,
		UseDate:      t,
	}

	err = eduClassSessionService.CreateEduClassSession(&classSession)

	if err != nil {
		return errors.New("创建课时操作记录失败")
	}

	return nil
}

// AddSession 为用户的课程添加课时
func (eduEnrollmentService *EduEnrollmentService) AddSession(userID, courseID, sessionsToAdd int, reason string, useData string) error {
	var enrollment edu_user_course.EduEnrollment
	db := global.GVA_DB

	// 查询用户在指定课程中的报名信息
	result := db.Where("user_id = ? AND course_id = ?", userID, courseID).Preload("EduCourse").First(&enrollment)

	// 如果找不到报名信息或发生错误，则返回错误
	if result.Error != nil {
		return errors.New("报名信息未找到或查询错误")
	}

	// 增加总课时
	*enrollment.TotalSessions += sessionsToAdd

	// 增加剩余课时
	if enrollment.RemainingSessions == nil {
		enrollment.RemainingSessions = new(int)
	}
	*enrollment.RemainingSessions += sessionsToAdd

	t, err := time.Parse("2006-01-02", useData)
	if err != nil {
		return errors.New("日期转换失败:" + err.Error())
	}

	// 更新数据库中的报名信息
	result = db.Save(&enrollment)

	if result.Error != nil {
		return errors.New("更新报名信息失败")
	}

	enrollmentId := int(enrollment.ID)
	// 记录课时操作
	classSession := edu_user_course.EduClassSession{
		EnrollmentId: &enrollmentId,
		Action:       "add",
		Reason:       reason,
		NumSessions:  &sessionsToAdd,
		CourseName:   enrollment.EduCourse.CourseName,
		UseDate:      t,
	}

	result = db.Create(&classSession)

	if result.Error != nil {
		return errors.New("创建课时操作记录失败")
	}

	return nil
}

// GetEduEnrollmentByUser 根据用户ID和课程ID获取报名信息
func (eduEnrollmentService *EduEnrollmentService) GetEduEnrollmentByUser(userId uint, courseId uint) (eduEnrollment edu_user_course.EduEnrollment, err error) {
	err = global.GVA_DB.Where("user_id = ? and course_id=?", userId, courseId).First(&eduEnrollment).Error
	return
}
