package jwt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Validate
// Confirm
// Inspect
type (
	IJwt interface {
		InspectGoogleIdToken(accessToken string) (*Claim, error)
	}
	Jwt struct{}
)

func NewJwt() IJwt {
	return &Jwt{}
}

func (j *Jwt) InspectGoogleIdToken(accessToken string) (*Claim, error) {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/tokeninfo", nil)
	q := req.URL.Query()
	q.Add("id_token", accessToken)
	req.URL.RawQuery = q.Encode()
	//fmt.Printf("debug: %v", req.RequestURI)
	resp, err := http.Get(req.URL.String())
	if err != nil {
		//TODO: implements error handling
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &Claim{}
	if err := json.Unmarshal(body, &result); err != nil {
		//TODO: implements error handling
		return nil, err
	}
	return result, nil
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
