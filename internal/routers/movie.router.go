package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/mt-w10d1/internal/handlers"
	"github.com/metgag/mt-w10d1/internal/repositories"
)

func InitMovieRouter(r *gin.Engine, dbpool *pgxpool.Pool) {
	mr := repositories.NewMovieRepository(dbpool)
	mh := handlers.NewMovieHandler(mr)

	r.PATCH("/movie/:id", mh.UpdateMovie)
}
