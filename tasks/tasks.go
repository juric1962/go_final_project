package tasks

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task ...
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
	/*
	   ID      string `json:"id,omitempty"`
	   Date    string `json:"date,omitempty"`
	   Title   string `json:"title,omitempty"`
	   Comment string `json:"comment,omitempty"`
	   Repeat  string `json:"repeat,omitempty"`
	*/
}

// RepID ...
type RepID struct {
	ID string `json:"id"`
}

// RepJSON ...
type RepJSON struct {
	Token string `json:"token"`
}

// RepErr ...
type RepErr struct {
	Error string `json:"error"`
}

// Dbinstance ...
type Dbinstance struct {
	Db *sql.DB
}

// Ta ...
type Ta struct {
	Tasks []Task `json:"tasks"`
}

type PSW struct {
	Password string `json:"password"`
}

var DB Dbinstance
