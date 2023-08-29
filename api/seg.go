package api

import (
	"github.com/Pedroxer/avito/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createSegRequest struct {
	Name string `json:"seg_name" binding:"required"`
}

func (serv *Server) createSeg(ctx *gin.Context) {
	var req createSegRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	seg, err := serv.db.CreateSeg(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, seg)

}

type updateSegRequest struct {
	Id      int    `json:"seg_id" binding:"required,alphanum"`
	NewName string `json:"new_name" binding:"required,alphanum"`
}

func (serv *Server) updateSeg(ctx *gin.Context) {
	var req updateSegRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.UpdateSegParams{
		ID:   int64(req.Id),
		Name: req.NewName,
	}
	err := serv.db.UpdateSeg(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	seg, err := serv.db.GetSeg(ctx, arg.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, seg)
}

type deleteSegRequest struct {
	id int `json:"seg_id" binding:"required"`
}

func (serv *Server) deleteSeg(ctx *gin.Context) {
	var req deleteSegRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := serv.db.DeleteSeg(ctx, int64(req.id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
