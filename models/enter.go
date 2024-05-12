package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primaryKey" json:"id" structs:"-"`
	CreatedAt time.Time `json:"create_at" structs:"-"`
	UpdatedAt time.Time `json:"-" `
}
type RemoveRequest struct {
	IDList []uint `json:"id_list" binding:"required" msg:"请输入id"`
}
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
