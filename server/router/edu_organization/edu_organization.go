package edu_organization

import (
	"github.com/KeSilent/study-hub/server/api/v1"
	"github.com/KeSilent/study-hub/server/middleware"
	"github.com/gin-gonic/gin"
)

type EduOrganizationRouter struct {
}

// InitEduOrganizationRouter 初始化 EduOrganization 路由信息
func (s *EduOrganizationRouter) InitEduOrganizationRouter(Router *gin.RouterGroup) {
	eduOrganizationRouter := Router.Group("eduOrganization").Use(middleware.OperationRecord())
	eduOrganizationRouterWithoutRecord := Router.Group("eduOrganization")
	var eduOrganizationApi = v1.ApiGroupApp.Edu_organizationApiGroup.EduOrganizationApi
	{
		eduOrganizationRouter.POST("createEduOrganization", eduOrganizationApi.CreateEduOrganization)             // 新建EduOrganization
		eduOrganizationRouter.DELETE("deleteEduOrganization", eduOrganizationApi.DeleteEduOrganization)           // 删除EduOrganization
		eduOrganizationRouter.DELETE("deleteEduOrganizationByIds", eduOrganizationApi.DeleteEduOrganizationByIds) // 批量删除EduOrganization
		eduOrganizationRouter.PUT("updateEduOrganization", eduOrganizationApi.UpdateEduOrganization)              // 更新EduOrganization
	}
	{
		eduOrganizationRouterWithoutRecord.GET("findEduOrganization", eduOrganizationApi.FindEduOrganization)       // 根据ID获取EduOrganization
		eduOrganizationRouterWithoutRecord.GET("getEduOrganizationList", eduOrganizationApi.GetEduOrganizationList) // 获取EduOrganization列表
	}
}
