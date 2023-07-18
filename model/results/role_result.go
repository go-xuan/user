package results

import (
	"github.com/quanxiaoxuan/quanx/utils/timex"
)

// 角色信息
type RoleInfoList []*RoleInfo
type RoleInfo struct {
	RoleId       int64      `json:"roleId" comment:"角色ID"`
	RoleCode     string     `json:"roleCode" comment:"角色编码"`
	RoleName     string     `json:"roleName" comment:"角色名称"`
	Remark       string     `json:"remark" comment:"备注"`
	CreateUserId int64      `json:"createUserId" comment:"创建人"`
	CreateTime   timex.Time `json:"createTime" comment:"创建时间"`
	UpdateUserId int64      `json:"updateUserId" comment:"更新人"`
	UpdateTime   timex.Time `json:"updateTime" comment:"更新时间"`
}

// 角色详情
type RoleDetail struct {
	RoleInfo       *RoleInfo      `json:"roleInfo" comment:"角色信息"`
	RoleMemberList RoleMemberList `json:"roleMemberList" comment:"角色成员列表"`
}

// 角色成员列表
type RoleMemberList []*RoleMember
type RoleMember struct {
	UserId     int64      `json:"userId" comment:"用户ID"`
	UserName   string     `json:"userName" comment:"姓名"`
	ValidStart timex.Time `json:"validStart" comment:"有效期始"`
	ValidEnd   timex.Time `json:"validEnd" comment:"有效期止"`
	Remark     string     `json:"remark" comment:"备注"`
}

// 角色简要信息
type RoleSimpleList []*RoleSimple
type RoleSimple struct {
	RoleId   int64  `json:"roleId" comment:"角色ID"`
	RoleCode string `json:"roleCode" comment:"角色编码"`
	RoleName string `json:"roleName" comment:"角色名称"`
}
