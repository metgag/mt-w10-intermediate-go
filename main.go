package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// var users = []User{}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Printf("unable to create connection pool: %s\n", err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return
	} else {
		log.Printf("connected to database %s as user %s\n", os.Getenv("DB_NAME"), os.Getenv("DB_USER"))
	}

	r := gin.Default()

	r.POST("/register", func(ctx *gin.Context) {
		var reg = User{}

		if err := ctx.ShouldBindJSON(&reg); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		validate := validator.New()
		if err := validate.Struct(reg); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		sql := `
			INSERT INTO users (email, password)
			VALUES ($1, $2)
			RETURNING id
		`
		var id uint16
		if err := dbpool.QueryRow(ctx, sql, reg.Email, reg.Password).Scan(&id); err != nil {
			log.Printf("error creating user: %s\n", err)
		}

		if id == 0 {
			ctx.JSON(http.StatusConflict, gin.H{
				"error":   "duplicate email addresses",
				"success": false,
			})
			return
		}

		// users = append(users, reg)
		// fmt.Println(users)

		ctx.JSON(http.StatusOK, gin.H{
			"result":  fmt.Sprintf("created user w/ ID: %d", id),
			"success": true,
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var login = Login{}

		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

		sql := `
			SELECT id, email, password
			FROM users;
		`

		rows, err := dbpool.Query(ctx, sql)
		if err != nil {
			log.Printf("internal server error: %s\n", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"result":  []any{},
			})
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
				log.Printf("query failed: %s\n", err.Error())
				return
			}
			users = append(users, user)
		}

		if len(users) <= 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"result":  []any{},
			})
			return
		}

		for _, v := range users {
			if v.Email == login.Email {
				if v.Password == login.Password {
					ctx.JSON(http.StatusOK, gin.H{
						"status": "logged in succesfully",
					})
				} else {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"status": "password tidak sesuai",
					})
				}
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "email tidak terdaftar",
		})
	})

	r.PATCH("/movie", func(ctx *gin.Context) {
		idParam, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   fmt.Sprintf("invalid id parameter: %s", err),
				"success": false,
			})
			return
		}

		var updateMovie Foo
		if err := ctx.ShouldBindJSON(&updateMovie); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}

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
		// fmt.Println(foo)

		dbpool.Exec(ctx, foo)
		ctx.JSON(http.StatusCreated, gin.H{
			"result":  fmt.Sprintf("update ID: %d", idParam),
			"success": true,
		})
	})

	r.Run()
}

type Foo struct {
	Title       string `db:"title" json:"title"`
	Director    string `db:"director" json:"director"`
	ReleaseYear string `db:"release_year" json:"release_year"`
}

type Movie struct {
	ID          uint16    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Director    string    `db:"director" json:"director"`
	ReleaseYear string    `db:"release_year" json:"release_year"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Login struct {
	Email    string `db:"email" json:"email" validate:"email"`
	Password string `db:"password" json:"password"`
}

type User struct {
	ID       uint16 `db:"id" json:"id"`
	Email    string `db:"email" json:"email" validate:"email"`
	Password string `db:"password" json:"password" validate:"min=8,containsany=!@#$%^&*,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}
