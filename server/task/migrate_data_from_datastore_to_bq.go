package task

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
)

func InitMigrateDataFromDatastoreToBQ(g *gin.RouterGroup) {

}

// DatastoreからBigQueryのTableに移し替える
func migrateDataFromDatastoreToBQ(c *gin.Context) {
	util.InfoLog(c.Request, "migrateDataFromDatastoreToBQ is called")

	// Datastoreからデータを取得する

}
