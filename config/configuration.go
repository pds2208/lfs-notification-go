package config

type Brokers struct {
	BrokerAddress BrokerAddress
}

type BrokerAddress struct {
	Addresses []string
}

type configuration struct {
	LoggingTopic  string
	ProgressTopic string
	AuditTopic    string
	Brokers       []string
}
