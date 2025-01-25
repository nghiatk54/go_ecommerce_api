package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Id       uint64 `gorm:"column:id; type:int; primaryKey; not null; autoIncrement; comment:'Primary key is Id'"`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
