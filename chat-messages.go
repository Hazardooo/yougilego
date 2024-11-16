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

type YGChatMessageService struct {
	Key string `json:"key"`
}

func (chatMessageService *YGChatMessageService) UseKey() string {
	return fmt.Sprintf("Bearer %s", chatMessageService.Key)
}

func (chatMessageService *YGChatMessageService) GetHistoryChat(chatId string) (err error, response ListResponse[MessageData]) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/chats/%s/messages", chatId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", chatMessageService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetHistoryChat StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (chatMessageService *YGChatMessageService) SendMessage(chatId string, message Message) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/chats/%s/messages", chatId)
	payloadByte, _ := json.Marshal(message)
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", chatMessageService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		err = errors.New(fmt.Sprintf("SendMessage StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (chatMessageService *YGChatMessageService) GetMessageById(chatId string, messageId string) (err error, response MessageData) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/chats/%s/messages/%s", chatId, messageId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", chatMessageService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("GetMessageById StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

func (chatMessageService *YGChatMessageService) EditMessage(chatId string, messageId string, deleteMessageRequest DeleteMessageRequest) (err error, response IDResponse) {
	url := fmt.Sprintf("https://ru.yougile.com/api-v2/chats/%s/messages/%s", chatId, messageId)
	payloadByte, _ := json.Marshal(deleteMessageRequest)
	req, _ := http.NewRequest("PUT", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", chatMessageService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("EditMessage StatusCode: %s", strconv.Itoa(res.StatusCode)))
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return
}

type DeleteMessageRequest struct {
	Deleted bool     `json:"deleted"`
	Label   string   `json:"label"`
	React   []string `json:"react"`
}

type MessageData struct {
	Deleted       bool   `json:"deleted"`
	Id            int64  `json:"id"`
	FromUserId    string `json:"fromUserId"`
	Text          string `json:"text"`
	TextHtml      string `json:"textHtml"`
	Label         string `json:"label"`
	EditTimestamp int64  `json:"editTimestamp"`
	Reactions     struct {
		A8Bff2F6B0425A9F893A003B8Eb039 []struct {
			Smiley    string `json:"smiley"`
			Timestamp int64  `json:"timestamp"`
		} `json:"18a8bff2-f6b0-425a-9f89-3a003b8eb039"`
	} `json:"reactions"`
}

type Message struct {
	Text     string `json:"text"`
	TextHtml string `json:"textHtml"`
	Label    string `json:"label"`
}
