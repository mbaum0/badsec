package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var sm = NewSecretManager(60 * time.Second)

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.Use(func(c *gin.Context) {
		c.Set("real_ip", c.Request.Header.Get("X-Real-IP"))
		c.Next()
	})
	router.Use(gin.Recovery())

	// trustedProxies := []string{"127.0.0.1"}
	// router.SetTrustedProxies(trustedProxies)

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.POST("/api/secret", GetSecretByKey)
	router.PUT("/api/secret", UpdateSecret)
	router.DELETE("/api/secret", DeleteSecretByKey)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run()
}

func GetSecretByKey(c *gin.Context) {
	var reqSecret Secret
	if err := c.ShouldBindJSON(&reqSecret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secret := sm.GetSecret(reqSecret.Key)
	c.JSON(http.StatusOK, gin.H{"content": secret})
}

func UpdateSecret(c *gin.Context) {
	var reqSecret Secret
	if err := c.ShouldBindJSON(&reqSecret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secret := sm.UpdateSecret(reqSecret.Key, reqSecret.Content)
	c.JSON(http.StatusOK, gin.H{"content": secret})
}

func DeleteSecretByKey(c *gin.Context) {
	var reqSecret Secret
	if err := c.ShouldBindJSON(&reqSecret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secret := sm.DeleteSecret(reqSecret.Key)
	c.JSON(http.StatusOK, gin.H{"content": secret})
}
