package store

import (
	"context"
)

const _UserKind = "User"

// User用のUserStore
type UserStore struct {
	ctx context.Context
}


// context.ContextからUserStoreを新規発行する。
func NewUserStoreWithContext(ctx context.Context) UserStore {
	return UserStore{ctx: ctx}
}