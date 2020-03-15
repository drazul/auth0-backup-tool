package pkg

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Job struct {
	Type         string              `json:"type"`
	Status       string              `json:"status"`
	ConnectionId []map[string]string `json:"connection_id"`
	Format       string              `json:"format"`
	Fields       []map[string]string `json:"fields"`
	Connection   string              `json:"connection"`
	CreatedAt    string              `json:"created_at"`
	Id           string              `json:"id"`
	Location     string              `json:"location"`
}

func (c client) getJobStatus(j Job) Job {
	req, _ := http.NewRequest("GET",
		"https://"+c.Domain+"/api/v2/jobs/"+j.Id,
		nil)
	req.Header.Add("authorization", "Bearer "+c.login())
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("job response: %+v\n", string(body))
	var job = Job{}
	_ = json.Unmarshal(body, &job)
	return job
}

func (c client) waitUntilJobFinish(j Job) Job {
	for ok := true; ok; ok = j.Status == "pending" {
		fmt.Println("Job status " + j.Status + " waiting 5 seconds.")
		time.Sleep(5 * time.Second)
		j = c.getJobStatus(j)
	}
	return j
}

func (c client) downloadFile(j Job, filename string) {
	fmt.Println("Downloading ", j.Location, " to ", filename)

	resp, err := http.Get(j.Location)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	zip, _ := gzip.NewReader(resp.Body)
	zip.Multistream(false)

	f, err := os.Create(filename)
	_, err = io.Copy(f, zip)

	if err != nil {
		panic(err)
	}
}
