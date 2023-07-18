package params

import (
	"github.com/quanxiaoxuan/quanx/common/paramx"
)

// 角色分页参数
type RolePage struct {
	SearchKey string       `json:"searchKey" comment:"关键字"`
	PageParam *paramx.Page `json:"pageParam" comment:"分页参数"`
}

// 角色新增
type RoleCreate struct {
	RoleId       int64  `json:"roleId" comment:"角色ID"`
	RoleCode     string `json:"roleCode" comment:"角色编码"`
	RoleName     string `json:"roleName" comment:"角色名称"`
	Remark       string `json:"remark" comment:"备注"`
	CreateUserId int64  `json:"createUserId" comment:"创建人"`
}

// 角色新增
type RoleUpdate struct {
	RoleId       int64  `json:"roleId" comment:"角色ID"`
	RoleCode     string `json:"roleCode" comment:"角色编码"`
	RoleName     string `json:"roleName" comment:"角色名称"`
	Remark       string `json:"remark" comment:"备注"`
	UpdateUserId int64  `json:"updateUserId" comment:"更新人"`
}

// 角色成员新增
type RoleMemberAdd struct {
	RoleId         int64          `json:"roleId" comment:"角色ID"`
	RoleMemberList RoleMemberList `json:"roleMemberList" comment:"新增角色成员列表"`
	CreateUserId   int64          `json:"createUserId" comment:"创建人"`
}

// 新增角色成员列表
type RoleMemberList []*RoleMember
type RoleMember struct {
	UserId     int64  `json:"userId" comment:"用户ID"`
	ValidStart string `json:"validStart" comment:"账号有效期始"`
	ValidEnd   string `json:"validEnd" comment:"账号有效期止"`
	Remark     string `json:"remark" comment:"备注"`
}
