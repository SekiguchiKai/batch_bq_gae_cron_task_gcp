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
func NewGCSClientWrapperReader(r *http.Request, bucketName, path string) (io.ReadCloser, error) {
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

// 新規にGCSClientWrapperWriterを作成する
func NewGCSClientWrapperWriter(r *http.Request, bucketName, path, contentType string) (io.WriteCloser, error) {
	ctx := appengine.NewContext(r)
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	bucket := client.Bucket(bucketName)

	wc := bucket.Object(path).NewWriter(ctx)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		wc.ContentType = contentType
	}

	return &GCSClientWrapperWriter{Writer: wc, client: client}, nil
}
