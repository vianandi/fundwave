package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model       // Ini akan menambahkan ID, CreatedAt, UpdatedAt, dan DeletedAt
	Name             string
	Occupation       string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Token            string
}
