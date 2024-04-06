package productrepo

import (
	"context"
	"database/sql"
	"newestcdd/exception"
	"newestcdd/model/domain"
)

type RepositoryImpl struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		Db: db,
	}
}

func (s *RepositoryImpl) GetAllProduct(ctx context.Context) ([]domain.Product, error) {
	rows, err := s.Db.QueryContext(ctx, "select * from produk")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	p, err := rowsScan(rows)
	if err != nil {
		return nil, err
	}
	if len(p) < 1 || p == nil {
		return nil, exception.ErrNotFound
	}
	return p, nil
}

func (s *RepositoryImpl) GetByProduct(ctx context.Context, input string) ([]domain.Product, error) {

	st := ("select * from produk where produkname in  ( select produkname from produk where produkname like ? )  or category in( select category from produk where category like ?  );")

	rows, err := s.Db.QueryContext(ctx, st, input+"%", input+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	p, err := rowsScan(rows)
	if err != nil {
		return nil, err
	}
	if len(p) < 1 || p == nil {
		return nil, exception.ErrNotFound
	}
	return p, nil
}

func (s *RepositoryImpl) CreateProduct(ctx context.Context, p domain.Product) error {
	result, err := s.Db.ExecContext(ctx, "insert into produk (produkname,deskripsi,category,userid,harga,quantity) values(?,?,?,?,?,?)", p.ProdukName, p.Deskripsi, p.Category, p.Userid, p.Harga, p.Quantity)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.Id = int(id)
	return nil
}

// rows scan only for product
func rowsScan(rows *sql.Rows) ([]domain.Product, error) {
	p := &domain.Product{}
	var pe []domain.Product
	for rows.Next() {
		err := rows.Scan(&p.Id, &p.ProdukName, &p.Deskripsi, &p.Category, &p.Userid, &p.Harga, &p.Quantity)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, exception.ErrNoRows
			}
			return nil, err
		}
		pe = append(pe, *p)

	}
	return pe, nil
}
