package model

import "github.com/go-xuan/quanx/common/modelx"

type Page struct {
	Keyword string       `json:"keyword" comment:"关键字"`
	Page    *modelx.Page `json:"page" comment:"分页参数"`
}
