package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/core/ginx"

	"user/internal/controller"
)

// BindGinRouter 绑定路由
func BindGinRouter(router *gin.RouterGroup) {
	// 鉴权方式
	auth := ginx.AuthValidate().Token
	//auth := ginx.AuthValidate().SetCookie

	router.POST("login", ginx.CheckIP, controller.UserLogin) // 用户登录
	router.GET("logout", controller.UserLogout)              // 用户登出
	router.GET("encrypt", controller.Encrypt)                // 密码加密
	router.GET("decrypt", controller.Decrypt)                // 密码解密
	router.GET("check", auth, controller.CheckLogin)         // 校验登录
	// 用户管理
	user := router.Group("user").Use(auth)
	user.POST("page", controller.UserPage)    // 用户分页
	user.GET("list", controller.UserList)     // 用户列表
	user.POST("save", controller.UserSave)    // 用户保存
	user.GET("delete", controller.UserDelete) // 用户删除
	user.GET("detail", controller.UserDetail) // 用户明细
	// 角色管理
	role := router.Group("role").Use(auth)
	role.POST("page", controller.RolePage)    // 角色分页
	role.GET("list", controller.RoleList)     // 角色列表
	role.POST("save", controller.RoleSave)    // 角色保存
	role.GET("delete", controller.RoleDelete) // 角色删除
	role.GET("detail", controller.RoleDetail) // 角色明细
	// 群组管理
	group := router.Group("group").Use(auth)
	group.POST("page", controller.GroupPage)    // 群组分页
	group.POST("save", controller.GroupSave)    // 群组保存
	group.GET("delete", controller.GroupDelete) // 群组删除
	group.GET("detail", controller.GroupDetail) // 群组明细
}
