package project

import (
	"log"
	"net/http"
	"project-testing/pkg/response"
)

func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.projectSvc.GetAllUser(ctx)
	if err != nil {
		log.Println("[ERROR]Handler GetAllUser", err.Error())
		resp.Data = result
		return
	}

	resp.Data = result
}
