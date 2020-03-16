package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func (c client) ImportUsers(usersFile string, updateExistingUsers string) {
	var job = Job{}
	job = c.RequestImportUsers(usersFile, updateExistingUsers)
	job = c.WaitUntilJobFinish(job)
	fmt.Printf("Import users job: %+v\n", job)
	fmt.Printf("Import users job summary: %+v\n", job.Summary)
}

func (c client) RequestImportUsers(usersFile string, updateExistingUsers string) Job {
	file, _ := os.Open(usersFile)
	fileContents, _ := ioutil.ReadAll(file)
	fi, _ := file.Stat()
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("users", fi.Name())

	part.Write(fileContents)

	_ = writer.WriteField("connection_id", c.Connection)
	_ = writer.WriteField("external_id", c.CorrelationId)
	_ = writer.WriteField("upsert", updateExistingUsers)
	_ = writer.WriteField("send_completion_email", "false")

	_ = writer.Close()

	req, _ := http.NewRequest("POST",
		"https://"+c.Domain+"/api/v2/jobs/users-imports",
		body)

	req.Header.Add("authorization", "Bearer "+c.login())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(resBody))
	time.Sleep(1000 * time.Second)
	//job := Job{}
	//_ = json.Unmarshal(body, &job)

	return Job{}
}

func (c client) ReadFile(usersFile string) string {
	content, err := ioutil.ReadFile(usersFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
