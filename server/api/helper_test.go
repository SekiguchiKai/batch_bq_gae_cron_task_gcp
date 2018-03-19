package api

import (
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine"
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