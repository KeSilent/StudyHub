package edu_user_course

import (
	"net/http"
	"strconv"

	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/common/response"
	"github.com/KeSilent/study-hub/server/model/edu_user_course"
	edu_user_courseReq "github.com/KeSilent/study-hub/server/model/edu_user_course/request"
	"github.com/KeSilent/study-hub/server/service"
	"github.com/KeSilent/study-hub/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EduClassSessionApi struct {
}

var eduClassSessionService = service.ServiceGroupApp.Edu_user_courseServiceGroup.EduClassSessionService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

// CreateEduClassSession 创建EduClassSession
// @Tags EduClassSession
// @Summary 创建EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_user_course.EduClassSession true "创建EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduClassSession/createEduClassSession [post]
func (eduClassSessionApi *EduClassSessionApi) CreateEduClassSession(c *gin.Context) {
	var eduClassSession edu_user_course.EduClassSession
	err := c.ShouldBindJSON(&eduClassSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduClassSessionService.CreateEduClassSession(&eduClassSession); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEduClassSession 删除EduClassSession
// @Tags EduClassSession
// @Summary 删除EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_user_course.EduClassSession true "删除EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduClassSession/deleteEduClassSession [delete]
func (eduClassSessionApi *EduClassSessionApi) DeleteEduClassSession(c *gin.Context) {
	var eduClassSession edu_user_course.EduClassSession
	err := c.ShouldBindJSON(&eduClassSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduClassSessionService.DeleteEduClassSession(eduClassSession); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEduClassSessionByIds 批量删除EduClassSession
// @Tags EduClassSession
// @Summary 批量删除EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /eduClassSession/deleteEduClassSessionByIds [delete]
func (eduClassSessionApi *EduClassSessionApi) DeleteEduClassSessionByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduClassSessionService.DeleteEduClassSessionByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEduClassSession 更新EduClassSession
// @Tags EduClassSession
// @Summary 更新EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body edu_user_course.EduClassSession true "更新EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduClassSession/updateEduClassSession [put]
func (eduClassSessionApi *EduClassSessionApi) UpdateEduClassSession(c *gin.Context) {
	var eduClassSession edu_user_course.EduClassSession
	err := c.ShouldBindJSON(&eduClassSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := eduClassSessionService.UpdateEduClassSession(eduClassSession); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEduClassSession 用id查询EduClassSession
// @Tags EduClassSession
// @Summary 用id查询EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_user_course.EduClassSession true "用id查询EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduClassSession/findEduClassSession [get]
func (eduClassSessionApi *EduClassSessionApi) FindEduClassSession(c *gin.Context) {
	var eduClassSession edu_user_course.EduClassSession
	err := c.ShouldBindQuery(&eduClassSession)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reeduClassSession, err := eduClassSessionService.GetEduClassSession(eduClassSession.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reeduClassSession": reeduClassSession}, c)
	}
}

// GetEduClassSessionList 分页获取EduClassSession列表
// @Tags EduClassSession
// @Summary 分页获取EduClassSession列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query edu_user_courseReq.EduClassSessionSearch true "分页获取EduClassSession列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduClassSession/getEduClassSessionList [get]
func (eduClassSessionApi *EduClassSessionApi) GetEduClassSessionList(c *gin.Context) {
	var pageInfo edu_user_courseReq.EduClassSessionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := eduClassSessionService.GetEduClassSessionInfoList(pageInfo); err != nil {
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
 * @Description: 获取用户的消费列表
 * @param {*gin.Context} c
 * @return {*}
 */
func (eduClassSessionApi *EduClassSessionApi) GetEduClassSessionListByUser(c *gin.Context) {

	studentIdStr := c.Query("studentId")
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "学生ID无效"})
		return
	}

	classSessions, err := eduClassSessionService.GetStudentClassSessions(studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取学生课时记录失败"})
		return
	}

	response.OkWithDetailed(response.PageResult{
		List: classSessions,
	}, "获取成功", c)
}

/**
 * @Description: 获取剩余课时少于5节的学生列表
 * @param {*gin.Context} c
 * @return {*}
 */
func (eduClassSessionApi *EduClassSessionApi) GetStudentsWithLessThanFiveSessions(c *gin.Context) {
	userId := utils.GetUserID(c)

	ReqUser, err := userService.FindUserById(int(userId))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	students, err := eduClassSessionService.GetStudentsWithLessThanFiveSessions(ReqUser.EduOrganizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取学生课时记录失败"})
		return
	}

	response.OkWithDetailed(response.PageResult{
		List: students,
	}, "获取成功", c)
}
