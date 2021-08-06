package mst_user

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) LoginHandler(ctx *gin.Context) {
	var req request.UserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "Failed to login")
		return
	}

	resp, err := h.uc.Login(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to login")
		return
	}

	response.NewSuccessResponse(ctx, http.StatusOK, resp)
	return
}
