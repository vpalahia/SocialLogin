package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/SocialLogin/db"
	"github.com/vpalahia/SocialLogin/types"
)

func Linkedinhome(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")

	url := "https://www.linkedin.com/oauth/v2/accessToken?client_id=" + linkedclient_id + "&client_secret=" + linkedclient_secret + "&code=" + code + "&grant_type=authorization_code&redirect_uri=http://localhost:8084/linkedinhome"

	resp, err := http.Post(url, "", nil)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var res types.LinkedinResp
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}

	url = "https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+res.AccessToken)

	re, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer re.Body.Close()
	body, err = ioutil.ReadAll(re.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	var response types.LinkedUser
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}

	var userdata types.ThirdPartyResponse
	userdata.MetaData.Linkedin = true
	if len(response.Elements) > 0 {
		userdata.Email = response.Elements[0].Handle.Email
	}

	err = db.InsertIntoDB(userdata)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, "Hello "+userdata.Email)

}
