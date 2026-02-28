package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imwaddy/url-shortner/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", h.Shorten)
	r.GET("/:code", h.Redirect)
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	short, err := h.service.Create(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short": short})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	original, err := h.service.Resolve(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusFound, original)
}
