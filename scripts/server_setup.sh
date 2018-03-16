#! /bin/sh

go get -u github.com/golang/dep/cmd/dep
cd $GOPATH/src/github.com/SekiguchiKai/batch_bq_task_gcp/server
dep ensure