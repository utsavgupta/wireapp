package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utsavgupta/wireapp/handlers"
	"github.com/utsavgupta/wireapp/repo"
)

const (
	port = 8080
)

func main() {
	r := gin.Default()

	albumRepo := repo.NewAlbum(10)
	movieRepo := repo.NewMovie(100)

	albumHandler := handlers.NewAlbumHandlers(albumRepo)
	movieHandler := handlers.NewMovieHandlers(movieRepo)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Resource urls are [GET, POST] /albums and /movies")
	})

	// Album handlers
	r.GET("/albums", albumHandler.GetAllAlbums)
	r.POST("/albums", albumHandler.SaveAlbum)

	// Movie handlers
	r.GET("/movies", movieHandler.GetAllMovies)
	r.POST("/movies", movieHandler.SaveMovie)

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
