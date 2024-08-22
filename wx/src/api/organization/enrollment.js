import { apiResquest } from "@/utils/http";

/**
 * 更新学生选择的课程
 * @param {*} param 
 * @returns 
 */
export const updateEduEnrollment = (param) => {
  return apiResquest({
    url: '/eduEnrollment/updateCourseId',
    method: 'PUT',
    data: param
  })
}

/**
 * 消耗课时
 * @param {*} param 
 * @returns 
 */
export const consumptionClass = (param) => {
  return apiResquest({
    url: '/eduEnrollment/consumptionClass',
    method: 'POST',
    data: param
  })
}

/**
 * 增加课时
 * @param {*} param 
 * @returns 
 */
export const addSession = (param) => {
  return apiResquest({
    url: '/eduEnrollment/addSession',
    method: 'POST',
    data: param
  })
}

/**
 * 获取日志列表
 * @param {*} param 
 * @returns 
 */
export const getEduClassSessionList = (param) => {
  return apiResquest({
    url: '/eduClassSession/getEduClassSessionListByUser',
    method: 'GET',
    data: param
  })
}

/**
 * 获取剩余课时少于5节的学生列表
 * @returns 
 */
export const getStudentsWithLessThanFiveSessions = () => {
  return apiResquest({
    url: '/eduClassSession/getStudentsWithLessThanFiveSessions',
    method: 'GET'
  })
}