package aws

import (
	"encoding/json"
	"fmt"
	"github.com/ONSDigital/lfs-notification-go/types"
	"github.com/aws/aws-sdk-go/service/sns"
)

func (s *Connection) sendMessage(key, message []byte) error {
	t := string(key)
	m := string(message)
	result, err := s.sns.Publish(&sns.PublishInput{
		Message:  &m,
		TopicArn: &t,
	})

	fmt.Println(*result.MessageId)
	return err

}

func (s *Connection) SendProgressMessage(message *types.ProgressMessage) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.sendMessage([]byte("lfs-progress"), b)
}
