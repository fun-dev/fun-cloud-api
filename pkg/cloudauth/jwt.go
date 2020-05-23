package cloudauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	_googleOAuthTokeninfoURL       = "https://www.googleapis.com/oauth2/v3/tokeninfo"
	ErrRetrieveTokeninfoFromGoogle = errors.New("failed retrieve tokeninfo from google token")
	ErrUnmarshalGoogleResponseBody = errors.New("failed unmarshal google response body")
)

func InspectGoogleIdToken(token string) (*Claim, error) {
	req, _ := http.NewRequest("GET", _googleOAuthTokeninfoURL, nil)
	q := req.URL.Query()
	q.Add("id_token", token)
	req.URL.RawQuery = q.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, ErrRetrieveTokeninfoFromGoogle
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := Claim{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, ErrUnmarshalGoogleResponseBody
	}
	return &result, nil
}

type Claim struct {
	Issuer          string `json:"iss"`
	AuthorizedParty string `json:"azp"`
	Audience        string `json:"aud"`
	Subject         string `json:"sub"`
	Email           string `json:"email"`
	EmailVerified   string `json:"email_verified"`
	AtHash          string `json:"at_hash"`
	Name            string `json:"name"`
	Picture         string `json:"picture"`
	GivenName       string `json:"given_name"`
	FamilyName      string `json:"family_name"`
	Locale          string `json:"locale"`
	IssuedAt        string `json:"iat"`
	ExpirationTime  string `json:"exp"`
	JWTID           string `json:"jti"`
	Algorithm       string `json:"alg"`
	KeyID           string `json:"kid"`
	Type            string `json:"typ"`
}
