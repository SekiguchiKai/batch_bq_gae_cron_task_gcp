package store

import (
	"context"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

const UserKind = "User"

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

// Userを全て取得する
func (s UserStore) GetAllUsers(dst *[]model.User) error {
	q := datastore.NewQuery(UserKind).Limit(10000)
	if _, err := q.GetAll(s.ctx, &dst); err != nil {
		return err
	}

	return nil
}

// idで指定したUserをdstにloadする。
func (s UserStore) GetUser(id string, dst *model.User) (exists bool, e error) {
	if id == "" {
		return false, nil
	}

	key := s.newUserKey(id)
	if err := datastore.Get(s.ctx, key, dst); err != nil {
		if err != datastore.ErrNoSuchEntity {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

// DatastoreにUserを格納する。
func (s UserStore) PutUser(u model.User) error {
	util.InfoLogWithContext(s.ctx, "PutUser is called")

	key := s.newUserKey(u.ID)

	if _, err := datastore.Put(s.ctx, key, &u); err != nil {
		return err
	}

	return nil

}

// 与えられたIDのUserがDatastore内に存在するかどうかを確認する。
func (s UserStore) ExistsUser(id string) (bool, error) {
	var dst model.User
	return s.GetUser(id, &dst)
}

// 与えられたIDのUserを削除する
func (s UserStore) DeleteUser(id string) error {
	key := s.newUserKey(id)
	return datastore.Delete(s.ctx, key)
}

// UserKind用のdatastore.Keyを発行する。
func (s UserStore) newUserKey(id string) *datastore.Key {
	return datastore.NewKey(s.ctx, UserKind, id, 0, nil)
}
