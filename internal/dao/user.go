package dao

import (
	"errors"
	"strings"

	"github.com/go-xuan/quanx/server/gormx"

	"user/internal/model"
	"user/internal/model/entity"
)

const SelectUser = `
select u.id,
       u.account,
       u.name,
       u.phone,
       u.gender,
       u.birthday,
       u.email,
       u.address,
       u.remark,
       a.session_time,
       a.valid_start,
       a.valid_end,
       u.update_time
  from t_sys_user u
  left join t_sys_user_auth a
    on u.id = a.user_id`

// 用户分页查询
func UserPage(in model.UserPage) (result []*model.User, total int64, err error) {
	sql := strings.Builder{}
	sql.WriteString(SelectUser)
	if in.Keyword != "" {
		sql.WriteString(` where account like '%`)
		sql.WriteString(in.Keyword)
		sql.WriteString(`%' or name like '%`)
		sql.WriteString(in.Keyword)
		sql.WriteString(`%'`)
	}
	selectSql := strings.Builder{}
	selectSql.WriteString(`select * from`)
	selectSql.WriteString(` ( `)
	selectSql.WriteString(sql.String())
	selectSql.WriteString(` ) t order by update_time desc`)
	if in.Page != nil && in.Page.PageSize > 0 {
		selectSql.WriteString(in.Page.PgPageSql())
	}
	err = gormx.This().DB.Raw(selectSql.String()).Scan(&result).Error
	if err != nil {
		return
	}
	countSql := strings.Builder{}
	countSql.WriteString(`select count(*) from`)
	countSql.WriteString(` ( `)
	countSql.WriteString(sql.String())
	countSql.WriteString(` ) t `)
	err = gormx.This().DB.Raw(countSql.String()).Scan(&total).Error
	if err != nil {
		return
	}
	return
}

// 查询用户信息
func GetUserById(id int64) (user *model.User, err error) {
	user = &model.User{}
	sql := strings.Builder{}
	sql.WriteString(SelectUser)
	sql.WriteString(` where u.id = ?`)
	if err = gormx.This().DB.Raw(sql.String(), id).Scan(user).Error; err != nil {
		return
	}
	if user.Id == 0 {
		err = errors.New("此用户不存在")
	}
	return
}

// 查询用户信息
func GetUserByName(username string) (user *model.User, err error) {
	user = &model.User{}
	sql := strings.Builder{}
	sql.WriteString(SelectUser)
	sql.WriteString(` where u.phone = ? or u.account = ? `)
	if err = gormx.This().DB.Raw(sql.String(), username, username).Scan(user).Error; err != nil {
		return
	}
	if user.Id == 0 {
		err = errors.New("此用户不存在")
	}
	return
}

// 查询用户基本信息
func QueryUser(id int64) (user *entity.User, err error) {
	user.Id = id
	if err = gormx.This().DB.Find(user).Error; err != nil {
		return
	}
	return
}

// 查询用户身份信息
func QueryUserAuth(userId int64) (auth *entity.UserAuth, err error) {
	auth = &entity.UserAuth{}
	auth.UserId = userId
	if err = gormx.This().DB.Find(auth).Error; err != nil {
		return
	}
	return
}

// 用户列表查询
func UserList() (result []*model.User, err error) {
	if err = gormx.This().DB.Model(&entity.User{}).
		Select([]string{"id", "account", "name", "phone", "birthday", "gender", "email", "address"}).
		Order("id desc").Scan(&result).Error; err != nil {
		return
	}
	return
}

// 查询手机是否存在
func UserExist(in *model.UserSave) (count int64, err error) {
	if err = gormx.This().DB.Model(&entity.User{}).Where(`account = ? or phone = ?`, in.Account, in.Phone).Count(&count).Error; err != nil {
		return
	}
	return
}

// 用户新增
func UserCreate(user *entity.User, userAuth *entity.UserAuth) error {
	tx := gormx.This().DB.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(userAuth).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 用户修改
func UserUpdate(in *model.UserSave) error {
	tx := gormx.This().DB.Begin()
	// 更新用户表
	var userCols = []string{"update_user_id", "update_time"}
	var userAuthCols = []string{"update_user_id", "update_time"}
	if in.Name != "" {
		userCols = append(userCols, "name")
	}
	if in.Phone != "" {
		userCols = append(userCols, "phone")
	}
	if in.Gender != "" {
		userCols = append(userCols, "gender")
	}
	if in.Birthday != "" {
		userCols = append(userCols, "birthday")
	}
	if in.Email != "" {
		userCols = append(userCols, "email")
	}
	if in.Address != "" {
		userCols = append(userCols, "address")
	}
	if in.Remark != "" {
		userCols = append(userCols, "remark")
	}
	if in.Password != "" {
		userAuthCols = append(userAuthCols, "password", "salt")
	}
	if in.ValidStart != "" {
		userAuthCols = append(userAuthCols, "valid_start")
	}
	if in.ValidEnd != "" {
		userAuthCols = append(userAuthCols, "valid_end")
	}
	// 更新用户表
	if err := tx.Model(&entity.User{}).Select(userCols).Where("user_id = ?", in.Id).Updates(in.UserUpdate()).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新用户鉴权表
	if err := tx.Model(&entity.UserAuth{}).Select(userAuthCols).Where("user_id = ?", in.Id).Updates(in.UserAuthUpdate()).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 用户删除
func UserDelete(userId int64) error {
	tx := gormx.This().DB.Begin()
	err := tx.Delete(&entity.User{}, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Delete(&entity.UserAuth{}, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
