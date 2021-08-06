package session

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *sessionHandler) FindDetailHandler(ctx *gin.Context) {
	var req request.SessionRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to find session detail data")
		return
	}

	resp, err := h.uc.FindDetailUseCase(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to find session detail data")
		return
	}

	response.NewSuccessResponse(ctx, http.StatusOK, resp)
	return
}
