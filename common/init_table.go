package common

import (
	"github.com/quanxiaoxuan/go-builder/gormx"
	"quan-admin/model/entity"
)

// 初始化数据库表结构
func InitGormTable() {
	PGDB := gormx.Ctl.DB
	// 用户表
	if PGDB.Migrator().HasTable(&entity.SysUser{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysUser{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysUser{})
	}
	// 用户鉴权表
	if PGDB.Migrator().HasTable(&entity.SysUserAuth{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysUserAuth{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysUserAuth{})
	}
	// 角色表
	if PGDB.Migrator().HasTable(&entity.SysRole{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysRole{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysRole{})
	}
	// 用户角色表
	if PGDB.Migrator().HasTable(&entity.SysRoleMember{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysRoleMember{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysRoleMember{})
	}

	// 群组表
	if PGDB.Migrator().HasTable(&entity.SysGroup{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysGroup{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysGroup{})
	}
	// 群组成员
	if PGDB.Migrator().HasTable(&entity.SysGroupMember{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysGroupMember{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysGroupMember{})
	}
	// 群组所有角色
	if PGDB.Migrator().HasTable(&entity.SysGroupRole{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysGroupRole{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysGroupRole{})
	}

	// 系统日志
	if PGDB.Migrator().HasTable(&entity.SysLog{}) {
		_ = PGDB.Migrator().AutoMigrate(&entity.SysLog{})
	} else {
		_ = PGDB.Migrator().CreateTable(&entity.SysLog{})
	}
}
