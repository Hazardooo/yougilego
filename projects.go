package yougilego

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type YGProjectService struct {
	YGEngine            `json:"YGEngine"`
	BugTruckerProjectId string `json:"bugTruckerProjectId"`
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

func (projService *YGProjectService) SetBugTrackerProject(admins []string) (err error) {
	url := "https://ru.yougile.com/api-v2/projects"
	payload := CreateProject{
		Title: projService.Config.BugTruckerProjectName,
		Users: make(map[string]string),
	}
	for _, admin := range admins {
		payload.Users[admin] = "admin"
	}
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.YGEngine.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	newProjectId := SuccessResponse{}
	json.Unmarshal(body, &newProjectId)
	projService.BugTruckerProjectId = newProjectId.Id
	return err
}

func (projService *YGProjectService) GetProjeсts() (err error, projects ListResponse[ProjectResponse]) {
	url := "https://ru.yougile.com/api-v2/projects"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.YGEngine.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &projects)
	return
}

func (projService *YGProjectService) CheckBugTrackerProject(projects ListResponse[ProjectResponse]) bool {
	for _, project := range projects.Content {
		if project.Title == projService.Config.BugTruckerProjectName {
			projService.BugTruckerProjectId = project.Id
			return true
		}
	}
	return false
}
