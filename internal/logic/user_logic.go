package logic

import (
	"errors"
	"github.com/quanxiaoxuan/quanx/utils/encryptx"
	"github.com/quanxiaoxuan/quanx/utils/idx"
	"quan-admin/model"
	"time"

	"github.com/quanxiaoxuan/quanx/common/respx"
	"github.com/quanxiaoxuan/quanx/utils/randx"
	"github.com/quanxiaoxuan/quanx/utils/timex"
	log "github.com/sirupsen/logrus"

	"quan-admin/internal/dao"
	"quan-admin/model/table"
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
	resp = respx.BuildPageResp(in.Page.Page, resultList, total)
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
	userId = idx.SnowFlake().NewInt64()
	var user = table.User{
		Id:           userId,
		Account:      in.Account,
		Name:         in.Name,
		Phone:        in.Phone,
		Gender:       in.Gender,
		Birthday:     in.Birthday,
		Email:        in.Email,
		Address:      in.Address,
		Remark:       in.Remark,
		CreateUserId: in.CurrUserId,
		UpdateUserId: in.CurrUserId,
	}
	passWord := in.Password
	if len(passWord) < 32 {
		passWord = encryptx.MD5(passWord)
	}
	salt := randx.UUID()
	validStart := time.Now()
	if in.ValidStart != "" {
		validStart = timex.ToTime(in.ValidStart)
	}
	validEnd := validStart.AddDate(1, 0, 0)
	if in.ValidEnd != "" {
		validEnd = timex.ToTime(in.ValidEnd)
	}
	var userAuth = table.UserAuth{
		UserId:       userId,
		Salt:         salt,
		Password:     encryptx.PasswordSalt(passWord, salt),
		SessionTime:  3600,
		ValidStart:   validStart,
		ValidEnd:     validEnd,
		CreateUserId: in.CurrUserId,
		UpdateUserId: in.CurrUserId,
	}
	err = dao.UserCreate(user, userAuth)
	if err != nil {
		log.Error("用户新增失败")
		return
	}
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
