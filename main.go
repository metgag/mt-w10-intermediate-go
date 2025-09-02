package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var users = []Register{}

func main() {
	r := gin.Default()

	r.POST("/register", handleRegister)
	r.POST("/login", handleLogin)

	r.Run()
}

func handleLogin(ctx *gin.Context) {
	var log = Login{}

	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	for _, v := range users {
		if v.Email == log.Email {
			if v.Password == log.Password {
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

func handleRegister(ctx *gin.Context) {
	var reg = Register{}

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

	users = append(users, reg)
	fmt.Println(users)

	ctx.JSON(http.StatusOK, gin.H{
		"result":  reg,
		"success": true,
	})
}

type Login struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,containsany=!@#$%^&*,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}
