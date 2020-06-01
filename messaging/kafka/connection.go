package kafka

import (
	"github.com/ONSDigital/lfs-notification-go/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

type Connection struct {
	brokers []string
	conn    *kafka.Writer
}

func (s *Connection) Connect(topic string) error {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	s.LoadConfiguration()

	s.conn = kafka.NewWriter(kafka.WriterConfig{
		Brokers:          s.brokers,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Async:            false,
		CompressionCodec: snappy.NewCompressionCodec(),
	})

	log.Debug().
		Msg("Connecting to Kafka Broker(s)")

	return nil
}

func (s *Connection) Close() {
	if s.conn != nil {
		_ = s.conn.Close()
	}
}

func (s *Connection) LoadConfiguration() {
	s.brokers = config.Config.Brokers
	if s.brokers == nil {
		panic("addresses table configuration not set")
	}

}
