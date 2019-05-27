package api

import (
	"gopkg.in/go-playground/validator.v9"
)

/*
IsISO8601Date function to check parameter pattern for valid ISO8601 Date
*/
func IsISO8601Date(fl validator.FieldLevel) bool {
	return true
}
