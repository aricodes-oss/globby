package db

import (
	"gorm.io/gorm"
)

type Lobby struct {
	gorm.Model

	Name  string
	Users uint
}
