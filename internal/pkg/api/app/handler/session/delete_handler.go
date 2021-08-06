package session

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *sessionHandler) DeleteHandler(ctx *gin.Context) {
	var req request.SessionRequest

	claimemail := ctx.GetString("userEmail")

	req.Email = claimemail

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to delete session data")
		return
	}

	resp, err := h.uc.DeleteUseCase(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to delete session data")
		return
	}

	response.NewSuccessResponse(ctx, http.StatusOK, resp)
	return
}
