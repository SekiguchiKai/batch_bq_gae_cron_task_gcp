package service

import (
	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"io"
	"net/http"
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

// 新規にGCSClientWrapperReaderを作成する
func NewGCSClientWrapperReader(r *http.Request, path, bucketName string) (io.ReadCloser, error) {
	ctx := appengine.NewContext(r)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucket := client.Bucket(bucketName)

	rc, err := bucket.Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	return &GCSClientWrapperReader{Reader: rc, client: client}, nil

}
