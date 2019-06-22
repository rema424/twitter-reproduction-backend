// +build appengine

package logutil

import (
	"context"

	"google.golang.org/appengine/log"
)

func Debugf(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}

	log.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}
	log.Infof(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}
	log.Warningf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}
	// postSlack(c, fmt.Sprintf(format, args...))
	log.Errorf(ctx, format, args...)
}
