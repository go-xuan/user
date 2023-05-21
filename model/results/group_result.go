package results

import (
	"github.com/quanxiaoxuan/go-utils/timex"
)

// 群组明细
type GroupInfoList []*GroupInfo
type GroupInfo struct {
	GroupId      int64      `json:"groupId" comment:"群组ID"`
	GroupCode    string     `json:"groupCode" comment:"群组编码"`
	GroupName    string     `json:"groupName" comment:"群组名称"`
	Remark       string     `json:"remark" comment:"备注"`
	CreateUserId int64      `json:"createUserId" comment:"创建人"`
	CreateTime   timex.Time `json:"createTime" comment:"创建时间"`
	UpdateUserId int64      `json:"updateUserId" comment:"更新人"`
	UpdateTime   timex.Time `json:"updateTime" comment:"更新时间"`
}

// 群组明细
type GroupDetail struct {
	GroupInfo       *GroupInfo      `json:"groupInfo"  comment:"群组基本信息"`
	GroupRoleList   GroupRoleList   `json:"groupRoleList" comment:"群组角色"`
	GroupMemberList GroupMemberList `json:"groupUserList" comment:"群组成员"`
}

// 群组成员
type GroupMemberList []*GroupMember
type GroupMember struct {
	UserId     int64      `json:"userId" comment:"用户ID"`
	UserName   string     `json:"userName" comment:"姓名"`
	IsAdmin    bool       `json:"isAdmin" comment:"是否管理员"`
	ValidStart timex.Time `json:"validStart" comment:"有效期始"`
	ValidEnd   timex.Time `json:"validEnd" comment:"有效期止"`
	Remark     string     `json:"remark" comment:"备注"`
}

// 群组角色
type GroupRoleList []*GroupRole
type GroupRole struct {
	RoleId     int64      `json:"roleId" comment:"角色ID"`
	RoleCode   string     `json:"roleCode" comment:"角色编码"`
	RoleName   string     `json:"roleName" comment:"角色名称"`
	ValidStart timex.Time `json:"validStart" comment:"有效期始"`
	ValidEnd   timex.Time `json:"validEnd" comment:"有效期止"`
	Remark     string     `json:"remark" comment:"备注"`
}
