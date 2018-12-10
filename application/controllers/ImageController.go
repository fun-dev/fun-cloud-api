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
	// TODO: Remove Hard-Cord URL
	url := "https://registry.cloud.fun.aigis.pw/"
	username := "fun-ict-cloud"
	password := "BiWsPSmo7V6v4I"
	hub, err := registry.New(url, username, password)
	if err != nil {
		log.Fatal(err)
		return
	}
	repositories, err := hub.Repositories()
	if err != nil {
		log.Fatal(err)
		return
	}
	imageData := []viewmodels.Image{}
	for _, repositry := range repositories {
		manifest, err := hub.Manifest(repositry, "latest")
		if err != nil {
			log.Fatal(err)
			return
		}
		repositoryName := manifest.Name
		for _, history := range manifest.History {
			v1CompatibilityJson := strings.Replace(history.V1Compatibility, "\n", "", -1)
			var v1CompatibilityObject v1Compatibility
			err := json.Unmarshal([]byte(v1CompatibilityJson), &v1CompatibilityObject)
			if err != nil {
				log.Fatal(err)
				return
			}
			repositryDescription := v1CompatibilityObject.Container_Config.Labels.Description
			if repositryDescription != "" {
				repositoryImageData := viewmodels.Image{Name: repositoryName, Description: repositryDescription}
				imageData = append(imageData, repositoryImageData)
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
