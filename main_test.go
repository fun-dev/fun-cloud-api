package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/stretchr/testify/assert"
)

const (
	testJWT = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjhhYWQ2NmJkZWZjMWI0M2Q4ZGIyN2U2NWUyZTJlZjMwMTg3OWQzZTgiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNDMxMjk0OTY4MjgtcGh1OGQzOXIyZ2V0YWQ4OHEzbWZ0aWZpdTkwcjYzNG4uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0MzEyOTQ5NjgyOC1waHU4ZDM5cjJnZXRhZDg4cTNtZnRpZml1OTByNjM0bi5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwODg5ODkxMzIwNjgyODk4MTQxMCIsImVtYWlsIjoieW91dGFuYWdhaUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6IkRjU1JMM0VXYy1hXzQtaWdjclpMa0EiLCJuYW1lIjoieW90YSBuYWdhaSAo44KI44GG44Gh44KD44KTKSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vLXd1TjZ2Q0VrNjJRL0FBQUFBQUFBQUFJL0FBQUFBQUFBTFZJL1Q2VG81WjZZMllBL3M5Ni1jL3Bob3RvLmpwZyIsImdpdmVuX25hbWUiOiJ5b3RhIiwiZmFtaWx5X25hbWUiOiJuYWdhaSIsImxvY2FsZSI6ImphIiwiaWF0IjoxNTQ3MjgyNTk0LCJleHAiOjE1NDcyODYxOTQsImp0aSI6IjZkMTQ0ZWNiYjQ5YTlmNTgxY2QxZDgwNzFhODBiNGQ2MDhjOTA5MzQifQ.d88wuQRtkGtCs_BjiJ8X-gwHW-gXUIHAYn4cSOyEJM4HAL5ubqrXE-uI0jSqglIuNGEt_qwe-0eCEhRJvrqX59z4LjwAaXhUGg9KlB98rcFEicPyirZ7qYvrGMC-1xBESSD9qCVph1Em4r3Q9l_g_apqMXF2xkAG2_oCasgGEeLzNFTGy8Pv2J_t-WxtSg6W7uWSEiGXUjQRV5GLrpL8uwAWAdLVOMZpeHENDVho5bF5ppWns-WhKw8czFT2rHIH5nCewyg5Wy219vuf_qzr4Wh3BpGF2urRgLBEsCBdCLtZHWS2JUo2E9UKIs0jABF7rjGUT_XvR4Gz7lyRzccs5A"
)

func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestContainerGet(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/v1/containers", nil)

	req.Header.Set("Authorization", testJWT)

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var containers []models.Container

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	json.Unmarshal(data, &containers)
	fmt.Println(containers)
}
