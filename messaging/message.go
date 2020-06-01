package messaging

import (
	"fmt"
	"github.com/ONSDigital/lfs-notification-go/messaging/kafka"
	"github.com/ONSDigital/lfs-notification-go/types"
	"github.com/rs/zerolog/log"
	"sync"
)

type MessageService interface {
	Connect(topic string) error
	Close()

	// send a progress message
	SendProgressMessage(message *types.ProgressMessage) error
}

var cachedConnection MessageService
var connectionMux = &sync.Mutex{}

func GetDefaultMessagingImpl(topic string) (MessageService, error) {
	connectionMux.Lock()
	defer connectionMux.Unlock()

	if cachedConnection != nil {
		log.Info().
			Msg("Returning cached messaging service")
		return cachedConnection, nil
	}

	cachedConnection = &kafka.Connection{}

	if err := cachedConnection.Connect(topic); err != nil {
		log.Info().
			Err(err).
			Msg("Cannot connect to messaging service. Is it running?")
		cachedConnection = nil
		return nil, fmt.Errorf("cannot connect to messaging service")
	}

	return cachedConnection, nil
}

//WriteAuditMessage() error
//WriteLogMessage() error
//WriteProgressMessage(types.ProgressMessage) error
