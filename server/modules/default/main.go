package main

import (
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/api"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// エントリポイント
func init() {
	g := gin.New()
	initAPI(g)
	// gin.New()の戻り値のEngineは、ServeHTTP(ResponseWriter, *Request)メソッドを持っているので、
	// type Handler interfaceを満たす
	http.Handle("/", g)
}

// API群を初期登録する。
func initAPI(g *gin.Engine) {
	apiGin := g.Group(util.GetApiPath())
	api.InitUserAPI(apiGin)
}
