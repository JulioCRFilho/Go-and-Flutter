package repository

import (
	"context"
	"time"
)

var defaultTimeOut time.Duration = 45

func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeOut*time.Second)
}
