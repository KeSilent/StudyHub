import service from '@/utils/request'

// @Tags EduClassSession
// @Summary 创建EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduClassSession true "创建EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduClassSession/createEduClassSession [post]
export const createEduClassSession = (data) => {
  return service({
    url: '/eduClassSession/createEduClassSession',
    method: 'post',
    data
  })
}

// @Tags EduClassSession
// @Summary 删除EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduClassSession true "删除EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduClassSession/deleteEduClassSession [delete]
export const deleteEduClassSession = (data) => {
  return service({
    url: '/eduClassSession/deleteEduClassSession',
    method: 'delete',
    data
  })
}

// @Tags EduClassSession
// @Summary 删除EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eduClassSession/deleteEduClassSession [delete]
export const deleteEduClassSessionByIds = (data) => {
  return service({
    url: '/eduClassSession/deleteEduClassSessionByIds',
    method: 'delete',
    data
  })
}

// @Tags EduClassSession
// @Summary 更新EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EduClassSession true "更新EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eduClassSession/updateEduClassSession [put]
export const updateEduClassSession = (data) => {
  return service({
    url: '/eduClassSession/updateEduClassSession',
    method: 'put',
    data
  })
}

// @Tags EduClassSession
// @Summary 用id查询EduClassSession
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EduClassSession true "用id查询EduClassSession"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eduClassSession/findEduClassSession [get]
export const findEduClassSession = (params) => {
  return service({
    url: '/eduClassSession/findEduClassSession',
    method: 'get',
    params
  })
}

// @Tags EduClassSession
// @Summary 分页获取EduClassSession列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EduClassSession列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eduClassSession/getEduClassSessionList [get]
export const getEduClassSessionList = (params) => {
  return service({
    url: '/eduClassSession/getEduClassSessionList',
    method: 'get',
    params
  })
}
