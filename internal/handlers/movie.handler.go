package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/metgag/mt-w10d1/internal/models"
	"github.com/metgag/mt-w10d1/internal/repositories"
)

type MovieHandler struct {
	mr *repositories.MovieRepository
}

func NewMovieHandler(mr *repositories.MovieRepository) *MovieHandler {
	return &MovieHandler{mr: mr}
}

func (m *MovieHandler) UpdateMovie(ctx *gin.Context) {
	idParam, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   fmt.Sprintf("invalid id parameter: %s", err),
			"success": false,
		})
		return
	}

	var updateMovie models.Foo
	if err := ctx.ShouldBindJSON(&updateMovie); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	ctag, err := m.mr.UpdateMovieData(ctx, updateMovie, idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}
	if ctag.RowsAffected() == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   fmt.Sprintf("there is no movie w/ ID, %d", idParam),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"result":  fmt.Sprintf("successfully update movie w/ ID, %d", idParam),
		"success": true,
	})
}
