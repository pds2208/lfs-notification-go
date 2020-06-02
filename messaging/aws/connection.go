package aws

import (
	"fmt"
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/rs/zerolog"
	"os"
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

	result, err := s.sns.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, t := range result.Topics {
		fmt.Println(*t.TopicArn)
	}
	return nil
}

func (s *Connection) Close() {
	if s.sns != nil {

	}
}

func (s *Connection) LoadConfiguration() {

}
