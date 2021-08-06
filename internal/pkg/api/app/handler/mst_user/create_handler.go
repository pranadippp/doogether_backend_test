package mst_user

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) CreateHandler(ctx *gin.Context) {
	var req request.UserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err, "failed to register user data")
		return
	}

	resp, err := h.uc.CreateUseCase(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err, "failed to register user data")
		return
	}

	response.NewSuccessResponse(ctx, http.StatusOK, resp)
	return
}
