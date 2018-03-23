package util

import "errors"

// 複数の文字列からエラーを作成する。
func CreateErrMessage(args ...string) error {
	ms := ""
	for _, arg := range args {
		ms = ms + arg
	}
	return errors.New(ms)
}