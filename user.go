package belajargolanggorm

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;column:id;<-:create"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;<-:update"`
}

func (u *User) TableName() string {
	return "users"
}
