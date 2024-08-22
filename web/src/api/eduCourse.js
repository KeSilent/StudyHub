import service from '@/utils/request'

// @Tags EduCourse
// @Summary 创建EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduCourse true "创建EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduCourse/createEduCourse [post]
export const createEduCourse = (data) => {
  return service({
    url: '/eduCourse/createEduCourse',
    method: 'post',
    data
  })
}

// @Tags EduCourse
// @Summary 删除EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduCourse true "删除EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduCourse/deleteEduCourse [delete]
export const deleteEduCourse = (data) => {
  return service({
    url: '/eduCourse/deleteEduCourse',
    method: 'delete',
    data
  })
}

// @Tags EduCourse
// @Summary 删除EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduCourse/deleteEduCourse [delete]
export const deleteEduCourseByIds = (data) => {
  return service({
    url: '/eduCourse/deleteEduCourseByIds',
    method: 'delete',
    data
  })
}

// @Tags EduCourse
// @Summary 更新EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduCourse true "更新EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduCourse/updateEduCourse [put]
export const updateEduCourse = (data) => {
  return service({
    url: '/eduCourse/updateEduCourse',
    method: 'put',
    data
  })
}

// @Tags EduCourse
// @Summary 用id查询EduCourse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EduCourse true "用id查询EduCourse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduCourse/findEduCourse [get]
export const findEduCourse = (params) => {
  return service({
    url: '/eduCourse/findEduCourse',
    method: 'get',
    params
  })
}

// @Tags EduCourse
// @Summary 分页获取EduCourse列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EduCourse列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduCourse/getEduCourseList [get]
export const getEduCourseList = (params) => {
  return service({
    url: '/eduCourse/getEduCourseList',
    method: 'get',
    params
  })
}
