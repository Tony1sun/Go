package errors

import (
	"code.oldboyedu.com/demo/src/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeCommentSuccess = 200
)

var (
	ErrorCommentFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeCommentSuccess,
	)
)
