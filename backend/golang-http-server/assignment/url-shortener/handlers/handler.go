package handlers

import (
	"errors"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	if h.repo.Data[c.Param("path")] == "" {
		c.String(http.StatusNotFound, "Not found")
	}
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	url := entity.URL{}

	if url.LongURL == "" {
		resp, err := h.repo.Create(url.LongURL)
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
		}
		c.JSON(http.StatusOK, resp)
		
	}
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	
}
