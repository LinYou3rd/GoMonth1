package todo

import "github.com/jinzhu/gorm"

type (
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Context   string `json:"context"`
		Completed int    `json:"completed"`
	}

	fmtTodo struct {
		Id        uint   `json:"id"`
		Title     string `json:"title"`
		Context   string `json:"context"`
		Completed bool   `json:"completed"`
	}
)

func (todoModel) TableName() string {
	return "todos"
}
