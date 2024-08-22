import { apiResquest } from "@/utils/http";

/**
 * 获取组织列表
 * @param {page=1&pageSize=10} param 
 * @returns 
 */
export const getEduOrganizationList = (param) => {
  return apiResquest({
    url: '/base/getOrganizationList',
    method: 'GET',
    data: param
  })
}