package api

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
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

// EntityのKeyを取得する
func (h ApiTestHelper) GetEntityKey(entityName string, order string) *datastore.Key {
	ctx := appengine.NewContext(h.request())
	q := datastore.NewQuery(entityName).Order(order).KeysOnly().Limit(1)
	keys, _ := q.GetAll(ctx, nil)
	return keys[0]
}

// UpdatedAtが有効な値かどうかを確認する
// 1分以内ならば、有効なUpdatedAtとして考える
func (ApiTestHelper) IsValidUpdatedAt(updatedAt, now time.Time) bool {
	acceptedUpdatedAtInterval := time.Duration(1) * time.Minute

	interval := now.Sub(updatedAt)
	if interval < 0 {
		interval = -interval
	}

	return interval <= acceptedUpdatedAtInterval
}
