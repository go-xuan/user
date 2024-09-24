package dao

import (
	"fmt"

	"github.com/go-xuan/quanx/core/gormx"
	"github.com/go-xuan/quanx/os/errorx"

	"user/internal/model"
	"user/internal/model/entity"
)

// UserPage 用户分页查询
func UserPage(in model.UserPage) (result []*model.User, total int64, err error) {
	db := gormx.DB().Model(&entity.User{})
	if in.Keyword != "" {
		db = db.Where(fmt.Sprintf("account like '%%%s%%' or name like '%%%s%%'", in.Keyword, in.Keyword))
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if in.Page != nil && in.Page.PageSize > 0 {
		db.Limit(in.Page.PageSize).Offset(in.Page.Offset()).Order("update_time desc")
	}
	if err = db.Scan(&result).Error; err != nil {
		return
	}
	return
}

// QueryUser 查询用户信息
func QueryUser(id int64) (user *model.User, err error) {
	user = &model.User{}
	if err = gormx.DB().Model(&entity.User{}).Where("id = ?", id).Scan(user).Error; err != nil {
		return
	}
	if user.Id == 0 {
		err = errorx.New("此用户不存在")
	}
	return
}

// QueryUserByPhone 查询用户信息
func QueryUserByPhone(phone string) (user *model.User, err error) {
	user = &model.User{}
	if err = gormx.DB().Model(&entity.User{}).Where("phone = ?", phone).Find(user).Error; err != nil {
		return
	}
	if user.Id == 0 {
		err = errorx.New("此用户不存在")
	}
	return
}

// GetUserById 查询用户基本信息
func GetUserById(id int64) (user *entity.User, err error) {
	user = &entity.User{
		Id: id,
	}
	if err = gormx.DB().Find(user).Error; err != nil {
		return
	}
	return
}

// UserList 用户列表查询
func UserList() (result []*model.User, err error) {
	if err = gormx.DB().Model(&entity.User{}).
		Select([]string{"id", "account", "name", "phone", "birthday", "gender", "email", "address"}).
		Order("id desc").Scan(&result).Error; err != nil {
		return
	}
	return
}

// UserExist 查询手机是否存在
func UserExist(in *model.UserSave) (count int64, err error) {
	if err = gormx.DB().Model(&entity.User{}).Where(`phone = ? `, in.Phone).Count(&count).Error; err != nil {
		return
	}
	return
}

// UserCreate 用户新增
func UserCreate(user *entity.User) error {
	if err := gormx.DB().Create(user).Error; err != nil {
		return err
	}
	return nil
}

// UserUpdate 用户修改
func UserUpdate(in *model.UserSave) error {
	// 更新用户表
	var userCols = []string{"update_user_id", "update_time"}
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
		userCols = append(userCols, "password", "salt")
	}
	if in.ValidStart != "" {
		userCols = append(userCols, "valid_start")
	}
	if in.ValidEnd != "" {
		userCols = append(userCols, "valid_end")
	}
	// 更新用户表
	if err := gormx.DB().Model(&entity.User{}).Select(userCols).Where("id = ? ", in.Id).Updates(in.UserUpdate()).Error; err != nil {
		return err
	}
	return nil
}

// UserDelete 用户删除
func UserDelete(userId int64) error {
	err := gormx.DB().Delete(&entity.User{}, userId).Error
	if err != nil {
		return err
	}
	return nil
}
