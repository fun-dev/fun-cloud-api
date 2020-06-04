package store_model

type Scope struct {
	BaseModel
	Name  string
	Roles []*Role `gorm:"many2many:scopes_roles;"`
}

type Role struct {
	BaseModel
	Name   string
	Scopes []*Scope `gorm:"many2many:scopes_roles;"`
	Users  []*User  `gorm:"many2many:users_roles;"`
}
