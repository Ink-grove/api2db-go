package repository

import (
	"api2db-server/middleware/db"
)

func DbGetTask(userId int64) (TaskData, error) {
	db := db.GetDB()
	task := TaskData{}
	err := db.Where("id = ?", userId).First(&task).Error
	if err != nil {
		return TaskData{}, err
	}
	return task, nil
}

func DbCreatTask(input TaskData) error {
	db := db.GetDB()
	err := db.Model(&TaskData{}).Create(&input).Error
	if err != nil {
		return err
	}
	return nil
}
