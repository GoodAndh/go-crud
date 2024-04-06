package service

import (
	"context"
	"net/http"
	"newestcdd/exception"
	"newestcdd/helper"
	"newestcdd/model/domain"
	"newestcdd/model/web"
	productrepo "newestcdd/repository/productRepo"
	"newestcdd/repository/userRepo"
	"newestcdd/service/auth"

	"github.com/go-playground/validator/v10"
)

type ServiceImpl struct {
	uRepo    userRepo.Repository
	pRepo    productrepo.Repository
	Validate *validator.Validate
	w        http.ResponseWriter
}

func NewService(urepo userRepo.Repository, validate *validator.Validate, prepo productrepo.Repository) Service {
	return &ServiceImpl{
		uRepo:    urepo,
		Validate: validate,
		pRepo:    prepo,
	}
}

// if len(map[any][any]) > 0 or userLoginPayload is nil there`s error available
func (svc *ServiceImpl) Login(ctx context.Context, request web.UserLoginPayload) (*web.UserWeb, map[any]any) {
	ok, errorList := helper.ValidateCustomStruct(svc.Validate, request)
	if !ok || len(errorList) > 0 {
		return nil, errorList
	}

	u, err := svc.uRepo.GetUserByUsername(ctx, request.Username)
	if err != nil {
		errorList["error"] = err
		return nil, errorList
	}

	err = auth.ComparePassword(u.Password, []byte(request.Password))
	if err != nil {
		errorList["error"] = err
		return nil, errorList
	}

	return helper.ConverUserToWeb(u), errorList
}

// errorList must be len(errorList)=0 if len(errorList)>0 error is available
func (svc *ServiceImpl) RegisterUser(ctx context.Context, request web.UserRegisterPayload) map[any]any {
	ok, errorList := helper.ValidateCustomStruct(svc.Validate, request)
	if !ok || len(errorList) > 0 {
		errorList["error"] = errorList
		return errorList
	}

	u, err := svc.uRepo.GetUserByUsername(ctx, request.Username)
	if u == nil || err == nil {
		errorList["Errorusername"] = errorList
		return errorList
	}

	e, err := svc.uRepo.GetUserByEmail(ctx, request.Email)
	if err == nil || e == nil {
		errorList["Erroremail"] = errorList
		return errorList
	}

	user := &domain.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Name:     request.Name,
	}

	err = svc.uRepo.RegisterUser(ctx, *user)
	if err != nil {
		errorList["error"] = errorList
		return errorList
	}

	return errorList

}

func (svc *ServiceImpl) GetAllProduct(ctx context.Context) ([]web.ProductWeb, error) {
	p, err := svc.pRepo.GetAllProduct(ctx)
	if err != nil {
		if err == exception.ErrNoRows {
			return nil, exception.WriteErrorInternalServerError(svc.w, err.Error(), err)
		}
		return nil, err
	}
	return helper.ConvertPdkSlice(p), nil
}

func (svc *ServiceImpl) GetByProduct(ctx context.Context, input string) ([]web.ProductWeb, error) {
	p, err := svc.pRepo.GetByProduct(ctx, input)
	if err != nil {
		if err == exception.ErrNoRows {
			return nil, exception.WriteErrorInternalServerError(svc.w, err.Error(), err)
		}
		return nil, err
	}
	return helper.ConvertPdkSlice(p), nil
}
