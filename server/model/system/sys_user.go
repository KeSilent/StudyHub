/*
 * @Author: Yang
 * @Date: 2023-04-10 18:22:12
 * @Description: 用户模型
 */
package system

import (
	"github.com/KeSilent/study-hub/server/global"
	"github.com/KeSilent/study-hub/server/model/edu_organization"
	"github.com/KeSilent/study-hub/server/model/edu_user_course"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID              uuid.UUID                        `json:"uuid" gorm:"index;comment:用户UUID"`    // 用户UUID
	Username          string                           `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	Password          string                           `json:"-"  gorm:"comment:用户登录密码"`            // 用户登录密码
	WXOpenid          string                           `json:"-"  gorm:"comment:微信Openid"`          // 用户登录密码
	WXSessionKey      string                           `json:"-"  gorm:"comment:微信SessionKey"`
	EduOrganizationID uint                             `json:"eduOrganizationID" gorm:"comment:组织ID"`
	EduOrganization   edu_organization.EduOrganization `json:"eduOrganization" gorm:"foreignKey:ID;references:EduOrganizationID;comment:用户昵称"`       // 用户组织
	NickName          string                           `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode          string                           `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg         string                           `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor         string                           `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor       string                           `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId       uint                             `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority         SysAuthority                     `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities       []SysAuthority                   `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone             string                           `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email             string                           `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable            int                              `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	EduEnrollments    []edu_user_course.EduEnrollment  `json:"eduEnrollment" gorm:"foreignKey:UserId"`          // 用户科目
}

func (SysUser) TableName() string {
	return "sys_users"
}
