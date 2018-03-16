package model

import (
	"time"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
)

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

// IDを付与して、新しいUserを作成する
func NewUser(u User) User {
	u.ID = newUserID(u.UserName)
	return u
}

// UserのIDを発行する
func newUserID(userName string) string {
	return util.GetHash(userName)
}