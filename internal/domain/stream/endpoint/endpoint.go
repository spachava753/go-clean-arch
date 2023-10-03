package endpoint

import (
	"context"
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream"
	"github.com/spachava753/go-kit-clean-arch/internal/pkg"
)

type CreateStreamRequest struct {
	Stream stream.Stream
}
type CreateStreamResponse struct {
	CreateErr error
}

func MakeCreateStreamEndpoint(svc stream.StreamService) pkg.Endpoint[CreateStreamRequest, CreateStreamResponse] {
	return func(ctx context.Context, request CreateStreamRequest) (CreateStreamResponse, error) {
		err := svc.CreateStream(ctx, request.Stream)
		return CreateStreamResponse{err}, nil
	}
}

type GetStreamRequest struct {
	Name string
}

type GetStreamResponse struct {
	S            stream.Stream
	GetStreamErr error
}

func MakeGetStreamEndpoint(svc stream.StreamService) pkg.Endpoint[GetStreamRequest, GetStreamResponse] {
	return func(ctx context.Context, request GetStreamRequest) (GetStreamResponse, error) {
		fetchedStream, err := svc.GetStream(ctx, request.Name)
		return GetStreamResponse{fetchedStream, err}, nil
	}
}

type DeleteStreamRequest struct {
	Name string
}

type DeleteStreamResponse struct {
	DeleteStreamErr error
}

func MakeDeleteStreamEndpoint(svc stream.StreamService) pkg.Endpoint[DeleteStreamRequest, DeleteStreamResponse] {
	return func(ctx context.Context, request DeleteStreamRequest) (DeleteStreamResponse, error) {
		err := svc.DeleteStream(ctx, request.Name)
		return DeleteStreamResponse{err}, nil
	}
}
