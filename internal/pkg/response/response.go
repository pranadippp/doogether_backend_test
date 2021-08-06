package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type MetaData struct {
// 	Meta ErrorResponse `json:"meta"`
// 	Data interface{}   `json:"data"`
// }

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Ok     bool        `json:"ok"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

type PagedResponse struct {
	Ok     bool        `json:"ok"`
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Size   int         `json:"size"`
	Page   int         `json:"page"`
	Total  int         `json:"total"`
}

type PhotoResponse struct {
	PathPhoto string `json:"path_photo"`
}

func NewErrorResponse(c *gin.Context, code int, err error, message string) {

	resp := ErrorResponse{
		Ok:      false,
		Status:  "error",
		Code:    code,
		Error:   err.Error(),
		Message: message,
	}

	// sendData.Meta = responses

	c.JSON(code, resp)
}

func NewPagedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func NewSuccessResponse(c *gin.Context, code int, data interface{}) {

	resp := SuccessResponse{
		Ok:     true,
		Status: "success",
		Code:   code,
		Data:   data,
	}

	// sendData.Meta = responses
	// sendData.Data = data

	c.JSON(code, resp)
}
