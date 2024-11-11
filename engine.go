package yougilego

import (
	"fmt"
)

func NewYGEngine(key string) YGEngine {
	newEngine := YGEngine{key: key}
	newEngine.Initialization()
	return newEngine
}

type YGEngine struct {
	key     string            `json:"key"`
	Config  YGConfig          `json:"config"`
	User    *YGUsersService   `json:"user"`
	Project *YGProjectService `json:"project"`
	Board   *YGBoardsService  `json:"board"`
	Column  *YGColumnService  `json:"column"`
	Task    *YGTaskService    `json:"task"`
}

func (conn *YGEngine) UseKey() string {
	return fmt.Sprintf("Bearer %s", conn.key)
}

func (conn *YGEngine) Initialization() {
	conn.User = &YGUsersService{YGEngine: *conn}
	conn.Project = &YGProjectService{
		YGEngine: *conn,
	}
	conn.Board = &YGBoardsService{
		YGEngine: *conn,
	}
	conn.Column = &YGColumnService{
		YGEngine: *conn,
	}
	conn.Task = &YGTaskService{
		YGEngine: *conn,
	}
}
