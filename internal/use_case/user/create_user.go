package user

import (
	"context"

	"github.com/thenopholo/go-bid/internal/validator"
)

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Name), "name", "This field can not be empty")
	eval.CheckField(validator.NotBlank(req.Email), "email", "This field can not be empty")
	eval.CheckField(validator.NotBlank(string(req.Password)), "password", "This field can not be empty")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "This field can not be empty")

	eval.CheckField(validator.MinChars(req.Bio, 10) && validator.MaxChars(req.Bio, 255), "bio", "The bio must have between 10 and 255 caracters")
	eval.CheckField(validator.MinPasswordStrength(string(req.Password)), "password", "The password must follow de strength rules")
	eval.CheckField(validator.ValidEmail(req.Email), "email", "The field must have a valid email")

	return eval
}
