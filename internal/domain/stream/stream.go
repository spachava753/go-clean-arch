package stream

// Below is one way to create an enum
type StreamMessageFormat string

const (
	StreamMessageAvro       StreamMessageFormat = "avro"
	StreamMessageProtobuf   StreamMessageFormat = "protobuf"
	StreamMessageJsonSchema StreamMessageFormat = "jsonSchema"
	StreamMessageJson       StreamMessageFormat = "json"
	StreamMessageString     StreamMessageFormat = "string"
	StreamMessageByteArray  StreamMessageFormat = "byteArray"
)

// Here is another way to make enums

type StreamEnvironment int

const (
	StreamEnvironmentStaging StreamEnvironment = iota
	StreamEnvironmentProduction
)

type StreamProducerGeo int

const (
	StreamStreamProducerGeoNA StreamProducerGeo = iota
	StreamStreamProducerGeoEU
)

type Stream struct {
	Name               string
	Environment        StreamEnvironment
	Platform           string
	ProducerGeos       []StreamProducerGeo
	MessageFormat      StreamMessageFormat
	RetentionGigabytes int
	RetentionHours     int
	PartitionCount     int
}
