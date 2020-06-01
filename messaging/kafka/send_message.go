package kafka

import (
	"context"
	"encoding/json"
	"github.com/ONSDigital/lfs-notification-go/types"
	"github.com/segmentio/kafka-go"
)

func (s *Connection) sendMessage(key, message []byte) error {
	return s.conn.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: message,
		},
	)
}

func (s *Connection) SendProgressMessage(message *types.ProgressMessage) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.sendMessage([]byte("lfs-progress"), b)
}
