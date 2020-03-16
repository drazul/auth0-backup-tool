package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Summary      map[string]int16    `json:"summary"`
}

func (c client) GetJobStatus(j Job) Job {
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

func (c client) WaitUntilJobFinish(j Job) Job {
	for ok := true; ok; ok = j.Status == "pending" {
		fmt.Println("Job status " + j.Status + " waiting 5 seconds.")
		time.Sleep(5 * time.Second)
		j = c.GetJobStatus(j)
	}
	return j
}
