package api

import (
	"github.com/gin-gonic/gin"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"errors"
)

func createUser(c *gin.Context) {
	

}


// HTTPのリクエストボディのjsonデータUserに変換する。
func bindUserFromJson(c *gin.Context, dst *model.User) error {
	if err := c.BindJSON(dst); err != nil {
		return err
	}

	dst.ID = getUserID(c)
	return nil
}

// URIのIDを取得する。
func getUserID(c *gin.Context) string {
	return c.Param("id")
}


// 送信されて来たUserに必要なデータが存在するかどうかのバリデーションを行う。
func validateParamsForUser(u model.User) error {
	if u.UserName == "" {
		return util.CreateErrMessage("UserName", _RequiredErrMessage)
	}

	if u.MailAddress == "" {
		return util.CreateErrMessage("MailAddress", _RequiredErrMessage)
	}

	if u.Age < 0 {
		return errors.New("Age should be over 0 ")
	}

	if u.Gender == "" {
		return util.CreateErrMessage("Gender", _RequiredErrMessage)
	}

	if u.From == "" {
		return util.CreateErrMessage("From", _RequiredErrMessage)
	}

	return nil

}