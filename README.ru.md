# YougileGo

YougileGo — это Go-пакет, предоставляющий полный набор методов и структур для подключения к платформе YouGile и получения данных из нее.

## Возможности

- Аутентификация в YouGile.
- Получение списка компаний, пользователей, проектов, департаментов, досок, колонок, задач, стикеров и многого другого.
- Управление групповыми чатами и сообщениями.
- Работа со спринт-стикерами и текстовыми стикерами.
- Подписка на события вебхуков.

## Установка

Установите YougileGo с помощью команды `go get`:

```sh
go get github.com/Hazardooo/yougilego
```

## Использование

Чтобы лучше понять работу библиотеки, ознакомьтесь с [документацией API Yougile](https://ru.yougile.com/api-v2#/).

### Получение ключа компании

[Пример использования всех объектов](https://github.com/Hazardooo/yougilego/blob/main/example/all-objects-in-one/main.go)

```go
package main

import (
	"fmt"
	"github.com/Hazardooo/yougilego"
	"log"
	"os"
)

func main() {
	auth := yougilego.YGAuthService{
		Login:    os.Getenv("LOGIN"), 
		Password: os.Getenv("PASSWORD"),
	}

	err, company := auth.GetListCompany("Hazardooo")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(company)
	myCompany := company.Content[0]

	err, keysList := auth.GetKeysList(myCompany.Id)
	if err != nil {
		log.Println(err)
		return
	}

	var key string
	if len(keysList) == 0 {
		err, key = auth.CreateKey(myCompany.Id)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		key = keysList[0].Key
	}
	fmt.Println(key)
}
```

В дальнейшем ключ компании понадобится для работы с остальными данными YouGile.

| Объект               | Статус  |
|-----------------------|---------|
| users                | Support |
| projects             | Support |
| project-roles        | Support |
| departments          | Support |
| boards               | Support |
| columns              | Support |
| tasks                | Support |
| text-stickers        | Support |
| status-text-stickers | Support |
| sprint-sticker       | Support |
| status-sprint-sticker| Support |
| group-chats          | Support |
| chat-messages        | Support |
| event-subs           | Support |

### Получение задач из колонки

```go
	boardsService := yougilego.YGBoardsService{Key: key}
	err, boardList := boardsService.GetBoards()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(boardList)
	columnService := yougilego.YGColumnService{Key: key}
	err, columnList := columnService.GetColumnList()
	if err != nil {
		log.Println(err)
		return
	}
	myColumn := columnList.Content[0]
	taskService := yougilego.YGTaskService{Key: key}
	err, taskList := taskService.GetTasks(myColumn.Id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(taskList)
```

### Пример создания подписки на событие

[Пример создания подписки на событие](https://github.com/Hazardooo/yougilego/blob/main/example/sub-event/main.go)

Для создания подписки на событие необходимо получить ваш адрес вебхука через [webhook.site](https://webhook.site/).

```go
func main() {
	auth := yougilego.YGAuthService{
		Login:    os.Getenv("LOGIN"),
		Password: os.Getenv("PASSWORD"),
	}
	err, company := auth.GetListCompany("Hazardooo")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(company)
	myCompany := company.Content[0]
	err, keysList := auth.GetKeysList(myCompany.Id)
	if err != nil {
		log.Println(err)
		return
	}
	var key string
	if len(keysList) == 0 {
		err, key = auth.CreateKey(myCompany.Id)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		key = keysList[0].Key
	}
	fmt.Println(key)
	eventSubscribeService := yougilego.YGEventSubscribeService{Key: key}
	err, eventList := eventSubscribeService.GetSubscribeList(false)
	if err != nil {
		log.Println(err)
		return
	}
	var event = yougilego.SubscribeResponse{}
	if len(eventList) != 0 {
		event = eventList[0]
	} else {
		err, _ = eventSubscribeService.CreateSubscribe(yougilego.CreateSubs{
			Url:   "https://webhook.site/d19ddaf8-36d5-4dcd-acf1-51b91b6dd375",
			Event: "task-created",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	fmt.Println(event)
}
```

После создания подписки перейдите на Webhook.site — он мгновенно отобразит все события, произошедшие в вашем YouGile.

![изображение](https://github.com/user-attachments/assets/2b8fd496-9a38-4a54-9a4e-4e86100293ca)
![изображение](https://github.com/user-attachments/assets/78cdff49-326b-408d-8172-307142d639eb)

### Варианты событий подписки

| Объект          | События                                                    |
|------------------|-----------------------------------------------------------|
| project         | created, deleted, restored, moved, renamed, updated        |
| board           | created, deleted, restored, moved, renamed, updated        |
| column          | created, deleted, restored, moved, renamed, updated        |
| task            | created, deleted, restored, moved, renamed, updated        |
| sticker         | created, deleted, restored, moved, renamed, updated        |
| department      | created, deleted, restored, moved, renamed, updated        |
| group_chat      | created, deleted, restored, moved, renamed, updated        |
| chat_message    | created, deleted, restored, moved, renamed, updated        |
| user            | added, removed                                             |

Для подписки на все события определенного объекта используйте формат `<object>-*`. Например:

- `task-*` — все события для задач.
- `.*` — все события для всех объектов.

## Вклад в проект
Ваши предложения приветствуются! Вы можете отправить Pull Request или создать Issue для сообщения об ошибках или предложения новых функций.

## Лицензия
Этот проект лицензирован под [MIT License](https://github.com/Hazardooo/yougilego/blob/main/LICENSE).