package service

import (
	"cloud.google.com/go/storage"
	"context"
)

// storage.Clientをwrapする。
type GCSClientWrapper struct {
	ctx    context.Context
	client *storage.Client
}
