package handlers

import (
	"net/http"

	"github.com/utsavgupta/wireapp/ent"

	"github.com/gin-gonic/gin"
	"github.com/utsavgupta/wireapp/repo"
)

type AlbumHandlers struct {
	albumRepo repo.Album
}

func NewAlbumHandlers(ar repo.Album) *AlbumHandlers {
	return &AlbumHandlers{albumRepo: ar}
}

func (ah *AlbumHandlers) GetAllAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ah.albumRepo.LoadItems())
}

func (ah *AlbumHandlers) SaveAlbum(ctx *gin.Context) {

	var album ent.Album

	err := ctx.ShouldBindJSON(&album)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err = ah.albumRepo.SaveItem(&album)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusCreated, "Album saved")
}
