package models

import (
	"time"
)

type HeroModel struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	CreateDate time.Duration `json:"create_date"`
	UpdateDate time.Duration `json:"update_date"`
}
