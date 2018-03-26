package service

import (
	"cloud.google.com/go/bigquery"
	"context"
	"google.golang.org/api/iterator"
	"reflect"
)

// bigquery.Clientをwrapする。
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

// 指定したBigQueryのDataset.Tableにデータをアップロードする
func (bq *BQClientWrapper) PutData(dataset, table string, src interface{}) error {

	// BigQueryの指定したDataset.TableのUploaderを作成
	upl := bq.client.Dataset(dataset).Table(table).Uploader()
	// 不適切なデータが含まれていた場合に、無視する
	upl.SkipInvalidRows = true

	// アップロードする
	return upl.Put(bq.ctx, src)
}

// BigQueryにQueryを発行して、結果をロードする。
func (bq *BQClientWrapper) QueryAndLoad(sql string, dst interface{}) error {
	// SQLからQueryを発行する
	query := bq.client.Query(sql)
	// スタンダードSQLを使用する
	query.QueryConfig.UseStandardSQL = true

	// Queryを実行して、RowIteratorを取得
	it, err := query.Read(bq.ctx)
	if err != nil {
		return err
	}

	if err := loadBQResult(it, dst); err != nil {
		return err
	}

	return nil

}

// BigQueryのQueryの結果をロードする。
func loadBQResult(it *bigquery.RowIterator, dst interface{}) error {
	var bs []interface{}
	for {
		var vs []interface{}
		if err := it.Next(&vs); err == iterator.Done {
			break
		} else if err != nil {
			return err
		}

		bs = append(bs, vs...)
	}

	dst = &bs
	return nil
}
