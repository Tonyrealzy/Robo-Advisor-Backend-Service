package config

import (
	"gorm.io/gorm"
)

func UpdateOneFieldByID(db *gorm.DB, model interface{}, id interface{}, field string, value interface{}) error {
	result := db.Model(model).Where("id = ?", id).Update(field, value)
	return result.Error
}

func UpdateModel(db *gorm.DB, model interface{}) error {
	result := db.Save(model)
	return result.Error
}
