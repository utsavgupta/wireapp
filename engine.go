package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/utsavgupta/wireapp/handlers"
)

func NewEngine(albumHandlers *handlers.AlbumHandlers, movieHandlers *handlers.MovieHandlers) *gin.Engine {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Resource urls are [GET, POST] /albums and /movies")
	})

	// Album handlers
	r.GET("/albums", albumHandlers.GetAllAlbums)
	r.POST("/albums", albumHandlers.SaveAlbum)

	// Movie handlers
	r.GET("/movies", movieHandlers.GetAllMovies)
	r.POST("/movies", movieHandlers.SaveMovie)

	return r
}
