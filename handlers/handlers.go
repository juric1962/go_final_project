package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/juric1962/go_final_project/nextdate"
	"github.com/juric1962/go_final_project/tasks"
	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

// GoodTask
func GoodTask(task tasks.Task, w http.ResponseWriter, next string) {
	res, err := tasks.DB.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", next),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	task.Date = next
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	repId := tasks.RepID{ID: fmt.Sprintf("%v", id)}
	res1, err := json.Marshal(repId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	out := string(res1)
	w.Write([]byte(out))
}

// GoodUpdate
func GoodUpdate(task tasks.Task, w http.ResponseWriter, next string) {
	_, errc := strconv.Atoi(task.ID)
	proba := "SELECT * FROM scheduler WHERE id = " + task.ID
	row := tasks.DB.Db.QueryRow(proba)
	if err := row.Scan(&task.Date, &task.Title, &task.Comment, &task.Repeat); err == sql.ErrNoRows || errc != nil {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	_, err := tasks.DB.Db.Exec("UPDATE scheduler SET date = :date , title = :title, comment = :comment , repeat = :repeat WHERE id= :id",
		sql.Named("date", next),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		out := "{}"
		w.Write([]byte(out))
		return
	} else {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
}

// HandleAPINextDay
func HandleAPINextDay(w http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")
	start, _ := time.Parse("20060102", now)
	res, _ := nextdate.NextDate(start, date, repeat)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	out := string(res)
	w.Write([]byte(out))
}

// HandleMain
func HandleMain(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/index.html")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleLogin
func HandleLogin(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/login.html")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleScript
func HandleScript(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/js/scripts.min.js")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleAxios
func HandleAxios(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/js/axios.min.js")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleStyle
func HandleStyle(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/css/style.css")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleTheme
func HandleTheme(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/css/theme.css")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleIco
func HandleIco(w http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("./web/favicon.ico")
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		http.Error(w, " ошибка при чтении файла:", http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/html")
	w.WriteHeader(http.StatusOK)
	out := string(htmlFile)
	w.Write([]byte(out))
}

// HandleGetTasks
func HandleGetTasks(w http.ResponseWriter, req *http.Request) {
	var countId int
	var resso tasks.Ta
	res := []tasks.Task{}
	p := tasks.Task{}
	search := req.FormValue("search")
	if len(search) != 0 {
		var proba string
		start, err1 := time.Parse("02.01.2006", search)
		if err1 == nil {
			proba = "SELECT * FROM scheduler WHERE date LIKE '%" + start.Format("20060102") + "%'"
		} else {
			proba = "SELECT * FROM scheduler WHERE title LIKE '%" + search + "%'"
		}
		rows, err := tasks.DB.Db.Query(proba)
		if err != nil {
			fmt.Printf("ошибка при чтении BD: %s", err.Error())
			http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
			return
		}
		defer rows.Close()
		for rows.Next() {
			p := tasks.Task{}
			err := rows.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
			if err != nil {
				fmt.Printf("ошибка при чтении BD: %s", err.Error())
				http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
				return
			}
			res = append(res, p)
		}
		resso.Tasks = res
		res1, err := json.Marshal(resso)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		out := string(res1)
		w.Write([]byte(out))
		return
	}
	row := tasks.DB.Db.QueryRow("SELECT count(id) from scheduler")
	err := row.Scan(&countId)
	if err != nil {
		fmt.Printf("ошибка при чтении BD: %s", err.Error())
		http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
		return
	}
	if countId == 0 {
		resso.Tasks = res
		res1, err := json.Marshal(resso)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		out := string(res1)
		w.Write([]byte(out))
		return
	}
	var lastid int
	row = tasks.DB.Db.QueryRow("SELECT * from scheduler order by id desc limit 1")
	err = row.Scan(&lastid, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	var IdLow int
	if countId == 0 || err != nil {
		fmt.Printf("ошибка при чтении BD: %s", err.Error())
		http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
		return
	}
	if countId > 10 {
		IdLow = lastid - 10
	} else {
		IdLow = lastid - countId - 1
	}
	rows, err := tasks.DB.Db.Query("SELECT * FROM scheduler WHERE id between ? and ? order by date ", IdLow, lastid)
	if err != nil {
		fmt.Printf("ошибка при чтении BD: %s", err.Error())
		http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := tasks.Task{}
		err := rows.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
		if err != nil {
			fmt.Printf("ошибка при чтении BD: %s", err.Error())
			http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
			return
		}
		res = append(res, p)
	}
	resso.Tasks = res
	res1, err := json.Marshal(resso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	out := string(res1)
	w.Write([]byte(out))
}

// HandleApiTaskPost
func HandleApiTaskPost(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		repErr := tasks.RepErr{Error: "ошибка десериализации JSON"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	now := time.Now().Format(`20060102`)
	if len(task.Date) == 0 {
		task.Date = now
	}
	_, err = nextdate.NextDate(time.Now(), task.Date, "y")
	start, err1 := time.Parse("20060102", task.Date)
	if err1 != nil || len(task.Title) == 0 || err != nil {
		repErr := tasks.RepErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	if start.After(time.Now()) {
		next := task.Date
		GoodTask(task, w, next)
		return
	}
	if now == task.Date {
		next := task.Date
		GoodTask(task, w, next)
		return
	}
	if len(task.Repeat) == 0 {
		next := now
		GoodTask(task, w, next)
		return
	}
	next, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		repErr := tasks.RepErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	GoodTask(task, w, next)
}

// HandleApiTaskGet
func HandleApiTaskGet(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if len(id) == 0 {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	var p tasks.Task
	proba := "SELECT * FROM scheduler WHERE id = " + id
	row := tasks.DB.Db.QueryRow(proba)
	err := row.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	if err == nil {
		resJson, _ := json.Marshal(p)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		out := string(resJson)
		w.Write([]byte(out))
		return
	} else {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
}

// HandleApiTaskPut
func HandleApiTaskPut(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		repErr := tasks.RepErr{Error: "ошибка десериализации JSON"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	now := time.Now().Format(`20060102`)
	if len(task.Date) == 0 {
		task.Date = now
	}
	_, err = nextdate.NextDate(time.Now(), task.Date, "y")
	start, err1 := time.Parse("20060102", task.Date)
	if err1 != nil || len(task.Title) == 0 || err != nil {
		repErr := tasks.RepErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	if start.After(time.Now()) {
		next := task.Date
		GoodUpdate(task, w, next)
		return
	}
	if now == task.Date {
		next := task.Date
		GoodUpdate(task, w, next)
		return
	}
	if len(task.Repeat) == 0 {
		next := now
		GoodUpdate(task, w, next)
		return
	}
	next, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil || len(task.Title) == 0 {
		repErr := tasks.RepErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	GoodUpdate(task, w, next)
}

// HandleApiTaskDonePost
func HandleApiTaskDonePost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, errc := strconv.Atoi(id)
	if len(id) == 0 || errc != nil {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	var p tasks.Task
	proba := "SELECT * FROM scheduler WHERE id = " + id
	row := tasks.DB.Db.QueryRow(proba)
	err := row.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	if err != nil {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	if len(p.Repeat) == 0 {
		_, err = tasks.DB.Db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
		if err != nil {
			resJson, _ := json.Marshal(tasks.RepErr{Error: "can't delete task"})
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			out := string(resJson)
			w.Write([]byte(out))
			return
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusAccepted)
			out := "{}"
			w.Write([]byte(out))
			return
		}
	}
	next, err := nextdate.NextDate(time.Now(), p.Date, p.Repeat)
	if err != nil || len(p.Title) == 0 {
		repErr := tasks.RepErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(repErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	_, err = tasks.DB.Db.Exec("UPDATE scheduler SET date = :date  WHERE id= :id",
		sql.Named("date", next),
		sql.Named("id", p.ID))

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		out := "{}"
		w.Write([]byte(out))
		return
	} else {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "can't update task"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
}

// HandleApiTaskDelete
func HandleApiTaskDelete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, errc := strconv.Atoi(id)
	if len(id) == 0 || errc != nil {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	var p tasks.Task
	proba := "SELECT * FROM scheduler WHERE id = " + id
	row := tasks.DB.Db.QueryRow(proba)
	err := row.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
	if err != nil {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "Задача не найдена"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	_, err = tasks.DB.Db.Exec("DELETE FROM scheduler WHERE id = :id", sql.Named("id", id))
	if err != nil {
		resJson, _ := json.Marshal(tasks.RepErr{Error: "can't delete task"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		out := string(resJson)
		w.Write([]byte(out))
		return
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusAccepted)
		out := "{}"
		w.Write([]byte(out))
		return
	}
}
