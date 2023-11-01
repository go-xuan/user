package logic

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/quanxiaoxuan/quanx/public/authx"
	"github.com/quanxiaoxuan/quanx/public/redisx"
	"github.com/quanxiaoxuan/quanx/utils/encryptx"
	"github.com/quanxiaoxuan/quanx/utils/idx"
	log "github.com/sirupsen/logrus"

	"quan-user/internal/dao"
	"quan-user/model"
	"quan-user/model/table"
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
	sessionTime := userAuth.SessionTime
	userId := strconv.FormatInt(user.Id, 10)
	var token string
	token, err = authx.BuildAuthToken(&authx.User{
		Id:         userId,
		Account:    user.Account,
		Name:       user.Name,
		Phone:      user.Phone,
		LoginIp:    loginIp,
		ExpireTime: time.Now().Unix() + sessionTime,
	})
	if err != nil {
		log.Error("生成token失败")
		return
	}
	// token存入redis
	redisx.GetCmd("user").Set(context.TODO(), authx.RedisKeyPrefix+user.Account, token, time.Duration(sessionTime*1e9))
	var sysLog = table.Log{
		Id:           idx.SnowFlake().NewInt64(),
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
func UserLogout(user *authx.User, ip string) (userId int64, err error) {
	userId, _ = strconv.ParseInt(user.Id, 10, 64)
	var sysLog = table.Log{
		Id:           idx.SnowFlake().NewInt64(),
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
	redisx.GetCmd("user").Del(context.TODO(), user.RedisCacheKey())
	return
}
