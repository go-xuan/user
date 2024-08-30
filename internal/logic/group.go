package logic

import (
	"sync"

	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/os/errorx"
	"github.com/go-xuan/quanx/types/timex"
	"github.com/go-xuan/quanx/utils/idx"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/entity"
)

// GroupPage 群组分页
func GroupPage(in model.GroupPage) (*respx.PageResponse, error) {
	if rows, total, err := dao.GroupPage(in); err != nil {
		log.Error("群组分页查询失败")
		return nil, err
	} else {
		return respx.BuildPageResp(in.Page, rows, total), nil
	}

}

// GroupExist 群组编码校验
func GroupExist(in *model.GroupSave) (err error) {
	var count int64
	if count, err = dao.GroupExist(in); err != nil {
		return
	}
	if count > 0 {
		err = errorx.New("此群组编码已存在")
	}
	return
}

// GroupCreate 群组新增
func GroupCreate(in *model.GroupSave) (id int64, err error) {
	if err = GroupExist(in); err != nil {
		return
	}
	id = idx.SnowFlake().Int64()
	in.Id = id
	if err = dao.GroupCreate(in.Group()); err != nil {
		log.Error("群组新增失败")
		return
	}
	if len(in.UserList) > 0 {
		if err = GroupUserAdd(in); err != nil {
			log.Error("群组用户新增失败")
			return
		}
	}
	if len(in.RoleList) > 0 {
		if err = GroupRoleAdd(in); err != nil {
			log.Error("群组用户新增失败")
			return
		}
	}
	return
}

// GroupUpdate 群组修改
func GroupUpdate(in *model.GroupSave) (id int64, err error) {
	if err = dao.GroupUpdate(in); err != nil {
		log.Error("群组修改失败")
		return
	}
	id = in.Id
	return
}

// GroupDelete 群组删除
func GroupDelete(groupId int64) error {
	return dao.GroupDelete(groupId)
}

// GroupDetail 群组明细
func GroupDetail(id int64) (*model.GroupDetail, error) {
	// 查询群组信息
	var wg sync.WaitGroup
	wg.Add(3)
	groupChan := make(chan *model.Group)
	userListChan := make(chan []*model.GroupUser)
	roleListChan := make(chan []*model.GroupRole)

	var err error
	// 查询角色信息
	go func() {
		wg.Done()
		var group *model.Group
		if group, err = dao.GetGroup(id); err != nil {
			log.Error("角色信息查询失败")
		}
		groupChan <- group
	}()

	// 查询群组成员列表
	go func() {
		wg.Done()
		var list []*model.GroupUser
		if list, err = dao.GroupUserList(id); err != nil {
			log.Error("用户角色列表失败")
		}
		userListChan <- list
	}()

	// 查询群组角色列表
	go func() {
		wg.Done()
		var list []*model.GroupRole
		if list, err = dao.GroupRoleList(id); err != nil {
			log.Error("用户角色列表失败")
		}
		roleListChan <- list
	}()

	detail := &model.GroupDetail{
		Group:    <-groupChan,
		UserList: <-userListChan,
		RoleList: <-roleListChan,
	}
	wg.Wait()
	return detail, err
}

// GroupUserExist 群组成员校验
func GroupUserExist(id int64, userIds []int64) error {
	if count, err := dao.GroupUserCount(id, userIds); err != nil {
		return err
	} else if count > 0 {
		return errorx.New("新增群组成员已存在")
	}
	return nil
}

// GroupUserAdd 群组成员新增
func GroupUserAdd(in *model.GroupSave) error {
	var createList []*entity.GroupUser
	for _, item := range in.UserList {
		var create = entity.GroupUser{
			Id:           idx.SnowFlake().Int64(),
			GroupId:      in.Id,
			UserId:       item.Id,
			IsAdmin:      item.IsAdmin,
			ValidStart:   timex.Parse(item.ValidStart),
			ValidEnd:     timex.Parse(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: in.CurrUserId,
			UpdateUserId: in.CurrUserId,
		}
		createList = append(createList, &create)
	}
	return dao.GroupUserCreateBatch(createList)
}

// GroupRoleExist 群组角色校验
func GroupRoleExist(groupId int64, roleIds []int64) error {
	if count, err := dao.GroupRoleCount(groupId, roleIds); err != nil {
		return err
	} else if count > 0 {
		return errorx.New("新增群组角色已存在")
	}
	return nil
}

// GroupRoleAdd 群组角色新增
func GroupRoleAdd(in *model.GroupSave) (err error) {
	var createList []*entity.GroupRole
	for _, item := range in.RoleList {
		var create = entity.GroupRole{
			Id:           idx.SnowFlake().Int64(),
			GroupId:      in.Id,
			RoleId:       item.Id,
			ValidStart:   timex.Parse(item.ValidStart),
			ValidEnd:     timex.Parse(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: in.CurrUserId,
			UpdateUserId: in.CurrUserId,
		}
		createList = append(createList, &create)
	}
	return dao.GroupRoleCreateBatch(createList)
}
