package dbmodels

type Container struct {
	ID          int    `json:id`
	ImageID     int    `json:image_id`
	ConnectInfo string `json:connect_info`
	Status      string `json:status`
}
