package edu_organization

import (
	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/common/response"
	"github.com/KeSilent/study-hub/server/model/edu_organization"
	edu_organizationReq "github.com/KeSilent/study-hub/server/model/edu_organization/request"
	"github.com/KeSilent/study-hub/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EduCourseApi struct {
}

var eduCourseService = service.ServiceGroupApp.Edu_organizationServiceGroup.EduCourseService

// CreateEduCourse 创建EduCourse
// @Tags EduCourse
// @Summary 创建EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_organization.EduCourse true "创建EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduCourse/createEduCourse [post]
func (eduCourseApi *EduCourseApi) CreateEduCourse(c *gin.Context) {
	var eduCourse edu_organization.EduCourse
	err := c.ShouldBindJSON(&eduCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduCourseService.CreateEduCourse(&eduCourse); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEduCourse 删除EduCourse
// @Tags EduCourse
// @Summary 删除EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_organization.EduCourse true "删除EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduCourse/deleteEduCourse [delete]
func (eduCourseApi *EduCourseApi) DeleteEduCourse(c *gin.Context) {
	var eduCourse edu_organization.EduCourse
	err := c.ShouldBindJSON(&eduCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduCourseService.DeleteEduCourse(eduCourse); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEduCourseByIds 批量删除EduCourse
// @Tags EduCourse
// @Summary 批量删除EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /eduCourse/deleteEduCourseByIds [delete]
func (eduCourseApi *EduCourseApi) DeleteEduCourseByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduCourseService.DeleteEduCourseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEduCourse 更新EduCourse
// @Tags EduCourse
// @Summary 更新EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_organization.EduCourse true "更新EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduCourse/updateEduCourse [put]
func (eduCourseApi *EduCourseApi) UpdateEduCourse(c *gin.Context) {
	var eduCourse edu_organization.EduCourse
	err := c.ShouldBindJSON(&eduCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduCourseService.UpdateEduCourse(eduCourse); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEduCourse 用id查询EduCourse
// @Tags EduCourse
// @Summary 用id查询EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_organization.EduCourse true "用id查询EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduCourse/findEduCourse [get]
func (eduCourseApi *EduCourseApi) FindEduCourse(c *gin.Context) {
	var eduCourse edu_organization.EduCourse
	err := c.ShouldBindQuery(&eduCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reeduCourse, err := eduCourseService.GetEduCourse(eduCourse.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reeduCourse": reeduCourse}, c)
	}
}

// GetEduCourseList 分页获取EduCourse列表
// @Tags EduCourse
// @Summary 分页获取EduCourse列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_organizationReq.EduCourseSearch true "分页获取EduCourse列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduCourse/getEduCourseList [get]
func (eduCourseApi *EduCourseApi) GetEduCourseList(c *gin.Context) {
	var pageInfo edu_organizationReq.EduCourseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := eduCourseService.GetEduCourseInfoList(pageInfo); err != nil {
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
