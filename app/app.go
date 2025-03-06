package app

import (
	"TODO_LIST_Practice/app/handler"
	"TODO_LIST_Practice/app/model"
	"TODO_LIST_Practice/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

type App struct{
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) InitializeRoutes(config *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset,
	)
    db, err := gorm.Open(config.DB.Dialect, dsn)
	if err != nil{
		log.Fatal("Error connecting to database",err)
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters(){
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontView/")))
	// projects
	a.Get("/projects", a.handleRequest(handler.GetAllProjects))
	a.Post("/projects", a.handleRequest(handler.CreateProject))
	a.Get("/projects/{title}", a.handleRequest(handler.GetProject))
	a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
	a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
	a.Put("/projects/{title}/archive", a.handleRequest(handler.ArchiveProject))
	a.Delete("/projects/{title}/archive", a.handleRequest(handler.RestoreProject))

	// tasks
	a.Get("/projects/{title}/tasks", a.handleRequest(handler.GetAllTasks))
	a.Post("/projects/{title}/tasks", a.handleRequest(handler.CreateTask))
	a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
	a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
	a.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.UndoTask))
}


func (a *App) Get(path string, f func(http.ResponseWriter, *http.Request)){
	a.Router.HandleFunc(path,f).Methods("GET")
}

func (a *App) Post(path string, f func(http.ResponseWriter, *http.Request)){
	a.Router.HandleFunc(path,f).Methods("POST")
}
func (a *App) Put(path string, f func(http.ResponseWriter, *http.Request)){
	a.Router.HandleFunc(path,f).Methods("PUT")
}

func (a *App) Delete(path string, f func(http.ResponseWriter, *http.Request)){
	a.Router.HandleFunc(path,f).Methods("DELETE")
}

func (a *App) Run(host string){
	log.Fatal(http.ListenAndServe(host,a.Router))
}

func (a *App) handleRequest(handler func(db *gorm.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		handler(a.DB,w,r)
	}
}

