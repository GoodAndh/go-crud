package routes

import (
	"net/http"
	"newestcdd/exception"
	"newestcdd/model/web"
	"newestcdd/service/auth"
	"newestcdd/service/service"
	"newestcdd/views"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Store service.Service
}

func NewHandler(store service.Service) *Handler {
	return &Handler{
		Store: store,
	}
}

func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	router.GET("/login", h.handleLogin)
	router.POST("/login", h.handleLogin)

	router.GET("/register", h.handleRegister)
	router.POST("/register", h.handleRegister)

	router.POST("/", h.handleProduct)
	router.GET("/", h.handleProduct)

	router.NotFound = exception.NotFoundError()
	router.MethodNotAllowed = exception.MethodNotAllowedMethod()
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		err := views.TemplateExecuted(w, nil, "views/login.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		r.ParseForm()
		input := &web.UserLoginPayload{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		exception.ParseJson(r, input)

		u, errorList := h.Store.Login(r.Context(), *input)
		if u == nil || len(errorList) > 0 {
			err := views.TemplateExecuted(w, errorList, "views/login.html")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		ses, _ := auth.Session.Get(r, "lg-ses")
		ses.Values["auten"] = true
		ses.Values["username"] = u.Username
		ses.Options.MaxAge = 600
		time.Sleep(1 * time.Second)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		err := views.TemplateExecuted(w, nil, "views/register.html")
		if err != nil {
			exception.WriteErrorInternalServerError(w, err.Error(), err)
			return
		}
	case http.MethodPost:
		r.ParseForm()
		payload := &web.UserRegisterPayload{
			Username:  r.Form.Get("username"),
			Password:  r.Form.Get("password"),
			CPassword: r.Form.Get("cpassword"),
			Email:     r.Form.Get("email"),
			Name:      r.Form.Get("name"),
		}

		exception.ParseJson(r, payload)

		errorList := h.Store.RegisterUser(r.Context(), *payload)
		if len(errorList) > 1 {
			err := views.TemplateExecuted(w, errorList, "views/register.html")
			if err != nil {
				exception.WriteErrorInternalServerError(w, err.Error(), err)
				return
			}
		}

		errorList["pesan"] = true
		err := views.TemplateExecuted(w, errorList, "views/register.html")
		if err != nil {
			exception.WriteErrorInternalServerError(w, err.Error(), err)
			return
		}
	}
}

func (h *Handler) handleProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		errorList := make(map[any]any)
		p, err := h.Store.GetAllProduct(r.Context())
		if err != nil {
			errorList["error"] = err
			err = views.TemplateExecuted(w, errorList, "views/index.html")
			if err != nil {
				exception.WriteErrorInternalServerError(w, err.Error(), err)
				return
			}
			return
		}
		errorList = service.ForRangeProduct(p)
		err = views.TemplateExecuted(w, errorList, "views/index.html")
		if err != nil {
			exception.WriteErrorInternalServerError(w, err.Error(), err)
			return
		}
	case http.MethodPost:
		errorList := make(map[any]any)

		ses, err := auth.Session.Get(r, "lg-ses")
		if err != nil {
			errorList["error"] = exception.ErrSesNotFound
			errorList["errormsg"] = true
			err = views.TemplateExecuted(w, errorList, "views/index.html")
			if err != nil {
				exception.WriteErrorInternalServerError(w, err.Error(), err)
				return
			}
			return
		}

		if auten, ok := ses.Values["auten"].(bool); !ok || !auten {
			errorList["error"] = exception.ErrSesNotFound
			errorList["errormsg"] = true
			err = views.TemplateExecuted(w, errorList, "views/index.html")
			if err != nil {
				exception.WriteErrorInternalServerError(w, err.Error(), err)
				return
			}
			return
		}

		r.ParseForm()
		input := r.Form.Get("namaproduk")
		exception.ParseJson(r, input)
		p, err := h.Store.GetByProduct(r.Context(), input)
		if err != nil {
			errorList["error"] = err
			err = views.TemplateExecuted(w, errorList, "views/index.html")
			if err != nil {
				exception.WriteErrorInternalServerError(w, err.Error(), err)
				return
			}
			return
		}
		errorList = service.ForRangeProduct(p)
		err = views.TemplateExecuted(w, errorList, "views/index.html")
		if err != nil {
			exception.WriteErrorInternalServerError(w, err.Error(), err)
			return
		}
	}
}
