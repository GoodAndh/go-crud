package userRepo

import (
	"context"
	"newestcdd/model/domain"
)

type Repository interface {
	GetUserByUsername(ctx context.Context,username string)(*domain.User,error)
	GetUserByEmail(ctx context.Context,email string)(*domain.User,error)
	RegisterUser(ctx context.Context,user domain.User)error
}
