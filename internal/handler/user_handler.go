package handler

import (
	"errors"
	"net/http"

	"github.com/thenopholo/go-bid/internal/use_case/user"
	"github.com/thenopholo/go-bid/internal/utils"
)

func (h *Handler) UserSignup(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJSON[user.CreateUserReq](r)
	if err != nil {
		_ = utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := h.UserService.CreateUser(r.Context(), data.Name, data.Email, data.Password, data.Bio)
	if err != nil {
		if errors.Is(err, utils.ErrDuplicateUserNameOrPassword) {
			_ = utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, map[string]any{
        "error": "email or username already exists",
      })
			return
		}
	}

  _ = utils.EncodeJSON(w, r, http.StatusCreated, map[string]any{
    "user_id": id,
  })

}

func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) UserLogout(w http.ResponseWriter, r *http.Request) {}
