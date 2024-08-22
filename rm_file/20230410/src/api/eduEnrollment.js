import service from '@/utils/request'

// @Tags EduEnrollment
// @Summary 创建EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduEnrollment true "创建EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduEnrollment/createEduEnrollment [post]
export const createEduEnrollment = (data) => {
  return service({
    url: '/eduEnrollment/createEduEnrollment',
    method: 'post',
    data
  })
}

// @Tags EduEnrollment
// @Summary 删除EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduEnrollment true "删除EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduEnrollment/deleteEduEnrollment [delete]
export const deleteEduEnrollment = (data) => {
  return service({
    url: '/eduEnrollment/deleteEduEnrollment',
    method: 'delete',
    data
  })
}

// @Tags EduEnrollment
// @Summary 删除EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduEnrollment/deleteEduEnrollment [delete]
export const deleteEduEnrollmentByIds = (data) => {
  return service({
    url: '/eduEnrollment/deleteEduEnrollmentByIds',
    method: 'delete',
    data
  })
}

// @Tags EduEnrollment
// @Summary 更新EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduEnrollment true "更新EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduEnrollment/updateEduEnrollment [put]
export const updateEduEnrollment = (data) => {
  return service({
    url: '/eduEnrollment/updateEduEnrollment',
    method: 'put',
    data
  })
}

// @Tags EduEnrollment
// @Summary 用id查询EduEnrollment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EduEnrollment true "用id查询EduEnrollment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduEnrollment/findEduEnrollment [get]
export const findEduEnrollment = (params) => {
  return service({
    url: '/eduEnrollment/findEduEnrollment',
    method: 'get',
    params
  })
}

// @Tags EduEnrollment
// @Summary 分页获取EduEnrollment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EduEnrollment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduEnrollment/getEduEnrollmentList [get]
export const getEduEnrollmentList = (params) => {
  return service({
    url: '/eduEnrollment/getEduEnrollmentList',
    method: 'get',
    params
  })
}
