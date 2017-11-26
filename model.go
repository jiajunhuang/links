package main

import (
	"github.com/jinzhu/gorm"
)

// Links short links service model
type Links struct {
	gorm.Model
	Original string `gorm:"size:768; not null; unique"`
}
