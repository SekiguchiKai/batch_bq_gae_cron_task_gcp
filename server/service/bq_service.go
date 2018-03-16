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
