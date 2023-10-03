package api

import (
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/spachava753/go-kit-clean-arch/internal/api/gen"
	"net/http"
)

type streamApiImpl struct {
	create, get, delete http.Handler
}

func (s streamApiImpl) CreateStream(w http.ResponseWriter, r *http.Request) {
	s.create.ServeHTTP(w, r)
}

func (s streamApiImpl) DeleteStream(w http.ResponseWriter, r *http.Request, id types.UUID) {
	s.delete.ServeHTTP(w, r)
}

func (s streamApiImpl) GetStream(w http.ResponseWriter, r *http.Request, id types.UUID) {
	s.get.ServeHTTP(w, r)
}

func MakeServerInterface(create, get, delete http.Handler) gen.ServerInterface {
	return &streamApiImpl{create, get, delete}
}
