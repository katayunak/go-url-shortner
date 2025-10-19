package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"urlShortner/service"
)

type handler struct {
	service service.Service
}

type Handler interface {
	Health(ctx *gin.Context)
	Redirect(ctx *gin.Context)
	Create(ctx *gin.Context)
}

func NewHandler(service service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "healthy :)"})
}

func (h *handler) Redirect(ctx *gin.Context) {
	short := ctx.Param("short")
	if short == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid URL"})
	}

	resp, err := h.service.FindURL(ctx, service.FindRequest{Short: short})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	ctx.Redirect(http.StatusFound, resp.Long)
}

func (h *handler) Create(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	request := CreateRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	response, err := h.service.Create(ctx, service.CreateRequest{Long: request.Long})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"short": response.Short})
}

// remember:
// handler is responsible for validation of request not controller
// when you have handler and controller together, controller should not know anything about HTTP
