package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/mt-w10d1/internal/handlers"
	"github.com/metgag/mt-w10d1/internal/repositories"
)

func InitAuthRouter(router *gin.Engine, dbpool *pgxpool.Pool) {
	ar := repositories.NewAuthRepository(dbpool)
	ah := handlers.NewAuthHandler(ar)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", ah.AddUser)
		authRouter.POST("/login", ah.AddCurrUser)
	}
}
