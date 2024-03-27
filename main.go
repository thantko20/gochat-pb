package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
	// 	return nil
	// })

	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		log.Println(e.Record.CleanCopy())
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
