package twilio

import (
	"fmt"

	"github.com/twilio/twilio-go"
	conversations "github.com/twilio/twilio-go/rest/conversations/v1"
)

type Client interface {
	CreateConversation(name string) (string, error)
	CreateUser(username, name string) (string, error)
}

type client struct {
	client *twilio.RestClient
}

func NewClient() Client {
	return client{
		client: twilio.NewRestClient(),
	}
}

func (c client) CreateConversation(name string) (string, error) {
	params := &conversations.CreateConversationParams{}
	params.SetFriendlyName(name)

	resp, err := c.client.ConversationsV1.CreateConversation(params)
	if err != nil {
		return "", err
	}

	if resp.Sid == nil {
		return "", fmt.Errorf("CreateConversation request returned nil Sid")
	}

	return *resp.Sid, nil
}

func (c client) CreateUser(username, name string) (string, error) {
	params := &conversations.CreateUserParams{}
	params.SetIdentity(username)
	params.SetFriendlyName(name)

	resp, err := c.client.ConversationsV1.CreateUser(params)
	if err != nil {
		return "", err
	}

	if resp.Sid == nil {
		return "", fmt.Errorf("CreateUser request returned nil Sid")
	}

	return *resp.Sid, nil
}
