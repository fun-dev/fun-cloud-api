package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/fun-dev/cloud-api/application/viewmodels"
	"github.com/gin-gonic/gin"
	"github.com/heroku/docker-registry-client/registry"
)

type ImageController struct {
}

func (ctrl ImageController) Get(c *gin.Context) {
	// FIXME: Remove Hard-Cord URL
	url := "https://registry.cloud.fun.aigis.pw/"     //Docker-registryのURL
	username := "fun-ict-cloud"                       //Docker-registryのユーザ名
	password := "BiWsPSmo7V6v4I"                      //Docker-registryのパスワード
	hub, err := registry.New(url, username, password) //Docker-registryのクライアントとなるオブジェクト
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println(err)
		return
	}
	repositories, err := hub.Repositories() //Docker-registryにあるイメージの一覧を取得
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		log.Println(err)
		return
	}
	imageData := []viewmodels.Image{} //Docker-registry APIから取得したイメージを、クライアントが処理しやすい形式に変換した後の配列。これを返す。今は空
	for _, repositry := range repositories {
		manifest, err := hub.Manifest(repositry, "latest")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			log.Println(err)
			break
		}
		repositoryName := manifest.Name // Docker-registry APIから取得したJSONからImage名を取得
		for _, history := range manifest.History {
			//Docker-registry APIから取得したJSONからV1Compatibilityを抽出。ここにImageのLabelが入っている
			v1CompatibilityJson := strings.Replace(history.V1Compatibility, "\n", "", -1)
			var v1CompatibilityObject v1Compatibility
			err := json.Unmarshal([]byte(v1CompatibilityJson), &v1CompatibilityObject)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
				log.Println(err)
				break
			}
			repositryDescription := v1CompatibilityObject.Container_Config.Labels.Description //V1CompatibilityからImageのLabelを取得
			//V1Compatibilityの中にLabelが無い場合もあるので、その時はスキップ
			if repositryDescription != "" {
				repositoryImageData := viewmodels.Image{Name: repositoryName, Description: repositryDescription}
				imageData = append(imageData, repositoryImageData) //Image名とImageのLabelをviewmodels.Image型の配列にPush
				break
			}
		}
	}

	c.JSON(http.StatusOK, imageData)
}

type v1Compatibility struct {
	Container_Config struct {
		Labels struct {
			Description string `json:description`
		} `json:Labels`
	} `json:"container_config"`
}
