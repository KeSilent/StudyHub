/*
 * @Author: Yang
 * @Date: 2023-03-19 22:23:28
 * @Description: 系统操作类扩展
 */
package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/edu_user_course"
	"github.com/KeSilent/study-hub/server/model/system"
	systemReq "github.com/KeSilent/study-hub/server/model/system/request"
	"github.com/KeSilent/study-hub/server/model/system/response"
	"github.com/KeSilent/study-hub/server/plugin/wx"
	"github.com/KeSilent/study-hub/server/service/edu_organization"
	"github.com/KeSilent/study-hub/server/utils"
	"gorm.io/gorm"
)

type UserExtendService struct {
	UserService
}

/**
 * @description: 微信登陆
 * @param {system.SysUser} u
 * @return {*}
 */
func (userService *UserService) WXLogin(u systemReq.WXLoginReq) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	wxKey, err := wx.GetWXSession(u.Code)
	if err != nil {
		return userInter, err
	}

	userInter, err = SelectUser(wxKey.OpenId, wxKey.SessionKey)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//没有注册
		return userInter, errors.New("没有注册")
	} else {
		global.GVA_DB.Model(&userInter).Where("wx_openid = ?", wxKey.OpenId).Updates(system.SysUser{WXSessionKey: wxKey.SessionKey})
	}
	userInter, err = SelectUser(wxKey.OpenId, wxKey.SessionKey)

	MenuServiceApp.UserAuthorityDefaultRouter(userInter)

	return userInter, err
}

/**
 * @Description: 查询用户
 * @param {string} openId
 * @param {string} sessionKey
 * @return {*}
 */
func SelectUser(openId string, _ string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	err = global.GVA_DB.Preload("EduOrganization").Where("wx_openid = ?", openId).Preload("Authorities").Preload("Authority").First(&user).Error
	return &user, err
}

/**
 * @description: 获取微信OpenId
 * @param {string} code
 * @return {*}
 */
func GetWXSession(code string) (*response.WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, global.WXAppId, global.WXSecret, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := response.WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		errCodeStr := strconv.Itoa(wxResp.ErrCode)
		return nil, fmt.Errorf("ErrCode:%s  ErrMsg:%s", errCodeStr, wxResp.ErrMsg)
	}

	return &wxResp, nil
}

/**
 * @Description: 导入并注册用户
 * @param {systemReq.ImportUserInfoReq} u
 * @return {*}
 */
func (userService *UserService) ImportRegister(u systemReq.ImportUserInfoReq) (err error) {
	service := &edu_organization.EduCourseService{}
	courseList, err := service.GetEduCourseInfoListByOrganId(u.EduOrganizationID)

	rs := utils.EnalysisExcel(u.ImportFile, courseList)

	for _, r := range rs {

		var authorities []system.SysAuthority
		for _, v := range u.AuthorityIds {
			authorities = append(authorities, system.SysAuthority{
				AuthorityId: v,
			})
		}

		user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email, EduOrganizationID: u.EduOrganizationID}

		eduEnrollment := edu_user_course.EduEnrollment{}
		eduEnrollment.CourseId = &r.Course
		eduEnrollment.TotalSessions = &r.TotalSessions
		eduEnrollment.RemainingSessions = &r.RemainingSessions

		userService.RegisterStudent(*user, &eduEnrollment)
	}

	return err
}

/**
 * @description: 手机号码绑定微信登陆
 * @param {system.SysUser} u
 * @return {*}
 */
func (userService *UserService) WXLoginForPhone(login systemReq.WXLoginReq) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	err = global.GVA_DB.Preload("EduOrganization").Where("phone = ?", login.Phone).Preload("Authorities").Preload("Authority").First(&userInter).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//没有注册
		return userInter, errors.New("没有注册")
	} else {
		//如果通过手机号查询到用户后，则绑定微信信息

		wxKey, err := wx.GetWXSession(login.Code)
		if err != nil {
			return userInter, err
		}
		global.GVA_DB.Model(&userInter).Where("phone = ?", login.Phone).Updates(system.SysUser{WXSessionKey: wxKey.SessionKey, WXOpenid: wxKey.OpenId})

		userInter, err = SelectUser(wxKey.OpenId, wxKey.SessionKey)

		MenuServiceApp.UserAuthorityDefaultRouter(userInter)
	}

	return userInter, err
}
