package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/api"
)

const _APIPath = "/api"

// API群を初期登録する。
func initAPI(g *gin.Engine) {
	apiGin := g.Group(_APIPath)
	api.InitUserAPI(apiGin)
}