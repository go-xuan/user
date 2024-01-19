package logic

import (
	"errors"

	"github.com/go-xuan/quanx/commonx/respx"
	"github.com/go-xuan/quanx/utilx/snowflakex"
	"github.com/go-xuan/quanx/utilx/timex"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/entity"
)

// 用户分页查询
func RolePage(in model.RolePage) (*respx.PageResponse, error) {
	var list []*model.Role
	var total int64
	var err error
	list, total, err = dao.RolePage(in)
	if err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	}
	return respx.BuildPageResp(in.Page.Page, list, total), nil
}

// 角色列表
func RoleList() ([]*model.Role, error) {
	return dao.RoleList()
}

// 角色校验
func RoleExist(in *model.RoleSave) (err error) {
	var count int64
	count, err = dao.RoleExist(in)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("此角色编码已存在")
	}
	return
}

// 角色新增
func RoleCreate(in *model.RoleSave) (roleId int64, err error) {
	if err = RoleExist(in); err != nil {
		return
	}
	roleId = snowflakex.New().Int64()
	in.Id = roleId
	var role = in.Role()
	err = dao.RoleCreate(role)
	if err != nil {
		return
	}
	if len(in.UserList) > 0 {
		err = RoleUserAdd(in)
		if err != nil {
			return
		}
	}
	return
}

// 角色修改
func RoleUpdate(in *model.RoleSave) error {
	err := dao.RoleUpdate(in)
	if err != nil {
		log.Error("角色修改失败")
		return err
	}
	return nil
}

// 角色删除
func RoleDelete(id int64) error {
	return dao.RoleDelete(id)
}

// 角色详情
func RoleDetail(id int64) (detail model.RoleDetail, err error) {
	// 查询角色信息
	var role model.Role
	roleChan := make(chan model.Role)
	go func() {
		var info model.Role
		info, err = dao.RoleInfo(id)
		if err != nil {
			log.Error("角色信息查询失败")
		}
		roleChan <- info
	}()
	role = <-roleChan
	detail.Role = &role

	// 查询角色成员列表
	roleUserListChan := make(chan []*model.RoleUser)
	go func() {
		var list []*model.RoleUser
		list, err = dao.RoleUserList(id)
		if err != nil {
			log.Error("用户角色列表失败")
		}
		roleUserListChan <- list
	}()
	detail.UserList = <-roleUserListChan

	return
}

// 角色成员校验
func RoleUserExist(roleId int64, userIds []int64) (err error) {
	var count int64
	count, err = dao.RoleUserCount(roleId, userIds)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("新增角色成员已存在")
	}
	return
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
	err = dao.RoleUserCreateBatch(createList)
	if err != nil {
		return
	}
	return
}
