package models_db

type User struct {
	ID              uint64 `gorm:"primaryKey"`
	Username        string `gorm:"unique;not null;size:30"`
	DisplayUsername string `gorm:"size:30"`
	PasswordHash    string `gorm:"size:60;not null"`
	URLs            []URL  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
