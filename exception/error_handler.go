package exception

import (
	"encoding/json"
	"errors"
	"net/http"
	"newestcdd/model/web"
)

type ErrorWeb interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func ParseJson(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJson(w http.ResponseWriter, code int, status string, v ...any) error {
	w.Header().Add("Content-Type", "application/json")
	p := &web.WebResponse{
		Code:   code,
		Status: status,
		Data:   v,
	}

	return json.NewEncoder(w).Encode(p)
}

var (
	ErrNoRows   error = errors.New("produk kosong,internal server error")
	ErrNotFound error = errors.New("produk yang diminta tidak ada")
	ErrSesNotFound error=errors.New("login terlebih dahulu untuk melanjutkan")
)

func WriteErrorInternalServerError(w http.ResponseWriter, status string, v ...any) error {
	return WriteJson(w, http.StatusInternalServerError, status, v...)
}

func WriteErrorBadRequest(w http.ResponseWriter, status string, v ...any) error {
	return WriteJson(w, http.StatusBadRequest, status, v...)
}

func WriteNotFoundHandler(w http.ResponseWriter, status string, v ...any) {
	WriteJson(w, http.StatusNotFound, status, v...)
}

func WriteMethodError(w http.ResponseWriter, status string, v ...any) {
	WriteJson(w, http.StatusMethodNotAllowed, status, v...)
}
