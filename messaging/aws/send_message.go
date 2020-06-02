package aws

import (
	"encoding/json"
	"fmt"
	"github.com/ONSDigital/lfs-notification-go/types"
	"github.com/aws/aws-sdk-go/service/sns"
)

func (s *Connection) sendMessage(topic, message []byte) error {
	t := string(topic)
	m := string(message)
	result, err := s.sns.Publish(&sns.PublishInput{
		Message:  &m,
		TopicArn: &t,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(*result.MessageId)
	return err
}

func (s *Connection) SendProgressMessage(message *types.ProgressMessage) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.sendMessage([]byte("arn:aws:sns:eu-west-2:000704438865:LFS-Progress"), b)
}
