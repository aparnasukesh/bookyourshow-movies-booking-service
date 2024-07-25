package movies

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Duration    int       `gorm:"not null"`
	Genre       string    `gorm:"type:varchar(50)"`
	ReleaseDate time.Time `gorm:"not null"`
	Rating      float64   `gorm:"type:decimal(3,1)"`
	Language    string    `gorm:"type:varchar(100);not null"`
}
