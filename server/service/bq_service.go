package service

import (
	"context"
	"cloud.google.com/go/bigquery"
)

// bigquery.Clientをwrapする
type BQClientWrapper struct {
	ctx    context.Context
	client *bigquery.Client
}

// BQClientWrapperを生成する。
func NewBQClientWrapper(ctx context.Context, prjID string) (BQClientWrapper, error) {
	// context.Contextとappengine.AppID(project ID)からBigQueryのClientLibraryを作成した
	client, err := bigquery.NewClient(ctx, prjID)
	if err != nil {
		return BQClientWrapper{}, err
	}

	return BQClientWrapper{client: client}, nil
}