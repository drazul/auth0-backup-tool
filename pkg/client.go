package pkg

import (
	"github.com/google/uuid"
	"time"
)

var DefaultUserAttributes = []string{
	"user_id",
	"given_name",
	"family_name",
	"nickname",
	"name",
	"email",
	"email_verified",
	"created_at",
	"updated_at",
	"app_metadata",
	"user_metadata",
	"blocked",
	"last_password_reset",
	"logins_count",
	"last_login",
	"identities",
}

type Client interface {
	ExportUsers(filename string)
	ImportUsers(usersFile string, updateExistingUsers string)
}

type client struct {
	ClientId       string
	ClientSecret   string
	Domain         string
	UserAttributes []string
	Connection     string
	AuthToken      string
	ValidUntil     time.Time
	CorrelationId  string
}

func New(clientId string, clientSecret string, domain string, connection string, userAttributes []string) Client {
	client := client{
		ClientId:       clientId,
		ClientSecret:   clientSecret,
		Domain:         domain,
		UserAttributes: userAttributes,
		Connection:     connection,
		CorrelationId:  uuid.New().String(),
	}

	if client.UserAttributes[0] == "" {
		client.UserAttributes = DefaultUserAttributes
	}

	return client
}
