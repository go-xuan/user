package logic

import (
	"context"
	"errors"
	"github.com/go-xuan/quanx/importx/ginx"
	"strconv"
	"time"

	"github.com/go-xuan/quanx/importx/redisx"
	"github.com/go-xuan/quanx/utilx/encryptx"
	"github.com/go-xuan/quanx/utilx/snowflakex"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/table"
)

// 用户登录
func UserLogin(param model.Login, loginIp string) (result *model.LoginResult, err error) {
	var user *model.User
	user, err = dao.GetUserByName(param.Username)
	if err != nil {
		return
	}
	var userAuth *table.UserAuth
	userAuth, err = dao.QueryUserAuth(user.Id)
	if err != nil {
		return
	}
	var password string
	password, err = encryptx.RSA().Decrypt(param.Password)
	if err != nil {
		return
	}
	password = encryptx.MD5(password)
	// 校验密码
	if encryptx.PasswordSalt(password, userAuth.Salt) != userAuth.Password {
		err = errors.New("密码错误")
		return
	}
	var authUser = &ginx.User{
		Id:         strconv.FormatInt(user.Id, 10),
		Username:   user.Account,
		Name:       user.Name,
		Phone:      user.Phone,
		LoginIp:    loginIp,
		ExpireTime: time.Now().Unix() + userAuth.SessionTime,
	}
	var token string
	token, err = ginx.GetTokenByUser(authUser)
	if err != nil {
		log.Error("生成token失败")
		return
	}
	// token存入redis
	authUser.SetTokenCache(token, time.Duration(userAuth.SessionTime)*time.Second)
	// 记录日志
	var sysLog = table.Log{
		Id:           snowflakex.New().Int64(),
		Module:       "auth",
		Type:         "login",
		Content:      user.Name + "【" + user.Phone + "账号密码登录】",
		Ip:           loginIp,
		CreateUserId: userAuth.UserId,
	}
	err = dao.SysLogCreate(sysLog)
	if err != nil {
		log.Error("记录登录日志失败")
	}
	result = &model.LoginResult{User: user, Token: token}
	return
}

// 用户登出
func UserLogout(user *ginx.User, ip string) (userId int64, err error) {
	userId, _ = strconv.ParseInt(user.Id, 10, 64)
	var sysLog = table.Log{
		Id:           snowflakex.New().Int64(),
		Module:       "auth",
		Type:         "logout",
		Content:      user.Name + "【登出】",
		Ip:           ip,
		CreateUserId: userId,
	}
	err = dao.SysLogCreate(sysLog)
	if err != nil {
		log.Error("记录日志失败")
	}
	// 删除redis上用户token
	redisx.GetCmd("user").Del(context.TODO(), user.RedisKey())
	return
}
