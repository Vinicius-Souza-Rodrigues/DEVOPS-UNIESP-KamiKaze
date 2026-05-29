package handler

import (
	"net/http"

	"github.com/uniesp/desafio-devops/internal/service"
	"github.com/uniesp/desafio-devops/internal/utils"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(
	service *service.ProductService,
) *ProductHandler {

	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) FindAll(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	products, err := h.service.FindAll()

	if err != nil {

		utils.Error(
			w,
			http.StatusInternalServerError,
			"failed to fetch products",
		)

		return
	}

	utils.JSON(
		w,
		http.StatusOK,
		products,
	)
}