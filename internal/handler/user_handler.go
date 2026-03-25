package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rana-touseef11/go-chi-postgresql/internal/dto"
	"github.com/rana-touseef11/go-chi-postgresql/internal/response"
	"github.com/rana-touseef11/go-chi-postgresql/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, users)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invaild Request", http.StatusBadRequest)
		return
	}

	data, token, err := h.service.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := make(map[string]any)
	res["data"] = data
	res["token"] = token

	response.WriteJson(w, http.StatusOK, res)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invaild Request", http.StatusBadRequest)
		return
	}

	data, err := h.service.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, data)
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// native net/http
	id := r.PathValue("id")
	// id := r.URL.Query().Get("id")

	// id := chi.URLParam(r, "id")
	user, err := h.service.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req dto.UpdateUserRequest
	if r.Body == nil {
		http.Error(w, "There noting to update", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var err = h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, "Deleted")
}
