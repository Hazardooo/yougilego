package yougilego

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type YGColumnService struct {
	Key string `json:"key"`

	BugTrackerColumnID string           `json:"BugTrackerColumnID"`
	Board              *YGBoardsService `json:"Board"`
}

func (columnService *YGColumnService) UseKey() string {
	return fmt.Sprintf("Bearer %s", columnService.Key)
}

func (columnService *YGColumnService) GetColumn() (err error, columns ListResponse[ColumnResponse]) {
	url := "https://ru.yougile.com/api-v2/columns"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", columnService.UseKey())
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &columns)
	return
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
