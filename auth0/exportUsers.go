package auth0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type exportUsersRequest struct {
	ConnectionId string              `json:"connection_id"`
	Format       string              `json:"format"`
	Fields       []map[string]string `json:"fields"`
}

func (c client) ExportUsers(filename string) {
	fmt.Printf("filename: %+v\n", filename)
	var job = Job{}
	job = c.requestExportUsers()
	job = c.waitUntilJobFinish(job)
	c.downloadFile(job, filename)
}

func (c client) requestExportUsers() Job {
	b := exportUsersRequest{
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
