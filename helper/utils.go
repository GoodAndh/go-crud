package helper

import (
	"newestcdd/model/domain"
	"newestcdd/model/web"
)

func ConverUserToWeb(u *domain.User) *web.UserWeb {
	return &web.UserWeb{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Name:     u.Name,
	}
}

func ConvertPdkToWeb(p *domain.Product) *web.ProductWeb {
	return &web.ProductWeb{
		ProdukName: p.ProdukName,
		Deskripsi:  p.Deskripsi,
		Category:   p.Category,
		Userid:     p.Userid,
		Harga:      p.Harga,
		Quantity:   p.Quantity,
	}
}

func ConvertPdkSlice(p []domain.Product) []web.ProductWeb {
	d := []web.ProductWeb{}
	for _, v := range p {
		d = append(d, *ConvertPdkToWeb(&v))
	}
	return d
}
