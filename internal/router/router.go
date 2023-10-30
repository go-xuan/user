package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/quanx/public/middlewarex"

	"quan-user/internal/controller"
)

// 添加api接口函数路由
func BindGinRouter(router *gin.RouterGroup) {
	api := router.Group("api/v1")
	api.Use(
		middlewarex.WhiteList("/auth/login", "/user/save"), // 白名单
		middlewarex.Cooke(), // cookie鉴权
		//middlewarex.Auth(),  // auth鉴权
	)
	// 用户登录
	auth := api.Group("auth")
	auth.POST("login", controller.UserLogin)      // 用户登录
	auth.GET("logout", controller.UserLogout)     // 用户登出
	auth.GET("tokenParse", controller.TokenParse) // token解析
	auth.GET("encrypt", controller.Encrypt)       // 密码加密
	// 用户管理
	user := api.Group("user")
	user.POST("page", controller.UserPage)    // 用户分页
	user.POST("save", controller.UserSave)    // 用户保存
	user.GET("delete", controller.UserDelete) // 用户删除
	user.GET("detail", controller.UserDetail) // 用户明细
	user.GET("list", controller.UserList)     // 用户列表
	// 角色管理
	role := api.Group("role")
	role.POST("page", controller.RolePage)    // 角色分页
	role.GET("list", controller.RoleList)     // 角色列表
	role.POST("save", controller.RoleSave)    // 角色保存
	role.GET("delete", controller.RoleDelete) // 角色删除
	role.GET("detail", controller.RoleDetail) // 角色明细
	// 群组管理
	group := api.Group("group")
	group.POST("page", controller.GroupPage)    // 群组分页
	group.POST("save", controller.GroupSave)    // 群组保存
	group.GET("delete", controller.GroupDelete) // 群组删除
	group.GET("detail", controller.GroupDetail) // 群组明细
}
