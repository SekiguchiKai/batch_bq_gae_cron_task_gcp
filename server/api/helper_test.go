package api

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/store"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type ApiTestHelper struct {
	inst aetest.Instance
}

func NewApiTestHelper(inst aetest.Instance) ApiTestHelper {
	return ApiTestHelper{inst: inst}
}

// DatastoreのKindからEntityを削除する
func (h ApiTestHelper) ClearEntity(entityName string) {
	ctx := appengine.NewContext(h.request())
	q := datastore.NewQuery(entityName).KeysOnly()

	keys, _ := q.GetAll(ctx, nil)

	datastore.DeleteMulti(ctx, keys)
}

// 空のリクエストを作成する
func (h ApiTestHelper) request() *http.Request {
	r, _ := h.inst.NewRequest("GET", "", nil)
	return r
}

// Datastoreに格納されている最新のUserのEntityを取得する
func (h ApiTestHelper) GetLatestUser() model.User {
	key := h.GetEntityKey(store.UserKind, "-UpdatedAt")
	ctx := appengine.NewContext(h.request())
	var dst model.User
	datastore.Get(ctx, key, &dst)

	return dst
}

// EntityのKeyを取得する
func (h ApiTestHelper) GetEntityKey(entityName string, order string) *datastore.Key {
	ctx := appengine.NewContext(h.request())
	q := datastore.NewQuery(entityName).Order(order).KeysOnly().Limit(1)
	keys, _ := q.GetAll(ctx, nil)
	return keys[0]
}
