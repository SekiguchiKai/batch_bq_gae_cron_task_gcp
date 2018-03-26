package task

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/service"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

// BigQueryからデータを抽出し、抽出したデータからCSVを作成する
func createCsvFromBigquery(c *gin.Context, sql string) error {
	util.InfoLog(c.Request, "createCsvFromBigquery is called")

	ctx := appengine.NewContext(c.Request)
	prjID := appengine.AppID(ctx)

	bq, err := service.NewBQClientWrapper(ctx, prjID)
	if err != nil {
		return err
	}

	var users []model.User
	if err := bq.QueryAndLoad(sql, &users); err != nil {
		return err
	}

	util.InfoLog(c.Request, "users : %+v", users)

	for _, user := range users {
		s := model.TranslateStructToSlice(user)

		if err := util.WriteCsv(user.UserName, s); err != nil {
			return err
		}
	}

	return nil

}
