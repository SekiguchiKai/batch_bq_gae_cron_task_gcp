package api

import "google.golang.org/appengine/aetest"

type ApiTestHelper struct {
	inst aetest.Instance
}

func NewApiTestHelper(inst aetest.Instance) ApiTestHelper {
	return ApiTestHelper{inst: inst}
}