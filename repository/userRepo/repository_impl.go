package userRepo

import (
	"context"
	"database/sql"
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

func (s *RepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	rows, err := s.Db.QueryContext(ctx, "select * from user where username = ? ;", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u, err := rowsScan(rows)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *RepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	rows, err := s.Db.QueryContext(ctx, "select * from user where email = ? ;", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	u, err := rowsScan(rows)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *RepositoryImpl) RegisterUser(ctx context.Context, u domain.User) error {
	result, err := s.Db.ExecContext(ctx, "insert into user(username,password,email,name) values(?,?,?,?)", u.Username, u.Password, u.Email, u.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)
	return nil
}

// rows scan only for user table
func rowsScan(rows *sql.Rows) (*domain.User, error) {
	u := &domain.User{}
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}
