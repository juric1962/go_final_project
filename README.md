# Файлы для итогового задания
Проект «планировщик задач»
Выполнен в виде HTTP сервера. 
В даном проекте реализован бэкенд сервера.
Для хранения задач используется SQLite.
Данные хранятся в таблице scheduler.db.
Фронтенд сервера прописан в пакете web и является неизменной частью проекта. 

Структура проекта

# Пакет Handlers 
    r.Get("/", handlers.HandleMain)
    r.Get("/index.html", handlers.HandleMain)
    r.Get("/login.html", handlers.HandleLogin)
    r.Get("/js/scripts.min.js", handlers.HandleScript)
    r.Get("/js/axios.min.js", handlers.HandleAxios)
    r.Get("/css/style.css", handlers.HandleStyle)
    r.Get("/css/theme.css", handlers.HandleTheme)
    r.Get("/favicon.ico", handlers.HandleIco)
    r.Get("/api/nextdate", handlers.HandleAPINextDay)
    r.Post("/api/task", handlers.HandleApiTaskPost)
    r.Get("/api/task", handlers.HandleApiTaskGet)
    r.Put("/api/task", handlers.HandleApiTaskPut)
    r.Post("/api/task/done", handlers.HandleApiTaskDonePost)
    r.Delete("/api/task", handlers.HandleApiTaskDelete)
    r.Get("/api/tasks", handlers.HandleGetTasks)
    r.Post("/api/signin", auth.HandleApiAuthPost)
    r.Post("/api/signin/test", auth.HandleApiAuthPostTesting)
    обработчики http запросов

# Пакет nextdate 
  nextdate.go
  Рассчитывает следующую дату для задания.   
  Возвращает дату задания в формате «20060102»

# Пакет tasks
  Tasks.go
  Объявление структур 

# Пакет auth 
  Auth.go
  Обработчик ввода пароля. Принимает PUT запрос с   
  паролем и возвращает jwt токен. Пароль на сервере  
  хранится в переменой окружения TODO_PASSWORD.

# Пакет scheduler
  scheduler.sql
  Макет таблицы DB SQlite3
  По этому макету создается  таблица scheduler.db

# Пакет dbhandler
  dbhandler.go
  Методы для работы с BD
  таблица scheduler.db

# Пакет tests
  Программы для тестирования функциональности            
  сервера.
  Неизменная часть проекта.

# Пакет web
  Прописан фронтенд сервера.
  Неизменная часть проекта

# параметры для тестирования

var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = false
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzaGFzaCI6ImE2NjVhNDU5MjA0MjJmOWQ0MTdlNDg2N2VmZGM0ZmI4YTA0YTFmM2ZmZjFmYTA3ZTk5OGU4NmY3ZjdhMjdhZTMiLCJyb2xlcyI6InF3ZXJ0eSJ9.eFhwVmwqPwV4z_g7bcZAnSCSREkm8cblsV7aqMK7ytg`

# переменные окружения
TODO_PASSWORD= 123
TODO_PORT= 7540
TODO_DB = "./scheduler.db"
# сборка и запуск контейнера
docker build --tag go_final_project:v3.0 .
docker run -it -p ${TODO_PORT}:${TODO_PORT} --mount type=bind,source=.,target=/app go_final_project:v3.0
