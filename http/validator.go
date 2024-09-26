package http

import (
	"context"

	valid "github.com/go-playground/validator/v10"
)

type Validator struct {
	engine *valid.Validate
}

func NewValidator(ctx context.Context) *Validator {
	return &Validator{
		engine: valid.New(),
	}
}

func (v *Validator) Validate(ctx context.Context, data interface{}) error {
	return v.engine.StructCtx(ctx, data)
}
