package models

type User struct {
	ID string `gorm:"primary_key" json:"id"`
	UserName string `gorm:"type:varchar(38);not null;" json:"username"`
	PassWord string `gorm:"type:varchar(38);not null;" json:"password"`
	CreatedAt uint64 `gorm:"type:char(13);" json:"createdAt"`
}
