// Identity of our warehouse actions
package handler

import (
	"fmt"
	"net/http"
)

type WarehouseHandler struct{}

func NewWarehouseHandler() *WarehouseHandler {
	return &WarehouseHandler{}
}

func (h *WarehouseHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display List Warehouse")
}

func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating Warehouse")
}
