package logic

import (
	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/os/errorx"
	"github.com/go-xuan/quanx/utils/idx"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
)

// UserPage 用户分页
func UserPage(in model.UserPage) (*respx.PageResponse, error) {
	if rows, total, err := dao.UserPage(in); err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	} else {
		return respx.BuildPageResp(in.Page, rows, total), nil
	}
}

// UserList 用户列表
func UserList() ([]*model.User, error) {
	return dao.UserList()
}

// UserExist 用户校验
func UserExist(in *model.UserSave) error {
	if count, err := dao.UserExist(in); err != nil {
		return err
	} else if count > 0 {
		return errorx.Errorf("此手机号(%s)已被使用", in.Phone)
	}
	return nil
}

// UserCreate 用户新增
func UserCreate(in *model.UserSave) (userId int64, err error) {
	if err = UserExist(in); err != nil {
		return
	}
	in.Id = idx.SnowFlake().Int64()
	if err = dao.UserCreate(in.UserCreate()); err != nil {
		log.Error("用户新增失败")
		return
	}
	userId = in.Id
	return
}

// UserUpdate 用户修改
func UserUpdate(in *model.UserSave) error {
	return dao.UserUpdate(in)
}

// UserDelete 用户删除
func UserDelete(userId int64) error {
	return dao.UserDelete(userId)
}

// UserDetail 用户明细查询
func UserDetail(id int64) (*model.User, error) {
	return dao.QueryUser(id)
}
