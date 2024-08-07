package services

import (
	db "b30northwindapi/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryService struct {
	*db.Queries
}

// constructor
func NewCategoryService(dbConn *pgxpool.Conn) *CategoryService {
	return &CategoryService{
		Queries: db.New(dbConn),
	}
}
