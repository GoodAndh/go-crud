package exception

import "net/http"

func MethodNotAllowedMethod() ErrorWeb {
	return ErrMethod{}
}

type ErrMethod struct {
}

func (e ErrMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	WriteMethodError(w, "method tidak diperbolehkan,silahkan hubungi pemilik website")
}
