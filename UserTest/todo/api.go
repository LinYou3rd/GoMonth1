package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	JSON_SUCCESS int = 1
	JSON_ERROR   int = 0
)

func Add(context *gin.Context) {

	completed, _ := strconv.Atoi(context.PostForm("completed"))
	todo := todoModel{Title: context.PostForm("title"), Context: context.PostForm("context"), Completed: completed}
	db.Save(&todo)

	context.JSON(http.StatusOK, gin.H{
		"status":     JSON_SUCCESS,
		"message":    "创建成功",
		"resourceId": todo.ID,
	})

}

func ShowAll(context *gin.Context) {

	var todos []todoModel
	var _todos []fmtTodo
	db.Find(&todos)

	if len(todos) <= 0 {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "没有待做事项可展示",
		})
		return
	}

	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, fmtTodo{
			Id:        item.ID,
			Title:     item.Title,
			Context:   item.Context,
			Completed: completed,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "成功展示所有",
		"data":    _todos,
	})

}

func ShowFinish(context *gin.Context) {

	var todos []todoModel
	var _todos []fmtTodo
	db.Where("completed=?", 1).Find(&todos)

	if len(todos) <= 0 {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "没有已完成事项可展示",
		})
		return
	}

	for _, item := range todos {
		_todos = append(_todos, fmtTodo{
			Id:      item.ID,
			Title:   item.Title,
			Context: item.Context,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "成功展示所有已完成事项",
		"data":    _todos,
	})

}

func ShowNoFinish(context *gin.Context) {

	var todos []todoModel
	var _todos []fmtTodo
	db.Where("completed=?", 0).Find(&todos)

	if len(todos) <= 0 {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "没有未完成事项可展示",
		})
		return
	}

	for _, item := range todos {
		_todos = append(_todos, fmtTodo{
			Id:      item.ID,
			Title:   item.Title,
			Context: item.Context,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "成功展示所有未完成事项",
		"data":    _todos,
	})

}

func ShowTarget(context *gin.Context) {

	var todo todoModel
	todoTitle := context.PostForm("title")
	db.First(&todo, todoTitle)

	if todo.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := fmtTodo{
		Id:        todo.ID,
		Title:     todo.Title,
		Context:   todo.Context,
		Completed: completed,
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "这就是你要找的",
		"data":    _todo,
	})

}

func UpdateOne(context *gin.Context) {

	var todo todoModel
	todoId := context.Param("id")
	db.First(&todo, todoId)

	if todo.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  JSON_ERROR,
			"message": "不存在该事项呢",
		})
		return
	}

	completed, _ := strconv.Atoi(context.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "更新成功",
	})

}

func UpdateALL(context *gin.Context) {

	var todos []todoModel
	var completed, _ = strconv.Atoi(context.Param("completed"))

	db.Find(&todos)

	for _, item := range todos {
		db.Model(&item).Update("completed", completed)
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "修改成功",
	})

}

func DeleteOne(context *gin.Context) {

	var todo todoModel
	todoId := context.Param("id")
	db.First(&todo, todoId)

	if todo.ID == 0 {
		context.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "虚空删除不可取",
		})
		return
	}

	db.Delete(&todo)

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "成功删除",
	})

}

func DeletePart(context *gin.Context) {

	var todos []todoModel
	db.Where("completed=?", context.Param("completed")).Find(&todos)
	for _, item := range todos {
		db.Delete(&item)
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "成功删除部分符合条件的事项",
	})

}

func DeleteAll(context *gin.Context) {

	var todos []todoModel
	db.Find(&todos)

	for _, item := range todos {
		db.Delete(item)
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "成功删除所有事项",
	})

}
