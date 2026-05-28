package handler

import (
	"encoding/json"
	"net/http"

	"github.com/uniesp/desafio-devops/internal/service"
	"github.com/uniesp/desafio-devops/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {

	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		utils.Error(
			w,
			http.StatusBadRequest,
			"invalid body",
		)
		return
	}

	err = h.service.Create(body.Name)

	if err != nil {
		utils.Error(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	utils.JSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "user created",
		},
	)
}