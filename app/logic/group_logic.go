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

// 群组分页
func GroupPage(param params.GroupPage) (*respx.PageResponse, error) {
	var resultList results.GroupInfoList
	var total int64
	var err error
	resultList, total, err = mapper.GroupPage(param)
	if err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	}
	return respx.BuildPageData(param.PageParam, resultList, total), nil
}

// 群组编码校验
func GroupCodeExist(groupCode string) (exist bool, err error) {
	var count int64
	count, err = mapper.GroupCodeCount(groupCode)
	if err != nil {
		return
	}
	exist = count > 0
	return
}

// 群组新增
func GroupAdd(param params.GroupCreate) (int64, error) {
	var err error
	groupId := snowflakex.Generator.NewId()
	var group = entity.SysGroup{
		GroupId:      groupId,
		GroupCode:    param.GroupCode,
		GroupName:    param.GroupName,
		Remark:       param.Remark,
		CreateUserId: param.CreateUserId,
		UpdateUserId: param.CreateUserId,
	}
	err = mapper.GroupAdd(group)
	if err != nil {
		log.Error("群组新增失败")
		return 0, err
	}
	return groupId, nil
}

// 群组修改
func GroupUpdate(param params.GroupUpdate) error {
	err := mapper.GroupUpdate(param)
	if err != nil {
		log.Error("群组修改失败")
		return err
	}
	return nil
}

// 群组删除
func GroupDelete(groupId int64) error {
	return mapper.GroupDelete(groupId)
}

// 群组明细
func GroupDetail(groupId int64) (detail results.GroupDetail, err error) {
	// 查询群组信息
	var groupInfo results.GroupInfo
	infoChan := make(chan results.GroupInfo)
	go func() {
		var info results.GroupInfo
		info, err = mapper.GroupInfo(groupId)
		if err != nil {
			log.Error("角色信息查询失败")
		}
		infoChan <- info
	}()
	groupInfo = <-infoChan
	detail.GroupInfo = &groupInfo

	// 查询群组成员列表
	memberListChan := make(chan results.GroupMemberList)
	go func() {
		var list results.GroupMemberList
		list, err = mapper.GroupMemberList(groupId)
		if err != nil {
			log.Error("用户角色列表失败")
		}
		memberListChan <- list
	}()
	detail.GroupMemberList = <-memberListChan

	// 查询群组角色列表
	roleListChan := make(chan results.GroupRoleList)
	go func() {
		var list results.GroupRoleList
		list, err = mapper.GroupRoleList(groupId)
		if err != nil {
			log.Error("用户角色列表失败")
		}
		roleListChan <- list
	}()
	detail.GroupRoleList = <-roleListChan

	return
}

// 群组成员校验
func GroupMemberExist(groupId int64, userIds []int64) (exist bool, err error) {
	var count int64
	count, err = mapper.GroupMemberCount(groupId, userIds)
	if err != nil {
		return
	}
	exist = count > 0
	return
}
