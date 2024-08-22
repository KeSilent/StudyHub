package edu_user_course

import (
	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/common/response"
	"github.com/KeSilent/study-hub/server/model/edu_user_course"
	edu_user_courseReq "github.com/KeSilent/study-hub/server/model/edu_user_course/request"
	"github.com/KeSilent/study-hub/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EduEnrollmentApi struct {
}

var eduEnrollmentService = service.ServiceGroupApp.Edu_user_courseServiceGroup.EduEnrollmentService

// CreateEduEnrollment 创建EduEnrollment
// @Tags EduEnrollment
// @Summary 创建EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_user_course.EduEnrollment true "创建EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduEnrollment/createEduEnrollment [post]
func (eduEnrollmentApi *EduEnrollmentApi) CreateEduEnrollment(c *gin.Context) {
	var eduEnrollment edu_user_course.EduEnrollment
	err := c.ShouldBindJSON(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.CreateEduEnrollment(&eduEnrollment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEduEnrollment 删除EduEnrollment
// @Tags EduEnrollment
// @Summary 删除EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_user_course.EduEnrollment true "删除EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduEnrollment/deleteEduEnrollment [delete]
func (eduEnrollmentApi *EduEnrollmentApi) DeleteEduEnrollment(c *gin.Context) {
	var eduEnrollment edu_user_course.EduEnrollment
	err := c.ShouldBindJSON(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.DeleteEduEnrollment(eduEnrollment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEduEnrollmentByIds 批量删除EduEnrollment
// @Tags EduEnrollment
// @Summary 批量删除EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /eduEnrollment/deleteEduEnrollmentByIds [delete]
func (eduEnrollmentApi *EduEnrollmentApi) DeleteEduEnrollmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.DeleteEduEnrollmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEduEnrollment 更新EduEnrollment
// @Tags EduEnrollment
// @Summary 更新EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_user_course.EduEnrollment true "更新EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduEnrollment/updateEduEnrollment [put]
func (eduEnrollmentApi *EduEnrollmentApi) UpdateEduEnrollment(c *gin.Context) {
	var eduEnrollment edu_user_course.EduEnrollment
	err := c.ShouldBindJSON(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.UpdateEduEnrollment(eduEnrollment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEduEnrollment 用id查询EduEnrollment
// @Tags EduEnrollment
// @Summary 用id查询EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_user_course.EduEnrollment true "用id查询EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduEnrollment/findEduEnrollment [get]
func (eduEnrollmentApi *EduEnrollmentApi) FindEduEnrollment(c *gin.Context) {
	var eduEnrollment edu_user_course.EduEnrollment
	err := c.ShouldBindQuery(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reeduEnrollment, err := eduEnrollmentService.GetEduEnrollment(eduEnrollment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reeduEnrollment": reeduEnrollment}, c)
	}
}

// GetEduEnrollmentList 分页获取EduEnrollment列表
// @Tags EduEnrollment
// @Summary 分页获取EduEnrollment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_user_courseReq.EduEnrollmentSearch true "分页获取EduEnrollment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduEnrollment/getEduEnrollmentList [get]
func (eduEnrollmentApi *EduEnrollmentApi) GetEduEnrollmentList(c *gin.Context) {
	var pageInfo edu_user_courseReq.EduEnrollmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := eduEnrollmentService.GetEduEnrollmentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

/**
 * @Description: 更新课程选择
 * @param {*gin.Context} c
 * @return {*}
 */
func (eduEnrollmentApi *EduEnrollmentApi) UpdateCourseId(c *gin.Context) {
	var eduEnrollment edu_user_course.EduEnrollment
	err := c.ShouldBindJSON(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.UpdateCourseId(eduEnrollment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

/**
 * @Description: 消耗课时
 * @param {*gin.Context} c
 * @return {*}
 */
func (eduEnrollmentApi *EduEnrollmentApi) ConsumptionClass(c *gin.Context) {
	var eduEnrollment edu_user_course.ConsumptionClassResp
	err := c.ShouldBindJSON(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.ConsumeSession(eduEnrollment.UserId, eduEnrollment.CourseId, eduEnrollment.SessionsToConsume, eduEnrollment.Reason, eduEnrollment.UseDate); err != nil {
		global.GVA_LOG.Error("消耗失败!", zap.Error(err))
		response.FailWithMessage("消耗失败", c)
	} else {
		response.OkWithMessage("消耗成功", c)
	}
}

/**
 * @Description: 增加课时
 * @param {*gin.Context} c
 * @return {*}
 */
func (eduEnrollmentApi *EduEnrollmentApi) AddSession(c *gin.Context) {
	var eduEnrollment edu_user_course.AddSessionResp
	err := c.ShouldBindJSON(&eduEnrollment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduEnrollmentService.AddSession(eduEnrollment.UserId, eduEnrollment.CourseId, eduEnrollment.SessionsToAdd, eduEnrollment.Reason, eduEnrollment.UseDate); err != nil {
		global.GVA_LOG.Error("增加课时失败!", zap.Error(err))
		response.FailWithMessage("增加课时失败", c)
	} else {
		response.OkWithMessage("增加课时成功", c)
	}
}
