package edu_organization

import (
	"github.com/KeSilent/study-hub/server/api/v1"
	"github.com/KeSilent/study-hub/server/middleware"
	"github.com/gin-gonic/gin"
)

type EduCourseRouter struct {
}

// InitEduCourseRouter 初始化 EduCourse 路由信息
func (s *EduCourseRouter) InitEduCourseRouter(Router *gin.RouterGroup) {
	eduCourseRouter := Router.Group("eduCourse").Use(middleware.OperationRecord())
	eduCourseRouterWithoutRecord := Router.Group("eduCourse")
	var eduCourseApi = v1.ApiGroupApp.Edu_organizationApiGroup.EduCourseApi
	{
		eduCourseRouter.POST("createEduCourse", eduCourseApi.CreateEduCourse)             // 新建EduCourse
		eduCourseRouter.DELETE("deleteEduCourse", eduCourseApi.DeleteEduCourse)           // 删除EduCourse
		eduCourseRouter.DELETE("deleteEduCourseByIds", eduCourseApi.DeleteEduCourseByIds) // 批量删除EduCourse
		eduCourseRouter.PUT("updateEduCourse", eduCourseApi.UpdateEduCourse)              // 更新EduCourse
	}
	{
		eduCourseRouterWithoutRecord.GET("findEduCourse", eduCourseApi.FindEduCourse)       // 根据ID获取EduCourse
		eduCourseRouterWithoutRecord.GET("getEduCourseList", eduCourseApi.GetEduCourseList) // 获取EduCourse列表
	}
}
