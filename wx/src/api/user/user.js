import { apiResquest } from "../../utils/http";
/**
 * 通过角色ID和组织ID获取用户列表
 * @param {页码} page 
 * @param {每页数量} pageSize 
 * @param {角色ID} roleId 
 * @param {组织ID} orgId 
 * @returns 
 */
export const getUserListByRoleIdAndOrgId = (param) => {
  return apiResquest({
    url: '/user/getUserListByRoleIdAndOrgId',
    method: 'POST',
    data: param
  })
}

export const createUser = (name, course, totalSessions, remainingSessions, phone, roleId) => {
  const possible = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz"
  let text = ""
  for (var i = 0; i < 6; i++)
    text += possible.charAt(Math.floor(Math.random() * possible.length));

  var user = uni.getStorageSync('user')
  const param = {
    "username": text,
    "password": "123123",
    "phone": phone,
    "nickName": name,
    "headerImg": "https://qmplusimg.henrongyi.top/1576554439myAvatar.png",
    "authorityId": roleId,
    "authorityIds": [
      roleId
    ],
    "enable": 1,
    "userName": text,
    "eduOrganizationID": user.eduOrganizationID,
    "course": course,
    "totalSessions": Number(totalSessions),
    "remainingSessions": Number(remainingSessions)
  }
  return apiResquest({
    url: '/user/student_register',
    method: 'POST',
    data: param
  })
}

/**
 * 更新用户信息
 * @param {*} param 
 * @returns 
 */
export const setUserInfo = (param) => {
  return apiResquest({
    url: '/user/setUserInfo',
    method: 'PUT',
    data: param
  })
}

/**
 * 获取用户信息
 * @param {*} param 
 * @returns 
 */
export const getUserInfoById = (param) => {
  return apiResquest({
    url: '/user/getUserInfoById',
    method: 'GET',
    data: param
  })
}
/**
 * 删除学生
 * @param {学生ID} param 
 * @returns 
 */
export const deleteUser = (param) => {
  return apiResquest({
    url: '/user/deleteUser',
    method: 'DELETE',
    data: param
  })
}
