package errors

import (
	"code.oldboyedu.com/demo/src/share/config"
	"github.com/micro/go-micro/errors"
)

const (
	errorCodeFilmSuccess = 200
)

var (
	ErrorFilmFailed = errors.New(
		config.ServiceNameUser, "操作异常", errorCodeFilmSuccess,
	)
)
