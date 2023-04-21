package httpserver

import (
	"net/http"

	"github.com/00mohamad00/url-shortener/pkg/urlshortener"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	service urlshortener.Service
	router  *gin.Engine
}

func NewRouter(service urlshortener.Service) *gin.Engine {
	router := gin.Default()

	router.GET("/:shortURL", shortener.Redirect)
	router.POST("/shorten", shortener.CreateShortURL)

	return router
}

func (s *HttpServer) Redirect(c *gin.Context) {
	token := c.Param("token")
	targetURL, err := s.service.GetUrl(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}
	c.Redirect(http.StatusSeeOther, targetURL)
}

func (s *HttpServer) CreateShortURL(c *gin.Context) {
	var form struct {
		TargetURL string `form:"target_url" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := s.service.AddUrl(form.TargetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Short URL created", "token": token})
}
