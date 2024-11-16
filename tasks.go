package yougilego

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type YGTaskService struct {
	Key string `json:"key"`
}

func (taskService *YGTaskService) UseKey() string {
	return fmt.Sprintf("Bearer %s", taskService.Key)
}

func (taskService *YGTaskService) GetTasks(columnId string) (err error, tasks ListResponse[TaskResponse]) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/tasks?columnId=%s", columnId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetTasks StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &tasks)
	return
}

func (taskService *YGTaskService) CreateTask(createTaskRequest CreateTask) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/tasks"
	payloadByte, _ := json.Marshal(createTaskRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateTask StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (taskService *YGTaskService) GetTaskById(taskId string) (err error, response TaskResponse) {
	url := "https://ru.yougile.com/api-v2/tasks/" + taskId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetColumnById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (taskService *YGTaskService) EditTask(taskId string, editTaskRequest EditTaskRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/tasks/" + taskId
	payloadByte, _ := json.Marshal(editTaskRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditTask StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (taskService *YGTaskService) GetUserListTaskChat(taskId string) (err error, response UserListTaskChatResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/tasks/%s/chat-subscribers", taskId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetUserListTaskChat StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (taskService *YGTaskService) EditUserListTaskChat(taskId string, editRequest UserListTaskChatResponse) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/tasks/%s/chat-subscribers", taskId)
	payloadByte, _ := json.Marshal(editRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditUserListTaskChat StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type UserListTaskChatResponse struct {
	Content []string `json:"content"`
}

type EditTaskRequest struct {
	Deleted     bool     `json:"deleted"`
	Title       string   `json:"title"`
	ColumnId    string   `json:"columnId"`
	Description string   `json:"description"`
	Archived    bool     `json:"archived"`
	Completed   bool     `json:"completed"`
	Subtasks    []string `json:"subtasks"`
	Assigned    []string `json:"assigned"`
	Deadline    struct {
		Deadline  int64 `json:"deadline"`
		StartDate int64 `json:"startDate"`
		WithTime  bool  `json:"withTime"`
		Deleted   bool  `json:"deleted"`
	} `json:"deadline"`
	TimeTracking struct {
		Plan    int  `json:"plan"`
		Work    int  `json:"work"`
		Deleted bool `json:"deleted"`
	} `json:"timeTracking"`
	Checklists []struct {
		Title string `json:"title"`
		Items []struct {
			Title       string `json:"title"`
			IsCompleted bool   `json:"isCompleted"`
		} `json:"items"`
	} `json:"checklists"`
	Stickers struct {
		Fbc30A9B42D04Cf780C031Fb048346F9 string `json:"fbc30a9b-42d0-4cf7-80c0-31fb048346f9"`
		Ca1Ae84514914DC070351Dd905       string `json:"645250ca-1ae8-4514-914d-c070351dd905"`
	} `json:"stickers"`
	Color string `json:"color"`
	Timer struct {
		Running bool `json:"running"`
		Seconds int  `json:"seconds"`
		Deleted bool `json:"deleted"`
	} `json:"timer"`
	Stopwatch struct {
		Running bool `json:"running"`
		Deleted bool `json:"deleted"`
	} `json:"stopwatch"`
}

type CreateTask struct {
	Title       string   `json:"title"`
	ColumnId    string   `json:"columnId"`
	Description string   `json:"description"`
	Archived    bool     `json:"archived"`
	Completed   bool     `json:"completed"`
	Subtasks    []string `json:"subtasks"`
	Assigned    []string `json:"assigned"`
	Deadline    struct {
		Deadline  int64 `json:"deadline"`
		StartDate int64 `json:"startDate"`
		WithTime  bool  `json:"withTime"`
	} `json:"deadline"`
	TimeTracking struct {
		Plan int `json:"plan"`
		Work int `json:"work"`
	} `json:"timeTracking"`
	Checklists []struct {
		Title string `json:"title"`
		Items []struct {
			Title       string `json:"title"`
			IsCompleted bool   `json:"isCompleted"`
		} `json:"items"`
	} `json:"checklists"`
	Stickers  map[string]string `json:"stickers"`
	Stopwatch struct {
		Running bool `json:"running"`
	} `json:"stopwatch"`
	Timer struct {
		Running bool `json:"running"`
		Seconds int  `json:"seconds"`
	} `json:"timer"`
}

type TaskResponse struct {
	Id                 string   `json:"id"`
	Deleted            bool     `json:"deleted"`
	Title              string   `json:"title"`
	Timestamp          int64    `json:"timestamp"`
	ColumnId           string   `json:"columnId"`
	Description        string   `json:"description"`
	Archived           bool     `json:"archived"`
	ArchivedTimestamp  int64    `json:"archivedTimestamp"`
	Completed          bool     `json:"completed"`
	CompletedTimestamp int64    `json:"completedTimestamp"`
	Subtasks           []string `json:"subtasks"`
	Assigned           []string `json:"assigned"`
	CreatedBy          string   `json:"createdBy"`
	Deadline           struct {
		Deadline  int64 `json:"deadline"`
		StartDate int64 `json:"startDate"`
		WithTime  bool  `json:"withTime"`
	} `json:"deadline"`
	TimeTracking struct {
		Plan int `json:"plan"`
		Work int `json:"work"`
	} `json:"timeTracking"`
	Checklists []struct {
		Title string `json:"title"`
		Items []struct {
			Title       string `json:"title"`
			IsCompleted bool   `json:"isCompleted"`
		} `json:"items"`
	} `json:"checklists"`
	Stickers  map[string]string `json:"stickers"`
	Stopwatch struct {
		Running   bool  `json:"running"`
		Time      int   `json:"time"`
		Timestamp int64 `json:"timestamp"`
		Seconds   int   `json:"seconds"`
		AtMoment  int   `json:"atMoment"`
	} `json:"stopwatch"`
	Timer struct {
		Running   bool  `json:"running"`
		Seconds   int   `json:"seconds"`
		Timestamp int64 `json:"timestamp"`
		Since     int   `json:"since"`
	} `json:"timer"`
}
