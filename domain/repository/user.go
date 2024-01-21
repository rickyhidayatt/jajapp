package repository

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	db *sqlx.DB
}

type UserRepositoryInterface interface {
}
