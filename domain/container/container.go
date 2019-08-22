package container

type Container struct {
	UID         string `json:"uid"`
	ImageName   string `json:"image_name"`
	ConnectInfo string `json:"connect_info"`
	Status      string `json:"status"`
}
