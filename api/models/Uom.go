package models

import "github.com/jinzhu/gorm"

type Uom struct {
	Code           string `gorm:"primary_key;" json:"code"`
	Name           string `gorm:"not null" json:"name"`
	PluralName     string `gorm:"not null;unique" json:"pluralName"`
	Description    string `gorm:"" json:"description"`
	UnitSystemCode string `gorm:"" json:"unitSystemCode"`
	Localization   string `gorm:"" json:"localization"`
	DataType       string `gorm:"not null;" json:"dataType"`
	Classification string `gorm:"not null;" json:"classification"`
	ReadOnly       bool   `gorm:"not null;" json:"readOnly"`
}

func (u *Uom) FindAllUoms(db *gorm.DB) (*[]Uom, error) {
	var err error
	Uoms := []Uom{}
	err = db.Table("measurements.uom").Model(&Uom{}).Limit(1000).Find(&Uoms).Error
	if err != nil {
		return &[]Uom{}, err
	}
	return &Uoms, err
}
