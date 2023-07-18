package common

import (
	"github.com/quanxiaoxuan/quanx/middleware/gormx"
	"quan-admin/model/entity"
)

// 初始化数据库表结构
func InitGormTable() {
	// 用户表
	_ = gormx.CTL.InitTable(&entity.SysUser{})
	// 用户鉴权表
	_ = gormx.CTL.InitTable(&entity.SysUserAuth{})
	// 角色表
	_ = gormx.CTL.InitTable(&entity.SysRole{})
	// 角色成员表
	_ = gormx.CTL.InitTable(&entity.SysRoleMember{})
	// 群组表
	_ = gormx.CTL.InitTable(&entity.SysGroup{})
	// 群组成员表
	_ = gormx.CTL.InitTable(&entity.SysGroupMember{})
	// 群组角色表
	_ = gormx.CTL.InitTable(&entity.SysGroupRole{})
	// 系统日志
	_ = gormx.CTL.InitTable(&entity.SysLog{})
}
