package stream

import (
	"context"
	"errors"
	"fmt"
)

// Errors

// StreamDoesNotExistsErr is an error type for a stream that doesn't exist
type StreamDoesNotExistsErr struct {
	Name string
}

func (s StreamDoesNotExistsErr) Error() string {
	return fmt.Sprintf("stream with name %s does not exist", s.Name)
}

type StreamService interface {
	CreateStream(ctx context.Context, stream Stream) error
	GetStream(ctx context.Context, name string) (Stream, error)
	DeleteStream(ctx context.Context, name string) error
}

type streamServiceImpl struct {
	repository StreamRepository
}

func (s *streamServiceImpl) CreateStream(ctx context.Context, newStream Stream) error {
	// validate stream
	if validationErr := validateNewStream(newStream); validationErr != nil {
		return validationErr
	}
	_, err := s.repository.GetStream(ctx, newStream.Name)
	if err == nil {
		return fmt.Errorf("stream with name %s already exists", newStream.Name)
	}
	if !errors.As(err, &StreamDoesNotExistsErr{}) {
		return err
	}
	return s.repository.SaveStream(ctx, newStream)
}

func (s *streamServiceImpl) GetStream(ctx context.Context, name string) (Stream, error) {
	return s.repository.GetStream(ctx, name)
}

func (s *streamServiceImpl) DeleteStream(ctx context.Context, name string) error {
	fetchedStream, err := s.repository.GetStream(ctx, name)
	if err != nil {
		return err
	}
	return s.repository.DeleteStream(ctx, fetchedStream)
}

func MakeStreamService(repository StreamRepository) StreamService {
	if repository == nil {
		panic("stream repository passed in is nil")
	}
	return &streamServiceImpl{repository}
}

// validateNewStream validates that a newly created stream is correctly formed
func validateNewStream(stream Stream) error {
	return nil
}

type StreamRepository interface {
	SaveStream(ctx context.Context, stream Stream) error
	GetStream(ctx context.Context, name string) (Stream, error)
	DeleteStream(ctx context.Context, stream Stream) error
}
