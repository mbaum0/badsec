package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	controller *SecretController
	Router     *gin.Engine
}

type Response struct {
	Content string `json:"content"`
	Error   string `json:"error"`
}

func CreateServer() *Server {
	server := Server{controller: newSecretController(time.Second * 60)}
	server.initRouter()

	return &server
}

func (server *Server) Run() {
	server.Router.Run(":8080")
}

func (server *Server) initRouter() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(func(c *gin.Context) {
		c.Set("real_ip", c.Request.Header.Get("X-Real-IP"))
		c.Next()
	})
	router.Use(gin.Recovery())

	// trustedProxies := []string{"127.0.0.1"}
	// router.SetTrustedProxies(trustedProxies)

	router.LoadHTMLFiles("./templates/index.html")
	router.Static("/static", "./static")

	router.POST("/api/secret", server.GetSecretByKey)
	router.PUT("/api/secret", server.UpdateSecret)
	router.DELETE("/api/secret", server.DeleteSecretByKey)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	server.Router = router
}

func (server *Server) GetSecretByKey(c *gin.Context) {
	var reqSecret Secret
	if err := c.ShouldBindJSON(&reqSecret); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	secret := server.controller.GetSecret(reqSecret.Key)
	c.JSON(http.StatusOK, Response{Content: secret})
}

func (server *Server) UpdateSecret(c *gin.Context) {
	var reqSecret Secret
	if err := c.ShouldBindJSON(&reqSecret); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}
	secret := server.controller.UpdateSecret(reqSecret.Key, reqSecret.Content)
	c.JSON(http.StatusOK, Response{Content: secret})
}

func (server *Server) DeleteSecretByKey(c *gin.Context) {
	var reqSecret Secret
	if err := c.ShouldBindJSON(&reqSecret); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}
	secret := server.controller.DeleteSecret(reqSecret.Key)
	c.JSON(http.StatusOK, Response{Content: secret})
}
