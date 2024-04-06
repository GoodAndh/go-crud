package server

import (
	"database/sql"
	"net/http"
	productrepo "newestcdd/repository/productRepo"
	"newestcdd/repository/userRepo"
	"newestcdd/service/routes"
	"newestcdd/service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
	Validate *validator.Validate
}

func NewApiServer(addr string, db *sql.DB,validate *validator.Validate) *ApiServer {
	return &ApiServer{
		Addr: addr,
		Db:   db,
		Validate: validate,
	}
}

func (a *ApiServer) Run() error {
	router := httprouter.New()
	
	userRepo:=userRepo.NewRepository(a.Db)
	productRepo:=productrepo.NewRepository(a.Db)

	service:=service.NewService(userRepo,a.Validate,productRepo)
	
	handler:=routes.NewHandler(service)
	handler.RegisterRoutes(router)


	return http.ListenAndServe(a.Addr, router)
}
