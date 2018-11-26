package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
}

func (ctrl ImageController) Get(c *gin.Context) {
	// TODO: Remove Hard-Cord URL
	req, _ := http.NewRequest("GET", "https://registry.cloud.fun.aigis.pw/v2/_catalog", nil)
	req.Header.Add("authorization", "Basic ZnVuLWljdC1jbG91ZDpCaVdzUFNtbzdWNnY0SQ==")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	c.JSON(http.StatusOK, string(body))
}
