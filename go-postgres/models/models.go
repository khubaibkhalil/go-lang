package models

type Movie struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Watched bool   `gorm:"default:false"`
}
