package mapper

import (
	"strings"

	"github.com/quanxiaoxuan/go-builder/gormx"
	"quan-admin/model/entity"
	"quan-admin/model/params"
)

// 查询用户身份信息
func QueryUserAuth(userId int64) (userAuth entity.SysUserAuth, err error) {
	userAuth.UserId = userId
	err = gormx.Ctl.DB.Find(&userAuth).Error
	return
}

// 更新用户身份信息
func UserAuthUpdate(update params.UserAuthUpdate) (err error) {
	sql := strings.Builder{}
	sql.WriteString(`update sys_user_auth set update_time = now(), update_user_id = ? `)
	if update.Password != "" {
		sql.WriteString(`, password = '` + update.Password + `'`)
		sql.WriteString(`, salt = '` + update.Salt + `'`)
	}
	if update.ValidStart != "" {
		sql.WriteString(`, valid_start = '` + update.ValidStart + `'`)
	}
	if update.ValidEnd != "" {
		sql.WriteString(`, valid_end = '` + update.ValidEnd + `'`)
	}
	sql.WriteString(` where user_id = ?`)
	err = gormx.Ctl.DB.Exec(sql.String(), update.UpdateUserId, update.UserId).Error
	return
}
