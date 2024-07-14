package store

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/juric1962/go_final_project/tasks"
	_ "modernc.org/sqlite"
)

type TodoList struct {
	db *sql.DB
}

func NewTodoList(db *sql.DB) TodoList {
	return TodoList{db: db}
}

var Todo TodoList

func (s TodoList) Add(p tasks.Task, next string) (int64, error) {
	res, err := s.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", next),
		sql.Named("title", p.Title),
		sql.Named("comment", p.Comment),
		sql.Named("repeat", p.Repeat))
	p.Date = next
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	// идентификатор последней добавленной записи
	return id, nil
}

func (s TodoList) GetTask(id string) (tasks.Task, error) {
	var p tasks.Task
	_, errc := strconv.Atoi(id)
	if errc != nil {
		return p, errc
	}
	proba := "SELECT * FROM scheduler WHERE id = " + id
	row := s.db.QueryRow(proba)
	err := row.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, err
		} else {
			log.Fatal(err)
		}
	}
	return p, nil
}

func (s TodoList) UpdateDB(p tasks.Task) error {
	_, err := s.db.Exec("UPDATE scheduler SET date = :date , title = :title, comment = :comment , repeat = :repeat WHERE id= :id",
		sql.Named("date", p.Date),
		sql.Named("title", p.Title),
		sql.Named("comment", p.Comment),
		sql.Named("repeat", p.Repeat),
		sql.Named("id", p.ID))
	return err
}

func (s TodoList) Delete(id string) error {
	_, err := s.db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	return err
}

func (s TodoList) Find(p tasks.Task, search string) ([]tasks.Task, error) {
	res := []tasks.Task{}
	var proba string
	start, err1 := time.Parse("02.01.2006", search)
	if err1 == nil {
		proba = "SELECT * FROM scheduler WHERE date LIKE '%" + start.Format(tasks.TimeFormat) + "%'"
	} else {
		proba = "SELECT * FROM scheduler WHERE title LIKE '%" + search + "%'"
	}
	rows, err := s.db.Query(proba)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, err
		} else {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		p := tasks.Task{}
		err := rows.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
		if err != nil {
			if err = rows.Close(); err != nil {
				log.Println(err)
			}
			return res, err
		}
		res = append(res, p)
	}
	return res, nil
}

func (s TodoList) SelectTasks(p tasks.Task) ([]tasks.Task, error) {
	res := []tasks.Task{}
	rows, err := s.db.Query("SELECT * FROM scheduler order by date desc limit 10 ")
	if err != nil {
		if err == sql.ErrNoRows {
			return res, err
		} else {
			log.Fatal(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
		if err != nil {
			if err = rows.Close(); err != nil {
				log.Println(err)
			}
			return res, err
		}
		res = append(res, p)
	}
	return res, nil
}
