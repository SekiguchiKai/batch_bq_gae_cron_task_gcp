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

// BQClientWrapperを生成する
func NewBQClientWrapper(ctx context.Context, prjID string) (BQClientWrapper, error) {
	// context.Contextとappengine.AppID(project ID)からBigQueryのClientLibraryを作成した
	client, err := bigquery.NewClient(ctx, prjID)
	if err != nil {
		return BQClientWrapper{}, err
	}

	return BQClientWrapper{client: client}, nil
}

// 指定したBigQueryのDataset.Tableにデータをアップロードする。
func (bq *BQClientWrapper) PutData(dataset, table string, src interface{}) error {

	// BigQueryの指定したDataset.TableのUploaderを作成
	upl := bq.client.Dataset(dataset).Table(table).Uploader()
	// 不適切なデータが含まれていた場合に、無視する
	upl.SkipInvalidRows = true

	// アップロードする
	return upl.Put(bq.ctx, src)
}