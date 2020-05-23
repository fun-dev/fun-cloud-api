package cloudutil

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	//
	ContextKeyRequestId = "CONTEXT_KEY_REQUEST_ID"
	ContextKeyUserIp    = "CONTEXT_KEY_USER_IP"
	//
	_log     = logrus.New()
	_appName = os.Getenv("APP_NAME")
)

func InfoLog(ctx context.Context, msg string) {
	requestId := ctx.Value(ContextKeyRequestId)
	userIP := ctx.Value(ContextKeyUserIp)

	_log.
		WithField("appName", _appName).
		WithField("requestID", requestId).
		WithField("userIP", userIP).
		Info(msg)
}

func ErrorInfoLog(ctx context.Context, err error, msg string) {
	requestId := ctx.Value(ContextKeyRequestId)
	userIP := ctx.Value(ContextKeyUserIp)

	_log.
		WithField("appName", _appName).
		WithField("requestID", requestId).
		WithField("userIP", userIP).
		WithField("error", err).
		Info(msg)
}

func ErrorLogOnExit(ctx context.Context, msg error) {
	requestId := ctx.Value(ContextKeyRequestId)
	userIP := ctx.Value(ContextKeyUserIp)

	_log.
		WithField("appName", _appName).
		WithField("requestID", requestId).
		WithField("userIP", userIP).
		Fatal(msg.Error())
}
