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

type YGBoardsService struct {
	Key string `json:"key"`
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
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetBoards StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &boards)
	return
}

func (boardService *YGBoardsService) CreateBoard(createBoardRequest CreateBoardRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/boards"
	payloadByte, _ := json.Marshal(createBoardRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", boardService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateBoard StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (boardService *YGBoardsService) GetBoardById(boardId string) (err error, response BoardResponse) {
	url := "https://ru.yougile.com/api-v2/boards/" + boardId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", boardService.UseKey())
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

func (boardService *YGBoardsService) EditBoard(boardId string, editBoardRequest BoardResponse) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/boards/" + boardId
	payloadByte, _ := json.Marshal(editBoardRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", boardService.UseKey())
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

type CreateBoardRequest struct {
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
