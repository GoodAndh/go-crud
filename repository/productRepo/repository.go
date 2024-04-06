package productrepo

import (
	"context"
	"newestcdd/model/domain"
)

type Repository interface {
	GetAllProduct(ctx context.Context)([]domain.Product,error)
	GetByProduct(ctx context.Context,input string)([]domain.Product,error)
	CreateProduct(ctx context.Context,p domain.Product)error
}