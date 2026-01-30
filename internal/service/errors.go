package service

import "errors"

var (
	ErrCategoryNotFound  = errors.New("category not found")
	ErrCategoryNameEmpty = errors.New("name is required")
)

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrProductNameEmpty    = errors.New("product name is empty")
	ErrProductPriceInvalid = errors.New("product price must be greater than 0")
)
