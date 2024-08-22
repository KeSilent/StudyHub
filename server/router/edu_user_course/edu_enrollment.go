/*
 * @Author: Yang
 * @Date: 2023-04-10 18:35:56
 * @Description: 用户课程信息
 */
package edu_user_course

import (
	v1 "github.com/KeSilent/study-hub/server/api/v1"
	"github.com/KeSilent/study-hub/server/middleware"
	"github.com/gin-gonic/gin"
)

type EduEnrollmentRouter struct {
}

// InitEduEnrollmentRouter 初始化 EduEnrollment 路由信息
func (s *EduEnrollmentRouter) InitEduEnrollmentRouter(Router *gin.RouterGroup) {
	eduEnrollmentRouter := Router.Group("eduEnrollment").Use(middleware.OperationRecord())
	eduEnrollmentRouterWithoutRecord := Router.Group("eduEnrollment")
	var eduEnrollmentApi = v1.ApiGroupApp.Edu_user_courseApiGroup.EduEnrollmentApi
	{
		eduEnrollmentRouter.POST("createEduEnrollment", eduEnrollmentApi.CreateEduEnrollment)             // 新建EduEnrollment
		eduEnrollmentRouter.DELETE("deleteEduEnrollment", eduEnrollmentApi.DeleteEduEnrollment)           // 删除EduEnrollment
		eduEnrollmentRouter.DELETE("deleteEduEnrollmentByIds", eduEnrollmentApi.DeleteEduEnrollmentByIds) // 批量删除EduEnrollment
		eduEnrollmentRouter.PUT("updateEduEnrollment", eduEnrollmentApi.UpdateEduEnrollment)              // 更新EduEnrollment
	}
	{
		eduEnrollmentRouterWithoutRecord.GET("findEduEnrollment", eduEnrollmentApi.FindEduEnrollment)       // 根据ID获取EduEnrollment
		eduEnrollmentRouterWithoutRecord.GET("getEduEnrollmentList", eduEnrollmentApi.GetEduEnrollmentList) // 获取EduEnrollment列表
	}
	{
		eduEnrollmentRouter.PUT("updateCourseId", eduEnrollmentApi.UpdateCourseId)      // 更新课程选择
		eduEnrollmentRouter.POST("consumptionClass", eduEnrollmentApi.ConsumptionClass) // 消耗课时
		eduEnrollmentRouter.POST("addSession", eduEnrollmentApi.AddSession)             // 增加课时
	}
}
