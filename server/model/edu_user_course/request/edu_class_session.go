/*
 * @Author: Yang
 * @Date: 2023-04-10 18:33:08
 * @Description: 请填写简介
 */
package request

import (
	"time"

	"github.com/KeSilent/study-hub/server/model/common/request"
	"github.com/KeSilent/study-hub/server/model/edu_user_course"
)

type EduClassSessionSearch struct {
	edu_user_course.EduClassSession
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type EduClassSessionSearchByUser struct {
	UserId   uint `json:"userId" form:"userId"`
	CourseId uint `json:"courseId" form:"courseId"`
}
