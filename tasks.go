package yougilego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type YGTaskService struct {
	YGEngine `json:"YGEngine"`
	Column   *YGColumnService `json:"Column"`
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

func (taskService *YGTaskService) GetTasks() (err error, tasks ListResponse[TaskResponse]) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/tasks?columnId=%s", taskService.Column.BugTrackerColumnID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", taskService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &tasks)
	return
}
