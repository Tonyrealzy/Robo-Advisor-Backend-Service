package config

import (
	"fmt"

	"gorm.io/gorm"
)

func DeleteRecordFromDb(db *gorm.DB, record interface{}) error {
	tx := db.Delete(record)
	return tx.Error
}

func DeleteByID(db *gorm.DB, model interface{}, id interface{}) error {
	result := db.Delete(model, id)
	return result.Error
}


func DeleteSpecificRecord(db *gorm.DB, model interface{}, query string, args ...interface{}) error {
	if err := db.Where(query, args...).Delete(model).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRecordWithNoModel(db *gorm.DB, query string, args ...interface{}) error {
	if err := db.Exec(query, args...).Error; err != nil {
		return err
	}
	return nil
}

func HardDeleteRecordFromDb(db *gorm.DB, record interface{}) error {
	tx := db.Unscoped().Delete(record)
	return tx.Error
}

func HardDeleteSpecificRecord(db *gorm.DB, model interface{}, query string, args ...interface{}) error {
	if err := db.Unscoped().Where(query, args...).Delete(model).Error; err != nil {
		return fmt.Errorf("failed to hard delete records: %v", err)
	}
	return nil
}
