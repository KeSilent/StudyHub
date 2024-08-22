import service from '@/utils/request'

// @Tags EduOrganization
// @Summary 创建EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduOrganization true "创建EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduOrganization/createEduOrganization [post]
export const createEduOrganization = (data) => {
  return service({
    url: '/eduOrganization/createEduOrganization',
    method: 'post',
    data
  })
}

// @Tags EduOrganization
// @Summary 删除EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduOrganization true "删除EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduOrganization/deleteEduOrganization [delete]
export const deleteEduOrganization = (data) => {
  return service({
    url: '/eduOrganization/deleteEduOrganization',
    method: 'delete',
    data
  })
}

// @Tags EduOrganization
// @Summary 删除EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduOrganization/deleteEduOrganization [delete]
export const deleteEduOrganizationByIds = (data) => {
  return service({
    url: '/eduOrganization/deleteEduOrganizationByIds',
    method: 'delete',
    data
  })
}

// @Tags EduOrganization
// @Summary 更新EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduOrganization true "更新EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduOrganization/updateEduOrganization [put]
export const updateEduOrganization = (data) => {
  return service({
    url: '/eduOrganization/updateEduOrganization',
    method: 'put',
    data
  })
}

// @Tags EduOrganization
// @Summary 用id查询EduOrganization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EduOrganization true "用id查询EduOrganization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduOrganization/findEduOrganization [get]
export const findEduOrganization = (params) => {
  return service({
    url: '/eduOrganization/findEduOrganization',
    method: 'get',
    params
  })
}

// @Tags EduOrganization
// @Summary 分页获取EduOrganization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EduOrganization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduOrganization/getEduOrganizationList [get]
export const getEduOrganizationList = (params) => {
  return service({
    url: '/eduOrganization/getEduOrganizationList',
    method: 'get',
    params
  })
}
