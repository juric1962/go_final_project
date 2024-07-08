package tasks

import (
	_ "github.com/mattn/go-sqlite3"
)

// Task ...
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// ResID ...
type ResID struct {
	ID string `json:"id"`
}

// ResJSON ...
type ResJSON struct {
	Token string `json:"token"`
}

// ResErr ...
type ResErr struct {
	Error string `json:"error"`
}

// Tasks ...
type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// Password
type Password struct {
	Password string `json:"password"`
}
