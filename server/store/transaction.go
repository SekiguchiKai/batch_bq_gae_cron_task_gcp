package store

import (
	"context"
	"net/http"

	gcontext "golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// transaction内で引数で渡された関数fを実行する。
func RunInTransaction(r *http.Request, f func(context.Context) error) error {
	ctx := appengine.NewContext(r)

	return datastore.RunInTransaction(ctx, func(ctx gcontext.Context) error {

		return f(ctx)

	}, &datastore.TransactionOptions{XG: true})
}