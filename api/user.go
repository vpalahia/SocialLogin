package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/SocialLogin/db"
)

func User(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	resp, err := db.User(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, resp)
}
