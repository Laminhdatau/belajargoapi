package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username" sql:"not null"`
	Password   string `gorm:"type:string" json:"password"`
}
