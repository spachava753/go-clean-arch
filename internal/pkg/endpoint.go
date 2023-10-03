package pkg

import (
	"context"
	"net/http"
)

type Endpoint[Req any, Resp any] func(ctx context.Context, request Req) (response Resp, err error)

type DecodeRequestFunc[Req any] func(context.Context, *http.Request) (request Req, err error)

type EncodeResponseFunc[Resp any] func(context.Context, http.ResponseWriter, Resp) error
