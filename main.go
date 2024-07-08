package main

import (
	"database/sql"
	"errors"
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
	errPort := errors.New("не опрелен порт в переменной окружения TODO_PORT")
	errDB := errors.New("не определен путь к scheduler.db в переменной окружения TODO_DB")
	path := os.Getenv("TODO_DB")
	if len(path) == 0 {
		panic(errDB)
	}
	port := os.Getenv("TODO_PORT")
	if len(port) == 0 {
		panic(errPort)
	}

	db, err := sql.Open("sqlite3", "./scheduler.db")
	if err != nil {
		fmt.Println(" No data base !")
	}
	defer db.Close()
	///tasks.DB = tasks.Dbinstance{
	//	Db: db,
	//}
	dbhandler.Todo = dbhandler.NewTodoList(db)
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

	//err = http.ListenAndServe(":7540", r)
	fmt.Printf(" server start port =%s  \n path database =%s\n", port, path)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
