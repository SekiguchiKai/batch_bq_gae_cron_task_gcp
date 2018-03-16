package util

import (
	"strconv"
	"context"
	"runtime"
	"google.golang.org/appengine/log"
	"net/http"
	"google.golang.org/appengine"
)

// Errorレベルのログを吐き出す。
func ErrorLog(r *http.Request, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "nofile"
		line = -1
	}

	ctx := appengine.NewContext(r)
	log.Errorf(ctx, file+":"+strconv.Itoa(line)+":"+format, args...)
}

// Warningレベルのログを吐き出す。
func WarningLog(r *http.Request, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "nofile"
		line = -1
	}

	ctx := appengine.NewContext(r)
	log.Warningf(ctx, file+":"+strconv.Itoa(line)+":"+format, args...)
}

// Infoレベルのログを吐き出す。
func InfoLog(r *http.Request, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "nofile"
		line = -1
	}

	ctx := appengine.NewContext(r)
	log.Infof(ctx, file+":"+strconv.Itoa(line)+":"+format, args...)
}

// context.Contextを元にInfoレベルのログを吐き出す。
func InfoLogWithContext(c context.Context, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "nofile"
		line = -1
	}

	log.Infof(c, file+":"+strconv.Itoa(line)+":"+format, args...)
}