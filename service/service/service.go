package service

import (
	"context"
	"newestcdd/model/web"
	"strconv"
)

type Service interface {
	Login(ctx context.Context, request web.UserLoginPayload) (*web.UserWeb, map[any]any)
	RegisterUser(ctx context.Context, request web.UserRegisterPayload) map[any]any
	GetAllProduct(ctx context.Context) ([]web.ProductWeb, error)
	GetByProduct(ctx context.Context,input string) ([]web.ProductWeb, error)

}

func ForRangeProduct(data []web.ProductWeb) map[any]any {
	d := make(map[any]any)
	da := make(map[any]any)
	for i, v := range data {
		d1 := map[interface{}]interface{}{
			"nama":      v.ProdukName,
			"harga":     v.Harga,
			"quantity":  v.Quantity,
			"deskripsi": v.Deskripsi,
			"category":  v.Category,
		}
		d[strconv.Itoa(i+1)] = d1
	}
	da["data"] = d
	return da
}
