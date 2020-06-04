package store_model

type User struct {
	BaseModel
	Name           string  `db:"name"`
	Roles          []*Role `gorm:"many2many:users_roles;"`
	IconURL        string  `db:"icon_url"`
}
