package yougilego

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type YGProjectService struct {
	Key string `json:"key"`
}

func (projService *YGProjectService) UseKey() string {
	return fmt.Sprintf("Bearer %s", projService.Key)
}

func (projService *YGProjectService) GetProjeсts() (err error, response ListResponse[ProjectResponse]) {
	url := "https://ru.yougile.com/api-v2/projects"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetProjeсts StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (projService *YGProjectService) CreateProject(createProjectRequest ProjectRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/projects"
	payloadByte, _ := json.Marshal(createProjectRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateProject StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (projService *YGProjectService) GetProjectById(projectId string) (err error, response ProjectResponse) {
	url := "https://ru.yougile.com/api-v2/projects/" + projectId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetProjectById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (projService *YGProjectService) EditProject(projectId string, createProjectRequest ProjectRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/projects/" + projectId
	payloadByte, _ := json.Marshal(createProjectRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", projService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditProject StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type ProjectRequest struct {
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
