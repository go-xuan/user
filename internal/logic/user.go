package logic

import (
	"errors"

	"github.com/go-xuan/quanx/frame/snowflakex"
	"github.com/go-xuan/quanx/net/respx"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
)

// 用户分页
func UserPage(in model.UserPage) (resp *respx.PageResponse, err error) {
	var resultList []*model.User
	var total int64
	resultList, total, err = dao.UserPage(in)
	if err != nil {
		log.Error("用户分页查询失败")
		return
	}
	resp = respx.BuildPageResp(in.Page, resultList, total)
	return
}

// 用户列表
func UserList() ([]*model.User, error) {
	return dao.UserList()
}

// 用户校验
func UserExist(in *model.UserSave) (err error) {
	var count int64
	count, err = dao.UserExist(in)
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("此账号/手机号已被使用")
	}
	return
}

// 用户新增
func UserCreate(in *model.UserSave) (userId int64, err error) {
	if err = UserExist(in); err != nil {
		return
	}
	in.Id = snowflakex.New().Int64()
	err = dao.UserCreate(in.UserCreate(), in.UserAuthCreate())
	if err != nil {
		log.Error("用户新增失败")
		return
	}
	userId = in.Id
	return
}

// 用户修改
func UserUpdate(in *model.UserSave) error {
	err := dao.UserUpdate(in)
	if err != nil {
		log.Error("用户修改失败")
		return err
	}
	return nil
}

// 用户删除
func UserDelete(userId int64) error {
	return dao.UserDelete(userId)
}

// 用户明细查询
func UserDetail(id int64) (user *model.User, err error) {
	user, err = dao.GetUserById(id)
	if err != nil {
		log.Error("查询用户基本信息失败")
		return
	}
	return
}
