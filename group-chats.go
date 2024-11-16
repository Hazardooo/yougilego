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

type YGGroupChatService struct {
	Key string `json:"key"`
}

func (groupChatService *YGGroupChatService) UseKey() string {
	return fmt.Sprintf("Bearer %s", groupChatService.Key)
}

func (groupChatService *YGGroupChatService) GetGroupChatList() (err error, response ListResponse[GroupChat]) {
	url := "https://ru.yougile.com/api-v2/group-chats"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", groupChatService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetGroupChatList StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (groupChatService *YGGroupChatService) CreateGroupChat(createGroupChatRequest CreateGroupChatRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/group-chats"
	payloadByte, _ := json.Marshal(createGroupChatRequest)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", groupChatService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("CreateGroupChat StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (groupChatService *YGGroupChatService) GetGroupChatById(groupChatId string) (err error, response GroupChat) {
	url := "https://ru.yougile.com/api-v2/group-chats/" + groupChatId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", groupChatService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetGroupChatById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (groupChatService *YGGroupChatService) EditGroupChat(groupChatId string, editGroupChatRequest EditGroupChatRequest) (err error, response IDResponse) {
	url := "https://ru.yougile.com/api-v2/group-chats/" + groupChatId
	payloadByte, _ := json.Marshal(editGroupChatRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", groupChatService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditGroupChat StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type GroupChat struct {
	Deleted bool              `json:"deleted"`
	Id      string            `json:"id"`
	Title   string            `json:"title"`
	Users   map[string]string `json:"users"`
}

type CreateGroupChatRequest struct {
	Title string            `json:"title"`
	Users map[string]string `json:"users"`
}

type EditGroupChatRequest struct {
	Deleted bool              `json:"deleted"`
	Title   string            `json:"title"`
	Users   map[string]string `json:"users"`
}
