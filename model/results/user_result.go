package results

import (
	"github.com/quanxiaoxuan/go-utils/timex"

	"quan-admin/model/entity"
)

// 用户登录返回结果
type LoginResult struct {
	UserInfo UserInfo `json:"userInfo"`
	Token    string   `json:"token" comment:"token"`
}

// 用户信息
type UserInfoList []*UserInfo
type UserInfo struct {
	UserId      int64      `json:"userId" comment:"用户ID"`
	UserName    string     `json:"userName" comment:"姓名"`
	Phone       string     `json:"phone" comment:"手机"`
	Gender      string     `json:"gender" comment:"性别"`
	Birthday    timex.Date `json:"birthday" comment:"生日"`
	Email       string     `json:"email" comment:"邮箱"`
	Address     string     `json:"address" comment:"地址"`
	Remark      string     `json:"remark" comment:"备注"`
	SessionTime int64      `json:"sessionTime" comment:"会话有效期"`
	ValidStart  timex.Time `json:"validStart" comment:"账号有效期始"`
	ValidEnd    timex.Time `json:"validEnd" comment:"账号有效期止"`
}

// 用户简要信息
type UserSimpleList []*UserSimple
type UserSimple struct {
	UserId   int64  `json:"userId" comment:"用户ID"`
	UserName string `json:"userName" comment:"用户姓名"`
	Phone    string `json:"phone" comment:"手机"`
	Email    string `json:"email" comment:"邮箱"`
}

// 用户明细
type UserDetail struct {
	UserInfo *entity.SysUser     `json:"userInfo"` // 用户基本信息
	UserAuth *entity.SysUserAuth `json:"userAuth"` // 用户鉴权信息
}
