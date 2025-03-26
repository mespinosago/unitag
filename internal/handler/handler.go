package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mespinosago/unitag/internal/model"
	"github.com/mespinosago/unitag/internal/parser"
	"github.com/mespinosago/unitag/internal/service"
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --with-expecter --exported --name Service
type Service interface {
	GetURL(code string, options model.Options) (string, error)
}

type Handler struct {
	service Service
	parser  *parser.Parser
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
		parser:  parser.NewParser(),
	}
}

func (h *Handler) GetURL(c *gin.Context) {
	// Retrieve Code from URL parameter
	code := c.Param("code")
	if !isValidFormat(code) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid code"})
		return
	}
	// Find out options from header
	options := model.Options{
		Language: h.parser.GetLanguage(c.GetHeader("Accept-Language")),
		OS:       h.parser.GetOS(c.GetHeader("User-Agent")),
		Browser:  h.parser.GetBrowser(c.GetHeader("User-Agent")),
	}

	url, err := h.service.GetURL(code, options)
	if err != nil {
		if errors.Is(err, service.ErrCodeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, url)
}

func isValidFormat(code string) bool {
	return len(code) == 6
}
