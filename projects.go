package yougilego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type YGProjectService struct {
	Key                 string `json:"key"`
	BugTruckerProjectId string `json:"bugTruckerProjectId"`
}

func (projService *YGProjectService) UseKey() string {
	return fmt.Sprintf("Bearer %s", projService.Key)
}

func (projService *YGProjectService) GetProjeсts() (err error, projects ListResponse[ProjectResponse]) {
	url := "https://ru.yougile.com/api-v2/projects"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &projects)
	return
}

type CreateProject struct {
	Title string            `json:"title"`
	Users map[string]string `json:"users"`
}

type ProjectResponse struct {
	Deleted   bool              `json:"deleted"`
	Id        string            `json:"id"`
	Title     string            `json:"title"`
	Timestamp int64             `json:"timestamp"`
	Users     map[string]string `json:"users"`
}
