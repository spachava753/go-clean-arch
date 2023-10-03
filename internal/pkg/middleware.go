package pkg

import (
	"context"
	"fmt"
	"time"
)

type Middleware[Req any, Resp any] func(Endpoint[Req, Resp]) Endpoint[Req, Resp]

func Measure[Req any, Resp any]() Middleware[Req, Resp] {
	return func(next Endpoint[Req, Resp]) Endpoint[Req, Resp] {
		return func(ctx context.Context, request Req) (Resp, error) {
			now := time.Now()
			defer fmt.Printf("took %d milliseconds", time.Since(now).Milliseconds())
			return next(ctx, request)
		}
	}
}

func Annotate[Req any, Resp any](s string) Middleware[Req, Resp] {
	return func(next Endpoint[Req, Resp]) Endpoint[Req, Resp] {
		return func(ctx context.Context, request Req) (Resp, error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			return next(ctx, request)
		}
	}
}

func LogIt[Req fmt.Stringer, Resp fmt.Stringer]() Middleware[Req, Resp] {
	return func(next Endpoint[Req, Resp]) Endpoint[Req, Resp] {
		return func(ctx context.Context, request Req) (Resp, error) {
			fmt.Printf("endpoint middleware req: %s\n", request)
			return next(ctx, request)
		}
	}
}

func Chain[Req any, Resp any](outer Middleware[Req, Resp], others ...Middleware[Req, Resp]) Middleware[Req, Resp] {
	return func(next Endpoint[Req, Resp]) Endpoint[Req, Resp] {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}
