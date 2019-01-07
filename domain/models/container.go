package models

type Container struct {
	Id          int    `json:id`
	ImageId     int    `json:image_id`
	ConnectInfo string `json:connect_info`
	Status      string `json:status`
}
