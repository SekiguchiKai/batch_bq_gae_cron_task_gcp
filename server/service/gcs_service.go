package service

import (
	"cloud.google.com/go/storage"
	"context"
)

// storage.ClientをwrapするWriter
type GCSClientWrapperReader struct {
	ctx    context.Context
	client *storage.Client
}

// storage.ClientをwrapするReader
type GCSClientWrapperWriter struct {
	ctx    context.Context
	client *storage.Client
}
