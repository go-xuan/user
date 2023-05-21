package params

import (
	"github.com/quanxiaoxuan/go-builder/model/request"
)

// 群组分页查询参数
type GroupPage struct {
	SearchKey string        `json:"searchKey" comment:"关键字"`
	PageParam *request.Page `json:"pageParam" comment:"分页参数"`
}

// 群组信息新增
type GroupCreate struct {
	GroupCode    string `json:"groupCode" comment:"群组编码"`
	GroupName    string `json:"groupName" comment:"群组名称"`
	Remark       string `json:"remark" comment:"备注"`
	CreateUserId int64  `json:"createUserId" comment:"创建人"`
}

// 群组信息修改
type GroupUpdate struct {
	GroupId      int64  `json:"groupId" comment:"群组ID"`
	GroupName    string `json:"groupName" comment:"群组名称"`
	Remark       string `json:"remark" comment:"备注"`
	UpdateUserId int64  `json:"updateUserId" comment:"更新人"`
}

// 群组成员新增
type GroupMemberAdd struct {
	GroupId         int64           `json:"groupId" comment:"群组ID"`
	GroupMemberList GroupMemberList `json:"groupMemberList" comment:"新增群组成员列表"`
	CreateUserId    int64           `json:"createUserId" comment:"创建人"`
}

// 新增群组成员列表
type GroupMemberList []*GroupMember
type GroupMember struct {
	UserId     int64  `json:"userId" comment:"用户ID"`
	IsAdmin    bool   `json:"isAdmin" comment:"是否管理员"`
	ValidStart string `json:"validStart" comment:"有效期始"`
	ValidEnd   string `json:"validEnd" comment:"有效期止"`
	Remark     string `json:"remark" comment:"备注"`
}

// 群组角色新增
type GroupRoleAdd struct {
	GroupId       int64         `json:"groupId" comment:"群组ID"`
	GroupRoleList GroupRoleList `json:"groupRoleList" comment:"新增群组角色列表"`
	CreateUserId  int64         `json:"createUserId" comment:"创建人"`
}

// 新增群组角色列表
type GroupRoleList []*GroupRole
type GroupRole struct {
	RoleId     int64  `json:"roleId" comment:"角色ID"`
	ValidStart string `json:"validStart" comment:"有效期始"`
	ValidEnd   string `json:"validEnd" comment:"有效期止"`
	Remark     string `json:"remark" comment:"备注"`
}
