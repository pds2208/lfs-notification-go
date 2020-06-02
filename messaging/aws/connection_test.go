package aws_test

import (
	"github.com/ONSDigital/lfs-notification-go/messaging"
	"github.com/ONSDigital/lfs-notification-go/producer"
	"os"
	"testing"
)

var conn messaging.MessageService

func TestMain(m *testing.M) {
	setupConnectionTest()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

// setup
func setupConnectionTest() {
	var err error
	conn, err = messaging.GetSNSMessagingImpl("lfs-progress")
	if err != nil {
		panic("cannot connect to messaging service")
	}
}

// teardown
func shutdown() {
	if conn != nil {
		conn.Close()
	}
}

func TestSendProgressMessage(t *testing.T) {
	fileMessage := producer.NewFileProgress("survey")
	fileMessage.SetPercentage(5)
	err := conn.SendProgressMessage(fileMessage.Message)
	if err != nil {
		panic("cannot send message")
	}

}
