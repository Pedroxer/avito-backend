package api

import (
	"github.com/Pedroxer/avito/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type addUserToGroupeRequest struct {
	User_id int `json:"user_id" binding:"required"`
	Seg_id  int `json:"seg_id" binding:"required"`
}

func (serv *Server) addSegToUser(ctx *gin.Context) {
	var req addUserToGroupeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.AddSegToUserParams{

		int32(req.User_id),

		int32(req.Seg_id),
	}
	segToUsr, err := serv.db.AddSegToUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, segToUsr)
}

type getUserSegmentsRequest struct {
	User_id int `json:"user_id" binding:"required"`
}
type getUserSegmentsResponse struct {
	User_id  int      `json:"user_id"`
	Segments []string `json:"segments"`
}

func (serv *Server) getUserSegments(ctx *gin.Context) {
	var req getUserSegmentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	userSegments, err := serv.db.GetSegOfUser(ctx, int32(req.User_id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := getUserSegmentsResponse{
		User_id: req.User_id,
	}
	for i := 0; i < len(userSegments); i++ {
		seg, err := serv.db.GetSeg(ctx, int64(userSegments[i].SegID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		resp.Segments = append(resp.Segments, seg.Name)
	}
	ctx.JSON(http.StatusOK, resp)
}

type removeSegFromUserRequest struct {
	User_id int `json:"user_id" binding:"required"`
	Seg_id  int `json:"seg_id" binding:"required"`
}

func (serv *Server) removeSegFromUser(ctx *gin.Context) {
	var req removeSegFromUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := sqlc.RemoveSegFromUserParams{

		int32(req.User_id),

		int32(req.Seg_id),
	}
	err := serv.db.RemoveSegFromUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Removed")
}
