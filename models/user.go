package models

import "gorm.io/datatypes"

type User struct {
	ID        uint           `gorm:"size:50"`
	Name      string         `gorm:"size:50"`
	Fullname  string         `gorm:"size:50"`
	FirstName string         `gorm:"size:50"`
	LastName  string         `gorm:"size:50"`
	Email     string         `gorm:"size:50;unique"`
	Password  string         `gorm:"size:100" json:"password"`
	Phone     string         `gorm:"size:20"`
	BirthDate datatypes.Date `gorm:"default:NULL"`
	Gender    string         `gorm:"size:20"`
	Photo     string
	Status    string `gorm:"default: 1"`
	CreatedBy int
	UpdatedBy int
	Roles     []Role     `gorm:"many2many:user_roles; constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	UserRoles []UserRole `gorm:"foreignkey:UserId"`
}

type UserRole struct { // TODO: Untuk Pivot Table
	User   User
	UserId int
	Role   User
	RoleId int
}

func (UserRole) TableName() string {
	return "user_roles"
}
