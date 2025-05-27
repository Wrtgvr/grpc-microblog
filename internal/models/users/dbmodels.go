package usersmodels

type User struct {
	ID              uint64 `gorm:"primaryKey"`
	Username        string `gorm:"unique;size:30;not null"`
	DisplayUsername string `gorm:"size:30"`
	PasswordHash    string `gorm:"size:60; not null"`
	Email           string `gorm:"size:254;unique;not null"`
}
