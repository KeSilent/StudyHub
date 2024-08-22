package system

import (
	"context"

	sysModel "github.com/KeSilent/study-hub/server/model/system"
	"github.com/KeSilent/study-hub/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{ApiGroup: "base", Method: "POST", Path: "/base/login", Description: "用户登录(必选)"},

		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出，必选)"},

		{ApiGroup: "系统用户", Method: "DELETE", Path: "/user/deleteUser", Description: "删除用户"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/admin_register", Description: "用户注册"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/getUserList", Description: "获取用户列表"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setUserInfo", Description: "设置用户信息"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setSelfInfo", Description: "设置自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/user/getUserInfo", Description: "获取自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthorities", Description: "设置权限组"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/changePassword", Description: "修改密码（建议选择)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthority", Description: "修改用户角色(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/resetPassword", Description: "重置用户密码"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "创建api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "删除Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "更新Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "获取api列表"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "获取所有api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "获取api详细信息"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "批量删除api"},

		{ApiGroup: "角色", Method: "POST", Path: "/authority/copyAuthority", Description: "拷贝角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/createAuthority", Description: "创建角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/deleteAuthority", Description: "删除角色"},
		{ApiGroup: "角色", Method: "PUT", Path: "/authority/updateAuthority", Description: "更新角色信息"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/getAuthorityList", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/setDataAuthority", Description: "设置角色资源权限"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "更改角色api权限"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表"},

		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addBaseMenu", Description: "新增菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenu", Description: "获取菜单树(必选)"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "删除菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/updateBaseMenu", Description: "更新菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuById", Description: "根据id获取菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuList", Description: "分页获取基础menu列表"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuAuthority", Description: "获取指定角色menu"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系"},

		{ApiGroup: "分片上传", Method: "GET", Path: "/fileUploadAndDownload/findFile", Description: "寻找目标文件（秒传）"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "断点续传"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "断点续传完成"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "上传完成移除文件"},

		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "文件上传示例"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "文件名或者备注编辑"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表"},

		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getServerInfo", Description: "获取服务器信息"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getSystemConfig", Description: "获取配置文件内容"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/setSystemConfig", Description: "设置配置文件内容"},

		{ApiGroup: "客户", Method: "PUT", Path: "/customer/customer", Description: "更新客户"},
		{ApiGroup: "客户", Method: "POST", Path: "/customer/customer", Description: "创建客户"},
		{ApiGroup: "客户", Method: "DELETE", Path: "/customer/customer", Description: "删除客户"},
		{ApiGroup: "客户", Method: "GET", Path: "/customer/customer", Description: "获取单一客户"},
		{ApiGroup: "客户", Method: "GET", Path: "/customer/customerList", Description: "获取客户列表"},

		{ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getDB", Description: "获取所有数据库"},
		{ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getTables", Description: "获取数据库表"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/createTemp", Description: "自动化代码"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/preview", Description: "预览自动化代码"},
		{ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getColumn", Description: "获取所选table的所有字段"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/createPlug", Description: "自动创建插件包"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/installPlugin", Description: "安装插件"},

		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/createPackage", Description: "生成包(package)"},
		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/getPackage", Description: "获取所有包(package)"},
		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/delPackage", Description: "删除包(package)"},

		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/getMeta", Description: "获取meta信息"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/rollback", Description: "回滚自动生成代码"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/getSysHistory", Description: "查询回滚记录"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/delSysHistory", Description: "删除回滚记录"},

		{ApiGroup: "系统字典详情", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容"},
		{ApiGroup: "系统字典详情", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容"},
		{ApiGroup: "系统字典详情", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容"},
		{ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容"},
		{ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表"},

		{ApiGroup: "系统字典", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "新增字典"},
		{ApiGroup: "系统字典", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典"},
		{ApiGroup: "系统字典", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "更新字典"},
		{ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典"},
		{ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表"},

		{ApiGroup: "操作记录", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史"},

		{ApiGroup: "断点续传(插件版)", Method: "POST", Path: "/simpleUploader/upload", Description: "插件版分片上传"},
		{ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证"},
		{ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "发送测试邮件"},
		{ApiGroup: "email", Method: "POST", Path: "/email/emailSend", Description: "发送邮件示例"},

		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "设置按钮权限"},
		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "获取已有按钮权限"},
		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "删除按钮"},

		{ApiGroup: "万用表格", Method: "POST", Path: "/chatGpt/getTable", Description: "通过gpt获取内容"},
		{ApiGroup: "万用表格", Method: "POST", Path: "/chatGpt/createSK", Description: "录入sk"},
		{ApiGroup: "万用表格", Method: "GET", Path: "/chatGpt/getSK", Description: "获取sk"},
		{ApiGroup: "万用表格", Method: "DELETE", Path: "/chatGpt/deleteSK", Description: "删除sk"},

		{ApiGroup: "eduClassSession", Method: "POST", Path: "/eduClassSession/createEduClassSession", Description: "新增课程日志表"},
		{ApiGroup: "eduClassSession", Method: "DELETE", Path: "/eduClassSession/deleteEduClassSession", Description: "删除课程日志表"},
		{ApiGroup: "eduClassSession", Method: "DELETE", Path: "/eduClassSession/deleteEduClassSessionByIds", Description: "批量删除课程日志表"},
		{ApiGroup: "eduClassSession", Method: "PUT", Path: "/eduClassSession/updateEduClassSession", Description: "更新课程日志表"},
		{ApiGroup: "eduClassSession", Method: "GET", Path: "/eduClassSession/findEduClassSession", Description: "根据ID获取课程日志表"},
		{ApiGroup: "eduClassSession", Method: "GET", Path: "/eduClassSession/getEduClassSessionList", Description: "获取课程日志表列表"},
		{ApiGroup: "eduCourse", Method: "POST", Path: "/eduCourse/createEduCourse", Description: "新增课程表"},
		{ApiGroup: "eduCourse", Method: "DELETE", Path: "/eduCourse/deleteEduCourse", Description: "删除课程表"},
		{ApiGroup: "eduCourse", Method: "DELETE", Path: "/eduCourse/deleteEduCourseByIds", Description: "批量删除课程表"},
		{ApiGroup: "eduCourse", Method: "PUT", Path: "/eduCourse/updateEduCourse", Description: "更新课程表"},
		{ApiGroup: "eduCourse", Method: "GET", Path: "/eduCourse/findEduCourse", Description: "根据ID获取课程表"},
		{ApiGroup: "eduCourse", Method: "GET", Path: "/eduCourse/getEduCourseList", Description: "获取课程表列表"},
		{ApiGroup: "eduOrganization", Method: "GET", Path: "/eduOrganization/getEduOrganizationList", Description: "获取用户组织表列表"},
		{ApiGroup: "eduOrganization", Method: "POST", Path: "/eduOrganization/createEduOrganization", Description: "新增用户组织表"},
		{ApiGroup: "eduOrganization", Method: "DELETE", Path: "/eduOrganization/deleteEduOrganization", Description: "删除用户组织表"},
		{ApiGroup: "eduOrganization", Method: "DELETE", Path: "/eduOrganization/deleteEduOrganizationByIds", Description: "批量删除用户组织表"},
		{ApiGroup: "eduOrganization", Method: "PUT", Path: "/eduOrganization/updateEduOrganization", Description: "更新用户组织表"},
		{ApiGroup: "eduEnrollment", Method: "POST", Path: "/eduEnrollment/createEduEnrollment", Description: "新增用户课程表"},
		{ApiGroup: "eduEnrollment", Method: "DELETE", Path: "/eduEnrollment/deleteEduEnrollment", Description: "删除用户课程表"},
		{ApiGroup: "eduEnrollment", Method: "DELETE", Path: "/eduEnrollment/deleteEduEnrollmentByIds", Description: "批量删除用户课程表"},
		{ApiGroup: "eduEnrollment", Method: "PUT", Path: "/eduEnrollment/updateEduEnrollment", Description: "更新用户课程表"},
		{ApiGroup: "eduEnrollment", Method: "GET", Path: "/eduEnrollment/findEduEnrollment", Description: "根据ID获取用户课程表"},
		{ApiGroup: "eduEnrollment", Method: "GET", Path: "/eduEnrollment/getEduEnrollmentList", Description: "获取用户课程表列表"},
		{ApiGroup: "base", Method: "POST", Path: "/base/wxlogin", Description: "微信登录"},
		{ApiGroup: "base", Method: "GET", Path: "/base/getOrganizationList", Description: "获取组织机构列表"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/getUserListByRoleIdAndOrgId", Description: "通过角色ID和组织ID获取用户列表"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/student_register", Description: "注册学生账号"},
		{ApiGroup: "eduEnrollment", Method: "PUT", Path: "/eduEnrollment/updateCourseId", Description: "更新课程选择"},
		{ApiGroup: "eduEnrollment", Method: "POST", Path: "/eduEnrollment/consumptionClass", Description: "消耗课时"},
		{ApiGroup: "eduEnrollment", Method: "POST", Path: "/eduEnrollment/addSession", Description: "增加课时"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/user/getUserInfoById", Description: "通过用户ID获取信息"},
		{ApiGroup: "eduClassSession", Method: "GET", Path: "/eduClassSession/getEduClassSessionListByUser", Description: "获取用户的消费列表"},
		{ApiGroup: "eduClassSession", Method: "GET", Path: "/eduClassSession/getStudentsWithLessThanFiveSessions", Description: "获取剩余课时少于5节的学生列表"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/importExcelUser", Description: "导入用户"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
