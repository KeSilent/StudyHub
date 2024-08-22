/*
 * @Author: Yang
 * @Date: 2023-04-10 18:22:12
 * @Description: 请填写简介
 */
package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// 创建一个接受页码、每页大小、角色id和组织ID的结构
type PageInfoByRoleIdAndOrgId struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	RoleId   int    `json:"roleId" form:"roleId"`
	OrgId    int    `json:"orgId" form:"orgId"`
	UserName string `json:"userName" form:"userName"`
	Phone    string `json:"phone" form:"phone"`
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
