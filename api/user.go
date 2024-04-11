package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/pe-Gomes/short-url/infra/db/repository"
	"github.com/pe-Gomes/short-url/util"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateUserParams{
		Name:     req.Name,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newUserResponse(user))
}

func (server *Server) getUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := server.store.GetUserByEmail(ctx, email)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
}

type listUsersRequest struct {
	Page     int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=50"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no users found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []userResponse

	for _, user := range users {
		response = append(response, newUserResponse(user))
	}

	ctx.JSON(http.StatusOK, response)
}

type deleteUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteUserById(ctx *gin.Context) {
	var req deleteUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return

	}

	_, err := server.store.GetUser(ctx, req.ID)

	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = server.store.DeleteUser(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
