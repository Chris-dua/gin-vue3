package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"-"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
