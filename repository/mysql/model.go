package mysql

import "time"

type ActivityGroups struct {
	ID        int `gorm:"primaryKey"`
	Email     string
	Title     string
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type Todo struct {
	ID               int `gorm:"primaryKey"`
	ActivityGroupsID int `gorm:"index"`
	Title            string
	IsActive         bool      `gorm:"index"`
	Priority         string    `gorm:"index"`
	CreatedAt        time.Time `gorm:"index"`
	UpdatedAt        time.Time
	DeletedAt        *time.Time `gorm:"index"`
}
