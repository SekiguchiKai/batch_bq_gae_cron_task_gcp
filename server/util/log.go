package util

import (
	"strconv"
	"context"
	"runtime"
	"google.golang.org/appengine/log"
)

// context.Contextを元にInfoレベルのログを吐き出す。
func InfoLogWithContext(c context.Context, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "nofile"
		line = -1
	}

	log.Infof(c, file+":"+strconv.Itoa(line)+":"+format, args...)
}