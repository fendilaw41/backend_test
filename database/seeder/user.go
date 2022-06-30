package seeder

import (
	"backend_test/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

// TODO: Untuk Dummy User
func (s Seed) UserSeed() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

	user := []models.User{
		{
			Name:      "Fendi Ray",
			FirstName: "Fendi",
			LastName:  "Ray",
			Email:     "fendi@gmail.com",
			Password:  string(bytes),
			CreatedBy: 1,
			BirthDate: datatypes.Date(time.Now()),
		},
		{
			Name:      "Admin",
			FirstName: "Admin",
			LastName:  "istrator",
			Email:     "admin@gmail.com",
			Password:  string(bytes),
			CreatedBy: 1,
			BirthDate: datatypes.Date(time.Now()),
		},
		{
			Name:      "Customer",
			FirstName: "Customer",
			LastName:  " ",
			Email:     "customer@gmail.com",
			Password:  string(bytes),
			CreatedBy: 1,
			BirthDate: datatypes.Date(time.Now()),
		},
	}

	s.db.Create(&user)
}

// TODO: Untuk Dummy Role
func (s Seed) RoleSeed() {

	role := []models.Role{
		{
			Name:        "superadmin",
			Description: "superadmin",
		},
		{
			Name:        "admin",
			Description: "administrator",
		},
		{
			Name:        "customer",
			Description: "customer",
		},
	}

	s.db.Create(&role)
}

// TODO: Untuk Dummy Pivot User & Role
func (s Seed) UserRoleSeed() { // untuk pivot

	usersRole := []models.UserRole{
		{
			UserId: 1,
			RoleId: 1,
		},
		{
			UserId: 2,
			RoleId: 2,
		},
		{
			UserId: 3,
			RoleId: 3,
		},
	}

	s.db.Create(&usersRole)
}

// TODO: Untuk Dummy Permission
func (s Seed) PermissionSeed() {

	permission := []models.Permission{
		{
			Name:        "View Food",
			Description: "Akses Melihat Data Food",
		},
		{
			Name:        "Edit Food",
			Description: "Akses Mengubah Data Food",
		},
		{
			Name:        "Save Food",
			Description: "Akses Menyimpan Data Food",
		},
		{
			Name:        "Delete Food",
			Description: "Akses Menghapus Data Food",
		},
	}

	s.db.Create(&permission)
}

// TODO: Untuk Dummy Pivot Role & Permission
func (s Seed) RolePermissionSeed() { // untuk pivot

	rolePermission := []models.RolePermission{
		{
			RoleId:       1,
			PermissionId: 1,
		},
		{
			RoleId:       1,
			PermissionId: 2,
		},
		{
			RoleId:       1,
			PermissionId: 3,
		},
		{
			RoleId:       2,
			PermissionId: 1,
		},
		{
			RoleId:       3,
			PermissionId: 1,
		},
	}

	s.db.Create(&rolePermission)
}
