package project

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"project-testing/internal/entity/project"
	"project-testing/pkg/response"
	"strconv"
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

func (h *Handler) SearchUserByNameAndttl(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, metadata, err := h.projectSvc.SearchUserByNameAndttl(ctx, r.FormValue("nama"), r.FormValue("ttl"))
	if err != nil {
		log.Println("[ERROR]Handler GetAllUser", err.Error())
		resp.Data = result
		resp.Metadata = metadata
		return
	}

	resp.Data = result
	resp.Metadata = metadata
}

func (h *Handler) GetAllUserPagination(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	limit, _ := strconv.Atoi(r.FormValue("limit"))
	page, _ := strconv.Atoi(r.FormValue("page"))
	result, metadata, err := h.projectSvc.GetAllUserPagination(ctx, limit, page)
	if err != nil {
		log.Println("[ERROR]Handler GetAllUserPagination", err.Error())
		resp.Data = result
		resp.Metadata = metadata
		return
	}

	resp.Data = result
	resp.Metadata = metadata
}

func (h *Handler) SearchUserDataByName(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.projectSvc.SearchUserDataByName(ctx, r.FormValue("keyword"))
	if err != nil {
		log.Println("[ERROR]Handler SearchUserDataByName", err.Error())
		resp.Data = result
		return
	}

	resp.Data = result
}

func (h *Handler) SearchUserDataByKwn(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.projectSvc.SearchUserDataByKwn(ctx, r.FormValue("keyword"))
	if err != nil {
		log.Println("[ERROR]Handler SearchUserDataByKwn", err.Error())
		resp.Data = result
		return
	}

	resp.Data = result
}

func (h *Handler) SearchUserDataByKwnOrName(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.projectSvc.SearchUserDataByKwnOrName(ctx, r.FormValue("keyword"))
	if err != nil {
		log.Println("[ERROR]Handler SearchUserDataByKwnOrName", err.Error())
		resp.Data = result
		return
	}

	resp.Data = result
}

func (h *Handler) InsertDataUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	req := project.User{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	rBody, errJson := ioutil.ReadAll(r.Body)
	if errJson != nil {
		log.Println("err Readjson", errJson)
	}

	defer resp.RenderJSON(w, r)

	json.Unmarshal(rBody, &req)

	err := h.projectSvc.InsertDatauser(ctx, req)
	if err != nil {
		log.Println("[ERROR]Handler InsertDataUser", err.Error())
		return
	}

}
