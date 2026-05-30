package handler

import (
	"net/http"

	"github.com/uniesp/desafio-devops/internal/utils"
)

func Health(
	w http.ResponseWriter,
	r *http.Request,
) {

	utils.JSON(
		w,
		http.StatusOK,
		map[string]string{
			"status": "ok",
		},
	)
}