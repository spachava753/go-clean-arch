package transport

import (
	"fmt"
	"github.com/spachava753/go-kit-clean-arch/internal/pkg"
	"net/http"
)

func MakeStdHttpHandler[Req any, Resp any](path string, e pkg.Endpoint[Req, Resp],
	dec pkg.DecodeRequestFunc[Req],
	enc pkg.EncodeResponseFunc[Resp]) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		request, err := dec(ctx, r)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("err: %s", err)))
			return
		}

		response, err := e(ctx, request)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("err: %s", err)))
			return
		}

		if writeErr := enc(ctx, w, response); writeErr != nil {
			w.Write([]byte(fmt.Sprintf("err: %s", writeErr)))
			return
		}
	})
}
