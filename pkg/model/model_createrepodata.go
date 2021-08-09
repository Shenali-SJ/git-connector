package model

import (
	"gorm.io/gorm"
)

type Createrepodata struct {
	Model             `bson:"-"`
	MID_              string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Name              string `bson:"name" json:"name" xml:"name"`
	Autoinit          bool   `bson:"autoinit" json:"autoinit" xml:"autoinit"`
	Private           bool   `bson:"private" json:"private" xml:"private"`
	Gitignoretemplate string `bson:"gitignoretemplate" json:"gitignoretemplate" xml:"gitignoretemplate"`
}

func (Createrepodata) TableName() string {
	return "createrepodata"
}
func (m *Createrepodata) PreloadCreaterepodata(db *gorm.DB) *gorm.DB {
	return db
}
