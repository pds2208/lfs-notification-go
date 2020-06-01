package config_test

import (
	conf "github.com/ONSDigital/lfs-notification-go/config"
	"testing"
)

func TestAuditTopic(t *testing.T) {
	if conf.Config.AuditTopic != "lfs-audit" {
		t.Errorf("audit Topic, want %s", "lfs-audit")
	}
}

func TestLoggingTopic(t *testing.T) {
	if conf.Config.LoggingTopic != "lfs-logging" {
		t.Errorf("logging Topic, want %s", "lfs-logging")
	}
}

func TestProgressTopic(t *testing.T) {
	if conf.Config.ProgressTopic != "lfs-progress" {
		t.Errorf("progress Topic, want %s", "lfs-progress")
	}
}

func TestBroker(t *testing.T) {

	expectedBroker := "localhost:9094"
	brokers := conf.Config.Brokers

	if brokers == nil {
		t.Errorf("brokers is nil, want %s", expectedBroker)
	} else if len(brokers) != 1 {
		t.Errorf("brokers array size is %d, want %s", len(brokers), expectedBroker)
	} else if brokers[0] != expectedBroker {
		t.Errorf("brokers: want %s, got %s", expectedBroker, brokers[0])
	}

}
