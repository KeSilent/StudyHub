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

type EduOrganizationApi struct {
}

var eduOrganizationService = service.ServiceGroupApp.Edu_organizationServiceGroup.EduOrganizationService

// CreateEduOrganization 创建EduOrganization
// @Tags EduOrganization
// @Summary 创建EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_organization.EduOrganization true "创建EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduOrganization/createEduOrganization [post]
func (eduOrganizationApi *EduOrganizationApi) CreateEduOrganization(c *gin.Context) {
	var eduOrganization edu_organization.EduOrganization
	err := c.ShouldBindJSON(&eduOrganization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduOrganizationService.CreateEduOrganization(&eduOrganization); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEduOrganization 删除EduOrganization
// @Tags EduOrganization
// @Summary 删除EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_organization.EduOrganization true "删除EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduOrganization/deleteEduOrganization [delete]
func (eduOrganizationApi *EduOrganizationApi) DeleteEduOrganization(c *gin.Context) {
	var eduOrganization edu_organization.EduOrganization
	err := c.ShouldBindJSON(&eduOrganization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduOrganizationService.DeleteEduOrganization(eduOrganization); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEduOrganizationByIds 批量删除EduOrganization
// @Tags EduOrganization
// @Summary 批量删除EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /eduOrganization/deleteEduOrganizationByIds [delete]
func (eduOrganizationApi *EduOrganizationApi) DeleteEduOrganizationByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduOrganizationService.DeleteEduOrganizationByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEduOrganization 更新EduOrganization
// @Tags EduOrganization
// @Summary 更新EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_organization.EduOrganization true "更新EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduOrganization/updateEduOrganization [put]
func (eduOrganizationApi *EduOrganizationApi) UpdateEduOrganization(c *gin.Context) {
	var eduOrganization edu_organization.EduOrganization
	err := c.ShouldBindJSON(&eduOrganization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduOrganizationService.UpdateEduOrganization(eduOrganization); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEduOrganization 用id查询EduOrganization
// @Tags EduOrganization
// @Summary 用id查询EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_organization.EduOrganization true "用id查询EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduOrganization/findEduOrganization [get]
func (eduOrganizationApi *EduOrganizationApi) FindEduOrganization(c *gin.Context) {
	var eduOrganization edu_organization.EduOrganization
	err := c.ShouldBindQuery(&eduOrganization)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reeduOrganization, err := eduOrganizationService.GetEduOrganization(eduOrganization.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reeduOrganization": reeduOrganization}, c)
	}
}

// GetEduOrganizationList 分页获取EduOrganization列表
// @Tags EduOrganization
// @Summary 分页获取EduOrganization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_organizationReq.EduOrganizationSearch true "分页获取EduOrganization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduOrganization/getEduOrganizationList [get]
func (eduOrganizationApi *EduOrganizationApi) GetEduOrganizationList(c *gin.Context) {
	var pageInfo edu_organizationReq.EduOrganizationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := eduOrganizationService.GetEduOrganizationInfoList(pageInfo); err != nil {
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
