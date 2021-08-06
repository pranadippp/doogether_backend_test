package session

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *sessionHandler) UpdateHandler(ctx *gin.Context) {
	var req request.SessionRequest

	// claimId, _ := ctx.Get("userInfo")
	claimemail := ctx.GetString("userEmail")
	// req.UserID = claimId.(int)
	req.Email = claimemail

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to update session data")
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to update session data")
		return
	}

	resp, err := h.uc.UpdateUseCase(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to update session data")
		return
	}

	response.NewSuccessResponse(ctx, http.StatusOK, resp)
	return
}
