package model

import "time"

type User struct {
	BaseModel
	Email           string     `gorm:"column:email;type:varchar(255);unique;not null" json:"email"`
	Password        string     `gorm:"column:password;type:varchar(255);not null" json:"password"`
	ActivatedAt     *time.Time `gorm:"column:activated_at;type:timestamp" json:"activated_at"`
	FirstName       string     `gorm:"column:first_name;type:varchar(255);not null" json:"first_name"`
	LastName        string     `gorm:"column:last_name;type:varchar(255);not null" json:"last_name"`
	PhoneNumber     string     `gorm:"column:phone_number;type:varchar(255)" json:"phone_number"`
	Address         string     `gorm:"column:address;type:text" json:"address"`
	Address2        string     `gorm:"column:address2;type:text" json:"address2"`
	City            string     `gorm:"column:city;type:varchar(255)" json:"city"`
	State           string     `gorm:"column:state;type:varchar(255)" json:"state"`
	ZipCode         string     `gorm:"column:zip_code;type:varchar(255)" json:"zip_code"`
	ProfilePhotoURL string     `gorm:"column:profile_photo_url;type:text" json:"profile_photo_url"` // minio path
}

func (User) TableName() string {
	return "users"
}
