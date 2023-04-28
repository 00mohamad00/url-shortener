package httpserver

import (
	"net/http"

	"github.com/00mohamad00/url-shortener/pkg/urlshortener"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Addr string
}

type HttpServer interface {
	Run(addr string) error
}

type Impl struct {
	router  *gin.Engine
	service urlshortener.Service
}

func NewRouter(service urlshortener.Service) HttpServer {
	server := &Impl{
		router:  gin.Default(),
		service: service,
	}

	server.router.GET("/:token", server.Redirect)
	server.router.POST("/shorten", server.CreateShortURL)

	return server
}

func (i *Impl) Run(addr string) error {
	return i.router.Run(addr)
}

func (i *Impl) Redirect(c *gin.Context) {
	token := c.Param("token")
	targetURL, err := i.service.GetUrl(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}
	c.Redirect(http.StatusSeeOther, targetURL)
}

func (i *Impl) CreateShortURL(c *gin.Context) {
	var form struct {
		TargetURL string `form:"target_url" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := i.service.AddUrl(form.TargetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Short URL created", "token": token})
}
