package session

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *sessionHandler) FindAllHandler(ctx *gin.Context) {
	var req request.SessionRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to get sesion data")
		return
	}

	resp, err := h.uc.FindAllUseCase(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to get session data")
		return
	}

	response.NewPagedResponse(ctx, resp)
	return
}
