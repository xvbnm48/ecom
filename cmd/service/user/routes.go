package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/ecom/types"
	"github.com/xvbnm48/ecom/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// req json payload
	var payload types.UserRegister
	if err := utils.ParseJSON(r, payload); err != nil {

	}
}
