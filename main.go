package main

import (
	"auth0-backup-tool/pkg"
	"flag"
	"fmt"
	"gopkg.in/auth0.v3/management"
	"os"
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
	Action         string
}

var RequiredFlags = []string{
	"client-id",
	"client-secret",
	"domain",
	"action",
}

func parseFlags() Flags {
	flags := Flags{}

	flag.StringVar(&flags.ClientId, "client-id", "", "Client ID of an application with user management rights")
	flag.StringVar(&flags.ClientSecret, "client-secret", "", "Client secret of an application with user management rights")
	flag.StringVar(&flags.Domain, "domain", "", "Auth0 domain")
	flag.StringVar(&flags.Connection, "connection", "", "Auth0 connection ID. If it's empty a backup to all connections will be performed")
	flag.StringVar(&flags.UsersFile, "users-file", "users-export.json", "File path where to store the exported users or where to read the users to import. When export all connections this flag contains the folder name where to store all exported files")
	flag.StringVar(&flags.UserAttributes, "user-attributes", "", "List of user attributes to export. Format: attr1,attr2,attr3")
	flag.StringVar(&flags.Action, "action", "", "Action to perform. Can be 'import' or 'export'")

	flag.Parse()

	checkNeededFlags(flags)
	return flags
}

func abortWithCommandHelp() {
	fmt.Fprintf(os.Stderr, "\n\nUsage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2) // the same exit code flag.Parse uses
}

func checkNeededFlags(flags Flags) {
	seen := make(map[string]bool)

	var missingFlags = false
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range RequiredFlags {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			missingFlags = true
		}
	}
	if missingFlags {
		abortWithCommandHelp()
	}
}

func main() {
	flags := parseFlags()

	manager, _ := management.New(flags.Domain, flags.ClientId, flags.ClientSecret)

	switch flags.Action {
	case "export":
		if flags.Connection == "" {
			pkg.ExportFromAllConnections(manager, flags.UsersFile)
		} else {
			pkg.ExportUsers(manager.Job, flags.Connection, strings.Split(flags.UserAttributes, ","), flags.UsersFile)
		}
	case "import":
		pkg.ImportUsers(manager.Job, flags.Connection, flags.UsersFile, false)
	default:
		fmt.Fprintf(os.Stderr, "Wrong action '%s'", flags.Action)
		abortWithCommandHelp()
	}
}
