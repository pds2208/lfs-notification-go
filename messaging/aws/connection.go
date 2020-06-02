package aws

import (
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/rs/zerolog"
)

type Connection struct {
	sns *sns.SNS
}

func (s *Connection) Connect(topic string) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	s.LoadConfiguration()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	s.sns = sns.New(sess)

	return nil
}

func (s *Connection) Close() {
	if s.sns != nil {

	}
}

func (s *Connection) LoadConfiguration() {

}
