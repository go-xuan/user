package model

import (
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/types/timex"
	"user/internal/model/entity"
)

// Role 角色信息
type Role struct {
	Id           int64      `json:"id" comment:"角色ID"`
	Code         string     `json:"code" comment:"角色编码"`
	Name         string     `json:"name" comment:"角色名称"`
	Remark       string     `json:"remark" comment:"备注"`
	CreateUserId int64      `json:"createUserId" comment:"创建人"`
	UpdateUserId int64      `json:"updateUserId" comment:"更新人"`
	CreateTime   timex.Time `json:"createTime" comment:"创建时间"`
	UpdateTime   timex.Time `json:"updateTime" comment:"更新时间"`
}

// RolePage 角色分页参数
type RolePage struct {
	Keyword string `json:"keyword" comment:"关键字"`
	*modelx.Page
}

// RoleSave 角色保存
type RoleSave struct {
	Id         int64       `json:"id" comment:"角色ID"`
	Code       string      `json:"code" comment:"角色编码"`
	Name       string      `json:"name" comment:"角色名称"`
	Remark     string      `json:"remark" comment:"备注"`
	UserList   []*RoleUser `json:"userList" comment:"新增用户列表"`
	CurrUserId int64       `json:"currUserId" comment:"当前用户ID"`
}

func (r *RoleSave) Role() *entity.Role {
	return &entity.Role{
		Id:           r.Id,
		Code:         r.Code,
		Name:         r.Name,
		Remark:       r.Remark,
		CreateUserId: r.CurrUserId,
		UpdateUserId: r.CurrUserId,
	}
}

// RoleUser 角色成员列表
type RoleUser struct {
	Id         int64  `json:"id" comment:"用户ID"`
	Name       string `json:"name" comment:"姓名"`
	ValidStart string `json:"validStart" comment:"有效期始"`
	ValidEnd   string `json:"validEnd" comment:"有效期止"`
	Remark     string `json:"remark" comment:"备注"`
}

// RoleDetail 角色详情
type RoleDetail struct {
	Role     *Role       `json:"role" comment:"角色信息"`
	UserList []*RoleUser `json:"userList" comment:"角色成员列表"`
}
