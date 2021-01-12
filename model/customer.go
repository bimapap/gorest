package model

import "time"

type Customer struct {
	Id        int        `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
