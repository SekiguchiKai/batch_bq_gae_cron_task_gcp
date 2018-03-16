package api

import (
	"github.com/gin-gonic/gin"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
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

