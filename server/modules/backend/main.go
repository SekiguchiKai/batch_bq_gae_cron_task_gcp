package backend

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/task"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// エントリポイント
func init() {
	g := gin.New()
	initTaskAPI(g)
	http.Handle("/", g)
}

// Task API群を初期登録する。
func initTaskAPI(g *gin.Engine) {
	taskGin := g.Group(util.GetTaskPath())

	task.InitMigrateUserDataFromDatastoreToBQ(taskGin)
}
