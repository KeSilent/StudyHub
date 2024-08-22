/*
 * @Author: Yang
 * @Date: 2023-04-11 16:26:12
 * @Description: 登录
 */
import { apiResquest } from "../../utils/http";

/**
 * 微信登录
 * @param {登陆Code} param 
 * @returns 
 */
export const wxLogin = (param) => {
  return apiResquest({
    url: '/base/wxlogin',
    method: 'POST',
    data: { code: param }
  })
}

/**
 * 通过手机号码绑定微信账号
 * @param {微信code} code 
 * @param {手机号码} phone 
 * @returns 
 */
export const wxLoginForPhone = (code, phone) => {
  return apiResquest({
    url: '/base/wxLoginForPhone',
    method: 'POST',
    data: { code: code, phone: phone }
  })
}