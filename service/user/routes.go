package user

import (
	"enfermeria_go/config"
	"enfermeria_go/service/auth"
	"enfermeria_go/types"
	"enfermeria_go/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {

	//repositorio de store
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %s", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting user by email: %w", err))
		return
	}

	if u == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user with email %s not found", user.Email))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
		return
	}

	//generate token
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating token: %w", err))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]string{"toke": token})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %s", errors))
		return
	}

	// check if user exists
	getUser, err := h.store.GetUserByEmail(user.Email)

	if getUser != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	log.Println(3)

	// hash password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Println(4)

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		log.Println(5)
		return
	}

	_ = utils.WriteJSON(w, http.StatusCreated, "ok")

}

func (h *Handler) handleRegistroPaciente(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterPaciente
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %s", errors))
		return
	}
}
