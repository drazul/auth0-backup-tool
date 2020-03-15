package auth0

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
}

type client struct {
	ClientId       string
	ClientSecret   string
	Domain         string
	UserAttributes []string
	Connection     string
}

func New(clientId string, clientSecret string, domain string, connection string, userAttributes []string) Client {
	client := client{
		ClientId:       clientId,
		ClientSecret:   clientSecret,
		Domain:         domain,
		UserAttributes: userAttributes,
		Connection:     connection,
	}

	if client.UserAttributes[0] == "" {
		client.UserAttributes = DefaultUserAttributes
	}

	return client
}
