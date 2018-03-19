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
	"net/http/httptest"
	"io/ioutil"
)

// aetest.Instanceのwrapper。
type userTestHelper struct {
	inst aetest.Instance
}

// UserAPIを起動して、POSTでリクエストする。
// 回答として、Status Code、レスポンスBodyを返す。
func (h userTestHelper) requestPostToUserAPI(param model.User) (int, string) {
	path := util.GetApiPath() + "/user/" + "/new"

	// ResponseRecorderを作成
	w := httptest.NewRecorder()
	// リクエストを作成
	r, _ := h.inst.NewRequest("POST", path, h.newRequestBodyFromUserInstance(param))

	// gin.TestModeで、UserAPI起動し、serveしておく
	h.newInitializedHandler().ServeHTTP(w, r)

	// レスポンスBodyを読み込み
	b, _ := ioutil.ReadAll(w.Body)

	return w.Code, string(b)
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

		helper := userTestHelper{inst}


		t.Run("全てのパラメータが正常な場合は200OKになること", func(t *testing.T) {
			defer helper.clear()

			status, _ := helper.requestPostToUserAPI(helper.newUserParam())

			if status != http.StatusOK {
				t.Errorf("status = %d, wants = %d", status, http.StatusOK)
			}
		})



	})

}

// parameter用のmodel.Userを作成する
func (userTestHelper) newUserParam() model.User {
	return model.User{
		UserName   : "太郎",
		MailAddress : "sample@test.mail",
		Age         : 20,
		Gender      : model.Male,
		From        : "japan",
	}
}

// Datastore内のUser KindのEntityを全て削除する
func (h userTestHelper) clear() {
	adminHelper := NewApiTestHelper(h.inst)
	adminHelper.ClearEntity("User")
}