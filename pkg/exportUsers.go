package pkg

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ExportUsersRequest struct {
	ConnectionId string              `json:"connection_id"`
	Format       string              `json:"format"`
	Fields       []map[string]string `json:"fields"`
}

func (c client) ExportUsers(usersFile string) {
	fmt.Printf("filename: %+v\n", usersFile)
	var job = Job{}
	job = c.RequestExportUsers()
	job = c.WaitUntilJobFinish(job)
	c.DownloadFile(job, usersFile)
}

func (c client) RequestExportUsers() Job {
	b := ExportUsersRequest{
		ConnectionId: c.Connection,
		Format:       "json",
	}
	for _, value := range c.UserAttributes {
		b.Fields = append(b.Fields, map[string]string{"name": value})
	}

	jsonData, _ := json.Marshal(b)

	req, _ := http.NewRequest("POST",
		"https://"+c.Domain+"/api/v2/jobs/users-exports",
		strings.NewReader(string(jsonData)))

	req.Header.Add("authorization", "Bearer "+c.login())
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	job := Job{}
	_ = json.Unmarshal(body, &job)

	return job
}

func (c client) DownloadFile(j Job, usersFile string) {
	fmt.Println("Downloading ", j.Location, " to ", usersFile)

	resp, err := http.Get(j.Location)
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
