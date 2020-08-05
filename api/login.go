package api

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {

	c.HTML(200, "login.html", gin.H{})
}
