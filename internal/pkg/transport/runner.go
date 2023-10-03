package transport

import (
	"errors"
	"github.com/spachava753/go-kit-clean-arch/internal/api"
	"github.com/spachava753/go-kit-clean-arch/internal/api/gen"
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream"
	streamendpoint "github.com/spachava753/go-kit-clean-arch/internal/domain/stream/endpoint"
	streamtransport "github.com/spachava753/go-kit-clean-arch/internal/domain/stream/transport"
	"github.com/spachava753/go-kit-clean-arch/internal/pkg"
	"log"
	"net/http"
	"strconv"
)

func MakeServer(c pkg.AppConfig, streamSvc stream.StreamService) func() error {
	createStreamEndpoint := pkg.Chain[streamendpoint.CreateStreamRequest, streamendpoint.CreateStreamResponse](
		pkg.Annotate[streamendpoint.CreateStreamRequest, streamendpoint.CreateStreamResponse]("calling create stream endpoint"),
	)(streamendpoint.MakeCreateStreamEndpoint(streamSvc))

	getStreamEndpoint := pkg.Chain[streamendpoint.GetStreamRequest, streamendpoint.GetStreamResponse](
		pkg.Annotate[streamendpoint.GetStreamRequest, streamendpoint.GetStreamResponse]("calling get stream endpoint"),
	)(streamendpoint.MakeGetStreamEndpoint(streamSvc))

	deleteStreamEndpoint := pkg.Chain[streamendpoint.DeleteStreamRequest, streamendpoint.DeleteStreamResponse](
		pkg.Annotate[streamendpoint.DeleteStreamRequest, streamendpoint.DeleteStreamResponse]("calling delete stream endpoint"),
	)(streamendpoint.MakeDeleteStreamEndpoint(streamSvc))

	switch c.Server {
	case pkg.ServerRuntimeOpenApiKey:
		return func() error {
			MakeStdHttpHandler("/stream/create", createStreamEndpoint,
				streamtransport.DecodeCreateStreamRequest,
				streamtransport.EncodeCreateStreamResponse)

			MakeStdHttpHandler("/stream/get", getStreamEndpoint,
				streamtransport.DecodeGetStreamRequest,
				streamtransport.EncodeGetStreamResponse)

			MakeStdHttpHandler("/stream/delete", deleteStreamEndpoint,
				streamtransport.DecodeDeleteStreamRequest,
				streamtransport.EncodeDeleteStreamResponse)

			log.Println("starting open api server in port", c.Port)
			if err := http.ListenAndServe(":"+strconv.Itoa(c.Port), nil); !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		}
	case pkg.ServerRuntimeStdHttpKey:
		return func() error {
			createHttpHandle := CreateHttpHandlerFunc(createStreamEndpoint,
				streamtransport.DecodeCreateStreamRequest,
				streamtransport.EncodeCreateStreamResponse)

			getHttpHandle := CreateHttpHandlerFunc(getStreamEndpoint,
				streamtransport.DecodeGetStreamRequest,
				streamtransport.EncodeGetStreamResponse)

			deleteHttpHandle := CreateHttpHandlerFunc(deleteStreamEndpoint,
				streamtransport.DecodeDeleteStreamRequest,
				streamtransport.EncodeDeleteStreamResponse)

			serverIface := api.MakeServerInterface(createHttpHandle, getHttpHandle, deleteHttpHandle)

			handler := gen.Handler(serverIface)

			log.Println("starting std http server in port", c.Port)
			if err := http.ListenAndServe(":"+strconv.Itoa(c.Port), handler); !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		}
	default:
		panic("unknown server runtime")
	}
}
