package api

import (
	"google.golang.org/appengine/aetest"
	"testing"

	"bytes"
	"encoding/json"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/store"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"
)

// aetest.Instanceのwrapper。
type userTestHelper struct {
	inst aetest.Instance
}

// UserAPIを起動して、POSTでリクエストする。
// 回答として、Status Code、レスポンスBodyを返す。
func (h userTestHelper) requestPostToUserAPI(param model.User) (int, string) {
	path := util.GetApiPath() + "/user/" + "new"

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
	InitUserAPI(g.Group(util.GetApiPath()))

	return g
}

func TestCreateUser(t *testing.T) {
	t.Run("User登録", func(t *testing.T) {
		inst, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true}) // strongly consistentにする
		if err != nil {
			t.Fatalf("Failed to create instance: %v", err)
		}
		defer inst.Close()

		apiHelper := NewApiTestHelper(inst)

		helper := userTestHelper{inst}

		t.Run("リクエストボディの全パラメータが正常な場合は200OKになること", func(t *testing.T) {
			defer helper.clear()

			status, body := helper.requestPostToUserAPI(helper.newUserParam())

			if status != http.StatusOK {
				t.Log("helper.newUserParam = %+v", helper.newUserParam())
				t.Errorf("status = %d, wants = %d", status, http.StatusOK)
				t.Errorf("response = %s", body)

			}
		})

		t.Run("既に登録済みのuserNameの場合はNotUniqueErrMessageエラーになること", func(t *testing.T) {
			defer helper.clear()

			// 既に登録済みのuserNameの状態を作成するために、まず1回登録する
			_, _ = helper.requestPostToUserAPI(helper.newUserParam())
			// 2回目の登録
			status, msg := helper.requestPostToUserAPI(helper.newUserParam())
			if status != http.StatusBadRequest {
				t.Errorf("status = %d, wants = %d", status, http.StatusBadRequest)
			}
			if msg != util.CreateErrMessage(_NotUniqueErrMessage).Error() {
				t.Errorf("Body = %s, wants = %s", msg, util.CreateErrMessage(_NotUniqueErrMessage).Error())
			}
		})

		t.Run("必須項目が空の場合は400エラーになり、エラーメッセージが適切に返却されること", func(t *testing.T) {
			defer helper.clear()

			requiredParams := []string{
				_UserName,
				_MailAddress,
				_Age,
				_Gender,
				_From,
			}
			for _, requiredParam := range requiredParams {
				t.Run(requiredParam, func(t *testing.T) {
					param := helper.newUserParam()
					param = helper.deleteProperty(param, requiredParam)
					status, msg := helper.requestPostToUserAPI(param)

					if status != http.StatusBadRequest {
						t.Errorf("status = %d, wants = %d", status, http.StatusBadRequest)
					}
					if requiredParam == _Age && msg != util.CreateErrMessage(requiredParam, _ShouldBeOver, strconv.Itoa(0)).Error() {
						t.Errorf("wants = %s, actual = %s", util.CreateErrMessage(requiredParam, _ShouldBeOver, strconv.Itoa(0)).Error(), msg)
					} else if requiredParam != _Age && msg != util.CreateErrMessage(requiredParam, _RequiredErrMessage).Error() {
						t.Errorf("wants = %s, actual = %s", util.CreateErrMessage(requiredParam, _RequiredErrMessage), msg)
					}
				})
			}
		})

		t.Run("各項目にパラメータで渡された値が設定されること", func(t *testing.T) {
			defer helper.clear()

			param := helper.newUserParam()
			helper.requestPostToUserAPI(param)

			u := helper.getLatestUser()

			t.Run(_UserName, func(t *testing.T) {
				if u.UserName != param.UserName {
					t.Errorf("want = %s, actual = %s", param.UserName, u.UserName)
				}
			})

			t.Run(_MailAddress, func(t *testing.T) {
				if u.MailAddress != param.MailAddress {
					t.Errorf("want = %s, actual = %s", param.MailAddress, u.MailAddress)
				}
			})

			t.Run(_Age, func(t *testing.T) {
				if u.Age != param.Age {
					t.Errorf("want = %d, actual = %d", param.Age, u.Age)
				}
			})

			t.Run(_Gender, func(t *testing.T) {
				if u.Gender != param.Gender {
					t.Errorf("want = %s, actual = %s", param.Gender, u.Gender)
				}
			})

			t.Run(_From, func(t *testing.T) {
				if u.From != param.From {
					t.Errorf("want = %s, actual = %s", param.From, u.From)
				}
			})

		})

		t.Run("UpdatedAtに適切なシステム日時が設定されること", func(t *testing.T) {
			defer helper.clear()

			helper.requestPostToUserAPI(helper.newUserParam())

			now := time.Now().UTC()
			a := helper.getLatestUser()
			if !apiHelper.IsValidUpdatedAt(a.UpdatedAt, now) {
				t.Errorf("UpdatedAt = %v, now = %v", a.UpdatedAt, now)
			}
		})

	})

}

// parameter用のmodel.Userを作成する
func (userTestHelper) newUserParam() model.User {
	return model.User{
		UserName:    "太郎",
		MailAddress: "sample@test.mail",
		Age:         20,
		Gender:      model.Female,
		From:        "japan",
	}
}

// Datastore内のUser KindのEntityを全て削除する
func (h userTestHelper) clear() {
	apiHelper := NewApiTestHelper(h.inst)
	apiHelper.ClearEntity(store.UserKind)
}

// 指定された構造体のインスタンスのPropertyを削除する
func (userTestHelper) deleteProperty(param model.User, propName string) model.User {
	switch propName {
	case _UserName:
		param.UserName = _Empty
	case _MailAddress:
		param.MailAddress = _Empty
	case _Age:
		param.Age = -1
	case _Gender:
		param.Gender = _Empty
	case _From:
		param.From = _Empty
	}
	return param
}

// Datastoreに格納されている最新のUserのEntityを取得する
func (h userTestHelper) getLatestUser() model.User {
	apiHelper := NewApiTestHelper(h.inst)
	key := apiHelper.GetEntityKey(store.UserKind, "-UpdatedAt")

	ctx := appengine.NewContext(apiHelper.request())
	var dst model.User
	datastore.Get(ctx, key, &dst)

	return dst
}
