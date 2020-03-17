package pkg

import (
	"fmt"
	"gopkg.in/auth0.v3/management"
	"os"
	"strings"
)

func ExportFromAllConnections(manager *management.Management, usersFile string) {
	connectionManager := manager.Connection
	connectionList, _ := connectionManager.List()

	folder := strings.ReplaceAll(usersFile, ".json", "")
	os.MkdirAll(folder, os.ModePerm)

	for _, connection := range connectionList.Connections {
		if *connection.Strategy == "auth0" {
			fmt.Printf("\nExporting users from connection: %+v\n", *connection.Name)
			ExportUsers(manager.Job, *connection.ID, []string{""}, folder+"/"+*connection.Name+".json")
		}
	}
}
