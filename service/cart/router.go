package cart

import (
	"net/http"

	"github.com/fmelihh/crud-api-go/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {

}
