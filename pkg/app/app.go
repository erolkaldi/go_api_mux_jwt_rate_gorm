package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erolkaldi/agency/pkg/api"
	"github.com/erolkaldi/agency/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	Db     *gorm.DB
}

func (app *App) InitializeDB() bool {
	er := godotenv.Load()
	if er != nil {
		panic("Environment not loaded")
	}
	connectionString := getConnectionString()
	println(connectionString)
	var err error
	var dial = sqlserver.Open(connectionString)

	app.Db, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}

}

func getConnectionString() string {
	return fmt.Sprintf("server=" + os.Getenv("SERVER") + ";user id=" + os.Getenv("USER_ID") + ";password=" + os.Getenv("PASSWORD") + ";database=" + os.Getenv("DB") + ";")
}
func (a *App) Routes() {
	a.Router = mux.NewRouter()
	userApi := api.InitializeUserApi(a.Db)
	a.Router.Handle("/user/{id:[0-9]+}", middleware.RateCheckLimit(middleware.AuthMiddleware(userApi.GetUserById()))).Methods("GET")
	a.Router.Handle("/user", middleware.AuthMiddleware(userApi.CreateUser())).Methods("POST")
	a.Router.HandleFunc("/login", userApi.Login()).Methods("POST")
}

func (a *App) Run() {
	port := os.Getenv("PORT")
	fmt.Printf("Server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}
