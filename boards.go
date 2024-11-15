package yougilego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type YGBoardsService struct {
	Key               string            `json:"key"`
	Project           *YGProjectService `json:"project"`
	BugTruckerBoardId string            `json:"bugTruckerBoardId"`
}

func (boardService *YGBoardsService) UseKey() string {
	return fmt.Sprintf("Bearer %s", boardService.Key)
}

func (boardService *YGBoardsService) GetBoards() (err error, boards ListResponse[BoardResponse]) {
	url := "https://ru.yougile.com/api-v2/boards"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", boardService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &boards)
	return
}

type BoardResponse struct {
	Deleted   bool   `json:"deleted"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	ProjectId string `json:"projectId"`
	Stickers  struct {
		Timer        bool              `json:"timer"`
		Deadline     bool              `json:"deadline"`
		Stopwatch    bool              `json:"stopwatch"`
		TimeTracking bool              `json:"timeTracking"`
		Assignee     bool              `json:"assignee"`
		Repeat       bool              `json:"repeat"`
		Custom       map[string]string `json:"custom"`
	} `json:"stickers"`
}

type CreateBoard struct {
	Title     string `json:"title"`
	ProjectId string `json:"projectId"`
	Stickers  struct {
		Timer        bool              `json:"timer"`
		Deadline     bool              `json:"deadline"`
		Stopwatch    bool              `json:"stopwatch"`
		TimeTracking bool              `json:"timeTracking"`
		Assignee     bool              `json:"assignee"`
		Repeat       bool              `json:"repeat"`
		Custom       map[string]string `json:"custom"`
	} `json:"stickers"`
}
