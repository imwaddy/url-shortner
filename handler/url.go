package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imwaddy/url-shortner/model"
	"github.com/imwaddy/url-shortner/service"
)

type URLHandler struct {
	service *service.URLService
	baseURL string
}

func NewURLHandler(s *service.URLService, baseURL string) *URLHandler {
	return &URLHandler{
		service: s,
		baseURL: baseURL,
	}
}

func (h *URLHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/v1/shorten", h.Shorten)
	r.GET("/api/v1/:code", h.Redirect)
	r.GET("/health", h.Health)
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req model.ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   "invalid_request",
			Message: "Request body is invalid or missing required fields",
		})
		return
	}

	// Validate URL
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   "invalid_url",
			Message: err.Error(),
		})
		return
	}

	shortCode, err := h.service.Create(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error:   "server_error",
			Message: "Failed to create short URL",
		})
		return
	}

	response := model.ShortenResponse{
		ShortCode: shortCode,
		ShortURL:  fmt.Sprintf("%s/%s", h.baseURL, shortCode),
		Original:  req.URL,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *URLHandler) Redirect(c *gin.Context) {
	ctx := c.Request.Context()
	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   "invalid_code",
			Message: "Short code is required",
		})
		return
	}

	original, err := h.service.Resolve(ctx, code)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   "not_found",
			Message: "Short URL not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, original)
}

func (h *URLHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "url-shortener",
	})
}
