package yougilego

import "fmt"

type YGEventSubscribeService struct {
	Key string `json:"key"`
}

func (eventSubscribeService *YGEventSubscribeService) UseKey() string {
	return fmt.Sprintf("Bearer %s", eventSubscribeService.Key)
}
