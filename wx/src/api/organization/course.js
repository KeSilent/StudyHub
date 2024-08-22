import { apiResquest } from "@/utils/http";

/**
 * 获取课程列表
 * @param {page=1&pageSize=10} param 
 * @returns 
 */
export const getEduCourseList = (param) => {

  var user = uni.getStorageSync('user')
  param = {
    organizationId: user.eduOrganizationID
  }
  return apiResquest({
    url: '/eduCourse/getEduCourseList',
    method: 'GET',
    data: param
  })
}
