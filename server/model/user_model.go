package model

import "time"

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

// UserのModel
type User struct {
	ID string `json:"id"`
	// これは各ユーザーで一意に定めるものとする
	// 変更不可
	UserName    string    `json:"userName"`
	MailAddress string    `json:"mailAddress"`
	Age         int       `json:"age"`
	Gender      Gender    `json:"gender"`
	From        string    `json:"from"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
