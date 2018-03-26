package task

import (
	"context"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/service"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
	"net/http"
	"net/url"
)

// CreateCsvFromBigQuery Taskのエントリポイント
func InitCreateCsvFromBigQuery(g *gin.RouterGroup) {
	g.POST("/createCsvFromBigQuery", createCsvFromBigQuery)

}

// CreateCsvFromBigQueryを開始する
func StartCreateCsvFromBigQuery(ctx context.Context, sql string) error {
	return startCreateCsvFromBigQuery(ctx, CreateCsvFromBigQueryQueue, sql)

}

// createCsvFromBigQueryをTaskQueueのTaskに追加する
func startCreateCsvFromBigQuery(ctx context.Context, queueName, sql string) error {
	util.InfoLogWithContext(ctx, "startCreateCsvFromBigQuery is called")

	taskPath := util.GetTaskPath()
	t := taskqueue.NewPOSTTask(taskPath+"/createCsvFromBigQuery", url.Values{
		"": []string{sql},
	})

	if _, err := taskqueue.Add(ctx, t, queueName); err != nil {
		return err
	}

	return nil

}

// BigQueryからデータを抽出し、抽出したデータからCSVを作成する
func createCsvFromBigQuery(c *gin.Context) {
	util.InfoLog(c.Request, "createCsvFromBigQuery is called")

	ctx := appengine.NewContext(c.Request)
	prjID := appengine.AppID(ctx)

	bq, err := service.NewBQClientWrapper(ctx, prjID)
	if err != nil {
		util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
	}

	var users []model.User
	sql := c.PostForm("sql")

	if err := bq.QueryAndLoad(sql, &users); err != nil {
		util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
	}

	util.InfoLog(c.Request, "users : %+v", users)

	for _, user := range users {
		s := model.TranslateStructToSlice(user)

		if err := util.WriteCsv(user.UserName, s); err != nil {
			util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
		}
	}

	c.JSON(http.StatusOK, nil)

}
