package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/SocialLogin/db"
)

func Search(c *gin.Context) {
	search := c.Request.URL.Query().Get("value")
	if search == "" {
		return
	}
	resp, err := db.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
