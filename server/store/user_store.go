package store

import "context"

const _UserKind = "User"

// User用のUserStore
type UserStore struct {
	ctx context.Context
}
