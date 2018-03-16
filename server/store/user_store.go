package store

import (
	"context"
	"google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine/datastore"
)

const _UserKind = "User"

// User用のUserStore
type UserStore struct {
	ctx context.Context
}

// http.RequestからUserStoreを新規発行する。
func NewUserStore(r *http.Request) UserStore {
	return NewUserStoreWithContext(appengine.NewContext(r))
}


// context.ContextからUserStoreを新規発行する。
func NewUserStoreWithContext(ctx context.Context) UserStore {
	return UserStore{ctx: ctx}
}

// UserKind用のdatastore.Keyを発行する。
func (s UserStore) newUserStoreKey(id string) *datastore.Key {
	return datastore.NewKey(s.ctx, _UserKind, id, 0, nil)
}