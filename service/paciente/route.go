package paciente

import (
	"enfermeria_go/types"
	"enfermeria_go/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {

	//repositorio de store
	//store types
	store types.PacienteStore
}

func NewHandlerPaciente(store types.PacienteStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/registro-paciente", h.handleRegisterPacinete).Methods(http.MethodPost)
	router.HandleFunc("/get-paciente-documento/{documento}", h.handleGetPaciente).Methods(http.MethodGet)
	router.HandleFunc("/update-paciente", h.handleUpdatePaciente).Methods(http.MethodPut)
	router.HandleFunc("/delete-paciente/{documento}", h.handleDeletePaciente).Methods(http.MethodDelete)
}

func (h *Handler) handleGetPaciente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["documento"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	documento, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	paciente, err := h.store.GetPacientePorDocumento(documento)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, paciente)
}

func (h *Handler) handleRegisterPacinete(w http.ResponseWriter, r *http.Request) {
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

	// check if user exists
	getUser, err := h.store.GetPacientePorDocumento(user.Documento)

	if getUser != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with Documento %s already exists", user.Documento))
		return
	}

	log.Println(3)

	err = h.store.CrearPaciente(types.RegisterPaciente{
		Documento:        user.Documento,
		Tipo:             user.Tipo,
		Primer_nombre:    user.Primer_nombre,
		Segundo_nombre:   user.Segundo_nombre,
		Primer_apellido:  user.Primer_apellido,
		Segundo_apellido: user.Segundo_apellido,
		Sexo:             user.Sexo,
		Eps:              user.Eps,
		Fecha_nacimiento: user.Fecha_nacimiento,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		log.Println(5)
		return
	}

	_ = utils.WriteJSON(w, http.StatusCreated, "ok")
}

func (h *Handler) handleUpdatePaciente(w http.ResponseWriter, r *http.Request) {

	var paciente types.RegisterPaciente
	if err := utils.ParseJSON(r, &paciente); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetPacientePorDocumento(paciente.Documento)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with Documento %s does not exists", paciente.Documento))
		return
	}

	err = h.store.UpdatePaciente(paciente)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusOK, "ok")
}

func (h *Handler) handleDeletePaciente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["documento"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing document ID"))
		return
	}

	documento, err := strconv.Atoi(str)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid document ID"))
		return
	}

	err = h.store.DeltePaciente(documento)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusOK, "ok")
}
