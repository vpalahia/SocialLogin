package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/SocialLogin/db"
)

func Allusers(c *gin.Context) {
	resp, err := db.Allusers()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
