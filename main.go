package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juric1962/go_final_project/auth"
	"github.com/juric1962/go_final_project/handlers"
	"github.com/juric1962/go_final_project/tasks"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./scheduler.db")
	if err != nil {
		fmt.Println(" No data base !")
	}
	defer db.Close()
	tasks.DB = tasks.Dbinstance{
		Db: db,
	}
	r := chi.NewRouter()
	r.Get("/", handlers.HandleMain)
	r.Get("/index.html", handlers.HandleMain)
	r.Get("/login.html", handlers.HandleLogin)
	r.Get("/js/scripts.min.js", handlers.HandleScript)
	r.Get("/js/axios.min.js", handlers.HandleAxios)
	r.Get("/css/style.css", handlers.HandleStyle)
	r.Get("/css/theme.css", handlers.HandleTheme)
	r.Get("/favicon.ico", handlers.HandleIco)
	r.Get("/api/nextdate", handlers.HandleAPINextDay)

	r.Post("/api/task", auth.AuthCookie(handlers.HandleApiTaskPost))
	r.Get("/api/task", auth.AuthCookie(handlers.HandleApiTaskGet))
	r.Put("/api/task", auth.AuthCookie(handlers.HandleApiTaskPut))
	r.Post("/api/task/done", auth.AuthCookie(handlers.HandleApiTaskDonePost))
	r.Delete("/api/task", auth.AuthCookie(handlers.HandleApiTaskDelete))
	r.Get("/api/tasks", auth.AuthCookie(handlers.HandleGetTasks))

	r.Post("/api/signin", auth.HandleApiAuthPost)

	r.Post("/api/signin/test", auth.HandleApiAuthPostTestingCookie)
	err = http.ListenAndServe(":7540", r)
	if err != nil {
		panic(err)
	}
}
