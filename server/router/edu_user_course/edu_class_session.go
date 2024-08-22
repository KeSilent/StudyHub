/*
 * @Author: Yang
 * @Date: 2023-04-10 18:33:08
 * @Description: 学生交易记录
 */
package edu_user_course

import (
	v1 "github.com/KeSilent/study-hub/server/api/v1"
	"github.com/KeSilent/study-hub/server/middleware"
	"github.com/gin-gonic/gin"
)

type EduClassSessionRouter struct {
}

// InitEduClassSessionRouter 初始化 EduClassSession 路由信息
func (s *EduClassSessionRouter) InitEduClassSessionRouter(Router *gin.RouterGroup) {
	eduClassSessionRouter := Router.Group("eduClassSession").Use(middleware.OperationRecord())
	eduClassSessionRouterWithoutRecord := Router.Group("eduClassSession")
	var eduClassSessionApi = v1.ApiGroupApp.Edu_user_courseApiGroup.EduClassSessionApi
	{
		eduClassSessionRouter.POST("createEduClassSession", eduClassSessionApi.CreateEduClassSession)             // 新建EduClassSession
		eduClassSessionRouter.DELETE("deleteEduClassSession", eduClassSessionApi.DeleteEduClassSession)           // 删除EduClassSession
		eduClassSessionRouter.DELETE("deleteEduClassSessionByIds", eduClassSessionApi.DeleteEduClassSessionByIds) // 批量删除EduClassSession
		eduClassSessionRouter.PUT("updateEduClassSession", eduClassSessionApi.UpdateEduClassSession)              // 更新EduClassSession
	}
	{
		eduClassSessionRouterWithoutRecord.GET("findEduClassSession", eduClassSessionApi.FindEduClassSession)                                 // 根据ID获取EduClassSession
		eduClassSessionRouterWithoutRecord.GET("getEduClassSessionList", eduClassSessionApi.GetEduClassSessionList)                           // 获取EduClassSession列表
		eduClassSessionRouterWithoutRecord.GET("getEduClassSessionListByUser", eduClassSessionApi.GetEduClassSessionListByUser)               // 获取用户的消费列表
		eduClassSessionRouterWithoutRecord.GET("getStudentsWithLessThanFiveSessions", eduClassSessionApi.GetStudentsWithLessThanFiveSessions) // 获取剩余课时少于5节的学生列表
	}
}
