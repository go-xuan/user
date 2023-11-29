package logic

import (
	"errors"

	"github.com/go-xuan/quanx/common/respx"
	"github.com/go-xuan/quanx/utils/idx"
	"github.com/go-xuan/quanx/utils/timex"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/table"
)

// 群组分页
func GroupPage(in model.GroupPage) (resp *respx.PageResponse, err error) {
	var rows []*model.Group
	var total int64
	rows, total, err = dao.GroupPage(in)
	if err != nil {
		log.Error("群组分页查询失败")
		return
	}
	resp = respx.BuildPageResp(in.Page.Page, rows, total)
	return
}

// 群组编码校验
func GroupExist(in *model.GroupSave) (err error) {
	var count int64
	count, err = dao.GroupExist(in)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("此群组编码已存在")
	}
	return
}

// 群组新增
func GroupCreate(in *model.GroupSave) (id int64, err error) {
	if err = GroupExist(in); err != nil {
		return
	}
	id = idx.SnowFlake().NewInt64()
	in.Id = id
	err = dao.GroupCreate(in.Group())
	if err != nil {
		log.Error("群组新增失败")
		return
	}
	if len(in.UserList) > 0 {
		err = GroupUserAdd(in)
		if err != nil {
			log.Error("群组用户新增失败")
			return
		}
	}
	if len(in.RoleList) > 0 {
		err = GroupRoleAdd(in)
		if err != nil {
			log.Error("群组用户新增失败")
			return
		}
	}
	return
}

// 群组修改
func GroupUpdate(in *model.GroupSave) (id int64, err error) {
	err = dao.GroupUpdate(in)
	if err != nil {
		log.Error("群组修改失败")
		return
	}
	id = in.Id
	return
}

// 群组删除
func GroupDelete(groupId int64) error {
	return dao.GroupDelete(groupId)
}

// 群组明细
func GroupDetail(id int64) (detail model.GroupDetail, err error) {
	// 查询群组信息
	var groupInfo model.Group
	groupChan := make(chan model.Group)
	go func() {
		var group model.Group
		group, err = dao.GetGroup(id)
		if err != nil {
			log.Error("角色信息查询失败")
		}
		groupChan <- group
	}()
	groupInfo = <-groupChan
	detail.Group = &groupInfo

	// 查询群组成员列表
	userListChan := make(chan []*model.GroupUser)
	go func() {
		var list []*model.GroupUser
		list, err = dao.GroupUserList(id)
		if err != nil {
			log.Error("用户角色列表失败")
		}
		userListChan <- list
	}()
	detail.UserList = <-userListChan

	// 查询群组角色列表
	roleListChan := make(chan []*model.GroupRole)
	go func() {
		var list []*model.GroupRole
		list, err = dao.GroupRoleList(id)
		if err != nil {
			log.Error("用户角色列表失败")
		}
		roleListChan <- list
	}()
	detail.RoleList = <-roleListChan

	return
}

// 群组成员校验
func GroupUserExist(id int64, userIds []int64) (err error) {
	var count int64
	count, err = dao.GroupUserCount(id, userIds)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("新增群组成员已存在")
	}
	return
}

// 群组成员新增
func GroupUserAdd(in *model.GroupSave) (err error) {
	var createList []*table.GroupUser
	for _, item := range in.UserList {
		var create = table.GroupUser{
			Id:           idx.SnowFlake().NewInt64(),
			GroupId:      in.Id,
			UserId:       item.Id,
			IsAdmin:      item.IsAdmin,
			ValidStart:   timex.ToTime(item.ValidStart),
			ValidEnd:     timex.ToTime(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: in.CurrUserId,
			UpdateUserId: in.CurrUserId,
		}
		createList = append(createList, &create)
	}
	err = dao.GroupUserCreateBatch(createList)
	if err != nil {
		return
	}
	return
}

// 群组角色校验
func GroupRoleExist(groupId int64, roleIds []int64) (err error) {
	var count int64
	count, err = dao.GroupRoleCount(groupId, roleIds)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("新增群组角色已存在")
	}
	return
}

// 群组角色新增
func GroupRoleAdd(in *model.GroupSave) (err error) {
	var createList []*table.GroupRole
	for _, item := range in.RoleList {
		var create = table.GroupRole{
			Id:           idx.SnowFlake().NewInt64(),
			GroupId:      in.Id,
			RoleId:       item.Id,
			ValidStart:   timex.ToTime(item.ValidStart),
			ValidEnd:     timex.ToTime(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: in.CurrUserId,
			UpdateUserId: in.CurrUserId,
		}
		createList = append(createList, &create)
	}
	return dao.GroupRoleCreateBatch(createList)
}
