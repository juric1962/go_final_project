package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/juric1962/go_final_project/auth"
	"github.com/juric1962/go_final_project/dbhandler"
	"github.com/juric1962/go_final_project/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	path := os.Getenv("TODO_DB")
	if len(path) == 0 {
		fmt.Println("не определен путь к scheduler.db в переменной окружения TODO_DB")
		return
	}
	port := os.Getenv("TODO_PORT")
	if len(port) == 0 {
		fmt.Println("не опрелен порт в переменной окружения TODO_PORT")
		return
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Println(" No data base !")
	}
	defer db.Close()
	dbhandler.Todo = dbhandler.NewTodoList(db)
	auth.Pass = []byte(os.Getenv("TODO_PASSWORD"))
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
	r.Post("/api/signin", handlers.HandleApiAuthPost)

	r.Post("/api/task", auth.AuthCookie(handlers.HandleApiTaskPost))
	r.Get("/api/task", auth.AuthCookie(handlers.HandleApiTaskGet))
	r.Put("/api/task", auth.AuthCookie(handlers.HandleApiTaskPut))
	r.Post("/api/task/done", auth.AuthCookie(handlers.HandleApiTaskDonePost))
	r.Delete("/api/task", auth.AuthCookie(handlers.HandleApiTaskDelete))
	r.Get("/api/tasks", auth.AuthCookie(handlers.HandleGetTasks))
	r.Post("/api/signin/test", auth.HandleApiAuthPostTestingCookie)
	fmt.Printf(" server start port =%s  \n path database =%s\n", port, path)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
