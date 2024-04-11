package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/pe-Gomes/short-url/infra/db/repository"
)

type getShortURLResponse struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

func createShortURLRsponse(shortLink db.ShortLink) getShortURLResponse {
	return getShortURLResponse{
		URL:  shortLink.Url,
		Slug: shortLink.Slug,
	}
}

type createShortURLRequest struct {
	UserId int64  `json:"user_id" binding:"required"`
	URL    string `json:"url" binding:"required,url"`
	Slug   string `json:"slug" binding:"required"`
}

func (server *Server) createShortURL(ctx *gin.Context) {
	var req createShortURLRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateShortURLParams{
		UserID: req.UserId,
		Url:    req.URL,
		Slug:   req.Slug,
	}

	shortURL, err := server.store.CreateShortURL(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createShortURLRsponse(shortURL))
}

type getShortURLBySlugRequest struct {
	Slug string `uri:"slug" binding:"required"`
}

func (server *Server) getShortURLBySlug(ctx *gin.Context) {
	var req getShortURLBySlugRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	shortLink, err := server.store.GetShortURLBySlug(ctx, req.Slug)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.String(http.StatusNotFound, "Not Found")
			return

		}
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, shortLink.Url)
}

type listURLByUserRequest struct {
	UserID   int64 `form:"user_id" binding:"required"`
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
}

func (server *Server) listURLByUser(ctx *gin.Context) {
	var req listURLByUserRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListShortURLsParams{
		UserID: req.UserID,
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	}

	shortLinks, err := server.store.ListShortURLs(ctx, arg)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []getShortURLResponse

	for _, shortLink := range shortLinks {
		response = append(response, createShortURLRsponse(shortLink))
	}

	ctx.JSON(http.StatusOK, response)
}
