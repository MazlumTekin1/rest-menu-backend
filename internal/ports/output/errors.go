package ports

//PATH: internal/ports/output/errors.go

import "errors"

var (
	ErrMenuNotFound     = errors.New("menu not found")
	ErrMenuExists       = errors.New("menu already exists")
	ErrProductExists    = errors.New("product already exists")
	ErrProductNotFound  = errors.New("product not found")
	ErrCategoryExists   = errors.New("category already exists")
	ErrCategoryNotFound = errors.New("category not found")
	ErrInvalidID        = errors.New("invalid ID")
	ErrInvalidRequest   = errors.New("invalid request")
	ErrInternalServer   = errors.New("internal server error")
	ErrUserNotFound     = errors.New("user not found")
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	SwaggerErrMenuNotFound     = ErrorResponse{Code: 404, Message: ErrMenuNotFound.Error()}
	SwaggerErrMenuExists       = ErrorResponse{Code: 409, Message: ErrMenuExists.Error()}
	SwaggerErrProductExists    = ErrorResponse{Code: 409, Message: ErrProductExists.Error()}
	SwaggerErrProductNotFound  = ErrorResponse{Code: 404, Message: ErrProductNotFound.Error()}
	SwaggerErrCategoryExists   = ErrorResponse{Code: 409, Message: ErrCategoryExists.Error()}
	SwaggerErrCategoryNotFound = ErrorResponse{Code: 404, Message: ErrCategoryNotFound.Error()}
	SwaggerErrInvalidID        = ErrorResponse{Code: 400, Message: ErrInvalidID.Error()}
	SwaggerErrInvalidRequest   = ErrorResponse{Code: 400, Message: ErrInvalidRequest.Error()}
	SwaggerErrInternalServer   = ErrorResponse{Code: 500, Message: ErrInternalServer.Error()}
)
