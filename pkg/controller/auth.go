package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Respuesta struct{
	Access_token string
	Token_type string
	Expires_in int
	Scope string
	User_id int
	Refresh_token string

}

func Auth(c *gin.Context){
	code := c.Query("code")
	client := http.Client{}
 	requestBody, _ := json.Marshal(map[string]string{
		"grant_type": "authorization_code",
		"client_id" : "3589105139307129",
		"client_secret": "knzqc1moyoKYFMlnHidSm7NlRMl16Dtj",
		"code" : code,
		"redirect_uri": "http://localhost:8080/auth",
	})
	request, _ := http.NewRequest("POST", "https://api.mercadolibre.com/oauth/token",bytes.NewBuffer(requestBody))
	request.Header.Set("accept", "application/json")
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	resp, _ := client.Do( request )
	defer resp.Body.Close()
	var res Respuesta
	data, _ := ioutil.ReadAll(resp.Body)
	er := json.Unmarshal(data, &res)
	if er != nil {
		println("There has been an error with the Unmarshal of the json")
		fmt.Println(er)
	}
	c.JSON(200, res)
}
