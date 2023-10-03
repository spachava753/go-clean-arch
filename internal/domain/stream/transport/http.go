package transport

import (
	"context"
	"encoding/json"
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream"
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream/endpoint"
	"net/http"
)

type errorJson struct {
	Err error `json:"error"`
}

type nameJson struct {
	Name string `json:"name"`
}

type jsonStream struct {
	Name               string   `json:"name"`
	Environment        string   `json:"environment"`
	Platform           string   `json:"platform"`
	ProducerGeos       []string `json:"producerGeos"`
	MessageFormat      string   `json:"messageFormat"`
	RetentionGigabytes int      `json:"retentionGigabytes"`
	RetentionHours     int      `json:"retentionHours"`
	PartitionCount     int      `json:"partitionCount"`
}

func streamToJsonStream(s stream.Stream) jsonStream {
	var env string
	switch s.Environment {
	case stream.StreamEnvironmentStaging:
		env = "staging"
	case stream.StreamEnvironmentProduction:
		env = "production"
	}

	var producerGeos []string
	for i := range s.ProducerGeos {
		switch s.ProducerGeos[i] {
		case stream.StreamStreamProducerGeoNA:
			producerGeos = append(producerGeos, "NA")
		case stream.StreamStreamProducerGeoEU:
			producerGeos = append(producerGeos, "EU")
		}
	}
	return jsonStream{
		Name:               s.Name,
		Environment:        env,
		Platform:           s.Platform,
		ProducerGeos:       producerGeos,
		MessageFormat:      string(s.MessageFormat),
		RetentionGigabytes: s.RetentionGigabytes,
		PartitionCount:     s.PartitionCount,
	}
}

func jsonStreamToStream(s jsonStream) (stream.Stream, error) {
	var env stream.StreamEnvironment
	switch s.Environment {
	case "staging":
		env = stream.StreamEnvironmentStaging
	case "production":
		env = stream.StreamEnvironmentProduction
	}

	var producerGeos []stream.StreamProducerGeo
	for i := range s.ProducerGeos {
		switch s.ProducerGeos[i] {
		case "NA":
			producerGeos = append(producerGeos, stream.StreamStreamProducerGeoNA)
		case "EU":
			producerGeos = append(producerGeos, stream.StreamStreamProducerGeoEU)
		}
	}

	var msgFormat stream.StreamMessageFormat
	switch s.MessageFormat {
	case string(stream.StreamMessageAvro),
		string(stream.StreamMessageProtobuf),
		string(stream.StreamMessageJsonSchema),
		string(stream.StreamMessageJson),
		string(stream.StreamMessageString),
		string(stream.StreamMessageByteArray):
		msgFormat = stream.StreamMessageFormat(s.MessageFormat)
	}

	return stream.Stream{
		Name:               s.Name,
		Environment:        env,
		Platform:           s.Platform,
		ProducerGeos:       producerGeos,
		MessageFormat:      msgFormat,
		RetentionGigabytes: s.RetentionGigabytes,
		PartitionCount:     s.PartitionCount,
	}, nil
}

func DecodeCreateStreamRequest(_ context.Context, r *http.Request) (endpoint.CreateStreamRequest, error) {
	var j jsonStream
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		return endpoint.CreateStreamRequest{}, err
	}
	s, decodeErr := jsonStreamToStream(j)
	if decodeErr != nil {
		return endpoint.CreateStreamRequest{}, decodeErr
	}
	return endpoint.CreateStreamRequest{Stream: s}, nil
}

func EncodeCreateStreamResponse(_ context.Context, w http.ResponseWriter, response endpoint.CreateStreamResponse) error {
	if response.CreateErr != nil {
		w.WriteHeader(500)
		return json.NewEncoder(w).Encode(errorJson{Err: response.CreateErr})
	}
	_, writeErr := w.Write([]byte("{}"))
	return writeErr
}

func DecodeGetStreamRequest(_ context.Context, r *http.Request) (endpoint.GetStreamRequest, error) {
	var n nameJson
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		return endpoint.GetStreamRequest{}, err
	}
	return endpoint.GetStreamRequest{Name: n.Name}, nil
}

func EncodeGetStreamResponse(_ context.Context, w http.ResponseWriter, response endpoint.GetStreamResponse) error {
	if response.GetStreamErr != nil {
		w.WriteHeader(500)
		return json.NewEncoder(w).Encode(errorJson{Err: response.GetStreamErr})
	}
	return json.NewEncoder(w).Encode(streamToJsonStream(response.S))
}

func DecodeDeleteStreamRequest(_ context.Context, r *http.Request) (endpoint.DeleteStreamRequest, error) {
	var n nameJson
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		return endpoint.DeleteStreamRequest{}, err
	}
	return endpoint.DeleteStreamRequest{Name: n.Name}, nil
}

func EncodeDeleteStreamResponse(_ context.Context, w http.ResponseWriter, response endpoint.DeleteStreamResponse) error {
	if response.DeleteStreamErr != nil {
		w.WriteHeader(500)
		return json.NewEncoder(w).Encode(errorJson{Err: response.DeleteStreamErr})
	}
	_, writeErr := w.Write([]byte("{}"))
	return writeErr
}
