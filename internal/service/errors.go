package service

import "errors"

var (
	ErrCategoryNotFound  = errors.New("category not found")
	ErrCategoryNameEmpty = errors.New("name is required")
)
