package models

import "time"

// type User struct {
// 	ID           uint
// 	Name         string
// 	Email        *string
// 	Age          uint8
// 	Birthday     *time.Time
// 	MemberNumber sql.NullString
// 	ActivatedAt  sql.NullTime
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
//   }

type Book struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description int64  `gorm:"type:varchar(255)" json:"description"`
	Stock       int64  `gorm:"type:integer" json:"stock"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
