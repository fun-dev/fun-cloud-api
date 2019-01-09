package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fun-dev/cloud-api/infrastructure/dbmodels"
	"github.com/fun-dev/cloud-api/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// TokenAuthMiddleware is
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(400, gin.H{"message": "missing token"})
			c.Abort()
		}
		_, err := JWTValidate(token)
		if err != nil {
			c.JSON(401, gin.H{"message": "token expired"})
			c.Abort()
		}
		_, err = FindByToken(token)
		if err != nil {
			c.JSON(401, gin.H{"message": "token invailed"})
			c.Abort()
		}
		c.Next()
	}
}

func FindByToken(token string) (user dbmodels.User, err error) {
	engine, err := repositories.NewEngine()
	_, err = engine.Where("access_token = ?", token).Get(&user)
	if err != nil {
		return
	}
	return
}

func JWTValidate(token string) (claim Claim, err error) {
	endPoint := os.Getenv("GOOGLE_TOKEN_VALIDATE") + token
	response, err := http.Get(endPoint)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &claim)
	return
}

func JWTInvalid(token string) (result bool, err error) {

	buf := strings.Split(token, ".")

	header := buf[0]
	decodeHeader, err := base64.StdEncoding.DecodeString(header)

	jsonHeader := Header{}
	decodeHeader = []byte(string(decodeHeader) + "}")
	json.Unmarshal(decodeHeader, &jsonHeader)
	fmt.Println("ヘッダー")
	fmt.Println(jsonHeader)

	claim := buf[1]
	decodeClaim, err := base64.StdEncoding.DecodeString(claim)
	fmt.Println("個人情報")
	fmt.Println(string(decodeClaim))

	signature := buf[2]
	fmt.Println("署名")
	fmt.Println((signature))

	keys, _ := GetGoogleCredential()
	fmt.Println("Google Pubkey")
	fmt.Println(keys)
	for _, key := range keys.Keys {
		if jsonHeader.KID == key.KID {
			fmt.Println("鍵IDの一致")
			lenSignature := len(signature)
			lenGoogleSignature := len(key.N)
			if lenSignature == lenGoogleSignature {
				fmt.Println("バイト数の一致")
			}
		}
	}
	return
}

func GetGoogleCredential() (keys Keys, err error) {
	url := "https://www.googleapis.com/oauth2/v3/certs"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &keys)
	return
}

type Key struct {
	KID string `json: "kid"`
	E   string `json: "e"`
	KTY string `json: "kty"`
	Alg string `json: "alg"`
	N   string `json: "n"`
	USE string `json: "use"`
}

type Keys struct {
	Keys []Key `json: "keys"`
}

type Header struct {
	Alg string `json: "alg"`
	KID string `json: "kid"`
	Typ string `json: "typ"`
}

type Claim struct {
	Iss           string `json: "iss"`
	Azp           string `json: "azp"`
	Aud           string `json: "aud"`
	Sub           string `json: "sub"`
	Email         string `json: "email"`
	EmailVerified bool   `json: "email_verified"`
	AtHash        string `json: "at_hash"`
	Name          string `json: "name"`
	Picture       string `json: "picture"`
	GivenName     string `json: "given_name"`
	FamilyName    string `json: "family_name"`
	Locale        string `json: "locale"`
	Iat           string `json: "iat"`
	Exp           string `json: "exp"`
	Jti           string `json: "jti"`
	Alg           string `json: "alg"`
	Kid           string `json: "kid"`
	Typ           string `json: "typ"`
}
