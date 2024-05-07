package logic

import (
	"errors"
	"sync"

	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/types/timex"
	"github.com/go-xuan/quanx/utils/snowflakex"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/entity"
)

// 用户分页查询
func RolePage(in model.RolePage) (*respx.PageResponse, error) {
	if rows, total, err := dao.RolePage(in); err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	} else {
		return respx.BuildPageResp(in.Page, rows, total), nil
	}
}

// 角色列表
func RoleList() ([]*model.Role, error) {
	return dao.RoleList()
}

// 角色校验
func RoleExist(in *model.RoleSave) error {
	if count, err := dao.RoleExist(in); err != nil {
		return err
	} else if count > 0 {
		return errors.New("此角色编码已存在")
	}
	return nil
}

// 角色新增
func RoleCreate(in *model.RoleSave) (roleId int64, err error) {
	if err = RoleExist(in); err != nil {
		return
	}
	roleId = snowflakex.New().Int64()
	in.Id = roleId
	if err = dao.RoleCreate(in.Role()); err != nil {
		return
	}
	if len(in.UserList) > 0 {
		if err = RoleUserAdd(in); err != nil {
			return
		}
	}
	return
}

// 角色修改
func RoleUpdate(in *model.RoleSave) error {
	return dao.RoleUpdate(in)
}

// 角色删除
func RoleDelete(id int64) error {
	return dao.RoleDelete(id)
}

// 角色详情
func RoleDetail(id int64) (detail *model.RoleDetail, err error) {

	var wg sync.WaitGroup
	wg.Add(2)
	roleChan := make(chan *model.Role)
	roleUserListChan := make(chan []*model.RoleUser)

	// 查询角色信息
	go func() {
		defer wg.Done()
		var info *model.Role
		if info, err = dao.RoleInfo(id); err != nil {
			log.Error("角色信息查询失败")
		}
		roleChan <- info
	}()
	// 查询角色成员列表
	go func() {
		defer wg.Done()
		var list []*model.RoleUser
		if list, err = dao.RoleUserList(id); err != nil {
			log.Error("用户角色列表失败")
		}
		roleUserListChan <- list
	}()
	detail = &model.RoleDetail{
		Role:     <-roleChan,
		UserList: <-roleUserListChan,
	}
	wg.Wait()
	return
}

// 角色成员校验
func RoleUserExist(roleId int64, userIds []int64) error {
	if count, err := dao.RoleUserCount(roleId, userIds); err != nil {
		return err
	} else if count > 0 {
		return errors.New("新增角色成员已存在")
	}
	return nil
}

func RoleUserAdd(in *model.RoleSave) (err error) {
	var createList []*entity.RoleUser
	for _, item := range in.UserList {
		var create = entity.RoleUser{
			Id:           snowflakex.New().Int64(),
			RoleId:       in.Id,
			UserId:       item.Id,
			ValidStart:   timex.ToTime(item.ValidStart),
			ValidEnd:     timex.ToTime(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: in.CurrUserId,
			UpdateUserId: in.CurrUserId,
		}
		createList = append(createList, &create)
	}
	if err = dao.RoleUserCreateBatch(createList); err != nil {
		return
	}
	return
}
