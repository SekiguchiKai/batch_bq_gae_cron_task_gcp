package task

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/service"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/store"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

const (
	_Dataset = "batch_bq_task_gcp"
	_Table   = "user"
)

// MigrateUserDataFromDatastoreToBQを初期化する
func InitMigrateUserDataFromDatastoreToBQ(g *gin.RouterGroup) {
	g.POST("migrateUserDataFromDatastoreToBQ", migrateUserDataFromDatastoreToBQ)
}

// DatastoreからBigQueryのTableに移し替える
func migrateUserDataFromDatastoreToBQ(c *gin.Context) {
	util.InfoLog(c.Request, "migrateDataFromDatastoreToBQ is called")

	// Datastoreからデータを取得する
	var users []model.User
	s := store.NewUserStore(c.Request)
	if err := s.GetAllUsers(&users); err != nil {
		util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := appengine.NewContext(c.Request)
	prjID := appengine.AppID(ctx)
	bq, err := service.NewBQClientWrapper(ctx, prjID)

	if err != nil {
		util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	// DatastoreのデータをBQに詰め込む
	for _, user := range users {
		if err := bq.PutData(_Dataset, _Table, user); err != nil {
			util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
			return
		}
		// Datastoreの余分なデータを消す
		if err := s.DeleteUser(user.ID); err != nil {
			util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
			return
		}

	}

	c.JSON(http.StatusOK, nil)

}
