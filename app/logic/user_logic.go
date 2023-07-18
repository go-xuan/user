package logic

import (
	"time"

	"github.com/quanxiaoxuan/quanx/common/respx"
	"github.com/quanxiaoxuan/quanx/common/snowflakex"
	"github.com/quanxiaoxuan/quanx/utils/randx"
	"github.com/quanxiaoxuan/quanx/utils/stringx"
	"github.com/quanxiaoxuan/quanx/utils/timex"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 用户分页
func UserPage(param params.UserPage) (*respx.PageResponse, error) {
	var resultList results.UserInfoList
	var total int64
	var err error
	resultList, total, err = mapper.UserPage(param)
	if err != nil {
		log.Error("用户分页查询失败")
		return nil, err
	}
	return respx.BuildPageData(param.PageParam, resultList, total), nil
}

// 用户列表
func UserList() (results.UserSimpleList, error) {
	return mapper.UserList()
}

// 用户手机校验
func UserPhoneExist(phone string) (exist bool, err error) {
	var count int64
	count, err = mapper.UserPhoneExist(phone)
	if err != nil {
		return
	}
	exist = count > 0
	return
}

// 用户新增
func UserAdd(param params.UserCreate) (int64, error) {
	var err error
	userId := snowflakex.Generator.NewId()
	var user = entity.SysUser{
		UserId:       userId,
		UserName:     param.UserName,
		Phone:        param.Phone,
		Gender:       param.Gender,
		Birthday:     param.Birthday,
		Email:        param.Email,
		Address:      param.Address,
		Remark:       param.Remark,
		CreateUserId: param.CreateUserId,
		UpdateUserId: param.CreateUserId,
	}
	passWord := param.Password
	if len(passWord) < 32 {
		passWord = stringx.MD5(passWord)
	}
	salt := randx.UUID()
	validStart := time.Now()
	if param.ValidStart != "" {
		validStart = timex.ToTime(param.ValidStart)
	}
	validEnd := validStart.AddDate(1, 0, 0)
	if param.ValidEnd != "" {
		validEnd = timex.ToTime(param.ValidEnd)
	}
	var userAuth = entity.SysUserAuth{
		UserId:       userId,
		Salt:         salt,
		Password:     stringx.PasswordSalt(passWord, salt),
		SessionTime:  3600,
		ValidStart:   validStart,
		ValidEnd:     validEnd,
		CreateUserId: param.CreateUserId,
		UpdateUserId: param.CreateUserId,
	}
	err = mapper.UserAdd(user, userAuth)
	if err != nil {
		log.Error("用户新增失败")
		return 0, err
	}
	return userId, nil
}

// 用户修改
func UserUpdate(param params.UserUpdate) error {
	err := mapper.UserUpdate(param)
	if err != nil {
		log.Error("用户修改失败")
		return err
	}
	return nil
}

// 用户删除
func UserDelete(userId int64) error {
	return mapper.UserDelete(userId)
}

// 用户明细查询
func UserDetail(userId int64) (*results.UserDetail, error) {
	var detail results.UserDetail
	var err error
	var userInfo entity.SysUser
	userInfo, err = mapper.UserSelect(userId)
	if err != nil {
		log.Error("查询用户基本信息失败")
		return nil, err
	}
	detail.UserInfo = &userInfo
	var userAuth entity.SysUserAuth
	userAuth, err = mapper.QueryUserAuth(userId)
	if err != nil {
		log.Error("查询用户身份信息失败")
		return nil, err
	}
	detail.UserAuth = &userAuth
	return &detail, nil
}
