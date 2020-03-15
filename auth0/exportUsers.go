package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type exportUsersRequest struct {
	ConnectionId string              `json:"connection_id"`
	Format       string              `json:"format"`
	Fields       []map[string]string `json:"fields"`
}

func (c client) ExportUsers(filename string) {
	fmt.Printf("filename: %+v\n", filename)
	job := c.requestExportUsers()
	fmt.Printf("job: %+v\n", job)
	result := c.waitUntilJobFinish(job)
	fmt.Printf("result: %+v\n", result)
}

func (c client) requestExportUsers() string {
	b := exportUsersRequest{
		ConnectionId: c.Connection,
		Format:       "json",
	}
	for _, value := range c.UserAttributes {
		b.Fields = append(b.Fields, map[string]string{"name": value})
	}

	jsonData, _ := json.Marshal(b)

	client := &http.Client{}
	req, err := http.NewRequest(
		"Post",
		"https://"+c.Domain+"/api/v2/jobs/users-exports",
		bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Add("authorization", "Bearer "+c.login())
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(responseBody))

	fmt.Printf("result: %+v\n", string(jsonData))

	return "job 123fdasf"
}

func (c client) waitUntilJobFinish(job string) string {
	return "result 12314"
}
