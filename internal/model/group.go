package model

import (
	"time"

	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/types/timex"

	"user/internal/model/entity"
)

// Group 群组明细
type Group struct {
	Id           int64      `json:"id" comment:"群组ID"`
	Code         string     `json:"code" comment:"群组编码"`
	Name         string     `json:"name" comment:"群组名称"`
	Remark       string     `json:"remark" comment:"备注"`
	CreateUserId int64      `json:"createUserId" comment:"创建人"`
	CreateTime   timex.Time `json:"createTime" comment:"创建时间"`
	UpdateUserId int64      `json:"updateUserId" comment:"更新人"`
	UpdateTime   timex.Time `json:"updateTime" comment:"更新时间"`
}

// GroupPage 群组分页查询参数
type GroupPage struct {
	Keyword string `json:"keyword" comment:"关键字"`
	*modelx.Page
}

// GroupSave 群组信息新增
type GroupSave struct {
	Id         int64        `json:"id" comment:"群组ID"`
	Code       string       `json:"code" comment:"群组编码"`
	Name       string       `json:"name" comment:"群组名称"`
	Remark     string       `json:"remark" comment:"备注"`
	UserList   []*GroupUser `json:"userList" comment:"群组成员列表"`
	RoleList   []*GroupRole `json:"roleList" comment:"群组角色列表"`
	CurrUserId int64        `json:"currUserId" comment:"当前用户"`
}

func (g *GroupSave) Group() *entity.Group {
	return &entity.Group{
		Id:           g.Id,
		Code:         g.Code,
		Name:         g.Name,
		Remark:       g.Remark,
		CreateUserId: g.CurrUserId,
		UpdateUserId: g.CurrUserId,
		UpdateTime:   time.Now(),
	}
}

// GroupUserAdd 群组成员新增
type GroupUserAdd struct {
	GroupId    int64        `json:"groupId" comment:"群组ID"`
	UserList   []*GroupUser `json:"userList" comment:"新增群组成员列表"`
	RoleList   []*GroupRole `json:"roleList" comment:"新增群组角色列表"`
	CurrUserId int64        `json:"currUserId" comment:"当前用户ID"`
}

// GroupRoleAdd 群组角色新增
type GroupRoleAdd struct {
	Id         int64        `json:"id" comment:"群组ID"`
	RoleList   []*GroupRole `json:"roleList" comment:"新增群组角色列表"`
	CurrUserId int64        `json:"currUserId" comment:"当前用户ID"`
}

// GroupDetail 群组明细
type GroupDetail struct {
	Group    *Group       `json:"group"  comment:"群组基本信息"`
	RoleList []*GroupRole `json:"roleList" comment:"群组角色"`
	UserList []*GroupUser `json:"userList" comment:"群组成员"`
}

// GroupUser 群组用户
type GroupUser struct {
	Id         int64  `json:"id" comment:"用户ID"`
	Account    string `json:"account" comment:"用户账号"`
	Name       string `json:"name" comment:"用户姓名"`
	IsAdmin    bool   `json:"isAdmin" comment:"是否管理员"`
	ValidStart string `json:"validStart" comment:"有效期始"`
	ValidEnd   string `json:"validEnd" comment:"有效期止"`
	Remark     string `json:"remark" comment:"备注"`
}

// GroupRole 群组角色
type GroupRole struct {
	Id         int64  `json:"id" comment:"角色ID"`
	Code       string `json:"code" comment:"角色编码"`
	Name       string `json:"name" comment:"角色名称"`
	ValidStart string `json:"validStart" comment:"有效期始"`
	ValidEnd   string `json:"validEnd" comment:"有效期止"`
	Remark     string `json:"remark" comment:"备注"`
}
