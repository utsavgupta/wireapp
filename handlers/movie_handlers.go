package handlers

import (
	"net/http"

	"github.com/utsavgupta/wireapp/ent"

	"github.com/gin-gonic/gin"
	"github.com/utsavgupta/wireapp/repo"
)

type MovieHandlers struct {
	movieRepo repo.Movie
}

func NewMovieHandlers(ar repo.Movie) *MovieHandlers {
	return &MovieHandlers{movieRepo: ar}
}

func (ah *MovieHandlers) GetAllMovies(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ah.movieRepo.LoadItems())
}

func (ah *MovieHandlers) SaveMovie(ctx *gin.Context) {

	var movie ent.Movie

	err := ctx.ShouldBindJSON(&movie)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err = ah.movieRepo.SaveItem(&movie)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusCreated, "Movie saved")
}
