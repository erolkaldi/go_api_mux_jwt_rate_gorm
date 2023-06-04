package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erolkaldi/agency/pkg/api"
	"github.com/erolkaldi/agency/pkg/middleware"
	"github.com/erolkaldi/agency/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	Db     *gorm.DB
	Config *models.Config
}

func (app *App) InitializeDB() bool {
	app.Config = &models.Config{}
	app.Config.GetConfigValues()
	connectionString := getConnectionString(app.Config)
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

func getConnectionString(config *models.Config) string {
	return fmt.Sprintf("server=" + config.SqlServer.Server + ";user id=" + config.SqlServer.User + ";password=" + config.SqlServer.Password + ";database=" + config.SqlServer.DbName + ";")
}
func (a *App) Routes() {
	a.Router = mux.NewRouter()
	userApi := api.InitializeUserApi(a.Db)
	a.Router.Handle("/user/{id:[0-9]+}", middleware.RateCheckLimit(middleware.AuthMiddleware(userApi.GetUserById()))).Methods("GET")
	a.Router.Handle("/user", middleware.AuthMiddleware(userApi.CreateUser())).Methods("POST")
	a.Router.HandleFunc("/login", userApi.Login()).Methods("POST")
}

func (a *App) Run() {
	fmt.Printf("Server started at %s\n", a.Config.Api.Port)
	log.Fatal(http.ListenAndServe(a.Config.Api.Port, a.Router))
}
