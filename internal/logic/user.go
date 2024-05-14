package logic

import (
	"errors"

	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/utils/idx"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
)

// 用户分页
func UserPage(in model.UserPage) (*respx.PageResponse, error) {
	if rows, total, err := dao.UserPage(in); err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	} else {
		return respx.BuildPageResp(in.Page, rows, total), nil
	}
}

// 用户列表
func UserList() ([]*model.User, error) {
	return dao.UserList()
}

// 用户校验
func UserExist(in *model.UserSave) error {
	if count, err := dao.UserExist(in); err != nil {
		return err
	} else if count > 0 {
		return errors.New("此账号/手机号已被使用")
	}
	return nil
}

// 用户新增
func UserCreate(in *model.UserSave) (userId int64, err error) {
	if err = UserExist(in); err != nil {
		return
	}
	in.Id = idx.SnowFlake().Int64()
	if err = dao.UserCreate(in.UserCreate(), in.UserAuthCreate()); err != nil {
		log.Error("用户新增失败")
		return
	}
	userId = in.Id
	return
}

// 用户修改
func UserUpdate(in *model.UserSave) error {
	return dao.UserUpdate(in)
}

// 用户删除
func UserDelete(userId int64) error {
	return dao.UserDelete(userId)
}

// 用户明细查询
func UserDetail(id int64) (*model.User, error) {
	return dao.GetUserById(id)
}
