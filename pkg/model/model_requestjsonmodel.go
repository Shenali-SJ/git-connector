package model

import (
	"gorm.io/gorm"
)

type Requestjsonmodel struct {
	Model     `bson:"-"`
	MID_      string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Operation string `bson:"operation" json:"operation" xml:"operation"`
	Data      string `bson:"data" json:"data" xml:"data"`
	Token     string `bson:"token" json:"token" xml:"token"`
}

func (Requestjsonmodel) TableName() string {
	return "requestjsonmodel"
}
func (m *Requestjsonmodel) PreloadRequestjsonmodel(db *gorm.DB) *gorm.DB {
	return db
}
