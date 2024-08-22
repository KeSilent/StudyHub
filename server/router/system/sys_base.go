/*
 * @Author: Yang
 * @Date: 2023-04-10 18:22:12
 * @Description: 请填写简介
 */
package system

import (
	v1 "github.com/KeSilent/study-hub/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("wxlogin", baseApi.WXLogin)
		baseRouter.POST("wxLoginForPhone", baseApi.WXLoginForPhone)
	}
	return baseRouter
}
