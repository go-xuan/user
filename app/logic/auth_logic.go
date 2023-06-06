package logic

import (
	"errors"
	"strconv"
	"time"

	"github.com/quanxiaoxuan/go-builder/authx"
	"github.com/quanxiaoxuan/go-builder/redisx"
	"github.com/quanxiaoxuan/go-builder/snowflakex"
	"github.com/quanxiaoxuan/go-utils/stringx"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 用户登录
func UserLogin(param params.UserLogin, loginIp string) (*results.LoginResult, error) {
	var userInfo results.UserInfo
	var err error
	userInfo, err = mapper.QueryUserInfo(param.Phone)
	if err != nil {
		log.Error("查询用户信息失败")
		return nil, err
	}
	if userInfo.UserId == 0 {
		return nil, errors.New("此用户未注册")
	}
	var userAuth entity.SysUserAuth
	userAuth, err = mapper.QueryUserAuth(userInfo.UserId)
	if err != nil {
		log.Error("查询用户信息失败")
		return nil, err
	}
	passWord := param.Password
	if len(passWord) < 32 {
		passWord = stringx.MD5(passWord)
	}
	if stringx.PasswordSalt(passWord, userAuth.Salt) == userAuth.Password {
		sessionTime := userAuth.SessionTime
		userIdStr := strconv.FormatInt(userInfo.UserId, 10)
		var tokenParam = authx.Param{
			UserId:     userIdStr,
			UserName:   userInfo.UserName,
			Phone:      userInfo.Phone,
			LoginIp:    loginIp,
			ExpireTime: time.Now().Unix() + sessionTime,
		}
		var token string
		token, err = authx.BuildAuthToken(&tokenParam)
		if err != nil {
			log.Error("生成token失败")
			return nil, err
		}
		// token存入redis
		redisx.CTL.Set("token_"+userIdStr, token, time.Duration(sessionTime*1e9))
		var sysLog = entity.SysLog{
			Id:           snowflakex.NewSnow.NewId(),
			Module:       "auth",
			Type:         "login",
			Content:      userInfo.UserName + "【" + userInfo.Phone + "账号密码登录】",
			Ip:           loginIp,
			CreateUserId: userAuth.UserId,
		}
		err = mapper.SysLogAdd(sysLog)
		if err != nil {
			log.Error("记录登录日志失败")
		}
		result := results.LoginResult{
			UserInfo: userInfo,
			Token:    token,
		}
		return &result, nil
	} else {
		return nil, errors.New("密码错误")
	}
}

// 验证token
func TokenParse(token string) (*authx.Param, error) {
	if token == "" {
		return nil, errors.New("未携带token")
	}
	tokenParam, err := authx.ParseAuthToken(token)
	if err != nil {
		return nil, errors.New("token解析失败")
	}
	return tokenParam, nil
}

// 用户登出
func UserLogout(token, ip string) (userId int64, err error) {
	var tokenParam *authx.Param
	tokenParam, err = TokenParse(token)
	if err != nil {
		return
	}
	userId, _ = strconv.ParseInt(tokenParam.UserId, 10, 64)
	var sysLog = entity.SysLog{
		Id:           snowflakex.NewSnow.NewId(),
		Module:       "auth",
		Type:         "logout",
		Content:      tokenParam.UserName + "【登出】",
		Ip:           ip,
		CreateUserId: userId,
	}
	err = mapper.SysLogAdd(sysLog)
	if err != nil {
		log.Error("记录登录日志失败")
	}
	// 删除redis上用户token
	redisx.CTL.Del("token_" + tokenParam.UserId)
	return
}
