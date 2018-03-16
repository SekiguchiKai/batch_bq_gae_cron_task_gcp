package api

import (
	"google.golang.org/appengine/aetest"
	"testing"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SekiguchiKai/GAE_Go_Cursor/api"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"io"
	"bytes"
	"encoding/json"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
)

// aetest.Instanceのwrapper。
type userTestHelper struct {
	inst aetest.Instance
}


// 引数で与えられたUserのインスタンスをjsonにし、io.Readerにして返す。
func (h userTestHelper) newRequestBodyFromUserInstance(param model.User) io.Reader {
	b, _ := json.Marshal(param)

	return bytes.NewReader(b)
}

// gin.TestModeで、UserAPI起動する。
func (userTestHelper) newInitializedHandler() http.Handler {
	gin.SetMode(gin.TestMode)
	g := gin.New()
	api.InitUserAPI(g.Group(util.GetApiPath()))

	return g
}




func TestCreateUser(t *testing.T) {
	t.Run("User登録", func(t *testing.T) {
		inst, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true}) // strongly consistentにする
		if err != nil {
			t.Fatalf("Failed to create instance: %v", err)
		}
		defer inst.Close()





	})

}