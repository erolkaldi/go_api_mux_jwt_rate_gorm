package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/erolkaldi/agency/pkg/models"
	"github.com/erolkaldi/agency/pkg/repository"
	"github.com/erolkaldi/agency/pkg/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserApi struct {
	service *service.UserService
	db      *gorm.DB
}

func InitializeUserApi(db *gorm.DB) *UserApi {
	repository := repository.Initialize(db)
	service := service.InitializeUserService(repository)
	userApi := UserApi{service: service}
	userApi.Migrate()
	return &userApi
}

func (api *UserApi) Migrate() {
	err := api.service.MigrateUser()
	if err != nil {
		log.Println(err)
	}
}
func (api *UserApi) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		post, err := api.service.GetUserById(id)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, post)
	}
}
func (api *UserApi) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var loginDto models.LoginDto

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&loginDto); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		user, err := api.service.GetUserByEmail(loginDto.Email)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		if user.Password != loginDto.Password {
			RespondWithError(w, http.StatusBadRequest, "Password is wrong")
			return
		}

		tokenDto, er := service.GenerateToken(*user)
		if er != nil {
			RespondWithError(w, http.StatusInternalServerError, "Could not create token")
			return
		}

		RespondWithJSON(w, http.StatusOK, tokenDto)
	}
}
func (api *UserApi) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdUser, err := api.service.SaveUser(&user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, createdUser)
	}
}
