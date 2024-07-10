package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/juric1962/go_final_project/dbhandler"
	"github.com/juric1962/go_final_project/nextdate"
	"github.com/juric1962/go_final_project/tasks"
	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func SendHttp(load []byte, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	out := string(load)
	w.Write([]byte(out))
}

// AddTask
func AddTask(task tasks.Task, w http.ResponseWriter, next string) {
	id, err := dbhandler.Todo.Add(task, next)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resId := tasks.ResID{ID: fmt.Sprintf("%v", id)}
	result, err := json.Marshal(resId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendHttp(result, w, http.StatusCreated)
}

// UpdateTask
func UpdateTask(task tasks.Task, w http.ResponseWriter, next string) {
	_, err := dbhandler.Todo.GetTask(task.ID)
	if err != nil {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	task.Date = next
	err = dbhandler.Todo.UpdateDB(task)
	if err == nil {
		resJson := []byte("{}")
		SendHttp(resJson, w, http.StatusOK)
		return
	} else {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
}

// HandleAPINextDay
func HandleAPINextDay(w http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")
	start, _ := time.Parse(tasks.TimeFormat, now)
	res, _ := nextdate.NextDate(start, date, repeat)
	SendHttp([]byte(res), w, http.StatusOK)
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
		replyErr := tasks.ResErr{Error: "ошибка десериализации JSON"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	now := time.Now().Format(tasks.TimeFormat)
	if len(task.Date) == 0 {
		task.Date = now
	}
	_, err = nextdate.NextDate(time.Now(), task.Date, "y")
	start, err1 := time.Parse(tasks.TimeFormat, task.Date)
	if err1 != nil || len(task.Title) == 0 || err != nil {
		replyErr := tasks.ResErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	if start.After(time.Now()) {
		next := task.Date
		AddTask(task, w, next)
		return
	}
	if now == task.Date {
		next := task.Date
		AddTask(task, w, next)
		return
	}
	if len(task.Repeat) == 0 {
		next := now
		AddTask(task, w, next)
		return
	}
	next, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		replyErr := tasks.ResErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	AddTask(task, w, next)
}

// HandleApiTaskGet
func HandleApiTaskGet(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if len(id) == 0 {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	var p tasks.Task
	p, err := dbhandler.Todo.GetTask(id)
	if err == nil {
		resJson, _ := json.Marshal(p)
		SendHttp(resJson, w, http.StatusOK)
		return
	} else {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
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
		replyErr := tasks.ResErr{Error: "ошибка десериализации JSON"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	now := time.Now().Format(tasks.TimeFormat)
	if len(task.Date) == 0 {
		task.Date = now
	}
	_, err = nextdate.NextDate(time.Now(), task.Date, "y")
	start, err1 := time.Parse(tasks.TimeFormat, task.Date)
	if err1 != nil || len(task.Title) == 0 || err != nil {
		replyErr := tasks.ResErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	if start.After(time.Now()) {
		next := task.Date
		UpdateTask(task, w, next)
		return
	}
	if now == task.Date {
		next := task.Date
		UpdateTask(task, w, next)
		return
	}
	if len(task.Repeat) == 0 {
		next := now
		UpdateTask(task, w, next)
		return
	}
	next, err := nextdate.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil || len(task.Title) == 0 {
		replyErr := tasks.ResErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	UpdateTask(task, w, next)
}

// HandleApiTaskDonePost
func HandleApiTaskDonePost(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, errc := strconv.Atoi(id)
	if len(id) == 0 || errc != nil {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	var p tasks.Task
	p, err := dbhandler.Todo.GetTask(id)
	if err != nil {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	if len(p.Repeat) == 0 {
		err = dbhandler.Todo.Delete(id)
		if err != nil {
			resJson, _ := json.Marshal(tasks.ResErr{Error: "can't delete task"})
			SendHttp(resJson, w, http.StatusInternalServerError)
			return
		} else {
			resJson := []byte("{}")
			SendHttp(resJson, w, http.StatusAccepted)
			return
		}
	}
	next, err := nextdate.NextDate(time.Now(), p.Date, p.Repeat)
	if err != nil || len(p.Title) == 0 {
		replyErr := tasks.ResErr{Error: "правило указано в неправильном формате"}
		resJson, _ := json.Marshal(replyErr)
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	p.Date = next
	err = dbhandler.Todo.UpdateDB(p)
	if err == nil {
		resJson := []byte("{}")
		SendHttp(resJson, w, http.StatusOK)
		return
	} else {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "can't update task"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
}

// HandleApiTaskDelete
func HandleApiTaskDelete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, errc := strconv.Atoi(id)
	if len(id) == 0 || errc != nil {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	_, err := dbhandler.Todo.GetTask(id)
	if err != nil {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Задача не найдена"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	}
	err = dbhandler.Todo.Delete(id)
	if err != nil {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "can't delete task"})
		SendHttp(resJson, w, http.StatusInternalServerError)
		return
	} else {
		resJson := []byte("{}")
		SendHttp(resJson, w, http.StatusAccepted)
		return
	}
}

// HandleGetTasks
func HandleGetTasks(w http.ResponseWriter, req *http.Request) {
	var resso tasks.Tasks
	p := tasks.Task{}
	search := req.FormValue("search")
	if len(search) != 0 {
		res, err := dbhandler.Todo.Find(p, search)
		if err != nil {
			fmt.Printf("ошибка при чтении BD: %s", err.Error())
			http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
			return
		}
		resso.Tasks = res
		resJson, err := json.Marshal(resso)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendHttp(resJson, w, http.StatusOK)
		return
	}
	// read list of tasks
	res, err := dbhandler.Todo.SelectTasks(p)
	if err != nil {
		fmt.Printf("ошибка при чтении BD: %s", err.Error())
		http.Error(w, " ошибка при чтении BD", http.StatusNoContent)
		return
	}
	resso.Tasks = res
	resJson, err := json.Marshal(resso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendHttp(resJson, w, http.StatusOK)
}

// HandleApiAuthPost
// возвращат подписаный токен в формате json
func HandleApiAuthPost(w http.ResponseWriter, r *http.Request) {
	// получаем пароль
	var psw tasks.Password

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &psw); err != nil {
		replyErr := tasks.ResErr{Error: "ошибка десериализации JSON"}
		resJson, _ := json.Marshal(replyErr)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	if len(psw.Password) == 0 {
		resJson, _ := json.Marshal(tasks.ResErr{Error: "Authentification required"})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)
		out := string(resJson)
		w.Write([]byte(out))
		return
	}
	crc := sha256.Sum256([]byte(psw.Password))
	hashString := hex.EncodeToString(crc[:])
	// создаём payload
	claims := jwt.MapClaims{
		"passhash": hashString,
		"roles":    "qwerty",
	}
	// создаём jwt и указываем payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// получаем подписанный токен
	signedToken, err := jwtToken.SignedString([]byte(psw.Password))
	if err != nil {
		fmt.Printf("failed to sign jwt: %s\n", err)
	}
	// возвращаем токен в формате json. {"token" : signedToken}
	resJson, _ := json.Marshal(tasks.ResJSON{Token: signedToken})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	out := string(resJson)
	w.Write([]byte(out))
}
