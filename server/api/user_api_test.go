package api

import "google.golang.org/appengine/aetest"

// aetest.Instanceのwrapper。
type userTestHelper struct {
	inst aetest.Instance
}