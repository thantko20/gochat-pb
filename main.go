package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

type MessageRequestBody struct {
	Receiver string `json:"receiver" default:""`
	Chat     string `json:"chat" default:""`
	Sender   string `json:"sender" default:""`
}

func main() {
	app := pocketbase.New()

	app.OnRecordBeforeCreateRequest("messages").Add(onBeforeCreateMessageRecord(app))
	app.OnRecordAfterCreateRequest("messages").Add(onAfterCreateMessageRecord(app))

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func onBeforeCreateMessageRecord(app *pocketbase.PocketBase) hook.Handler[*core.RecordCreateEvent] {
	return func(e *core.RecordCreateEvent) error {

		bodyReader := e.HttpContext.Request().Body
		defer bodyReader.Close()

		var body MessageRequestBody
		decoder := json.NewDecoder(bodyReader)
		if err := decoder.Decode(&body); err != nil {
			log.Println(err)
			return err
		}

		if body.Chat == "" && body.Receiver == "" {
			return errors.New("something went wrong")
		}

		if body.Chat == "" {
			chatRecords, err := app.Dao().FindRecordsByFilter("chats", "users ~ {:receiver} && users ~ {:sender} && type = {:type}",
				"-created",
				1,
				0,
				dbx.Params{
					"receiver": body.Receiver,
					"sender":   body.Sender,
					"type":     "normal",
				})
			if err != nil {
				return err
			}

			if len(chatRecords) == 0 {
				users := []string{body.Receiver, body.Sender}
				chatsCollection, err := app.Dao().FindCollectionByNameOrId("chats")
				if err != nil {
					return err
				}
				record := models.NewRecord(chatsCollection)

				log.Println("users: ", users)

				record.Set("type", "normal")
				record.Set("users", users)

				if err := app.Dao().SaveRecord(record); err != nil {
					return err
				}

				e.Record.Set("chat", record.Id)

				return nil
			}
			e.Record.Set("chat", chatRecords[0].Id)
		}

		return nil
	}
}

func onAfterCreateMessageRecord(app *pocketbase.PocketBase) hook.Handler[*core.RecordCreateEvent] {
	return func(e *core.RecordCreateEvent) error {
		record, err := app.Dao().FindRecordById("chats", e.Record.Get("chat").(string))
		if err != nil {
			return err
		}

		record.Set("lastMessage", e.Record.Id)

		if err := app.Dao().SaveRecord(record); err != nil {
			return err
		}

		return nil
	}
}
