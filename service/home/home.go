package home

import (
	"enfermeria_go/utils"
	"enfermeria_go/views"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", handlerHome).Methods(http.MethodGet)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	err := views.Home().Render(r.Context(), w)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %s", err))
		return
	}
}
