package logic

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/app"
	"github.com/go-xuan/quanx/app/ginx"
	"github.com/go-xuan/quanx/os/errorx"
	"github.com/go-xuan/quanx/utils/idx"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/entity"
)

// UserLogin 用户登录
func UserLogin(ctx *gin.Context, in model.Login, loginIp string) (result *model.LoginResult, err error) {
	var user *model.User
	if user, err = dao.QueryUserByPhone(in.Username); err != nil {
		err = errorx.Wrap(err, "查询用户失败")
		return
	}
	if !user.CheckPassword(in.Password) {
		err = errorx.Wrap(err, "密码验证错误")
		return
	}
	var token string
	if token, err = ginx.SetToken(ctx, &ginx.JwtUser{
		Id:      user.Id,
		Account: user.Account,
		Name:    user.Name,
		Phone:   user.Phone,
		Ip:      loginIp,
		Domain:  ctx.Request.Host,
		TTL:     user.SessionTime,
	}); err != nil {
		err = errorx.Wrap(err, "设置token失败")
		return
	}
	// 记录日志
	if err = dao.AddLog(&entity.Log{
		Id:           idx.SnowFlake().Int64(),
		Service:      app.GetServer().Name,
		Type:         "request",
		Url:          ctx.Request.URL.Path,
		Operation:    "账号登录",
		Content:      fmt.Sprintf("%s通过%s进行账号密码登录", user.Name, user.Phone),
		Ip:           loginIp,
		CreateUserId: user.Id,
	}); err != nil {
		err = errorx.Wrap(err, "记录登录日志失败")
		return
	}
	result = &model.LoginResult{Token: token, User: user}
	return
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context, user ginx.AuthUser, ip string) error {
	if err := dao.AddLog(&entity.Log{
		Id:           idx.SnowFlake().Int64(),
		Service:      app.GetServer().Name,
		Type:         "request",
		Url:          ctx.Request.URL.Path,
		Operation:    "退出登录",
		Content:      user.Username() + "登出",
		Ip:           ip,
		CreateUserId: user.UserId(),
	}); err != nil {
		return errorx.Wrap(err, "记录登录日志失败")
	}
	// 移除token
	ginx.RemoveToken(ctx, user.Username())
	return nil
}
