package yougilego

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type YGColumnService struct {
	YGEngine           `json:"YGEngine"`
	BugTrackerColumnID string           `json:"BugTrackerColumnID"`
	Board              *YGBoardsService `json:"Board"`
}

type CreateColumn struct {
	Title   string `json:"title"`
	Color   int    `json:"color"`
	BoardId string `json:"boardId"`
}

type ColumnResponse struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Title   string `json:"title"`
	Color   int    `json:"color"`
	BoardId string `json:"boardId"`
}

func (columnService *YGColumnService) GetColumn() (err error, columns ListResponse[ColumnResponse]) {
	url := "https://ru.yougile.com/api-v2/columns"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.YGEngine.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &columns)
	return
}

func (columnService *YGColumnService) CheckBugTrackerColumn(columns ListResponse[ColumnResponse]) bool {
	for _, colum := range columns.Content {
		if colum.BoardId == columnService.Board.BugTruckerBoardId && colum.Title == columnService.Config.BugTruckerColumnName {
			columnService.BugTrackerColumnID = colum.Id
			return true
		}
	}
	return false
}

func (columnService *YGColumnService) SetBugTrackerColumn() (err error) {
	url := "https://ru.yougile.com/api-v2/columns"
	payload := CreateColumn{
		Title:   columnService.Config.BugTruckerColumnName,
		Color:   10,
		BoardId: columnService.Board.BugTruckerBoardId,
	}
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return
	}
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(payloadByte)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	newColumnId := SuccessResponse{}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &newColumnId)
	return
}
