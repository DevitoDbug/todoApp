package models

import (
	"github.com/DevitoDbug/todoApp/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Name     string `gorm:"" json:"name"`
	Status   string `json:"status"`
	Priority string `json:"priority"` //can be high,low or medium
	Notes    string `json:"notes"`
}

func (t *Task) CreateTask() (createdTask *Task) {
	db.NewRecord(t)
	db.Create(&t)
	return
}

func init() {
	config.ConnectToDB()
	db = config.GetDBConnection()
	db.AutoMigrate(&Task{})
	return
}

func GetAllTasks() *[]Task {
	var tasks []Task
	db.Find(&tasks)
	return &tasks
}

func GetTaskBYID(ID int64) (searchedTask *Task, dbInfo *gorm.DB) {
	var task Task
	db := db.Where("ID=?", ID).Find(&task)
	return &task, db
}

func DeleteTask(ID int64) (deletedTask *Task) {
	var task Task
	db.Where("ID=?", ID).Delete(&deletedTask)
	return &task
}
