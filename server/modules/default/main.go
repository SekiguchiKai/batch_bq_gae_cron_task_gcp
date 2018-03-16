package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/api"
)

const _APIPath = "/api"

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
	apiGin := g.Group(_APIPath)
	api.InitUserAPI(apiGin)
}