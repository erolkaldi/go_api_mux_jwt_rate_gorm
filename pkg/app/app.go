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
	Router      *mux.Router
	Db          *gorm.DB
	Config      *models.Config
	RateLimiter *middleware.RateLimiterStore
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
	}

	app.RateLimiter = middleware.NewRateLimiterStore()

	return true

}

func getConnectionString(config *models.Config) string {
	return fmt.Sprintf("server=" + config.SqlServer.Server + ";user id=" + config.SqlServer.User + ";password=" + config.SqlServer.Password + ";database=" + config.SqlServer.DbName + ";")
}
func (a *App) Routes() {
	a.Router = mux.NewRouter()
	userApi := api.InitializeUserApi(a.Db, &a.Config.Smtp)
	a.Router.Handle("/user/{id:[0-9]+}", a.authorizeRequest(userApi.GetUserById(), userApi, true)).Methods("GET")
	a.Router.Handle("/user", a.authorizeRequest(userApi.CreateUser(), userApi, true)).Methods("POST")
	a.Router.Handle("/login", a.authorizeRequest(userApi.Login(), userApi, false)).Methods("POST")
	a.Router.Handle("/register", a.authorizeRequest(userApi.RegisterUser(), userApi, false)).Methods("POST")
}

func (a *App) authorizeRequest(next http.Handler, userApi *api.UserApi, tokened bool) http.Handler {
	if tokened {
		return a.RateLimiter.RateCheckLimit(middleware.AppKeyAuthorization(middleware.AuthMiddleware(next), &a.Config.Api))
	} else {
		return a.RateLimiter.RateCheckLimit(middleware.AppKeyAuthorization(next, &a.Config.Api))
	}
}

func (a *App) Run() {
	fmt.Printf("Server started at %s\n", a.Config.Api.Port)
	log.Fatal(http.ListenAndServe(a.Config.Api.Port, a.Router))
}
