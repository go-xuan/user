package mapper

import (
	"strings"

	"github.com/quanxiaoxuan/go-builder/gormx"
	"github.com/quanxiaoxuan/go-utils/randx"
	"github.com/quanxiaoxuan/go-utils/stringx"
	log "github.com/sirupsen/logrus"

	"quan-admin/model/entity"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

const UserBaseQuery = `
select u.user_id,
       u.user_name,
       u.phone,
       u.gender,
       u.birthday,
       u.email,
       u.address,
       u.remark,
       a.session_time,
       a.valid_start,
       a.valid_end
  from sys_user u
  left join sys_user_auth a
    on u.user_id = a.user_id`

// 用户分页查询
func UserPage(param params.UserPage) (resultList results.UserInfoList, total int64, err error) {
	sql := strings.Builder{}
	sql.WriteString(UserBaseQuery)
	if param.SearchKey != "" {
		sql.WriteString(` where user_name like '%` + param.SearchKey + `%'`)
	}
	selectSql := strings.Builder{}
	selectSql.WriteString(`select * from`)
	selectSql.WriteString(` ( `)
	selectSql.WriteString(sql.String())
	selectSql.WriteString(`) t order by user_id `)
	if param.PageParam != nil || param.PageParam.PageSize > 0 {
		selectSql.WriteString(param.PageParam.GetPgPageSql())
	}
	err = gormx.GormDB.Raw(selectSql.String()).Scan(&resultList).Error
	if err != nil {
		log.Error("查询用户信信息失败 ： ", err)
		return
	}
	countSql := strings.Builder{}
	countSql.WriteString(`select count(*) from`)
	countSql.WriteString(` ( `)
	countSql.WriteString(sql.String())
	countSql.WriteString(`) t `)
	err = gormx.GormDB.Raw(countSql.String()).Scan(&total).Error
	if err != nil {
		log.Error("查询用户信信息失败 ： ", err)
		return
	}
	return
}

// 查询用户信息
func QueryUserInfo(userKey string) (result results.UserInfo, err error) {
	sql := strings.Builder{}
	sql.WriteString(UserBaseQuery)
	sql.WriteString(` where u.phone = ?`)
	err = gormx.GormDB.Raw(sql.String(), userKey).Scan(&result).Error
	return
}

// 查询用户基本信息
func UserSelect(userId int64) (sysUser entity.SysUser, err error) {
	sysUser.UserId = userId
	err = gormx.GormDB.Find(&sysUser).Error
	return
}

// 用户列表查询
func UserList() (resultList results.UserSimpleList, err error) {
	err = gormx.GormDB.Model(&entity.SysUser{}).Select("user_id", "user_name", "phone", "email").Order("user_id").Scan(&resultList).Error
	return
}

// 查询群组编码是否存在
func UserPhoneExist(phone string) (count int64, err error) {
	err = gormx.GormDB.Model(&entity.SysUser{}).Where(`phone = ? `, phone).Count(&count).Error
	return
}

// 用户新增
func UserAdd(user entity.SysUser, userAuth entity.SysUserAuth) error {
	tx := gormx.GormDB.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		log.Error("写入用户表失败 ： ", err)
		return err
	}
	err = tx.Create(&userAuth).Error
	if err != nil {
		tx.Rollback()
		log.Error("写入用户鉴权表失败 ： ", err)
		return err
	}
	tx.Commit()
	return nil
}

// 用户修改
func UserUpdate(param params.UserUpdate) error {
	tx := gormx.GormDB.Begin()
	// 更新用户表
	userSql := strings.Builder{}
	userSql.WriteString(`update sys_user set update_time = now(), update_user_id = ? `)
	if param.UserName != "" {
		userSql.WriteString(`, user_name = '` + param.UserName + `'`)
	}
	if param.Phone != "" {
		userSql.WriteString(`, phone = '` + param.Phone + `'`)
	}
	if param.Gender != "" {
		userSql.WriteString(`, gender = '` + param.Gender + `'`)
	}
	if param.Birthday != "" {
		userSql.WriteString(`, birthday = '` + param.Birthday + `'`)
	}
	if param.Email != "" {
		userSql.WriteString(`, email = '` + param.Email + `'`)
	}
	if param.Address != "" {
		userSql.WriteString(`, address = '` + param.Address + `'`)
	}
	if param.Remark != "" {
		userSql.WriteString(`, remark = '` + param.Remark + `'`)
	}
	userSql.WriteString(` where user_id = ?`)
	err := tx.Exec(userSql.String(), param.UpdateUserId, param.UserId).Error
	if err != nil {
		tx.Rollback()
		log.Error("更新用户表失败 ： ", err)
		return err
	}
	// 更新用户鉴权表
	authSql := strings.Builder{}
	authSql.WriteString(`update sys_user_auth set update_time = now(), update_user_id = ? `)
	if param.Password != "" {
		passWord := param.Password
		salt := randx.UUID()
		if len(passWord) < 32 {
			passWord = stringx.MD5(passWord)
		}
		passWord = stringx.PasswordSalt(passWord, salt)
		authSql.WriteString(`, password = '` + passWord + `'`)
		authSql.WriteString(`, salt = '` + salt + `'`)
	}
	if param.ValidStart != "" {
		authSql.WriteString(`, valid_start = '` + param.ValidStart + `'`)
	}
	if param.ValidEnd != "" {
		authSql.WriteString(`, valid_end = '` + param.ValidEnd + `'`)
	}
	authSql.WriteString(` where user_id = ?`)
	err = tx.Exec(authSql.String(), param.UpdateUserId, param.UserId).Error
	if err != nil {
		tx.Rollback()
		log.Error("更新用户鉴权表失败 ： ", err)
		return err
	}
	tx.Commit()
	return nil
}

// 用户删除
func UserDelete(userId int64) error {
	tx := gormx.GormDB.Begin()
	err := gormx.GormDB.Delete(&entity.SysUser{}, userId).Error
	if err != nil {
		tx.Rollback()
		log.Error("删除用户表失败 ： ", err)
		return err
	}
	err = gormx.GormDB.Delete(&entity.SysUserAuth{}, userId).Error
	if err != nil {
		tx.Rollback()
		log.Error("删除用户鉴权表失败 ： ", err)
		return err
	}
	tx.Commit()
	return nil
}
