package handler

import (
	"net/http"

	"github.com/rana-touseef11/go-chi-postgresql/internal/dto"
	"github.com/rana-touseef11/go-chi-postgresql/internal/middleware"
	"github.com/rana-touseef11/go-chi-postgresql/internal/response"
	"github.com/rana-touseef11/go-chi-postgresql/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// GetAll
// @Description  /users
// @Tags         Users
// @Router       /users [get]
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, users)
}

// Login
// @Description  /auth/login
// @Tags         Auth
// @Param        payload  body  dto.UserLoginRequest false " "
// @Router       /auth/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req, ok = middleware.GetBody[dto.UserLoginRequest](r)
	if !ok {
		http.Error(w, "Invaild Context", http.StatusBadRequest)
		return
	}

	data, token, err := h.service.Login(r.Context(), *req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := make(map[string]any)
	res["data"] = data
	res["token"] = token

	response.WriteJson(w, http.StatusOK, res)
}

// Create
// @Description  /users
// @Tags         Users
// @Security     BearerAuth
// @Param        payload  body  dto.CreateUserRequest false " "
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req, ok = middleware.GetBody[dto.CreateUserRequest](r)
	if !ok {
		http.Error(w, "Invaild Context", http.StatusBadRequest)
		return
	}

	data, err := h.service.Create(r.Context(), *req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, data)
}

// GetById
// @Description  /users/{id}
// @Tags         Users
// @Param        id    path     string  true "User Id"
// @Success      200
// @Router       /users/{id} [get]
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

// Update
// @Description  /users/{id}
// @Tags         Users
// @Security     BearerAuth
// @Param        id  path  string true "User Id"
// @Param        payload  body  dto.UpdateUserRequest false " "
// @Router       /users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req, ok = middleware.GetBody[dto.UpdateUserRequest](r)
	if !ok {
		http.Error(w, "Invaild Context", http.StatusBadRequest)
		return
	}

	user, err := h.service.Update(r.Context(), id, *req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, user)
}

// Delete
// @Description  /users/{id}
// @Tags         Users
// @Security     BearerAuth
// @Param        id  path  string true "User Id"
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var err = h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJson(w, http.StatusOK, "Deleted")
}
