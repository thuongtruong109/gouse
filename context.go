package gouse

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var CtxBg = context.Background()

func SetCtx(ctx context.Context) {
	CtxBg = ctx
}

func GetCtx() context.Context {
	return CtxBg
}

func CtxCancel() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}
