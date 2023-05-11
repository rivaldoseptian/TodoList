package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/config"
	"server/helpers"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	var todo []models.Todo
	activity_group_id := r.URL.Query().Get("activity_group_id")

	if activity_group_id == "" {
		if err := config.DB.Find(&todo).Error; err != nil {
			helpers.Response(w, 500, err.Error(), nil)
		}
	} else {
		if err := config.DB.Where("activity_group_id = ?", activity_group_id).Find(&todo).Error; err != nil {
			helpers.Response(w, 500, err.Error(), nil)
		}
	}
	helpers.Response(w, 200, "Success", todo)
}

func GetOneTodo(w http.ResponseWriter, r *http.Request) {
	idTodo := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTodo)

	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, fmt.Sprintf("Todo with ID %v Not Found", id), nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 200, "Success", todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if todo.ActivityGroupID == 0 {
		helpers.Response(w, 400, "activity_group_id cannot be null", nil)
		return
	}

	if todo.Title == "" {
		helpers.Response(w, 400, "title cannot be null", nil)
		return
	}
	if todo.Priority == "" {
		helpers.Response(w, 400, "priority cannot be null", nil)
		return
	}

	if err := config.DB.Create(&todo).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 201, "Success", todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idTodo := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idTodo)

	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, fmt.Sprintf("Todo with ID %v Not Found", id), nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if todo.ActivityGroupID == 0 {
		helpers.Response(w, 400, "activity_group_id cannot be null", nil)
		return
	}

	if todo.Title == "" {
		helpers.Response(w, 400, "title cannot be null", nil)
		return
	}
	if todo.Priority == "" {
		helpers.Response(w, 400, "priority cannot be null", nil)
		return
	}

	if err := config.DB.Where("todo_id = ? ", id).Updates(&todo).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 200, "Success", todo)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idTodo := mux.Vars(r)["id"]
	todo_id, _ := strconv.Atoi(idTodo)

	var todo models.Todo

	res := config.DB.Delete(&todo, todo_id)

	if res.Error != nil {
		helpers.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, 404, fmt.Sprintf("Todo with ID %v Not Found", todo_id), nil)
		return
	}
	helpers.Response(w, 200, "Success", nil)
}
