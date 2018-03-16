package api

import (
	"google.golang.org/appengine/aetest"
	"testing"
)

// aetest.Instanceのwrapper。
type userTestHelper struct {
	inst aetest.Instance
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