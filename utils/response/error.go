package response

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrNotFound = errors.New("Your requested item is not found")
	ErrConflict = errors.New("Your item already exist")
	ErrCustom   = errors.New("Custom")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrCustom:
		return http.StatusOK

	default:
		return http.StatusInternalServerError
	}
}
