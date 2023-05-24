package params

import (
	"github.com/quanxiaoxuan/go-builder/paramx/request"
)

// 用户分页参数
type UserPage struct {
	SearchKey string        `json:"searchKey" comment:"关键字"`
	PageParam *request.Page `json:"pageParam" comment:"分页参数"`
}

// 用户登录参数
type UserLogin struct {
	UserId   int64  `json:"userId" comment:"用户ID"`
	Phone    string `json:"phone" comment:"手机"`
	Password string `json:"password" comment:"密码"`
}

// 用户新增参数
type UserCreate struct {
	UserName     string `json:"userName" comment:"姓名"`
	Password     string `json:"password" comment:"密码"`
	Phone        string `json:"phone" comment:"手机"`
	Gender       string `json:"gender" comment:"性别"`
	Birthday     string `json:"birthday" comment:"生日"`
	Email        string `json:"email" comment:"邮箱"`
	Address      string `json:"address" comment:"地址"`
	Remark       string `json:"remark" comment:"备注"`
	ValidStart   string `json:"validStart" comment:"账号有效期始"`
	ValidEnd     string `json:"validEnd" comment:"账号有效期止"`
	CreateUserId int64  `json:"createUserId" comment:"创建人"`
}

// 用户修改参数
type UserUpdate struct {
	UserId       int64  `json:"userId" comment:"用户ID"`
	UserName     string `json:"userName" comment:"姓名"`
	Password     string `json:"password" comment:"密码"`
	Phone        string `json:"phone" comment:"手机"`
	Gender       string `json:"gender" comment:"性别"`
	Birthday     string `json:"birthday" comment:"生日"`
	Email        string `json:"email" comment:"邮箱"`
	Address      string `json:"address" comment:"地址"`
	Remark       string `json:"remark" comment:"备注"`
	ValidStart   string `json:"validStart" comment:"账号有效期始"`
	ValidEnd     string `json:"validEnd" comment:"账号有效期止"`
	UpdateUserId int64  `json:"updateUserId"  comment:"更新人"`
}

// 用户鉴权表
type UserAuthUpdate struct {
	UserId       int64  `json:"userId" `
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	SessionTime  int64  `json:"sessionTime"`
	ValidStart   string `json:"validStart"`
	ValidEnd     string `json:"validEnd"`
	UpdateUserId int64  `json:"updateUserId"`
}
