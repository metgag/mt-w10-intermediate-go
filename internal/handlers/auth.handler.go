package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/metgag/mt-w10d1/internal/models"
	"github.com/metgag/mt-w10d1/internal/repositories"
)

type AuthHandler struct {
	ar *repositories.AuthRepository
}

func NewAuthHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

func (a *AuthHandler) AddUser(ctx *gin.Context) {
	var reg = models.User{}

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

	id, err := a.ar.AddNewUser(ctx.Request.Context(), reg)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error":   "email address is already registered. Please use a different email or log in",
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":  fmt.Sprintf("created user w/ ID: %d", id),
		"success": true,
	})
}

func (a *AuthHandler) AddCurrUser(ctx *gin.Context) {
	var login = models.Login{}

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	rows, err := a.ar.GetUsersData(ctx)
	if err != nil {
		log.Printf("internal server error: %s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"result":  []any{},
		})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
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
}
