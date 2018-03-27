package model

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"strconv"
	"time"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)
const (
	IDIndex = iota
	UserNameIndex
	MailAddressIndex
	AgeIndex
	GenderIndex
	FromIndex
	CreatedAtIndex
	UpdatedAtIndex
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

type UserNameField struct {
	Signal string
	Value  string
}

type MailAddressField struct {
	Signal string
	Value  string
}

type AgeField struct {
	Signal string
	Value  int
}

type GenderField struct {
	Signal string
	Value  string
}

type FromField struct {
	Signal string
	Value  string
}

type CreatedAtField struct {
	Signal string
	Value  string
}

type UpdatedAtField struct {
	Signal string
	Value  string
}

// BigQueryの解析用
type UserForAnalysis struct {
	UserNameField
	MailAddressField
	AgeField
	GenderField
	FromField
	CreatedAtField
	UpdatedAtField
}

// IDを付与して、新しいUserを作成する
func NewUser(u User) User {
	u.ID = newUserID(u.UserName)
	return u
}

// Userの情報を更新する
func UpdateUser(original, param User) User {
	original.MailAddress = param.MailAddress
	original.Age = param.Age
	original.Gender = param.Gender
	original.From = param.From

	return original
}

// UserのIDを発行する
func newUserID(userName string) string {
	return util.GetHash(userName)
}

// Userを構造体からSliceにする
func TranslateStructToSlice(u User) []string {

	s := make([]string, 8)
	s[IDIndex] = u.ID
	s[UserNameIndex] = u.UserName
	s[MailAddressIndex] = u.MailAddress
	s[AgeIndex] = strconv.Itoa(u.Age)
	s[GenderIndex] = string(u.Gender)
	s[FromIndex] = u.From
	s[CreatedAtIndex] = u.CreatedAt.String()
	s[UpdatedAtIndex] = u.UpdatedAt.String()

	return s
}
