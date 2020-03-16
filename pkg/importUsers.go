package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/auth0.v3/management"
	"os"
	"time"
)

func ImportUsers(jobManager *management.JobManager, connection string, usersFile string, updateExistingUsers bool) {

	format := "json"
	sendCompletionEmail := false
	job := management.Job{
		ConnectionID:        &connection,
		Format:              &format,
		Upsert:              &updateExistingUsers,
		SendCompletionEmail: &sendCompletionEmail,
		Users:               ReadUsersFile(usersFile),
	}

	err := jobManager.ImportUsers(&job)
	if err != nil {
		panic(err)
	}

	j, err := jobManager.Read(*job.ID)
	for ok := true; ok; ok = *j.Status == "pending" {
		fmt.Println("Job status " + *j.Status + " waiting 5 seconds.")
		time.Sleep(5 * time.Second)
		j, _ = jobManager.Read(*job.ID)
	}

	fmt.Printf("Response: %+v\n", j)
}

func ReadUsersFile(filename string) []map[string]interface{} {
	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		panic(err)
	}
	var userList []map[string]interface{}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		var user map[string]interface{}
		json.Unmarshal(fileScanner.Bytes(), &user)
		userList = append(userList, user)
	}
	return userList
}
