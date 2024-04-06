package exception

import "net/http"



func NotFoundError()ErrorWeb  {
	return ErrNotFoundHandler{}
}

type ErrNotFoundHandler struct {
}

func (e ErrNotFoundHandler)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	WriteNotFoundHandler(w,"halaman yang anda cari tidak valid")
}
