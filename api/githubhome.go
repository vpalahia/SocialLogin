package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/SocialLogin/db"
	"github.com/vpalahia/SocialLogin/types"
)

const (
	githubclient_id     = "69115e30be634bf2b4d1"
	githubclinet_secret = "11ac96143a2fff8a5e696256676235a3eef82211"
	linkedclient_id     = "78fsjtrjq6qlv5"
	linkedclient_secret = "ueEGQESPo8g9i4k8"
)

func Githubhome(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	url := "https://github.com/login/oauth/access_token?client_id=" + githubclient_id + "&client_secret=" + githubclinet_secret + "&code=" + code
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	s := string(body)
	tokenarr := strings.Split(s, "\u0026")
	token := strings.Split(tokenarr[0], "=")
	fmt.Println(token[1])
	url = "https://api.github.com/user"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "token "+token[1])

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response types.ThirdPartyResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	response.MetaData.Github = true
	err = db.InsertIntoDB(response)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, "Hello "+response.Name)
}
