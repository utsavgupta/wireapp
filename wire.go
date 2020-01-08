//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/utsavgupta/wireapp/handlers"
	"github.com/utsavgupta/wireapp/repo"
)

func InitializeMovieHandler(a repo.MovieBuffer) *handlers.MovieHandlers {
	wire.Build(handlers.NewMovieHandlers, repo.NewMovie)

	return &handlers.MovieHandlers{}
}

func InitializeAlbumHandler(a repo.AlbumBuffer) *handlers.AlbumHandlers {
	wire.Build(handlers.NewAlbumHandlers, repo.NewAlbum)

	return &handlers.AlbumHandlers{}
}

func IntializeEngine(a repo.MovieBuffer, b repo.AlbumBuffer) *gin.Engine {
	wire.Build(NewEngine, InitializeMovieHandler, InitializeAlbumHandler)

	return &gin.Engine{}
}
