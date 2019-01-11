package viewmodels

type User struct {
	IconUrl     string `json:"icon_url"`
	GoogleName  string `json:"google_name"`
	AccessToken string `json:"access_token"`
	Email       string `json:"-"`
}
