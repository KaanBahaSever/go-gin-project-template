package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (hc *HomeController) Index(c *gin.Context) {
	userID := c.MustGet("userID")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"userID": userID,
	})
}
