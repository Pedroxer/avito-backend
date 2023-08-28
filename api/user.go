package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRequest struct {
	Name string `json:"user_name" binding:"required,alphanum"` // TODO: Сделать валадицию только на цифры
}

func (serv *Server) createUser(ctx *gin.Context) {
	var req UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := serv.db.CreateUser(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (serv *Server) deleteUser(ctx *gin.Context) {
	var req UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := serv.db.DeleteUser(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Delete")
}
