package session

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *sessionHandler) CreateHandler(ctx *gin.Context) {
	var req request.SessionRequest

	// claimid := ctx.GetInt("userId")
	claimname := ctx.GetString("userName")
	claimemail := ctx.GetString("userEmail")

	// req.UserID = claimid
	req.Name = claimname
	req.Email = claimemail

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to register session data")
		return
	}

	resp, err := h.uc.CreateUseCase(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to register session data")
		return
	}

	response.NewSuccessResponse(ctx, http.StatusOK, resp)
	return
}
