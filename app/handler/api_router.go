package handler

import (
	"github.com/gin-gonic/gin"
)

// 添加api接口函数路由
func AddHandlerFunc(router *gin.RouterGroup) {
	// 用户登录
	authRouter := router.Group("auth")
	authRouter.POST("login", UserLogin)      // 用户登录
	authRouter.GET("logout", UserLogout)     // 用户登出
	authRouter.GET("tokenParse", TokenParse) // token解析
	// 用户管理
	userRouter := router.Group("user")
	userRouter.POST("page", UserPage)     // 用户分页
	userRouter.GET("list", UserList)      // 用户列表
	userRouter.POST("add", UserAdd)       // 用户新增
	userRouter.POST("update", UserUpdate) // 用户修改
	userRouter.GET("delete", UserDelete)  // 用户删除
	userRouter.GET("detail", UserDetail)  // 用户明细
	// 角色管理
	roleRouter := router.Group("role")
	roleRouter.POST("page", RolePage)     // 角色分页
	roleRouter.GET("list", RoleList)      // 角色列表
	roleRouter.POST("add", RoleAdd)       // 角色新增
	roleRouter.POST("update", RoleUpdate) // 角色修改
	roleRouter.GET("delete", RoleDelete)  // 角色删除
	roleRouter.GET("detail", RoleDetail)  // 角色明细
	// 角色成员
	roleMemberRouter := roleRouter.Group("member")
	roleMemberRouter.POST("add", RoleMemberAdd)    // 角色成员新增
	roleMemberRouter.POST("delete", RoleMemberAdd) // 角色成员删除
	// 群组管理
	groupRouter := router.Group("group")
	groupRouter.POST("page", GroupPage)     // 群组分页
	groupRouter.POST("add", GroupAdd)       // 群组新增
	groupRouter.POST("update", GroupUpdate) // 群组修改
	groupRouter.GET("delete", GroupDelete)  // 群组删除
	groupRouter.GET("detail", GroupDetail)  // 群组明细
	// 群组成员
	groupMemberRouter := groupRouter.Group("member")
	groupMemberRouter.POST("add", GroupMemberAdd)    // 群组成员新增
	groupMemberRouter.POST("delete", GroupMemberAdd) // 群组成员新增
	// 群组角色
	groupRoleRouter := groupRouter.Group("role")
	groupRoleRouter.POST("add", GroupRoleAdd)    // 群组角色新增
	groupRoleRouter.POST("delete", GroupRoleAdd) // 群组角色新增
	// 日志管理
	logRouter := router.Group("log")
	logRouter.POST("add", LogAdd)       // 群组新增
	logRouter.POST("delete", LogDelete) // 群组新增
}
