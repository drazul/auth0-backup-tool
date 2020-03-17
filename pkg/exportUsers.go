package pkg

import (
	"compress/gzip"
	"fmt"
	"gopkg.in/auth0.v3/management"
	"io"
	"net/http"
	"os"
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

func ExportUsers(jobManager *management.JobManager, connection string, usersAttributes []string, usersFile string) {
	format := "json"
	job := management.Job{
		ConnectionID: &connection,
		Format:       &format,
	}

	AddUserAttributes(&job, usersAttributes)

	err := jobManager.ExportUsers(&job)
	if err != nil {
		panic(err)
	}

	j, err := jobManager.Read(*job.ID)
	for ok := true; ok; ok = *j.Status == "pending" {
		fmt.Println("Job status " + *j.Status + " waiting 5 seconds.")
		time.Sleep(5 * time.Second)
		j, _ = jobManager.Read(*job.ID)
	}

	DownloadFile(*j.Location, usersFile)
}

func AddUserAttributes(job *management.Job, usersAttributes []string) {
	attributes := usersAttributes
	if attributes[0] == "" {
		attributes = DefaultUserAttributes
	}

	for _, value := range attributes {
		field := map[string]interface{}{"name": value}
		job.Fields = append(job.Fields, field)
	}
}

func DownloadFile(url string, usersFile string) {
	fmt.Println("Downloading ", url, " to ", usersFile)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	zip, _ := gzip.NewReader(resp.Body)
	zip.Multistream(false)

	f, err := os.Create(usersFile)
	_, err = io.Copy(f, zip)

	if err != nil {
		panic(err)
	}
}
