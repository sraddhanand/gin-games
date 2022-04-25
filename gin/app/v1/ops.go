package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sraddhanand/profitGames/utils/app"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, "up")
}

func GetEchoREST(c *gin.Context) {
	appG := app.Gin{C: c}
	msg := c.Param("msg")
	url := fmt.Sprintf("http://%s:%s/echo/%s", os.Getenv("ECHO_HOST"), os.Getenv("ECHO_PORT"), msg)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var usr_resp struct {
		StatusCode string `json:"code"`
		Message    string `json:"msg"`
		Data       string `json:"data"`
	}

	json.Unmarshal(body, &usr_resp)
	appG.Response(http.StatusOK, usr_resp.Data)
}
