package handlers

import (
	"encoding/json"
	diSvc "neolist-backend/internal/di/services"
	"neolist-backend/internal/dto"
	"neolist-backend/internal/utils"
	"net/http"
)

type UserHandler struct {
	service diSvc.IUserService
}

func NewUserHandler(service diSvc.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}

	user_service_res, err := h.service.Register(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := dto.RegisterResponse{
		ID:   user_service_res.ID,
		Name: user_service_res.Name,
	}

	utils.ResponseWriter(w, http.StatusCreated, resp)
}

func (h *UserHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, res)
}

func (h *UserHandler) FindByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	res, err := h.service.FindByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, res)
}

func (h *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.ID == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	res, err := h.service.Update(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, res)

}

func (h *UserHandler) SoftDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.SoftDelete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, "User Deleted Successfully")

}

func (h *UserHandler) ForceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.ForceDelete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.ResponseWriter(w, http.StatusOK, "User Deleted Successfully")
}
