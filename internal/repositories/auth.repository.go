package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/mt-w10d1/internal/models"
)

type AuthRepository struct {
	dbpool *pgxpool.Pool
}

func NewAuthRepository(dbpool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{dbpool: dbpool}
}

func (a *AuthRepository) AddNewUser(ctx context.Context, reg models.User) (uint16, error) {
	sql := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`
	var id uint16
	if err := a.dbpool.QueryRow(ctx, sql, reg.Email, reg.Password).Scan(&id); err != nil {
		log.Printf("error creating user: %s\n", err)
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("duplicate email addresses")
	}

	// users = append(users, reg)
	// fmt.Println(users)

	return id, nil
}

func (a *AuthRepository) GetUsersData(ctx *gin.Context) (pgx.Rows, error) {
	sql := `
		SELECT id, email, password
		FROM users;
	`
	return a.dbpool.Query(ctx, sql)
}
