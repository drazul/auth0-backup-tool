package main

import (
	"auth0-user-management-cli/pkg"
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	ConfigFile     string
	ClientId       string
	ClientSecret   string
	Domain         string
	UsersFile      string
	UserAttributes string
	Connection     string
}

func parseFlags() Flags {
	flags := Flags{}

	flag.StringVar(&flags.ConfigFile, "config-file", "", "Json file where to read the client configuration. "+
		"Configuration flags with have preferences if we use together")
	flag.StringVar(&flags.ClientId, "client-id", "", "Client ID of an application with user management rights")
	flag.StringVar(&flags.ClientSecret, "client-secret", "", "Client secret of an application with user management rights")

	flag.StringVar(&flags.Domain, "domain", "", "Auth0 domain")
	flag.StringVar(&flags.Connection, "connection", "", "Auth0 connection ID")
	flag.StringVar(&flags.UsersFile, "users-file", "users-export.json", "File path where to store the exported users or where to read the users to import")
	flag.StringVar(&flags.UserAttributes, "user-attributes", "", "List of user attributes to export. Format: attr1,attr2,attr3")

	flag.Parse()
	return flags
}

func main() {
	flags := parseFlags()

	auth0 := pkg.New(flags.ClientId, flags.ClientSecret, flags.Domain, flags.Connection, strings.Split(flags.UserAttributes, ","))
	auth0.ExportUsers(flags.UsersFile)

	fmt.Printf("Flags: %+v\n", flags)
	fmt.Printf("Auth0 Client: %+v\n", auth0)
}
