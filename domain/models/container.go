package models

type Container struct {
	UID         string `json:id`
	ImageName   string `json:image_id`
	ConnectInfo string `json:connect_info`
	Status      string `json:status`
}
