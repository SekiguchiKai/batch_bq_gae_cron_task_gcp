package service

import (
	"cloud.google.com/go/storage"
)

// storage.ClientをwrapするWriter
type GCSClientWrapperReader struct {
	*storage.Reader
	client *storage.Client
}

// storage.ClientをwrapするReader
type GCSClientWrapperWriter struct {
	*storage.Writer
	client *storage.Client
}

// ClientとReaderをCloseする
func (r *GCSClientWrapperReader) Close() error {
	defer r.client.Close()
	return r.Reader.Close()
}
