package controllers

import (
	"net/http"
	"server/config"
	"server/helpers"
	"server/models"
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
