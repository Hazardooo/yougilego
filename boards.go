package yougilego

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type YGBoardsService struct {
	YGEngine          `json:"YGEngine"`
	Project           *YGProjectService `json:"project"`
	BugTruckerBoardId string            `json:"bugTruckerBoardId"`
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

func (boardService *YGBoardsService) CheckBugTrackerBoard(boards ListResponse[BoardResponse]) bool {
	for _, board := range boards.Content {
		if board.ProjectId == boardService.Project.BugTruckerProjectId && board.Title == boardService.Config.BugTruckerBoardName {
			boardService.BugTruckerBoardId = board.Id
			return true
		}
	}
	return false
}

func (boardService *YGBoardsService) SetBugTrackerBoard() (err error) {
	url := "https://ru.yougile.com/api-v2/boards"
	payload := CreateBoard{
		Title:     boardService.Config.BugTruckerBoardName,
		ProjectId: boardService.Project.BugTruckerProjectId,
		Stickers: struct {
			Timer        bool              `json:"timer"`
			Deadline     bool              `json:"deadline"`
			Stopwatch    bool              `json:"stopwatch"`
			TimeTracking bool              `json:"timeTracking"`
			Assignee     bool              `json:"assignee"`
			Repeat       bool              `json:"repeat"`
			Custom       map[string]string `json:"custom"`
		}{
			Timer:        false,
			Deadline:     false,
			Stopwatch:    false,
			TimeTracking: false,
			Assignee:     false,
			Repeat:       false,
			Custom:       make(map[string]string),
		},
	}
	payloadByte, err := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", boardService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	newBoardId := SuccessResponse{}
	json.Unmarshal(body, &newBoardId)
	boardService.BugTruckerBoardId = newBoardId.Id
	return err
}
