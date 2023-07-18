package logic

import (
	"github.com/quanxiaoxuan/quanx/common/respx"
	"github.com/quanxiaoxuan/quanx/common/snowflakex"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 用户分页查询
func RolePage(param params.RolePage) (*respx.PageResponse, error) {
	var resultList results.RoleInfoList
	var total int64
	var err error
	resultList, total, err = mapper.RolePage(param)
	if err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	}
	return respx.BuildPageData(param.PageParam, resultList, total), nil
}

// 角色列表
func RoleList() (results.RoleSimpleList, error) {
	return mapper.RoleList()
}

// 角色成员校验
func RoleCodeExist(roleCode string) (exist bool, err error) {
	var count int64
	count, err = mapper.RoleCodeExist(roleCode)
	if err != nil {
		return
	}
	exist = count > 0
	return
}

// 角色新增
func RoleAdd(param params.RoleCreate) (int64, error) {
	var err error
	roleId := snowflakex.Generator.NewId()
	var role = entity.SysRole{
		RoleId:       roleId,
		RoleCode:     param.RoleCode,
		RoleName:     param.RoleName,
		Remark:       param.Remark,
		CreateUserId: param.CreateUserId,
		UpdateUserId: param.CreateUserId,
	}
	err = mapper.RoleAdd(role)
	if err != nil {
		log.Error("角色新增失败")
		return 0, err
	}
	return roleId, nil
}

// 角色修改
func RoleUpdate(param params.RoleUpdate) error {
	err := mapper.RoleUpdate(param)
	if err != nil {
		log.Error("角色修改失败")
		return err
	}
	return nil
}

// 角色删除
func RoleDelete(roleId int64) error {
	return mapper.RoleDelete(roleId)
}

// 角色详情
func RoleDetail(roleId int64) (detail results.RoleDetail, err error) {
	// 查询角色信息
	var roleInfo results.RoleInfo
	roleInfoChan := make(chan results.RoleInfo)
	go func() {
		var info results.RoleInfo
		info, err = mapper.RoleInfo(roleId)
		if err != nil {
			log.Error("角色信息查询失败")
		}
		roleInfoChan <- info
	}()
	roleInfo = <-roleInfoChan
	detail.RoleInfo = &roleInfo

	// 查询角色成员列表
	roleUserListChan := make(chan results.RoleMemberList)
	go func() {
		var list results.RoleMemberList
		list, err = mapper.RoleMemberList(roleId)
		if err != nil {
			log.Error("用户角色列表失败")
		}
		roleUserListChan <- list
	}()
	detail.RoleMemberList = <-roleUserListChan

	return
}

// 角色成员校验
func RoleMemberExist(roleId int64, userIds []int64) (exist bool, err error) {
	var count int64
	count, err = mapper.RoleMemberCount(roleId, userIds)
	if err != nil {
		return
	}
	exist = count > 0
	return
}
