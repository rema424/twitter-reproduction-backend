// +build !appengine

package logutil

import (
	"context"
	"log"
)

// Debugf ...
func Debugf(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}

	log.Printf(format, args...)
}

// Infof ...
func Infof(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}

	log.Printf(format, args...)
}

// Warningf ...
func Warningf(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}

	log.Printf(format, args...)
}

// Errorf ...
func Errorf(ctx context.Context, format string, args ...interface{}) {
	if ctx == nil {
		return
	}

	// postSlack(c, fmt.Sprintf(format, args...))
	log.Printf(format, args...)
}
