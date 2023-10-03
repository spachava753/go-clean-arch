package repository

import (
	"context"
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream"
)

type hardCodedStreamRepo struct{}

func (h *hardCodedStreamRepo) SaveStream(ctx context.Context, stream stream.Stream) error {
	return nil
}

func (h *hardCodedStreamRepo) GetStream(ctx context.Context, name string) (stream.Stream, error) {
	return stream.Stream{
		Name:        "awesome-stream",
		Environment: stream.StreamEnvironmentProduction,
		Platform:    "Near Real-Time Data",
		ProducerGeos: []stream.StreamProducerGeo{
			stream.StreamStreamProducerGeoNA,
		},
		MessageFormat:      stream.StreamMessageJson,
		RetentionGigabytes: 10,
		RetentionHours:     2,
		PartitionCount:     2,
	}, nil
}

func (h *hardCodedStreamRepo) DeleteStream(ctx context.Context, stream stream.Stream) error {
	return nil
}

func MakeHardCodedStreamRepository() stream.StreamRepository {
	return &hardCodedStreamRepo{}
}
