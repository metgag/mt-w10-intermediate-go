package repositories

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/metgag/mt-w10d1/internal/models"
)

type MovieRepository struct {
	dbpool *pgxpool.Pool
}

func NewMovieRepository(dbpool *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{dbpool: dbpool}
}

func (m *MovieRepository) UpdateMovieData(ctx *gin.Context, updateMovie models.Foo, idParam int) (pgconn.CommandTag, error) {
	rt := reflect.TypeOf(updateMovie)
	rv := reflect.ValueOf(updateMovie)

	foo := "UPDATE movies SET "

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		if value.String() == "" {
			continue
		}

		foo += fmt.Sprintf("%s = '%s'", field.Tag.Get("db"), value)
		foo += ", "
	}

	foo += fmt.Sprintf(" updated_at = current_timestamp WHERE id = %d;", idParam)

	return m.dbpool.Exec(ctx, foo)
}
