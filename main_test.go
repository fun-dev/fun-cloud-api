package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fun-dev/cloud-api/application/viewmodels"

	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/stretchr/testify/assert"
)

const (
	testJWT = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjhhYWQ2NmJkZWZjMWI0M2Q4ZGIyN2U2NWUyZTJlZjMwMTg3OWQzZTgiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNDMxMjk0OTY4MjgtcGh1OGQzOXIyZ2V0YWQ4OHEzbWZ0aWZpdTkwcjYzNG4uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0MzEyOTQ5NjgyOC1waHU4ZDM5cjJnZXRhZDg4cTNtZnRpZml1OTByNjM0bi5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwODg5ODkxMzIwNjgyODk4MTQxMCIsImVtYWlsIjoieW91dGFuYWdhaUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6ImxtMXhuLXJodmdYRlRKTFA4RnhUT1EiLCJuYW1lIjoieW90YSBuYWdhaSAo44KI44GG44Gh44KD44KTKSIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vLXd1TjZ2Q0VrNjJRL0FBQUFBQUFBQUFJL0FBQUFBQUFBTFZJL1Q2VG81WjZZMllBL3M5Ni1jL3Bob3RvLmpwZyIsImdpdmVuX25hbWUiOiJ5b3RhIiwiZmFtaWx5X25hbWUiOiJuYWdhaSIsImxvY2FsZSI6ImphIiwiaWF0IjoxNTQ3MjkxNjUwLCJleHAiOjE1NDcyOTUyNTAsImp0aSI6IjQwNGMwNzExZjdhY2Y2NWE2NTk1YTFkZmNmNmNiZTU1MWNkYzNkYjEifQ.rDbOORUj1z7bGQrHlZjhZ6bBbcAcpz6nAQ--zQz3BX5T8DMz4nVXuwcL56suaY6LU6QnZMtwVlLewFGvBsjH4-Vm7Xs4KGg5XhTXvN0a-PQ5kCKH83FK3nojka-fAq91i6XwKlEZj_-VvS3mAMHSifPsotvKtK9Yh010m8u27d5HE5iRHUaYJIU-8blc7JoYMs8xirJGQCfmbzH5reC1_Lj-zm8K16I4ToPxTqmt6E2E6PILZL8NOyqWVemStSkBCpCQIc9EsfjBzk0HWm3XlyDFk3W9WkLfLvi8ZpxvOD3FnQXqERM-NWLcPbOqqWrv4zBa_KA7HSO_URSLYa5tUQ"
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

func TestContainerPost(t *testing.T) {
	router := setupRouter()

	image := viewmodels.ContainerImage{
		ImageName: "nginx:latest",
	}

	data, err := json.Marshal(image)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/containers", bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", testJWT)

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}
